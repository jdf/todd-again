package tuning

import (
	_ "embed"
	"fmt"

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
	imgui.CurrentIO().SetFontGlobalScale(3.0)
}

var (
	counter   = 0
	floatVal  = float32(0.3)
	name      = "Todd"
	someColor = [3]float32{0.3, 0.3, 0.3}
)

func (ui *UI) UpdatePhysics(s *engine.UpdateState) {}

func (ui *UI) UpdateInput(s *engine.UpdateState) {
	if !ui.Showing {
		return
	}
	ui.mgr.Update(float32(s.DeltaSeconds))
	ui.mgr.BeginFrame()
	{
		imgui.Text("ภาษาไทย测试조선말")                      // To display these, you'll need to register a compatible font
		imgui.Text("Hello, world!")                     // Display some text
		imgui.SliderFloat("float", &floatVal, 0.0, 1.0) // Edit 1 float using a slider from 0.0f to 1.0f
		imgui.ColorEdit3("clear color", &someColor)     // Edit 3 floats representing a color

		//imgui.Checkbox("Demo Window", &showDemoWindow) // Edit bools storing our window open/close state
		//imgui.Checkbox("Go Demo Window", &showGoDemoWindow)
		//imgui.Checkbox("Another Window", &showAnotherWindow)

		if imgui.Button("Button") { // Buttons return true when clicked (most widgets return true when edited/activated)
			counter++
		}
		imgui.SameLine()
		imgui.Text(fmt.Sprintf("counter = %d", counter))

		imgui.InputText("Name", &name)
		imgui.Spacing()
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
