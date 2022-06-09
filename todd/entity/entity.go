package entity

import (
	"github.com/jakecoffman/cp"
	"github.com/jdf/todd-again/todd/camera"
	"github.com/jdf/todd-again/todd/frame"
	"github.com/jdf/todd-again/todd/geometry"
	"github.com/jdf/todd-again/todd/graphics"
)

// Entity is a thing in the game that can be updated and drawn.
type Entity interface {
	// Absolute time for animations; dt for physics.
	Update(frameState *frame.State, dt float64)
	Draw(*graphics.Context, *camera.Camera)
	Bounds() *geometry.Rect

	RemoveFromSpace(space *cp.Space)
}
