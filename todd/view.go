package todd

import "fmt"

// Dimension is a unitless dimension.
type Dimension struct {
	Width, Height float64
}

func (d *Dimension) String() string {
	return fmt.Sprintf("%f x %f", d.Width, d.Height)
}

// View encapsulate both a camera and a target destination on a display.
// The camera is specified by a rectangular region of the world and a
// (usually 0) rotation.
// The display is specified by a rectangle in an abstract space that
// goes from (0, 0) in the lower left to (1, 1) in the upper right. That
// rectangle is mapped to the physical display elsewhere.
type View struct {
	center   *Vec2
	size     *Dimension
	rotation float64

	target *Rect

	// Cached transforms from world to display space and back.
	worldToDisplay *Affine
	displayToWorld *Affine
}

func (v *View) String() string {
	return "View[center=" + v.center.String() +
		", size=" + v.size.String() +
		", rotation=" + fmt.Sprintf("%f", v.rotation) + "]"
}

// NewView constructs a View mapped to the complete display area.
func NewView(rect *Rect) *View {
	return &View{
		center: rect.Center(),
		size:   rect.Size(),
		target: NewRect(0, 0, 1, 1),
	}
}

func (v *View) getTransform() *Affine {
	if v.worldToDisplay == nil {
		v.worldToDisplay = Compose(
			Translation(v.center.Negate()),
			Scale(&Vec2{
				v.target.Size().Width / v.size.Width,
				v.target.Size().Height / v.size.Height,
			}),
			Translation(v.target.Center()),
		)
	}
	return v.worldToDisplay
}

// WorldToDisplayP converts a point in world space to a point in display space.
func (v *View) WorldToDisplayP(worldPos *Vec2) *Vec2 {
	return v.getTransform().TransformVec2(worldPos)
}

// DisplayToWorldP converts a point in display space to a point in world space.
func (v *View) DisplayToWorldP(displayPos *Vec2) *Vec2 {
	return nil
}

// WorldToDisplayR converts a rectangle in world space to a rectangle in display space.
func (v *View) WorldToDisplayR(rect *Rect) *Rect {
	return v.getTransform().TransformRect(rect)
}

// DisplayToWorldR converts a rectangle in display space to a rectangle in world space.
func (v *View) DisplayToWorldR(Rect *Rect) *Rect {
	return nil
}
