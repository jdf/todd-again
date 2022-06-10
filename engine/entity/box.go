package entity

import (
	"image/color"

	"github.com/jakecoffman/cp"
	"github.com/jdf/todd-again/engine/camera"
	"github.com/jdf/todd-again/engine/frame"
	"github.com/jdf/todd-again/engine/geometry"
	"github.com/jdf/todd-again/engine/graphics"
)

type Box struct {
	body  *cp.Body
	shape *cp.Shape
}

func NewBox(space *cp.Space, pos *geometry.Vec2, size *geometry.Vec2) *Box {
	body := space.AddBody(cp.NewBody(size.X*size.Y, cp.INFINITY))
	body.SetPosition(cp.Vector{pos.X, pos.Y})

	shape := space.AddShape(cp.NewBox(body, size.X, size.Y, 0))
	shape.SetElasticity(.35)
	shape.SetFriction(0.7)

	return &Box{
		body:  body,
		shape: shape,
	}
}

func (box *Box) Update(frameState *frame.State, dt float64) {}

func (box *Box) Draw(g *graphics.Context, camera *camera.Camera) {
	g.SetColor(color.White)
	g.FillRect(camera, box.Bounds())
}

func (box *Box) Bounds() *geometry.Rect {
	bb := box.shape.BB()
	return geometry.NewRect(bb.L, bb.B, bb.R, bb.T)
}

func (box *Box) RemoveFromSpace(space *cp.Space) {
	space.RemoveShape(box.shape)
	space.RemoveBody(box.body)
}
