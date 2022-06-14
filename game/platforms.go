package game

import (
	"image/color"

	"github.com/jdf/todd-again/engine"
)

type Platform struct {
	bounds *engine.Rect
	color  color.Color
}

func (platform *Platform) Draw(ctx *engine.Graphics, cam *engine.Camera) {
	ctx.SetColor(platform.color)
	ctx.FillRect(cam, platform.bounds)
}
