package todd

// WorldPoint is a point in world space.
type WorldPoint Point

// DisplayPoint is a point in display space.
type DisplayPoint Point

// WorldRect is a rectangle in world space.
type WorldRect Rect

// DisplayRect is a rectangle in display space.
type DisplayRect Rect

// Dimension is a unitless dimension.
type Dimension struct {
	Width, Height float64
}

// View encapsulate both a camera and a target destination on a display.
// The camera is specified by a rectangular region of the world and a
// (usually 0) rotation.
// The display is specified by a rectangle in an abstract space that
// goes from (0, 0) in the lower left to (1, 1) in the upper right. That
// rectangle is mapped to the physical display elsewhere.
type View struct {
	camera   *WorldRect
	target   *DisplayRect
	rotation float64

	// A cached transform from world to display space.
	affine *Affine
}

// WorldPointToDisplay converts a point in world space to a point in display space.
func (*View) WorldPointToDisplay(p *WorldPoint) *DisplayPoint {
	return nil
}

// DisplayPointToWorld converts a point in display space to a point in world space.
func (*View) DisplayPointToWorld(worldPos *DisplayPoint) *WorldPoint {
	return nil
}

// WorldRectToDisplay converts a rectangle in world space to a rectangle in display space.
func WorldRectToDisplay(worldRect *WorldRect) *DisplayRect {
	return nil
}

// DisplayRectToWorld converts a rectangle in display space to a rectangle in world space.
func DisplayRectToWorld(displayRect *DisplayRect) *WorldRect {
	return nil
}
