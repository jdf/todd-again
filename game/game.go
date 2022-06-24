package game

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/jdf/todd-again/engine"
	"github.com/jdf/todd-again/game/tuning"
	"github.com/tanema/gween"
)

var (
	CameraVerticalAnimation *gween.Tween
)

const DebugWorld = true

type ToddGame struct{}

func (ToddGame) Draw(img *ebiten.Image, ctx *engine.Graphics) {
	ctx.SetWorldToScreen(Camera.GetTransform())
	for _, plat := range Platforms {
		plat.Draw(img, ctx)
	}
	Todd.Draw(img, ctx)
}

func (ToddGame) Resize(w, h int) {
	ar := float64(w) / float64(h)
	Camera = engine.NewCamera(
		engine.NewRect(-200, -10, 200, 400.0/ar-10),
		engine.NewRect(0, 0, w, h),
		engine.FlipYAxis)
	Camera.SetCenterX(Todd.pos.X)
}

func AnimateCameraVertical() {
	b := Camera.WorldBounds()
	var target float64
	if Todd.pos.Y < 200 {
		target = -10
	} else {
		target = Todd.pos.Y - 20 - b.Height()/2
	}
	CameraVerticalAnimation = gween.New(
		float32(b.Bottom()),
		float32(target),
		float32(tuning.Instance.GetCameraTiltSeconds()),
		tuning.CameraTiltEasing)
}

func ControlCamera(s *engine.UpdateState) {
	Camera.SetCenterX(Todd.pos.X)
	if CameraVerticalAnimation != nil {
		y, done := CameraVerticalAnimation.Update(float32(s.DeltaSeconds))
		Camera.RelativelyPositionY(float64(y), 0)
		if done {
			CameraVerticalAnimation = nil
		}
	}
	if Camera.Bottom() > Todd.pos.Y-10 {
		Camera.RelativelyPositionY(Todd.pos.Y-10, 0)
	}
}

func (ToddGame) Update(s *engine.UpdateState) {
	Todd.Update(s)
	ControlCamera(s)
}
