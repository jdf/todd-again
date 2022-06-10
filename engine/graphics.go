package engine

import (
	"github.com/fogleman/gg"
)

// Graphics is a graphics context that knows how to manipulate
// coordinates with a camera.
type Graphics struct {
	gg.Context
}

// FillRoundRect fills a world-space round rectangle with the current color.
func (g *Graphics) FillRoundRect(camera *Camera, r *Rect, screenRadius float64) {
	sr := camera.ToScreenRect(r)

	g.DrawRoundedRectangle(sr.Left(), sr.Bottom(), sr.Width(), sr.Height(), screenRadius)
	g.Fill()
}

// FillRect fills a world-space rectangle with the current color.
func (g *Graphics) FillRect(camera *Camera, r *Rect) {
	g.FillRectScreen(camera.ToScreenRect(r))
}

// FillRectScreen fills a screen-space rectangle with the current color.
func (g *Graphics) FillRectScreen(r *Rect) {
	g.Context.DrawRectangle(r.Left(), r.Bottom(), r.Width(), r.Height())
	g.Fill()
}

// DrawText draws text at the given position in world space.
func (g *Graphics) DrawText(camera *Camera, s string, x, y float64) {
	xs, ys := camera.ToScreen(x, y)
	g.DrawTextScreen(s, xs, ys)
}

// DrawTextScreen draws text at the given position in screen space.
func (g *Graphics) DrawTextScreen(s string, x, y float64) {
	g.Context.DrawString(s, x, y)
}

// DrawLine draws a line from p1 to p2 in world space.
func (g *Graphics) DrawLine(camera *Camera, x1, y1, x2, y2 float64) {
	x1s, y1s := camera.ToScreen(x1, y1)
	x2s, y2s := camera.ToScreen(x2, y2)
	g.Context.DrawLine(x1s, y1s, x2s, y2s)
	g.Stroke()
}
