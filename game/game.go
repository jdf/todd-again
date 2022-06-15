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
		engine.NewRect(0, 0, float64(w)/10, float64(h)/10),
		engine.NewRect(0, 0, w, h),
		engine.FlipYAxis)
}

func (level *Level) Update(s *engine.UpdateState) {
}

func Level1() *Level {
	level := &Level{
		todd: &Todd{
			sideLength: 10,
			fillColor:  color.RGBA{R: 233, G: 180, B: 30, A: 255},
			pos:        engine.Vec2{X: 5, Y: 0},
		},
	}
	return level
}
