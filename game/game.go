package game

import (
	"github.com/jdf/todd-again/engine"
)

type Level struct {
	camera *engine.Camera
	todd   *Todd
}

func (level *Level) Draw(ctx *engine.Graphics) {
	level.todd.Draw(ctx, cam)
}

func (level *Level) Resize(w, h int) {
	ar := float64(w) / float64(h)
	camera = engine.NewCamera(
		engine.NewRect(-25, 0, 25, 50/ar),
		engine.NewRect(0, 0, w, h),
		engine.FlipYAxis)
}

func (level *Level) Update(s *engine.UpdateState) {
}

func Level1() *Level {
	space := standardSpace()

	level := &Level{
		space:    space,
		entities: []engine.Entity{NewBox(space, engine.Vec(0, 60), engine.Vec(3, 3))},
	}
	return level
}
