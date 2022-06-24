package tuning

import (
	"embed"
	"math"

	"github.com/jdf/todd-again/game/proto"
	"github.com/tanema/gween/ease"
	"google.golang.org/protobuf/encoding/prototext"
)

func init() {
	buf, err := assets.ReadFile("tuning.textproto")
	if err != nil {
		panic(err)
	}
	Instance = &proto.Tuning{}
	if err = prototext.Unmarshal(buf, Instance); err != nil {
		panic(err)
	}
}

var (
	Instance *proto.Tuning

	//go:embed *.textproto
	assets embed.FS

	CameraTiltEasing = ease.Linear
	Step1            = Instance.GetMaxVelocity() * .333
	Step2            = Instance.GetMaxVelocity() * .666
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
	return Instance.GetJumpImpulse() * JumpImpulseFactors[SpeedStepFunction(speed)]
}

// The platform has a width differing from its apparent width, depending on
// your speed. The slower you are, the narrower the platform is.
var PlatformMargins = []float64{-8, 0, 5}

func PlatformMargin(xVel float64) float64 {
	return PlatformMargins[SpeedStepFunction(xVel)]
}
