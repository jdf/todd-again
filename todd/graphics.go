package todd

import (
	"github.com/fogleman/gg"
)

// Graphics is a graphics context that knows how to manipulate
// coordinates with a camera.
type Graphics struct {
	gg.Context
}

// FillRect fills a rectangle with the current color.
func (g *Graphics) FillRect(camera *Camera, r *Rect) {
	tr := camera.ToScreenRect(r)
	g.Context.DrawRectangle(tr.Left(), tr.Bottom(), tr.Width(), tr.Height())
	g.Fill()
}

// DrawText draws text at the given position in world space.
func (g *Graphics) DrawText(camera *Camera, s string, x, y float64) {
	xs, ys := camera.ToScreen(x, y)
	g.Context.DrawString(s, xs, ys)
}

// DrawLine draws a line from p1 to p2 in world space.
func (g *Graphics) DrawLine(camera *Camera, x1, y1, x2, y2 float64) {
	x1s, y1s := camera.ToScreen(x1, y1)
	x2s, y2s := camera.ToScreen(x2, y2)
	g.Context.DrawLine(x1s, y1s, x2s, y2s)
	g.Stroke()
}
