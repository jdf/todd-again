package engine

import (
	"fmt"
)

type YAxisPolicy int

const (
	FlipYAxis YAxisPolicy = iota
	DoNotFlipYAxis
)

// Camera transforms between world and screen coordinates.
type Camera struct {
	worldRect  *Rect
	screenRect *Rect

	yAxisPolicy YAxisPolicy

	// Cached transforms from world to display space and back.
	worldToScreen *Affine
	screenToWorld *Affine

	// hack to limit zoom
	zoom float32
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

// New constructs a View mapped to the given display area.
func NewCamera(worldRect *Rect, screenRect *Rect, yAxisPolicy YAxisPolicy) *Camera {
	return &Camera{
		worldRect:   worldRect.Copy(),
		screenRect:  screenRect.Copy(),
		yAxisPolicy: yAxisPolicy,
		zoom:        1,
	}
}

func (c *Camera) WorldBounds() *Rect {
	return c.worldRect.Copy()
}

// Left returns the left edge of the world rectangle.
func (c *Camera) Left() float32 {
	return c.worldRect.Left()
}

// Right returns the right edge of the world rectangle.
func (c *Camera) Right() float32 {
	return c.worldRect.Right()
}

// Top returns the top edge of the world rectangle.
func (c *Camera) Top() float32 {
	return c.worldRect.Top()
}

// Bottom returns the bottom edge of the world rectangle.
func (c *Camera) Bottom() float32 {
	return c.worldRect.Bottom()
}

func (c *Camera) SetCenterX(x float32) {
	cur := c.worldRect.Center()
	c.Pan(x-cur.X, 0)
}

func (c *Camera) SetCenterY(y float32) {
	cur := c.worldRect.Center()
	c.Pan(0, y-cur.Y)
}

func (c *Camera) RelativelyPositionY(worldY, cameraY float32) {
	newBottom := worldY - cameraY
	c.worldRect.SetBottomPreservingSize(newBottom)
	c.invalidate()
}

// Pan moves the camera by the given amount.
func (c *Camera) Pan(x, y float32) {
	c.worldRect.AddToSelf(&Vec2{x, y})
	c.invalidate()
}

// Zoom scales the camera by the given factor, keeping the center of the camera fixed.
func (c *Camera) Zoom(factor float32) {
	c.ZoomInto(factor, c.worldRect.Center())
}

// ZoomInto scales the camera by the given factor, keeping the given point fixed.
func (c *Camera) ZoomInto(factor float32, center *Vec2) {
	newZoom := c.zoom * factor
	if newZoom < 0.1 || newZoom > 10 {
		return
	}
	c.zoom = newZoom
	zoomer := Compose(
		TranslationV(center.Negate()),
		UniformScale(factor),
		TranslationV(center),
	)
	c.worldRect = zoomer.TransformRect(c.worldRect)
	c.invalidate()
}

func (c *Camera) invalidate() {
	c.worldToScreen = nil
	c.screenToWorld = nil
}

// SetScreenRect sets the screen rectangle.
func (c *Camera) SetScreenRect(viewport *Rect) {
	c.screenRect = viewport
	c.invalidate()
}

func (c *Camera) GetTransform() *Affine {
	if c.worldToScreen == nil {
		c.worldToScreen = Compose(
			TranslationV(c.worldRect.Center().Negate()),
			Scale(c.screenRect.Size().Div(c.worldRect.Size())),
			TranslationV(c.screenRect.Center()),
		)
		if c.yAxisPolicy == FlipYAxis {
			c.worldToScreen = Compose(
				c.worldToScreen,
				Scale(&Vec2{1, -1}),
				Translation(0, c.screenRect.Height()),
			)
		}
	}
	return c.worldToScreen
}

func (c *Camera) getInverseTransform() *Affine {
	if c.screenToWorld == nil {
		c.screenToWorld = c.GetTransform().Inverse()
	}
	return c.screenToWorld
}

// CanSee returns true if the given rectangle is visible in the camera's world window.
func (c *Camera) CanSee(rect *Rect) bool {
	return c.worldRect.Intersects(rect)
}

// ToScreenVec2 converts a point in world space to a point in display space.
func (c *Camera) ToScreenVec2(worldPos *Vec2) *Vec2 {
	return c.GetTransform().TransformVec2(worldPos)
}

// ToScreen converts a point in world space to a point in display space.
func (c *Camera) ToScreen(x, y float32) (float32, float32) {
	return c.GetTransform().Transform(x, y)
}

// ToScreenRect converts a rectangle in world space to a rectangle in display space.
func (c *Camera) ToScreenRect(rect *Rect) *Rect {
	return c.GetTransform().TransformRect(rect)
}

// ToWorldVec2 converts a point in display space to a point in world space.
func (c *Camera) ToWorldVec2(displayPos *Vec2) *Vec2 {
	return c.getInverseTransform().TransformVec2(displayPos)
}

// ToWorld converts a point in display space to a point in world space.
func (c *Camera) ToWorld(x, y float32) (float32, float32) {
	return c.getInverseTransform().Transform(x, y)
}

// ToWorldRect converts a rectangle in display space to a rectangle in world space.
func (c *Camera) ToWorldRect(rect *Rect) *Rect {
	return c.getInverseTransform().TransformRect(rect)
}

// Generic funcs to transform arbitrary numeric types.

// ScreenToWorld converts a point in display space to a point in world space.
func ScreenToWorld[T Numeric](c *Camera, x, y T) (float32, float32) {
	return c.ToWorld(float32(x), float32(y))
}

// WorldToScreen converts a point in world space to a point in display space.
func WorldToScreen[T Numeric](c *Camera, x, y T) (float32, float32) {
	return c.ToScreen(float32(x), float32(y))
}
