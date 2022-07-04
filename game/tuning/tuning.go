package tuning

import (
	"embed"

	"git.maze.io/go/math32"
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

func SpeedStepFunction(v float32) int {
	v = math32.Abs(v)
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
var JumpImpulseFactors = []float32{1, 1, 1.2}

func GetJumpImpulse(speed float32) float32 {
	return Instance.GetJumpImpulse() * JumpImpulseFactors[SpeedStepFunction(speed)]
}

// The platform has a width differing from its apparent width, depending on
// your speed. The slower you are, the narrower the platform is.
var PlatformMargins = []float32{-8, 0, 5}

func PlatformMargin(xVel float32) float32 {
	return PlatformMargins[SpeedStepFunction(xVel)]
}
