package game

import (
	"fmt"
	"image/color"
	"math"

	"github.com/jdf/todd-again/engine"
)

func nextTumbleHeight(height float64) float64 {
	for _, h := range TumbleLevels {
		if h < height {
			return h
		}
	}
	panic(fmt.Errorf("nowhere to tumble to from %v", height))
}

type TumbleAnimation struct {
	sign         float64
	startHeight  float64
	startAngle   float64
	targetHeight float64
}

type RotationDirection int

const (
	CounterClockwise = RotationDirection(1)
	Clockwise        = RotationDirection(-1)
)

func NewTumbleAnimation(direction RotationDirection, startHeight float64) *TumbleAnimation {
	return &TumbleAnimation{
		sign:         float64(direction),
		startHeight:  startHeight,
		startAngle:   0,
		targetHeight: nextTumbleHeight(startHeight),
	}
}

func (t *TumbleAnimation) AngleFor(height float64) float64 {
	if height >= t.startHeight {
		return t.startAngle
	}
	if height <= t.targetHeight {
		t.startHeight = t.targetHeight
		t.targetHeight = nextTumbleHeight(t.startHeight)
		t.startAngle += (math.Pi / 2.0) * float64(t.sign)
	}
	return t.startAngle + (height-t.startHeight)*t.sign
}

type Todd struct {
	sideLength float64
	fillColor  color.Color
	pos        engine.Vec2
	vel        engine.Vec2

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

/*

  // Maps an x-velocity to an integer in [0,2].
  speedStepFunction(xVel) {
    xVel = Math.abs(xVel);
    if (xVel < Constants.maxvel * .333) {
      return 0;
    }
    if (xVel < Constants.maxvel * .666) {
      return 1;
    }
    return 2;
  }

  // The platform has a width differing from its apparent width, depending on
  // your speed. The slower you are, the narrower the platform is.
  platformMargin(xVel) {
    return [8, 0, -5][this.speedStepFunction(xVel)];
  }

  getJumpImpulse(speed) {
    return Constants.jumpImpulse * [1.0, 1.0, 1.2][this.speedStepFunction(
        speed)];
  }

  getGravity() {
    if (this.vel.y < 0) {
      // going up!
      return Constants.gravity * (World.controller.jump ? 0.55 : 1.0);
    }
    return Constants.gravity;
  }

  setPos(x, y) {
    this.pos = {
      x,
      y
    };
  }

  blink() {
    if (this.blinkCumulativeTime !== -1) {
      return;
    }
    this.blinkCumulativeTime = 0;
  }

  drawAt(x, y) {
    p.rectMode(p.CENTER);
    p.fill(this.fillColor);
    p.noStroke();

    const s = this.sideLength;
    const half = s / 2;
    const xsquish = this.vSquish * 0.8;
    const ysquish = this.vSquish * 1.6;

    p.push();
    p.translate(x, y);
    if (this.tumbleAnimation != null) {
      p.translate(0, -half);
      p.rotate(this.tumbleAnimation.angleFor(this.pos.y));
      p.translate(0, half);
    }
    p.rect(0, -(half + ysquish / 2.0), s - xsquish, s + ysquish, 3, 3);

    const eyeVCenter = -s + 8 - this.vSquish;
    const eyeOffset = p.lerp(0, half - 6,
        Math.abs(this.bearing / Constants.maxvel));
    const pupilOffset = p.lerp(0, half - 3,
        Math.abs(this.bearing / Constants.maxvel));
    let eyePos = {x: eyeOffset * Math.sign(this.bearing), y: eyeVCenter};
    let pupilPos = {x: pupilOffset * Math.sign(this.bearing), y: eyeVCenter};
    if (this.eyeCentering !== 0) {
      const center = {x: 0, y: -half};
      eyePos = lerpVec(eyePos, center, this.eyeCentering);
      pupilPos = lerpVec(pupilPos, center, this.eyeCentering);
    }
    p.fill(255, 255, 255);
    p.ellipse(eyePos.x, eyePos.y, 10, 10);
    p.fill(0);
    p.ellipse(pupilPos.x, pupilPos.y, 3, 3);

    p.rectMode(p.CORNERS);
    if (this.blinkCumulativeTime !== -1) {
      const blinkCycle = this.blinkCumulativeTime / Constants.blinkCycleSeconds;
      p.fill(this.fillColor);
      const lidTopY = eyePos.y - 6;
      const lidBottomY = lidTopY + 12 * Math.sin(Math.PI * blinkCycle);
      p.rect(-half + 3, lidTopY, half - 3, lidBottomY);
    }
    p.pop();
  }

  draw() {
    const {
      x,
      y
    } = this.pos;
    this.drawAt(x, y);
    const half = this.sideLength / 2;
    if (x < half) {
      this.drawAt(width + x, y);
    } else if (x > width - half) {
      this.drawAt(x - width, y);
    }
  }

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

    if (this.blinkCumulativeTime !== -1) {
      this.blinkCumulativeTime += dt;
    }
    if (this.blinkCumulativeTime >= Constants.blinkCycleSeconds) {
      this.blinkCumulativeTime = -1;
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
