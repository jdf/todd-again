package game

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/jdf/todd-again/engine"
)

type Platform struct {
	bounds *engine.Rect
	color  color.Color
}

func (platform *Platform) Draw(img *ebiten.Image, g *engine.Graphics) {
	g.SetColor(platform.color)
	g.DrawRoundedRect(img, platform.bounds, 5)
}
