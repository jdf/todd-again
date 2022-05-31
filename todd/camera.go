package todd

import "fmt"

// Camera transforms between world and screen coordinates.
type Camera struct {
	worldRect  *Rect
	screenRect *Rect

	// Cached transforms from world to display space and back.
	worldToDisplay *Affine
	displayToWorld *Affine
}

// Copy returns a deep copy of the camera.
func (c *Camera) Copy() *Camera {
	return &Camera{
		worldRect:  c.worldRect.Copy(),
		screenRect: c.screenRect.Copy(),
	}
}

func (c *Camera) String() string {
	return fmt.Sprintf("Camera[%s -> %s]", c.worldRect, c.screenRect)
}

// NewView constructs a View mapped to the complete display area.
func NewView(worldPort *Rect) *Camera {
	return &Camera{
		worldRect:  worldPort.Copy(),
		screenRect: NewRect(0, 0, 1, 1),
	}
}

func (c *Camera) invalidate() {
	c.worldToDisplay = nil
	c.displayToWorld = nil
}

func (c *Camera) SetScreenRect(viewport *Rect) {
	// TODO - check viewport aspect ratio and do something sensible.
	c.screenRect = viewport
	c.invalidate()
}

func (c *Camera) getTransform() *Affine {
	if c.worldToDisplay == nil {
		c.worldToDisplay = Compose(
			Translate(c.worldRect.Center().Negate()),
			Scale(c.screenRect.Size().Div(c.worldRect.Size())),
			Translate(c.screenRect.Center()),
		)
	}
	return c.worldToDisplay
}

// ToScreenVec2 converts a point in world space to a point in display space.
func (c *Camera) ToScreenVec2(worldPos *Vec2) *Vec2 {
	return c.getTransform().TransformVec2(worldPos)
}

// ToScreenRect converts a rectangle in world space to a rectangle in display space.
func (c *Camera) ToScreenRect(rect *Rect) *Rect {
	return c.getTransform().TransformRect(rect)
}
