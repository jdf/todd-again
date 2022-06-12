package game

// bg color
const BG = 0

// length of dude's side
const Side = 30

// A unitless constant that we apply to velocity while on the ground.
const Friction = 0.83

// A unitless constant that we apply to bearing when not accelerating.
const BearingFriction = 0.8

// The following constants are pixels per second.
const Gravity = 1400.0
const Maxvel = 240.0
const Accel = 900.0
const AirBending = 575.0
const BearingAccel = 1200.0
const JumpImpulse = -350.0

const MaxSquishVel = 80.0

// Max vertical velocity while holding down jump.
const JumpTerminalVelocity = 350.0
const TerminalVelocity = 550.0

// Blinking.
const BlinkOdds = 1 / 300.0
const BlinkCycleSeconds = 0.25

// Eye centering speed for tumbling/landing.
const EyeCenteringDurationSeconds = 0.25
