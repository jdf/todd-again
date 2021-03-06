package game

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/jdf/todd-again/engine"
	"github.com/jdf/todd-again/game/level"
	"github.com/tanema/gween"
)

var (
	CameraVerticalAnimation *gween.Tween
)

const DebugWorld = true

type toddGame struct {
	ui *level.UI
}

func NewToddGame() engine.GameModule {
	return &toddGame{
		ui: level.NewUI(),
	}
}

func (g *toddGame) Draw(img *ebiten.Image, ctx *engine.Graphics) {
	img.Fill(level.RGBA(level.Instance.World.GetBg()))
	ctx.SetWorldToScreen(Camera.GetTransform())
	for _, plat := range Platforms {
		plat.Draw(img, ctx)
	}
	Todd.Draw(img, ctx)
	g.ui.Draw(img, ctx)
}

func (g *toddGame) Resize(w, h int) {

	ar := float32(w-level.UIWidth) / float32(h)
	Camera = engine.NewCamera(
		engine.NewRect(-200, -10, 200, 400.0/ar-10),
		engine.NewRect(level.UIWidth, 0, w, h),
		engine.FlipYAxis)
	Camera.SetCenterX(Todd.pos.X)
	g.ui.Resize(w, h)
}

func AnimateCameraVertical() {
	b := Camera.WorldBounds()
	var target float32
	if Todd.pos.Y < 200 {
		target = -10
	} else {
		target = Todd.pos.Y - 20 - b.Height()/2
	}
	CameraVerticalAnimation = gween.New(
		float32(b.Bottom()),
		float32(target),
		float32(level.Instance.Camera.GetTiltSeconds()),
		level.CameraTiltEasing)
}

func ControlCamera(s *engine.UpdateState) {
	Camera.SetCenterX(Todd.pos.X)
	if CameraVerticalAnimation != nil {
		y, done := CameraVerticalAnimation.Update(float32(s.DeltaSeconds))
		Camera.RelativelyPositionY(float32(y), 0)
		if done {
			CameraVerticalAnimation = nil
		}
	}
	if Camera.Bottom() > Todd.pos.Y-10 {
		Camera.RelativelyPositionY(Todd.pos.Y-10, 0)
	}
}

func (g *toddGame) UpdateInput(s *engine.UpdateState) {
	// cap := imgui.CurrentIO().WantCaptureKeyboard()
	// if !(g.ui.Showing && cap) && inpututil.IsKeyJustPressed(ebiten.KeyU) {
	// 	g.ui.Showing = !g.ui.Showing
	// }
	g.ui.UpdateInput(s)
}

func (g *toddGame) UpdatePhysics(s *engine.UpdateState) {
	Todd.Update(s)
	ControlCamera(s)
}
