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

// NewCamera constructs a View mapped to the complete display area.
func NewCamera(worldRect *Rect, screenRect *Rect) *Camera {
	return &Camera{
		worldRect:  worldRect.Copy(),
		screenRect: screenRect.Copy(),
	}
}

// Left returns the left edge of the world rectangle.
func (c *Camera) Left() float64 {
	return c.worldRect.Left()
}

// Right returns the right edge of the world rectangle.
func (c *Camera) Right() float64 {
	return c.worldRect.Right()
}

// Pan moves the camera by the given amount.
func (c *Camera) Pan(v *Vec2) {
	c.worldRect.AddToSelf(v)
	c.invalidate()
}

func (c *Camera) invalidate() {
	c.worldToDisplay = nil
	c.displayToWorld = nil
}

// SetScreenRect sets the screen rectangle.
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
			Scale(&Vec2{1, -1}),
			Translate(&Vec2{0, c.screenRect.Height()}),
		)
	}
	return c.worldToDisplay
}

// CanSee returns true if the given rectangle is visible in the camera's world window.
func (c *Camera) CanSee(rect *Rect) bool {
	return c.worldRect.Intersects(rect)
}

// ToScreenVec2 converts a point in world space to a point in display space.
func (c *Camera) ToScreenVec2(worldPos *Vec2) *Vec2 {
	return c.getTransform().TransformVec2(worldPos)
}

// ToScreenXY converts a point in world space to a point in display space.
func (c *Camera) ToScreenXY(x, y float64) (float64, float64) {
	return c.getTransform().TransformXY(x, y)
}

// ToScreenRect converts a rectangle in world space to a rectangle in display space.
func (c *Camera) ToScreenRect(rect *Rect) *Rect {
	return c.getTransform().TransformRect(rect)
}
