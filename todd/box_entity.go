package todd

import (
	"image/color"

	"github.com/jakecoffman/cp"
)

type BoxEntity struct {
	body  *cp.Body
	shape *cp.Shape
}

func NewBoxEntity(space *cp.Space, pos *Vec2, size *Vec2) *BoxEntity {
	body := space.AddBody(cp.NewBody(size.X*size.Y, cp.INFINITY))
	body.SetPosition(cp.Vector{pos.X, pos.Y})

	shape := space.AddShape(cp.NewBox(body, size.X, size.Y, 0))
	shape.SetElasticity(.35)
	shape.SetFriction(0.7)

	return &BoxEntity{
		body:  body,
		shape: shape,
	}
}

func (box *BoxEntity) Update(frameState *FrameState, dt float64) {}

func (box *BoxEntity) Draw(g *Graphics, camera *Camera) {
	g.SetColor(color.White)
	g.FillRect(camera, box.Bounds())
}

func (box *BoxEntity) Bounds() *Rect {
	bb := box.shape.BB()
	return NewRect(bb.L, bb.B, bb.R, bb.T)
}

func (box *BoxEntity) RemoveFromSpace(space *cp.Space) {
	space.RemoveShape(box.shape)
	space.RemoveBody(box.body)
}
