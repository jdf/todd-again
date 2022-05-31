package todd

import (
	"github.com/fogleman/gg"
)

type Graphics struct {
	gg.Context
}

func (g *Graphics) FillRect(camera *Camera, r *Rect) {
	tr := camera.ToScreenRect(r)
	g.Context.DrawRectangle(tr.Left(), tr.Bottom(), tr.Width(), tr.Height())
	g.Fill()
}

func (g *Graphics) DrawText(camera *Camera, s string, p *Vec2) {
	pr := camera.ToScreenVec2(p)
	g.Context.DrawString(s, pr.X, pr.Y)
}
