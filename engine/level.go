package engine

import (
	"image/color"

	"github.com/jakecoffman/cp"
	"github.com/jdf/todd-again/engine/camera"
	"github.com/jdf/todd-again/engine/entity"
	"github.com/jdf/todd-again/engine/frame"
	"github.com/jdf/todd-again/engine/geometry"
	"github.com/jdf/todd-again/engine/graphics"
)

const debugSpace = true

const tick = 1 / 180.0

var accumulator float64

type Level struct {
	space    *cp.Space
	entities []entity.Entity
}

func (level *Level) Draw(g *graphics.Context, camera *camera.Camera) {
	if debugSpace {
		level.space.StaticBody.EachShape(func(shape *cp.Shape) {
			g.SetColor(color.RGBA{0xFF, 0, 0, 0xFF})
			bb := shape.BB()
			g.FillRect(camera, geometry.NewRect(bb.L-.01, bb.B-.01, bb.R+.01, bb.T+.01))
		})
	}
	for _, entity := range level.entities {
		entity.Draw(g, camera)
	}
}

func (level *Level) Update(frameState *frame.State) {
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
		entities: []entity.Entity{entity.NewBox(space, geometry.Vec(0, 20), geometry.Vec(10, 10))},
	}
	return level
}
