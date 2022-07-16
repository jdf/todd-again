package tuning

import (
	"flag"
	"image/color"

	"github.com/gabstv/ebiten-imgui/renderer"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/inkyblackness/imgui-go/v4"
	"github.com/jdf/todd-again/engine"
	"github.com/jdf/todd-again/game/proto"
)

var showUI = flag.Bool("show-ui", false, "Show the UI")

const (
	UIWidth = 600
)

type UI struct {
	mgr     *renderer.Manager
	Showing bool
	height  float32
}

func NewUI() *UI {
	return &UI{
		mgr: renderer.New(nil),
	}
}

func (ui *UI) Resize(w, h int) {
	ui.height = float32(h)
	ui.mgr.SetDisplaySize(float32(w), ui.height)
	imgui.CurrentIO().SetFontGlobalScale(2.5)
}

func (ui *UI) UpdatePhysics(s *engine.UpdateState) {}

func (ui *UI) UpdateInput(s *engine.UpdateState) {
	if !*showUI {
		return
	}
	ui.mgr.Update(float32(s.DeltaSeconds))
	ui.mgr.BeginFrame()
	imgui.SetNextWindowPos(imgui.Vec2{X: 0, Y: 0})
	imgui.SetNextWindowSize(imgui.Vec2{X: float32(UIWidth), Y: ui.height})
	imgui.Begin("Settings")
	proto.RenderTuning(Instance)
	imgui.End()
	ui.mgr.EndFrame()
}

func (ui *UI) Draw(screen *ebiten.Image, g *engine.Graphics) {
	if *showUI {
		g.SetColor(color.Black)
		g.DrawRectScreen(screen, 0, 0, UIWidth, int(ui.height))
		ui.mgr.Draw(screen)
	}
}
