package todd

import (
	"fmt"
)

// Camera transforms between world and screen coordinates.
type Camera struct {
	worldRect  *Rect
	screenRect *Rect

	// Cached transforms from world to display space and back.
	worldToDisplay *Affine
	displayToWorld *Affine

	invertY bool

	// hack to limit zoom
	zoom float64
}

// Copy returns a deep copy of the camera.
func (c *Camera) Copy() *Camera {
	return &Camera{
		worldRect:  c.worldRect.Copy(),
		screenRect: c.screenRect.Copy(),
	}
}

func (c *Camera) String() string {
	return fmt.Sprintf("Camera[%s -> %s @ %0.2f]", c.worldRect, c.screenRect, c.zoom)
}

// NewCamera constructs a View mapped to the given display area.
func NewCamera(worldRect *Rect, screenRect *Rect) *Camera {
	return &Camera{
		worldRect:  worldRect.Copy(),
		screenRect: screenRect.Copy(),
		zoom:       1,
	}
}

// SetInvertY sets whether the Y axis should be inverted in display space.
func (c *Camera) SetInvertY(invertY bool) {
	c.invertY = invertY
	c.invalidate()
}

// Left returns the left edge of the world rectangle.
func (c *Camera) Left() float64 {
	return c.worldRect.Left()
}

// Right returns the right edge of the world rectangle.
func (c *Camera) Right() float64 {
	return c.worldRect.Right()
}

// Top returns the top edge of the world rectangle.
func (c *Camera) Top() float64 {
	return c.worldRect.Top()
}

// Bottom returns the bottom edge of the world rectangle.
func (c *Camera) Bottom() float64 {
	return c.worldRect.Bottom()
}

// Pan moves the camera by the given amount.
func (c *Camera) Pan(v *Vec2) {
	c.worldRect.AddToSelf(v)
	c.invalidate()
}

// Zoom scales the camera by the given factor, keeping the center of the camera fixed.
func (c *Camera) Zoom(factor float64) {
	c.ZoomInto(factor, c.worldRect.Center())
}

// ZoomInto scales the camera by the given factor, keeping the given point fixed.
func (c *Camera) ZoomInto(factor float64, center *Vec2) {
	newZoom := c.zoom * factor
	if newZoom < 0.1 || newZoom > 10 {
		return
	}
	c.zoom = newZoom
	zoomer := Compose(
		Translation(center.Negate()),
		UniformScale(factor),
		Translation(center),
	)
	c.worldRect = zoomer.TransformRect(c.worldRect)
	c.invalidate()
}

func (c *Camera) invalidate() {
	c.worldToDisplay = nil
	c.displayToWorld = nil
}

// SetScreenRect sets the screen rectangle.
func (c *Camera) SetScreenRect(viewport *Rect) {
	c.screenRect = viewport
	c.invalidate()
}

func (c *Camera) getTransform() *Affine {
	if c.worldToDisplay == nil {
		c.worldToDisplay = Compose(
			Translation(c.worldRect.Center().Negate()),
			Scale(c.screenRect.Size().Div(c.worldRect.Size())),
			Translation(c.screenRect.Center()),
		)
		if c.invertY {
			c.worldToDisplay = Compose(
				c.worldToDisplay,
				Scale(&Vec2{1, -1}),
				Translation(&Vec2{0, c.screenRect.Height()}),
			)
		}
	}
	return c.worldToDisplay
}

func (c *Camera) getInverseTransform() *Affine {
	if c.displayToWorld == nil {
		c.displayToWorld = c.getTransform().Inverse()
	}
	return c.displayToWorld
}

// CanSee returns true if the given rectangle is visible in the camera's world window.
func (c *Camera) CanSee(rect *Rect) bool {
	return c.worldRect.Intersects(rect)
}

// ToScreenVec2 converts a point in world space to a point in display space.
func (c *Camera) ToScreenVec2(worldPos *Vec2) *Vec2 {
	return c.getTransform().TransformVec2(worldPos)
}

// ToScreen converts a point in world space to a point in display space.
func (c *Camera) ToScreen(x, y float64) (float64, float64) {
	return c.getTransform().Transform(x, y)
}

// ToScreenRect converts a rectangle in world space to a rectangle in display space.
func (c *Camera) ToScreenRect(rect *Rect) *Rect {
	return c.getTransform().TransformRect(rect)
}

// ToWorldVec2 converts a point in display space to a point in world space.
func (c *Camera) ToWorldVec2(displayPos *Vec2) *Vec2 {
	return c.getInverseTransform().TransformVec2(displayPos)
}

// ToWorld converts a point in display space to a point in world space.
func (c *Camera) ToWorld(x, y float64) (float64, float64) {
	return c.getInverseTransform().Transform(x, y)
}

// ToWorldRect converts a rectangle in display space to a rectangle in world space.
func (c *Camera) ToWorldRect(rect *Rect) *Rect {
	return c.getInverseTransform().TransformRect(rect)
}

// Generic functions to transform arbitrary numeric types.

// ScreenToWorld converts a point in display space to a point in world space.
func ScreenToWorld[T Numeric](c *Camera, x, y T) (float64, float64) {
	return c.ToWorld(float64(x), float64(y))
}

// WorldToScreen converts a point in world space to a point in display space.
func WorldToScreen[T Numeric](c *Camera, x, y T) (float64, float64) {
	return c.ToScreen(float64(x), float64(y))
}
