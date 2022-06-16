package game

import (
	"fmt"
	"image/color"
	"math"
	"math/rand"

	"github.com/jdf/todd-again/engine"
)

const debugTodd = true

var debugColor = color.RGBA{0, 255, 0, 255}

type Todd struct {
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
}

func (t *Todd) String() string {
	return fmt.Sprintf("Todd pos %v vel %v)", t.pos, t.vel)
}

func (t *Todd) Gravity() float64 {
	if t.vel.Y > 0 && World.JumpState == JumpStateJumping {
		return Gravity * JumpStateGravityFactor
	}
	return Gravity
}

func (t *Todd) Update(s *engine.UpdateState) {
	dt := s.DeltaSeconds

	accel := Accel
	if !t.IsInContactWithGround() {
		accel = AirBending
	}

	if World.Controller.Left() {
		t.AdjustBearing(-BearingAccel * dt)
		t.AccelX(-accel * dt)
	} else if World.Controller.Right() {
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

	if rand.Float32() < 0.001 {
		t.Blink()
	}

	wantJump := World.Controller.Jump()

	t.AccelY(t.Gravity() * dt)
	if t.IsInContactWithGround() {
		if wantJump && World.JumpState == JumpStateIdle {
			t.Jump()
		} else if !wantJump && World.JumpState == JumpStateLanded {
			World.JumpState = JumpStateIdle
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
		for _, plat := range World.Platforms {
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
	if World.JumpState == JumpStateJumping {
		if colliding {
			t.eyeCenteringAnimation = NewAnimation(t.eyeCentering, 0, s.NowSeconds, EyeCenteringDurationSeconds)
			// blink on hard landing
			if oldvel > TerminalVelocity*0.95 {
				t.Blink()
			}
			t.vSquishVel = oldvel / 5.0
			if wantJump {
				World.JumpState = JumpStateLanded
			} else {
				World.JumpState = JumpStateIdle
			}
		}
	} else if !colliding {
		World.JumpState = JumpStateJumping // we fell off a platform
		// Squish, but, if already squishing, squish in that direction.
		if MaxSquishVel > math.Abs(t.vSquishVel) {
			t.vSquishVel = math.Copysign(MaxSquishVel, t.vSquishVel)
		}
		dir := Clockwise
		if t.vel.X < 0 || World.Controller.Left() {
			dir = CounterClockwise
		}
		t.tumbleAnimation = NewTumbleAnimation(dir, t.pos.Y)
		t.eyeCenteringAnimation = NewAnimation(t.eyeCentering,
			1, s.NowSeconds, EyeCenteringDurationSeconds)
	}

	if math.Abs(t.vSquishVel+t.vSquish) < 0.2 {
		// Squish damping when the energy is below threshold.
		t.vSquishVel = 0
		t.vSquish = 0
	} else {
		// squish stiffness
		const k = 100.0
		const damping = 8.5

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

func (t *Todd) Blink() {
	if t.blinkCumulativeTime == -1 {
		t.blinkCumulativeTime = 0
	}
}

func (t *Todd) Draw(g *engine.Graphics) {
	x, y := t.pos.X, t.pos.Y
	s := t.sideLength
	half := s / 2.0

	xsquish := t.vSquish * 0.8
	ysquish := t.vSquish * 1.6

	g.Push()
	g.Translate(x, y)
	if t.tumbleAnimation != nil {
		g.RotateAround(t.tumbleAnimation.AngleFor(y), engine.Vec(0, half))
	}
	width := s - xsquish
	height := s + ysquish
	g.DrawRoundedRect(engine.NewRect(-width/2, 0, width/2, height), s/8)
	g.SetColor(t.fillColor)
	g.Fill()

	if debugTodd {
		g.SetColor(debugColor)
		g.DrawLine(-.5, -.5, .5, .5)
		g.Stroke()
		g.DrawLine(-.5, .5, .5, -.5)
		g.Stroke()
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
	g.DrawEllipse(engine.NewRect(eyePos.X-5, eyePos.Y-5, eyePos.X+5, eyePos.Y+5))
	g.SetColor(color.White)
	g.Fill()
	g.DrawEllipse(engine.NewRect(pupilPos.X-1.5, pupilPos.Y-1.5, pupilPos.X+1.5, pupilPos.Y+1.5))
	g.SetColor(color.Black)
	g.Fill()

	if t.blinkCumulativeTime != -1 {
		blinkCycle := t.blinkCumulativeTime / BlinkCycleSeconds
		g.SetColor(t.fillColor)
		lidTop := eyePos.Y + 6
		lidBottom := lidTop - 12*math.Sin(math.Pi*blinkCycle)
		g.DrawRect(engine.NewRect(-half+3, lidBottom, half-3, lidTop))
		g.Fill()
	}

	g.Pop()
}

func (t *Todd) AccelX(a float64) {
	t.vel.X = Clamp(t.vel.X+a, -Maxvel, Maxvel)
}

func (t *Todd) AccelY(a float64) {
	t.vel.Y = t.vel.Y + a
	maxVel := TerminalVelocity
	if World.Controller.Jump() {
		maxVel = JumpTerminalVelocity
	}
	if t.vel.Y < maxVel {
		t.vel.Y = maxVel
	}
}

func (t *Todd) AdjustBearing(a float64) {
	t.bearing = Clamp(t.bearing+a, -Maxvel, Maxvel)
}

func (t *Todd) Jump() {
	World.JumpState = JumpStateJumping
	t.initialJumpSpeed = math.Abs(t.vel.X)
	t.vel.Y = GetJumpImpulse(t.initialJumpSpeed)
	t.vSquishVel = MaxSquishVel
}

func (t *Todd) ApplyFriction() {
	t.vel.X *= Friction
	if math.Abs(t.vel.X) < .01 {
		t.vel.X = 0
	}
}

func (t *Todd) ApplyBearingFriction() {
	t.bearing *= BearingFriction
}

func (t *Todd) Left() float64 {
	return t.pos.X - t.sideLength/2
}

func (t *Todd) Right() float64 {
	return t.pos.X + t.sideLength/2
}

func (t *Todd) IsInContactWithGround() bool {
	return t.GetContactHeight() != -1
}

// Y value of surface top or -1 for not landing.
func (t *Todd) GetContactHeight() float64 {
	if t.pos.Y <= 0 {
		return 0
	}
	// can't land if we're moving up
	if t.vel.Y > 0 {
		return -1
	}
	margin := PlatformMargin(t.vel.X)
	for _, plat := range World.Platforms {
		if t.pos.Y >= plat.bounds.Bottom() &&
			t.pos.Y <= plat.bounds.Top() &&
			t.Right() >= plat.bounds.Left()+margin &&
			t.Left() <= plat.bounds.Right()-margin {
			return plat.bounds.Top()
		}
	}
	return -1
}
