package testgame

import (
	"image/color"

	"github.com/jakecoffman/cp"
	"github.com/jdf/todd-again/engine"
)

const debugSpace = true

type Level struct {
	space    *cp.Space
	entities []engine.Entity
	camera   *engine.Camera
}

func (level *Level) Resize(w, h int) {
	ar := float64(w) / float64(h)
	level.camera = engine.NewCamera(
		engine.NewRect(-50, 0, 50, 100/ar),
		engine.NewRect(0, 0, w, h),
		engine.FlipYAxis)
}

func (level *Level) Draw(g *engine.Graphics) {
	if debugSpace {
		level.space.StaticBody.EachShape(func(shape *cp.Shape) {
			g.SetColor(color.RGBA{0xFF, 0, 0, 0xFF})
			bb := shape.BB()
			g.FillRect(level.camera, engine.NewRect(bb.L-.01, bb.B-.01, bb.R+.01, bb.T+.01))
		})
	}
	for _, entity := range level.entities {
		entity.Draw(g, level.camera)
	}
}

func (level *Level) Update(s *engine.UpdateState) {
	level.space.Step(s.DeltaSeconds)
	for _, entity := range level.entities {
		entity.Update(s, s.DeltaSeconds)
	}
	box := level.entities[0].(*engine.Box)
	onGround := box.Bounds().Min.Y < .01
	if s.Input.Spacebar {
		jump := 5
		if onGround {
			jump = 1000
		}
		box.Impulse(engine.Vec(0, jump))
	}
	scootch := 8
	if onGround {
		scootch = 20
	}
	if s.Input.Left {
		box.Impulse(engine.Vec(-scootch, 0))
	}
	if s.Input.Right {
		box.Impulse(engine.Vec(scootch, 0))
	}
}

func standardSpace() *cp.Space {
	space := cp.NewSpace()
	space.SetGravity(cp.Vector{0, -100})

	walls := []cp.Vector{
		{-100, 0}, {-100, 100},
		{100, 0}, {100, 100},
		{-100, 0}, {100, 0},
		{-100, 100}, {100, 100},
	}
	for i := 0; i < len(walls)-1; i += 2 {
		shape := space.AddShape(cp.NewSegment(space.StaticBody, walls[i], walls[i+1], 0))
		shape.SetElasticity(1)
		shape.SetFriction(1)
		shape.SetFilter(cp.SHAPE_FILTER_ALL)
	}
	return space
}

func Level1() *Level {
	space := standardSpace()

	level := &Level{
		space:    space,
		entities: []engine.Entity{engine.NewBox(space, engine.Vec(0, 60), engine.Vec(5, 5))},
	}
	return level
}
