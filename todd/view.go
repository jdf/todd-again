package todd

import "github.com/ungerik/go3d/vec2"

// View encapsulate both a camera and a target destination on a display.
// The camera is specified by a rectangular region of the world and a
// (usually 0) rotation.
// The target region is specified by a rectangle in an abstract space that
// goes from (0, 0) in the lower left to (1, 1) in the upper right.
type View interface {
	WorldPointToDisplay(worldPos *vec2.T) vec2.T
	WorldRectToDisplay(worldRect *vec2.Rect) vec2.Rect
}
