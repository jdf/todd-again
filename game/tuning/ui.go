package tuning

import (
	"flag"

	"github.com/gabstv/ebiten-imgui/renderer"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/inkyblackness/imgui-go/v4"
	"github.com/jdf/todd-again/engine"
	"github.com/jdf/todd-again/game/proto"
)

var showUI = flag.Bool("show-ui", false, "Show the UI")

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
	imgui.CurrentIO().SetFontGlobalScale(2.5)
}

func (ui *UI) UpdatePhysics(s *engine.UpdateState) {}

func (ui *UI) UpdateInput(s *engine.UpdateState) {
	if !*showUI {
		return
	}
	ui.mgr.Update(float32(s.DeltaSeconds))
	ui.mgr.BeginFrame()
	proto.RenderTuning(Instance)
	ui.mgr.EndFrame()
}

func (ui *UI) Draw(screen *ebiten.Image, g *engine.Graphics) {
	if *showUI {
		ui.mgr.Draw(screen)
	}
}
