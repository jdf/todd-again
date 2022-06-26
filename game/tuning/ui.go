package tuning

import (
	_ "embed"

	"github.com/gabstv/ebiten-imgui/renderer"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/inkyblackness/imgui-go/v4"
	"github.com/jdf/todd-again/engine"
)

type UI struct {
	mgr     *renderer.Manager
	Showing bool
}

func NewUI() *UI {
	return &UI{
		mgr: renderer.New(nil),
	}
}

func (ui *UI) Resize(w, h int) {
	ui.mgr.SetDisplaySize(float32(w), float32(h))
	imgui.CurrentIO().SetFontGlobalScale(2.0)
}

func (ui *UI) UpdatePhysics(s *engine.UpdateState) {}

func (ui *UI) UpdateInput(s *engine.UpdateState) {
	if !ui.Showing {
		return
	}
	ui.mgr.Update(float32(s.DeltaSeconds))
	ui.mgr.BeginFrame()
	{
		imgui.CollapsingHeader("World Physics")
		if Instance.Gravity == nil {
			g := Instance.GetGravity()
			Instance.Gravity = &g
		}
		imgui.SliderFloat("Gravity", Instance.Gravity, -5000, 0)
		imgui.Dummy(imgui.Vec2{X: 0, Y: 60})
		imgui.SameLineV(0, imgui.ContentRegionAvail().X*.666)
		if imgui.Button("Hide") {
			ui.Showing = false
		}
	}
	ui.mgr.EndFrame()
}

func (ui *UI) Draw(screen *ebiten.Image, g *engine.Graphics) {
	if ui.Showing {
		ui.mgr.Draw(screen)
	}
}
