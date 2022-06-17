package game

import (
	"image/color"
	"sort"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/jdf/todd-again/engine"
)

func init() {
	sort.Slice(Platforms, func(i, j int) bool {
		return Platforms[i].bounds.Top() < Platforms[j].bounds.Top()
	})
	for _, plat := range Platforms {
		n := len(TumbleLevels)
		if n > 0 && TumbleLevels[n-1] == plat.bounds.Top() {
			continue
		}
		TumbleLevels = append(TumbleLevels, plat.bounds.Top())
	}
	TumbleLevels = append(TumbleLevels, 0)
}

var TumbleLevels []float64

var Platforms = []Platform{
	{engine.NewRect(100, 110, 250, 130), color.RGBA{190, 190, 255, 255}},
	{engine.NewRect(300, 210, 500, 230), color.RGBA{190, 255, 190, 255}},
}

type Platform struct {
	bounds *engine.Rect
	color  color.Color
}

func (platform *Platform) Draw(img *ebiten.Image, g *engine.Graphics) {
	g.SetColor(platform.color)
	g.DrawRoundedRect(img, platform.bounds, 5)
}
