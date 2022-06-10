package game

import (
	"image/color"

	"github.com/jakecoffman/cp"
	"github.com/jdf/todd-again/engine"
)

type Vec2 = engine.Vec2

const TODD_GROUND_COLLISION_TYPE = 837232

type Todd struct {
	body  *cp.Body
	shape *cp.Shape

	OnGround bool
}

func NewBox(space *cp.Space, pos *Vec2, size *Vec2) *Todd {
	body := space.AddBody(cp.NewBody(size.X*size.Y, cp.INFINITY))
	body.SetPosition(cp.Vector{pos.X, pos.Y})
	body.SetMass(25)

	shape := space.AddShape(cp.NewBox(body, size.X, size.Y, 0))
	shape.SetElasticity(0)
	shape.SetFriction(.2)

	box := &Todd{
		body:  body,
		shape: shape,
	}
	shape.UserData = box
	return box
}

func (t *Todd) Impulse(impulse *Vec2) {
	t.body.ApplyImpulseAtLocalPoint(cp.Vector{impulse.X, impulse.Y}, cp.Vector{0, 0})
}

func (box *Todd) Update(frameState *engine.UpdateState, dt float64) {}

func (box *Todd) Draw(g *engine.Graphics, camera *engine.Camera) {
	g.SetColor(color.White)
	g.FillRoundRect(camera, box.Bounds(), 5)
}

func (box *Todd) Bounds() *engine.Rect {
	bb := box.shape.BB()
	return engine.NewRect(bb.L, bb.B, bb.R, bb.T)
}

func (box *Todd) Velocity() *Vec2 {
	return engine.Vec(box.body.Velocity().X, box.body.Velocity().Y)
}

func (box *Todd) RemoveFromSpace(space *cp.Space) {
	space.RemoveShape(box.shape)
	space.RemoveBody(box.body)
}
