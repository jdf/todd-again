package entity

import (
	"github.com/jakecoffman/cp"
	"github.com/jdf/todd-again/engine/camera"
	"github.com/jdf/todd-again/engine/frame"
	"github.com/jdf/todd-again/engine/geometry"
	"github.com/jdf/todd-again/engine/graphics"
)

// Entity is a thing in the game that can be updated and drawn.
type Entity interface {
	// Absolute time for animations; dt for physics.
	Update(frameState *frame.State, dt float64)
	Draw(*graphics.Context, *camera.Camera)
	Bounds() *geometry.Rect

	RemoveFromSpace(space *cp.Space)
}
