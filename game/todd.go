package game

import (
	"context"
	"fmt"
	"image/color"
	"math"
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/jdf/todd-again/engine"
)

const debugTodd = false

var debugColor = color.RGBA{0, 255, 0, 255}

type Dude struct {
	sideLength float64
	fillColor  color.Color
	// Todd's bottom-center.
	pos engine.Vec2
	vel engine.Vec2

	// The direction the eye is facing.
	bearing float64

	// Defect in vertical size, world coordinates.
	// Horizontal gets inverse.
	vSquish float64
	// Per second.
	vSquishVel float64

	blinkCumulativeTime float64

	// A quantity >= 0 when jumpState is jumpStateJumping.
	// Used in calculating gravity during jump.
	initialJumpSpeed float64

	// Tumbling!
	tumbleAnimation *TumbleAnimation
	// During tumbling, this animates to 1, which means completely centered.
	// When tumbling ends, this animates to 0, which means "where it wants to be".
	eyeCentering          float64
	eyeCenteringAnimation Animation

	rnd *rand.Rand
}

func (t *Dude) String() string {
	return fmt.Sprintf("Todd pos %v vel %v)", t.pos, t.vel)
}

func (t *Dude) Update(s *engine.UpdateState) {
	dt := s.DeltaSeconds

	accel := Accel
	if !t.IsInContactWithGround() {
		accel = AirBending
	}

	if Controller.Left() {
		t.AdjustBearing(-BearingAccel * dt)
		t.AccelX(-accel * dt)
	} else if Controller.Right() {
		t.AdjustBearing(BearingAccel * dt)
		t.AccelX(accel * dt)
	} else {
		t.ApplyBearingFriction()
		if t.IsInContactWithGround() {
			t.ApplyFriction()
		}
	}

	if t.blinkCumulativeTime != -1 {
		t.blinkCumulativeTime += dt
		if t.blinkCumulativeTime >= BlinkCycleSeconds {
			t.blinkCumulativeTime = -1
		}
	}

	if t.rnd.Float64() < BlinkOdds {
		t.Blink()
	}

	wantJump := Controller.Jump()

	gravity := Gravity
	if t.vel.Y > 0 && wantJump {
		gravity *= JumpStateGravityFactor
	}

	t.AccelY(gravity * dt)
	if t.IsInContactWithGround() {
		if wantJump && JumpState == JumpStateIdle {
			t.Jump()
		} else if !wantJump && JumpState == JumpStateLanded {
			JumpState = JumpStateIdle
		}
	}

	if math.Abs(t.vel.X) > 0 {
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
		margin := PlatformMargin(t.vel.X)
		for _, plat := range Platforms {
			if currentY >= plat.bounds.Top() && t.pos.Y <= plat.bounds.Top() &&
				t.Right() >= plat.bounds.Left()+margin &&
				t.Left() <= plat.bounds.Right()-margin {
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
			t.eyeCenteringAnimation = NewAnimation(t.eyeCentering, 0, s.NowSeconds, EyeCenteringDurationSeconds)
			// blink on hard landing
			if oldvel > TerminalVelocity*0.95 {
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
		const FallingSquishVel = .5 * MaxSquishVel
		if FallingSquishVel > math.Abs(t.vSquishVel) {
			t.vSquishVel = math.Copysign(FallingSquishVel, t.vSquishVel)
		}
		dir := Clockwise
		if t.vel.X < 0 || Controller.Left() {
			dir = CounterClockwise
		}
		t.tumbleAnimation = NewTumbleAnimation(dir, t.pos.Y)
		t.eyeCenteringAnimation = NewAnimation(t.eyeCentering,
			1, s.NowSeconds, EyeCenteringDurationSeconds)
	}

	if math.Abs(t.vSquishVel+t.vSquish) < 0.01 {
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
			t.vSquishVel, -MaxSquishVel, MaxSquishVel)
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

	g.Push()
	if t.tumbleAnimation != nil {
		g.RotateAround(t.tumbleAnimation.AngleFor(y), engine.Vec(0, half))
	}
	g.Translate(x, y)
	width := s - xsquish
	height := s + ysquish
	g.SetColor(t.fillColor)
	g.DrawRoundedRect(img, engine.NewRect(-width/2, 0, width/2, height), s/8)

	if debugTodd {
		g.SetColor(debugColor)
		g.DrawLine(img, -.5, -.5, .5, .5)
		g.DrawLine(img, -.5, .5, .5, -.5)
	}

	speedRatio := math.Abs(t.bearing / Maxvel)
	eyeVCenter := half + 4 + t.vSquish
	eyeOffset := Lerp(0, half-6, speedRatio)
	pupilOffset := Lerp(0, half-3, speedRatio)

	eyePos := &engine.Vec2{X: math.Copysign(eyeOffset, t.bearing), Y: eyeVCenter}
	pupilPos := &engine.Vec2{X: math.Copysign(pupilOffset, t.bearing), Y: eyeVCenter}
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
		blinkCycle := t.blinkCumulativeTime / BlinkCycleSeconds
		lidTop := eyePos.Y + 6
		lidBottom := lidTop - 12*math.Sin(math.Pi*blinkCycle)
		g.SetColor(t.fillColor)
		g.DrawRect(img, engine.NewRect(-half+1, lidBottom, half-1, lidTop))
	}

	g.Pop()
}

func (t *Dude) AccelX(a float64) {
	t.vel.X = Clamp(t.vel.X+a, -Maxvel, Maxvel)
}

func (t *Dude) AccelY(a float64) {
	t.vel.Y = t.vel.Y + a
	maxVel := TerminalVelocity
	if Controller.Jump() {
		maxVel = JumpTerminalVelocity
	}
	if t.vel.Y < maxVel {
		t.vel.Y = maxVel
	}
}

func (t *Dude) AdjustBearing(a float64) {
	t.bearing = Clamp(t.bearing+a, -Maxvel, Maxvel)
}

func (t *Dude) Jump() {
	JumpState = JumpStateJumping
	t.initialJumpSpeed = math.Abs(t.vel.X)
	t.vel.Y = GetJumpImpulse(t.initialJumpSpeed)
	t.vSquishVel = MaxSquishVel
}

func (t *Dude) ApplyFriction() {
	t.vel.X *= Friction
	if math.Abs(t.vel.X) < .01 {
		t.vel.X = 0
	}
}

func (t *Dude) ApplyBearingFriction() {
	t.bearing *= BearingFriction
}

func (t *Dude) Left() float64 {
	return t.pos.X - t.sideLength/2
}

func (t *Dude) Right() float64 {
	return t.pos.X + t.sideLength/2
}

func (t *Dude) IsInContactWithGround() bool {
	return t.GetContactHeight() != -1
}

// Y value of surface top or -1 for not landing.
func (t *Dude) GetContactHeight() float64 {
	if t.pos.Y <= 0 {
		return 0
	}
	// can't land if we're moving up
	if t.vel.Y > 0 {
		return -1
	}
	margin := PlatformMargin(t.vel.X)
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
