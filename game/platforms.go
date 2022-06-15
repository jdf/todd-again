package game

import (
	"image/color"

	"github.com/jdf/todd-again/engine"
)

type Platform struct {
	bounds *engine.Rect
	color  color.Color
}

func (platform *Platform) Draw(g *engine.Graphics) {
	g.SetColor(platform.color)
	g.DrawRoundedRect(platform.bounds, 5)
	g.Fill()
}
