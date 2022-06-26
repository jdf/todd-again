package engine

import (
	"github.com/jakecoffman/cp"
)

// Entity is a thing in the game that can be updated and drawn.
type Entity interface {
	// Absolute time for animations; dt for physics.
	Update(frameState *UpdateState, dt float32)
	Draw(*Graphics, *Camera)

	Impulse(*Vec2)
	Bounds() *Rect
	Velocity() *Vec2

	RemoveFromSpace(space *cp.Space)
}
