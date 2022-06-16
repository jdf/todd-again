package game

import (
	"image/color"

	"github.com/jdf/todd-again/engine"
)

type Level struct {
	camera *engine.Camera
	todd   *Todd
}

func (level *Level) Draw(ctx *engine.Graphics) {
	ctx.SetWorldToScreen(level.camera.GetTransform())
	level.todd.Draw(ctx)
}

func (level *Level) Resize(w, h int) {
	//ar := float64(w) / float64(h)
	level.camera = engine.NewCamera(
		engine.NewRect(0, 0, float64(w)/2, float64(h)/2),
		engine.NewRect(0, 0, w, h),
		engine.FlipYAxis)
}

func (level *Level) Update(s *engine.UpdateState) {
	SetControllerState(s.Input)
	level.todd.Update(s)
}

func Level1() *Level {
	InitPlatforms()
	level := &Level{
		todd: &Todd{
			sideLength: 30,
			fillColor:  color.RGBA{R: 233, G: 180, B: 30, A: 255},
			pos:        engine.Vec2{X: 0, Y: 1},
		},
	}
	return level
}
