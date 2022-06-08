package todd

import (
	"image/color"

	"github.com/jakecoffman/cp"
)

const tick = 1 / 180.0

var accumulator float64

type Level struct {
	space    *cp.Space
	entities []Entity
}

func (level *Level) Draw(g *Graphics, camera *Camera) {
	level.space.StaticBody.EachShape(func(shape *cp.Shape) {
		g.SetColor(color.RGBA{0xFF, 0, 0, 0xFF})
		bb := shape.BB()
		g.FillRect(camera, NewRect(bb.L-.01, bb.B-.01, bb.R+.01, bb.T+.01))
	})
	for _, entity := range level.entities {
		entity.Draw(g, camera)
	}
}

func (level *Level) Update(frameState *FrameState) {
	for accumulator += frameState.DeltaT; accumulator >= tick; accumulator -= tick {
		level.space.Step(tick)
		for _, entity := range level.entities {
			entity.Update(frameState, tick)
		}
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
		entities: []Entity{NewBoxEntity(space, Vec(0, 20), Vec(10, 10))},
	}
	return level
}
