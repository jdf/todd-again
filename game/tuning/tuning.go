package tuning

import (
	"embed"
	"image/color"

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
	if Instance.Todd == nil {
		Instance.Todd = &proto.Todd{
			Blink: &proto.Blink{},
			Color: &proto.Color{
				C: []float32{233 / 255.0, 180 / 255.0, 30 / 255.0},
			},
		}
	}
	if Instance.World == nil {
		Instance.World = &proto.World{}
		Instance.World.Bg = &proto.Color{
			C: []float32{0, 0, 0},
		}
	}
	if Instance.Camera == nil {
		Instance.Camera = &proto.Camera{}
	}
	Step1 = Instance.Todd.GetMaxVelocity() * .333
	Step2 = Instance.Todd.GetMaxVelocity() * .666
}

var (
	Instance *proto.Tuning

	//go:embed *.textproto
	assets embed.FS

	CameraTiltEasing = ease.Linear
	Step1            float32
	Step2            float32
)

func RGBA(c *proto.Color) color.RGBA {
	return color.RGBA{
		uint8(c.C[0] * 255),
		uint8(c.C[1] * 255),
		uint8(c.C[2] * 255),
		255,
	}
}

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
	return Instance.Todd.GetJumpImpulse() * JumpImpulseFactors[SpeedStepFunction(speed)]
}

// The platform has a width differing from its apparent width, depending on
// your speed. The slower you are, the narrower the platform is.
var PlatformMargins = []float32{-8, 0, 5}

func PlatformMargin(xVel float32) float32 {
	return PlatformMargins[SpeedStepFunction(xVel)]
}
