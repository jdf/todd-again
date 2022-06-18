package game

import (
	"math"

	"github.com/tanema/gween/ease"
)

// bg color
const BG = 0

// length of dude's side
const ToddSideLength = 30.0

// A unitless constant that we apply to velocity while on the ground.
const Friction = 0.97

// A unitless constant that we apply to bearing when not accelerating.
const BearingFriction = 0.95

// The following constants are pixels per second.
const Gravity = -1200.0
const Maxvel = 240.0
const Accel = 900.0
const AirBending = 575.0
const BearingAccel = 1200.0
const JumpImpulse = 350.0

const MaxSquishVel = 60.0

// Max vertical velocity while holding down jump.
const JumpTerminalVelocity = -350.0
const TerminalVelocity = -550.0

// Blinking.
const BlinkOdds = 1 / 3000.0
const BlinkCycleSeconds = 0.25

// Eye centering speed for tumbling/landing.
const EyeCenteringDurationSeconds = 0.25

const JumpStateGravityFactor = 0.55

const CameraTiltSeconds = 0.2

var CameraTiltEasing = ease.Linear

const (
	Step1 = Maxvel * .333
	Step2 = Maxvel * .666
)

func SpeedStepFunction(v float64) int {
	v = math.Abs(v)
	switch {
	case v < Step1:
		return 0
	case v < Step2:
		return 1
	default:
		return 2
	}
}

// The faster you're going, the harder you jump.
var JumpImpulseFactors = []float64{1, 1, 1.2}

func GetJumpImpulse(speed float64) float64 {
	return JumpImpulse * JumpImpulseFactors[SpeedStepFunction(speed)]
}

// The platform has a width differing from its apparent width, depending on
// your speed. The slower you are, the narrower the platform is.
var PlatformMargins = []float64{8, 0, -5}

func PlatformMargin(xVel float64) float64 {
	return PlatformMargins[SpeedStepFunction(xVel)]
}
