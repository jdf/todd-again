package engine

import (
	"github.com/jakecoffman/cp"
)

// Entity is a thing in the game that can be updated and drawn.
type Entity interface {
	// Absolute time for animations; dt for physics.
	Update(frameState *FrameState, dt float64)
	Draw(*Context, *Camera)
	Bounds() *Rect

	RemoveFromSpace(space *cp.Space)
}
