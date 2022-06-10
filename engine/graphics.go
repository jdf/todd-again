package engine

import (
	"github.com/fogleman/gg"
)

// Context is a graphics context that knows how to manipulate
// coordinates with a camera.
type Context struct {
	gg.Context
}

// FillRect fills a world-space rectangle with the current color.
func (g *Context) FillRect(camera *Camera, r *Rect) {
	tr := camera.ToScreenRect(r)
	g.Context.DrawRectangle(tr.Left(), tr.Bottom(), tr.Width(), tr.Height())
	g.Fill()
}

// FillRectScreen fills a screen-space rectangle with the current color.
func (g *Context) FillRectScreen(camera *Camera, r *Rect) {
	g.Context.DrawRectangle(r.Left(), r.Bottom(), r.Width(), r.Height())
	g.Fill()
}

// DrawText draws text at the given position in world space.
func (g *Context) DrawText(camera *Camera, s string, x, y float64) {
	xs, ys := camera.ToScreen(x, y)
	g.Context.DrawString(s, xs, ys)
}

// DrawTextScreen draws text at the given position in screen space.
func (g *Context) DrawTextScreen(camera *Camera, s string, x, y float64) {
	g.Context.DrawString(s, x, y)
}

// DrawLine draws a line from p1 to p2 in world space.
func (g *Context) DrawLine(camera *Camera, x1, y1, x2, y2 float64) {
	x1s, y1s := camera.ToScreen(x1, y1)
	x2s, y2s := camera.ToScreen(x2, y2)
	g.Context.DrawLine(x1s, y1s, x2s, y2s)
	g.Stroke()
}
