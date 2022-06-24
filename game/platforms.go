package game

import (
	"image/color"
	"sort"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/jdf/todd-again/engine"
	"github.com/jdf/todd-again/game/tuning"
)

var (
	TumbleLevels []float64

	Platforms = []Platform{
		{engine.NewRect(100, 110, 250, 130), color.RGBA{190, 190, 255, 255}},
		{engine.NewRect(300, 210, 500, 230), color.RGBA{190, 255, 190, 255}},
	}
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

type Platform struct {
	bounds *engine.Rect
	color  color.Color
}

const DebugPlatforms = true

func (platform *Platform) Draw(img *ebiten.Image, g *engine.Graphics) {
	b := platform.bounds
	g.SetColor(platform.color)
	g.DrawRoundedRect(img, b, 5)
	if DebugPlatforms {
		margin := tuning.PlatformMargin(Todd.vel.X)
		g.SetColor(color.RGBA{0, 255, 0, 255})
		g.DrawLine(img, b.Left()-margin, b.Top(), b.Right()+margin, b.Top())
	}
}
