package engine

import (
	"fmt"
	"image/color"

	"github.com/jakecoffman/cp"
)

type Box struct {
	body  *cp.Body
	shape *cp.Shape

	onGround bool
}

func NewBox(space *cp.Space, pos *Vec2, size *Vec2) *Box {
	body := space.AddBody(cp.NewBody(size.X*size.Y, cp.INFINITY))
	body.SetPosition(cp.Vector{pos.X, pos.Y})

	shape := space.AddShape(cp.NewBox(body, size.X, size.Y, 0))
	shape.SetElasticity(0)
	shape.SetFriction(0.7)

	handler := space.NewWildcardCollisionHandler(0)
	handler.BeginFunc = func(arb *cp.Arbiter, space *cp.Space, data interface{}) bool {
		a, b := arb.Shapes()
		fmt.Println("collision %v %v", a, b)
		return true
	}
	handler.SeparateFunc = func(arb *cp.Arbiter, space *cp.Space, data interface{}) {
		a, b := arb.Shapes()
		fmt.Println("separation %v %v", a, b)
	}

	return &Box{
		body:  body,
		shape: shape,
	}
}

func (box *Box) Impulse(impulse *Vec2) {
	box.body.ApplyImpulseAtLocalPoint(cp.Vector{impulse.X, impulse.Y}, cp.Vector{0, 0})
}

func (box *Box) Update(frameState *UpdateState, dt float64) {}

func (box *Box) Draw(g *Graphics, camera *Camera) {
	g.SetColor(color.White)
	g.FillRect(camera, box.Bounds())
}

func (box *Box) Bounds() *Rect {
	bb := box.shape.BB()
	return NewRect(bb.L, bb.B, bb.R, bb.T)
}

func (box *Box) RemoveFromSpace(space *cp.Space) {
	space.RemoveShape(box.shape)
	space.RemoveBody(box.body)
}
