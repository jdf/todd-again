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

func (t *Todd) AccelX(a float64) {
	t.vel.X = Clamp(t.vel.X+a, -Maxvel, Maxvel)
}

func (t *Todd) AccelY(a float64) {
	t.vel.Y = t.vel.Y + a
	maxVel := TerminalVelocity
	if ToddController.Jump() {
		maxVel = JumpTerminalVelocity
	}
	if t.vel.Y > maxVel {
		t.vel.Y = maxVel
	}
}

func (t *Todd) AdjustBearing(a float64) {
	t.bearing = Clamp(t.bearing+a, -Maxvel, Maxvel)
}

func (t *Todd) Jump() {
	JumpState = JumpStateJumping
	World.setJumpState(World.JumpState.jumping)
	t.initialJumpSpeed = Math.abs(t.vel.x)
	t.vel.y = t.getJumpImpulse(t.initialJumpSpeed)
	t.vSquishVel = MaxSquishVel
}

/*

  func (t* Todd) applyFriction() {
    t.vel.x *= Constants.friction;
  }

 func (t* Todd)  applyBearingFriction() {
    t.bearing *= Constants.bearingFriction;
  }

  func (t* Todd) left() {
    return t.pos.x - t.sideLength / 2;
  }

  func (t* Todd) right() {
    return t.pos.x + t.sideLength / 2;
  }

  // Y value of surface top or -1 for not landing.
 func (t* Todd)  getContactHeight() {
    if (t.pos.y >= height) {
      return height;
    }
    // can't land if we're moving up
    if (t.vel.y < 0) {
      return -1;
    }
    const margin = t.platformMargin(t.vel.x);
    for (const plat of World.platforms) {
      if (t.pos.y >= plat.top &&
          t.pos.y <= plat.bottom &&
          t.right() >= plat.left + margin &&
          t.left() <= plat.right - margin) {
        return plat.top;
      }
    }
    return -1;
  }

 func (t* Todd)  move(dt) {
    t.yAccel(t.getGravity() * dt);
    if (t.isInContactWithGround()) {
      if (World.controller.jump && World.isJumpIdle()) {
        t.jump();
      } else if (!World.controller.jump && World.isJumpLanded()) {
        World.setJumpState(World.JumpState.idle);
      }
    }


    t.pos.x += t.vel.x * dt;
    if (t.pos.x > width) {
      t.pos.x = 0;
    } else if (t.pos.x < 0) {
      t.pos.x = width;
    }

    // Collisions.
    const currentY = t.pos.y;
    t.pos.y = currentY + t.vel.y * dt;
    let colliding = false;
    if (t.pos.y >= height) {
      colliding = true;
      t.pos.y = height;
    } else {
      const margin = t.platformMargin(t.vel.x);
      for (const plat of World.platforms) {
        if (currentY <= plat.top && t.pos.y >= plat.top &&
            t.right() >= plat.left + margin &&
            t.left() <= plat.right - margin) {
          colliding = true;
          t.pos.y = plat.top;
          break;
        }
      }
    }
    const oldvel = t.vel.y;
    if (colliding) {
      t.vel.y = 0;
      t.tumbleAnimation = null;
    }
    if (World.isJumpJumping()) {
      if (colliding) {
        t.eyeCenteringAnimation = new TimeBasedAnimation(t.eyeCentering,
            0, Constants.eyeCenteringDurationSeconds);
        // blink on hard landing
        if (oldvel > Constants.terminalVelocity * 0.95) {
          t.blink();
        }
        t.vSquishVel = -oldvel / 5.0;
        World.setJumpState(
            World.controller.jump ? World.JumpState.landed
                : World.JumpState.idle);
      }
    } else if (!colliding) {
      World.setJumpState(World.JumpState.jumping);  // we fell off a platform
      // Squish, but, if already squishing, squish in that direction.
      if (MaxSquishVel > Math.abs(t.vSquishVel)) {
        t.vSquishVel = MaxSquishVel * Math.sign(t.vSquishVel);
      }
      let sign = Math.sign(t.vel.x);
      if (World.controller.left) {
        sign = -1;
      } else if (World.controller.right) {
        sign = 1;
      }
      t.tumbleAnimation = new TumbleAnimation(sign, t.pos.y);
      t.eyeCenteringAnimation = new TimeBasedAnimation(t.eyeCentering,
          1, Constants.eyeCenteringDurationSeconds);
    }

    if (Math.abs(t.vSquishVel + t.vSquish) < 0.2) {
      // Squish damping when the energy is below threshold.
      t.vSquishVel = t.vSquish = 0;
    } else {
      // squish stiffness
      const k = 200.0;
      const damping = 8.5;

      const squishForce = -k * t.vSquish;
      const dampingForce = damping * t.vSquishVel;
      t.vSquishVel +=
          (squishForce - dampingForce) * dt;
      t.vSquishVel = clamp(
          t.vSquishVel, -MaxSquishVel, MaxSquishVel);
      t.vSquish += t.vSquishVel * dt;
    }

    if (t.eyeCenteringAnimation != null) {
      t.eyeCentering = t.eyeCenteringAnimation.value();
      if (t.eyeCenteringAnimation.isDone()) {
        t.eyeCenteringAnimation = null;
      }
    }
  }

  isInContactWithGround() {
    return t.getContactHeight() !== -1;
  }
}
*/
