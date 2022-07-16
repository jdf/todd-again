package game

import (
	"context"
	"fmt"
	"image/color"
	"math"
	"math/rand"

	"git.maze.io/go/math32"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/jdf/todd-again/engine"
	"github.com/jdf/todd-again/game/level"
)

const debugTodd = false

var (
	debugColor = color.RGBA{0, 255, 0, 255}

	jumpRequestTimerSeconds   float32 = 0.0
	groundingSlopTimerSeconds float32 = 0.0
)

type Dude struct {
	sideLength float32
	// Todd's bottom-center.
	pos engine.Vec2
	vel engine.Vec2

	// The direction the eye is facing.
	bearing float32

	// Defect in vertical size, world coordinates.
	// Horizontal gets inverse.
	vSquish float32
	// Per second.
	vSquishVel float32

	blinkCumulativeTime float32

	// A quantity >= 0 when jumpState is jumpStateJumping.
	// Used in calculating gravity during jump.
	initialJumpSpeed float32

	// Tumbling!
	tumbleAnimation *TumbleAnimation
	// During tumbling, this animates to 1, which means completely centered.
	// When tumbling ends, this animates to 0, which means "where it wants to be".
	eyeCentering          float32
	eyeCenteringAnimation Animation

	rnd *rand.Rand
}

func (t *Dude) String() string {
	return fmt.Sprintf("Todd pos %v vel %v)", t.pos, t.vel)
}

func (t *Dude) Update(s *engine.UpdateState) {
	dt := s.DeltaSeconds

	grounded := t.Grounded()

	if grounded {
		groundingSlopTimerSeconds = level.Instance.Todd.GetGroundingSlopSeconds()
	} else {
		groundingSlopTimerSeconds -= dt
	}

	canJump := groundingSlopTimerSeconds > 0

	accel := level.Instance.Todd.GetAcceleration()
	if !grounded {
		accel = level.Instance.Todd.GetAirBending()
	}

	if Controller.Left() {
		t.AdjustBearing(-level.Instance.Todd.GetBearingAcceleration() * dt)
		t.AccelX(-accel * dt)
	} else if Controller.Right() {
		t.AdjustBearing(level.Instance.Todd.GetBearingAcceleration() * dt)
		t.AccelX(accel * dt)
	} else {
		t.ApplyBearingFriction()
		t.ApplyFriction()
	}

	if t.blinkCumulativeTime != -1 {
		t.blinkCumulativeTime += dt
		if t.blinkCumulativeTime >= level.Instance.Todd.Blink.GetCycleSeconds() {
			t.blinkCumulativeTime = -1
		}
	}

	if t.rnd.Float32() < level.Instance.Todd.Blink.GetOdds() {
		t.Blink()
	}

	if Controller.Jump() {
		jumpRequestTimerSeconds = level.Instance.Todd.GetJumpRequestSlopSeconds()
	}
	wantJump := jumpRequestTimerSeconds > 0
	if wantJump {
		jumpRequestTimerSeconds = 0
	} else {
		jumpRequestTimerSeconds -= dt
	}

	gravity := level.Instance.World.GetGravity()
	if t.vel.Y > 0 && wantJump {
		gravity *= level.Instance.Todd.GetJumpStateGravityFactor()
	}

	t.AccelY(gravity * dt)
	if canJump {
		if wantJump && JumpState == JumpStateIdle {
			t.Jump()
		} else if !wantJump && JumpState == JumpStateLanded {
			JumpState = JumpStateIdle
		}
	}

	if math32.Abs(t.vel.X) > 0 {
		t.pos.X += t.vel.X * dt
		if t.Right() > WorldBounds.Right() {
			t.pos.X = WorldBounds.Right() - t.sideLength/2
		} else if t.Left() < WorldBounds.Left() {
			t.pos.X = t.sideLength / 2
		}
	}

	// Collisions.
	currentY := t.pos.Y
	t.pos.Y = currentY + t.vel.Y*dt
	colliding := false
	if t.pos.Y <= 0 {
		colliding = true
		t.pos.Y = 0
	} else {
		margin := level.PlatformMargin(t.vel.X)
		for _, plat := range Platforms {
			if currentY >= plat.bounds.Top() && t.pos.Y <= plat.bounds.Top() &&
				t.Right() >= plat.bounds.Left()-margin &&
				t.Left() <= plat.bounds.Right()+margin {
				colliding = true
				t.pos.Y = plat.bounds.Top()
				break
			}
		}
	}
	oldvel := t.vel.Y
	if colliding {
		t.vel.Y = 0
		t.tumbleAnimation = nil
	}
	if JumpState == JumpStateJumping {
		if colliding {
			t.eyeCenteringAnimation = NewAnimation(
				t.eyeCentering,
				0,
				s.NowSeconds,
				level.Instance.Todd.GetEyeCenteringDurationSeconds())
			// blink on hard landing
			if oldvel > level.Instance.Todd.GetTerminalVelocity()*0.95 {
				t.Blink()
			}
			t.vSquishVel = oldvel / 5.0
			if wantJump {
				JumpState = JumpStateLanded
			} else {
				JumpState = JumpStateIdle
			}
			Bus.Emit(context.Background(), ToddVerticalLevelChanged, t.pos)
		}
	} else if !colliding {
		JumpState = JumpStateJumping // we fell off a platform
		// Squish, but, if already squishing, squish in that direction.
		FallingSquishVel := .5 * level.Instance.Todd.GetMaxSquishVelocity()
		if FallingSquishVel > math32.Abs(t.vSquishVel) {
			t.vSquishVel = math32.Copysign(FallingSquishVel, t.vSquishVel)
		}
		dir := Clockwise
		if t.vel.X < 0 || Controller.Left() {
			dir = CounterClockwise
		}
		t.tumbleAnimation = NewTumbleAnimation(dir, t.pos.Y)
		t.eyeCenteringAnimation = NewAnimation(t.eyeCentering,
			1, s.NowSeconds, level.Instance.Todd.GetEyeCenteringDurationSeconds())
	}

	if math32.Abs(t.vSquishVel+t.vSquish) < 0.01 {
		// Squish damping when the energy is below threshold.
		t.vSquishVel = 0
		t.vSquish = 0
	} else {
		// squish stiffness
		const k = 100.0
		const damping = 8.0 //8.5

		squishForce := -k * t.vSquish
		dampingForce := damping * t.vSquishVel
		t.vSquishVel += (squishForce - dampingForce) * dt
		t.vSquishVel = Clamp(
			t.vSquishVel, -level.Instance.Todd.GetMaxSquishVelocity(), level.Instance.Todd.GetMaxSquishVelocity())
		t.vSquish += t.vSquishVel * dt
	}

	if t.eyeCenteringAnimation != nil {
		t.eyeCentering = t.eyeCenteringAnimation.Value(s.NowSeconds)
		if t.eyeCenteringAnimation.IsDone(s.NowSeconds) {
			t.eyeCenteringAnimation = nil
		}
	}
}

func (t *Dude) Blink() {
	if t.blinkCumulativeTime == -1 {
		t.blinkCumulativeTime = 0
	}
}

func (t *Dude) Draw(img *ebiten.Image, g *engine.Graphics) {
	x, y := t.pos.X, t.pos.Y
	s := t.sideLength
	half := s / 2.0

	xsquish := t.vSquish * 0.8
	ysquish := t.vSquish * 1.6

	fillColor := level.RGBA(level.Instance.Todd.GetColor())
	g.Push()
	if t.tumbleAnimation != nil {
		g.RotateAround(t.tumbleAnimation.AngleFor(y), engine.Vec(0, half))
	}
	g.Translate(x, y)
	width := s - xsquish
	height := s + ysquish
	g.SetColor(fillColor)
	g.DrawRoundedRect(img, engine.NewRect(-width/2, 0, width/2, height), s/8)

	if debugTodd {
		g.SetColor(debugColor)
		g.DrawLine(img, -.5, -.5, .5, .5)
		g.DrawLine(img, -.5, .5, .5, -.5)
	}

	speedRatio := math32.Abs(t.bearing / level.Instance.Todd.GetMaxVelocity())
	eyeVCenter := half + 4 + t.vSquish
	eyeOffset := Lerp(0, half-6, speedRatio)
	pupilOffset := Lerp(0, half-3, speedRatio)

	eyePos := &engine.Vec2{X: math32.Copysign(eyeOffset, t.bearing), Y: eyeVCenter}
	pupilPos := &engine.Vec2{X: math32.Copysign(pupilOffset, t.bearing), Y: eyeVCenter}
	if t.eyeCentering != 0 {
		center := &engine.Vec2{X: 0, Y: half}
		eyePos = LerpVec(eyePos, center, t.eyeCentering)
		pupilPos = LerpVec(pupilPos, center, t.eyeCentering)
	}
	g.SetColor(color.White)
	g.DrawEllipse(img, engine.NewRect(eyePos.X-5, eyePos.Y-5, eyePos.X+5, eyePos.Y+5))
	g.SetColor(color.Black)
	g.DrawEllipse(img, engine.NewRect(pupilPos.X-1.5, pupilPos.Y-1.5, pupilPos.X+1.5, pupilPos.Y+1.5))

	if t.blinkCumulativeTime != -1 {
		blinkCycle := t.blinkCumulativeTime / level.Instance.Todd.Blink.GetCycleSeconds()
		lidTop := eyePos.Y + 6
		lidBottom := lidTop - 12*math32.Sin(math.Pi*blinkCycle)
		g.SetColor(fillColor)
		g.DrawRect(img, engine.NewRect(-half+1, lidBottom, half-1, lidTop))
	}

	g.Pop()
}

func (t *Dude) AccelX(a float32) {
	maxvel := level.Instance.Todd.GetMaxVelocity()
	t.vel.X = Clamp(t.vel.X+a, -maxvel, maxvel)
}

func (t *Dude) AccelY(a float32) {
	t.vel.Y = t.vel.Y + a
	maxVel := level.Instance.Todd.GetTerminalVelocity()
	if Controller.Jump() {
		maxVel = level.Instance.Todd.GetJumpTerminalVelocity()
	}
	if t.vel.Y < maxVel {
		t.vel.Y = maxVel
	}
}

func (t *Dude) AdjustBearing(a float32) {
	maxvel := level.Instance.Todd.GetMaxVelocity()
	t.bearing = Clamp(t.bearing+a, -maxvel, maxvel)
}

func (t *Dude) Jump() {
	JumpState = JumpStateJumping
	t.initialJumpSpeed = math32.Abs(t.vel.X)
	t.vel.Y = level.GetJumpImpulse(t.initialJumpSpeed)
	t.vSquishVel = level.Instance.Todd.GetMaxSquishVelocity()
}

func (t *Dude) ApplyFriction() {
	if t.Grounded() {
		t.vel.X *= level.Instance.Todd.GetFriction()
	} else {
		t.vel.X *= level.Instance.Todd.GetAirFriction()
	}

	if math32.Abs(t.vel.X) < 1 {
		t.vel.X = 0
	}
}

func (t *Dude) ApplyBearingFriction() {
	t.bearing *= level.Instance.Todd.GetBearingFriction()
}

func (t *Dude) Left() float32 {
	return t.pos.X - t.sideLength/2
}

func (t *Dude) Right() float32 {
	return t.pos.X + t.sideLength/2
}

func (t *Dude) Grounded() bool {
	return t.GetContactHeight() != -1
}

// Y value of surface top or -1 for not landing.
func (t *Dude) GetContactHeight() float32 {
	if t.pos.Y <= 0 {
		return 0
	}
	// can't land if we're moving up
	if t.vel.Y > 0 {
		return -1
	}
	margin := level.PlatformMargin(t.vel.X)
	for _, plat := range Platforms {
		if t.pos.Y >= plat.bounds.Bottom() &&
			t.pos.Y <= plat.bounds.Top() &&
			t.Right() >= plat.bounds.Left()+margin &&
			t.Left() <= plat.bounds.Right()-margin {
			return plat.bounds.Top()
		}
	}
	return -1
}
