package todd

import "math"

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
	center   *Point
	size     *Dimension
	rotation float64

	target *Rect

	// Cached transforms from world to display space and back.
	affine  *Affine
	inverse *Affine
}

// NewView constructs a View mapped to the complete display area.
func NewView(rect *Rect) *View {
	return &View{
		center: rect.Center(),
		size:   rect.Size(),
	}
}

func (v *View) getTransform() *Affine {
	if v.affine == nil {
		cos := math.Cos(v.rotation)
		sin := math.Sin(v.rotation)
		tx := v.center.X*cos - v.center.Y*sin + v.center.X
		ty := v.center.X*sin - v.center.Y*cos + v.center.Y

		a := 2 / v.size.Width
		b := -2 / v.size.Height
		c := -a * v.center.X
		d := -b * v.center.Y

		v.affine = &Affine{
			a * cos, a * sin, a*tx + c,
			-b * sin, b * cos, b*ty + d,
		}
	}
	return v.affine
}

// PointToDisplay converts a point in world space to a point in display space.
func (v *View) PointToDisplay(p *Point) *Point {
	return v.getTransform().TransformPoint(p)
}

// PointToWorld converts a point in display space to a point in world space.
func (v *View) PointToWorld(worldPos *Point) *Point {
	return nil
}

// RectToDisplay converts a rectangle in world space to a rectangle in display space.
func (v *View) RectToDisplay(rect *Rect) *Rect {
	return v.getTransform().TransformRect(rect)
}

// RectToWorld converts a rectangle in display space to a rectangle in world space.
func (v *View) RectToWorld(Rect *Rect) *Rect {
	return nil
}
