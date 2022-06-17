package game

import (
	"image/color"
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/jdf/todd-again/engine"
)

type Level struct {
	camera *engine.Camera
	todd   *Todd
}

func (level *Level) Draw(img *ebiten.Image, ctx *engine.Graphics) {
	ctx.SetWorldToScreen(level.camera.GetTransform())
	for _, plat := range Platforms {
		plat.Draw(img, ctx)
	}
	level.todd.Draw(img, ctx)
}

func (level *Level) Resize(w, h int) {
	ar := float64(w) / float64(h)
	level.camera = engine.NewCamera(
		engine.NewRect(-200, 0, 200, 400.0/ar),
		engine.NewRect(0, 0, w, h),
		engine.FlipYAxis)
}

func (level *Level) Update(s *engine.UpdateState) {
	SetControllerState(s.Input)
	level.todd.Update(s)
	level.camera.CenterHorizontalOn(level.todd.pos.X)
}

func Level1() *Level {
	level := &Level{
		todd: &Todd{
			sideLength: ToddSideLength,
			fillColor:  color.RGBA{R: 233, G: 180, B: 30, A: 255},
			pos:        engine.Vec2{X: 0, Y: 0},
			rnd:        rand.New(rand.NewSource(time.Now().UnixNano())),
		},
	}
	return level
}
