package game

import (
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
	eyeCenteringAnimation *Animation
}

func (t *Todd) Gravity() float64 {
	if t.vel.Y > 0 && JumpState == JumpStateJumping {
		return Gravity * JumpStateGravityFactor
	}
	return Gravity
}

func (t *Todd) Update(s *engine.UpdateState) {

	if t.blinkCumulativeTime != -1 {
		t.blinkCumulativeTime += s.DeltaSeconds
		if t.blinkCumulativeTime >= BlinkCycleSeconds {
			t.blinkCumulativeTime = -1
		}
	}

	if rand.Float32() < 0.001 {
		t.Blink()
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
	eyeVCenter := half - t.vSquish
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

/*
  xAccel(a) {
    this.vel.x = clamp(this.vel.x + a, -Constants.maxvel, Constants.maxvel);
  }

  adjustBearing(a) {
    this.bearing =
        clamp(this.bearing + a, -Constants.maxvel, Constants.maxvel);
  }

  yAccel(a) {
    this.vel.y = this.vel.y + a;
    const term = World.controller.jump ? Constants.jumpTerminalVelocity
        : Constants.terminalVelocity;
    if (this.vel.y > term) {
      this.vel.y = term;
    }
  }

  jump() {
    World.setJumpState(World.JumpState.jumping);
    this.initialJumpSpeed = Math.abs(this.vel.x);
    this.vel.y = this.getJumpImpulse(this.initialJumpSpeed);
    this.vSquishVel = Constants.maxSquishVel;
  }

  applyFriction() {
    this.vel.x *= Constants.friction;
  }

  applyBearingFriction() {
    this.bearing *= Constants.bearingFriction;
  }

  left() {
    return this.pos.x - this.sideLength / 2;
  }

  right() {
    return this.pos.x + this.sideLength / 2;
  }

  // Y value of surface top or -1 for not landing.
  getContactHeight() {
    if (this.pos.y >= height) {
      return height;
    }
    // can't land if we're moving up
    if (this.vel.y < 0) {
      return -1;
    }
    const margin = this.platformMargin(this.vel.x);
    for (const plat of World.platforms) {
      if (this.pos.y >= plat.top &&
          this.pos.y <= plat.bottom &&
          this.right() >= plat.left + margin &&
          this.left() <= plat.right - margin) {
        return plat.top;
      }
    }
    return -1;
  }

  move(dt) {
    this.yAccel(this.getGravity() * dt);
    if (this.isInContactWithGround()) {
      if (World.controller.jump && World.isJumpIdle()) {
        this.jump();
      } else if (!World.controller.jump && World.isJumpLanded()) {
        World.setJumpState(World.JumpState.idle);
      }
    }


    this.pos.x += this.vel.x * dt;
    if (this.pos.x > width) {
      this.pos.x = 0;
    } else if (this.pos.x < 0) {
      this.pos.x = width;
    }

    // Collisions.
    const currentY = this.pos.y;
    this.pos.y = currentY + this.vel.y * dt;
    let colliding = false;
    if (this.pos.y >= height) {
      colliding = true;
      this.pos.y = height;
    } else {
      const margin = this.platformMargin(this.vel.x);
      for (const plat of World.platforms) {
        if (currentY <= plat.top && this.pos.y >= plat.top &&
            this.right() >= plat.left + margin &&
            this.left() <= plat.right - margin) {
          colliding = true;
          this.pos.y = plat.top;
          break;
        }
      }
    }
    const oldvel = this.vel.y;
    if (colliding) {
      this.vel.y = 0;
      this.tumbleAnimation = null;
    }
    if (World.isJumpJumping()) {
      if (colliding) {
        this.eyeCenteringAnimation = new TimeBasedAnimation(this.eyeCentering,
            0, Constants.eyeCenteringDurationSeconds);
        // blink on hard landing
        if (oldvel > Constants.terminalVelocity * 0.95) {
          this.blink();
        }
        this.vSquishVel = -oldvel / 5.0;
        World.setJumpState(
            World.controller.jump ? World.JumpState.landed
                : World.JumpState.idle);
      }
    } else if (!colliding) {
      World.setJumpState(World.JumpState.jumping);  // we fell off a platform
      // Squish, but, if already squishing, squish in that direction.
      if (Constants.maxSquishVel > Math.abs(this.vSquishVel)) {
        this.vSquishVel = Constants.maxSquishVel * Math.sign(this.vSquishVel);
      }
      let sign = Math.sign(this.vel.x);
      if (World.controller.left) {
        sign = -1;
      } else if (World.controller.right) {
        sign = 1;
      }
      this.tumbleAnimation = new TumbleAnimation(sign, this.pos.y);
      this.eyeCenteringAnimation = new TimeBasedAnimation(this.eyeCentering,
          1, Constants.eyeCenteringDurationSeconds);
    }

    if (Math.abs(this.vSquishVel + this.vSquish) < 0.2) {
      // Squish damping when the energy is below threshold.
      this.vSquishVel = this.vSquish = 0;
    } else {
      // squish stiffness
      const k = 200.0;
      const damping = 8.5;

      const squishForce = -k * this.vSquish;
      const dampingForce = damping * this.vSquishVel;
      this.vSquishVel +=
          (squishForce - dampingForce) * dt;
      this.vSquishVel = clamp(
          this.vSquishVel, -Constants.maxSquishVel, Constants.maxSquishVel);
      this.vSquish += this.vSquishVel * dt;
    }

    if (this.eyeCenteringAnimation != null) {
      this.eyeCentering = this.eyeCenteringAnimation.value();
      if (this.eyeCenteringAnimation.isDone()) {
        this.eyeCenteringAnimation = null;
      }
    }
  }

  isInContactWithGround() {
    return this.getContactHeight() !== -1;
  }
}
*/
