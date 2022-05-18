package todd

import (
	"fmt"
	"image"
	"log"
	"math"

	"github.com/fogleman/gg"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

const (
	worldViewWidth = 600
	screenWidth    = 600
	screenHeight   = 300
)

// Game is a game state.
type Game struct {
	ggCtx       *gg.Context
	frameBuffer *image.RGBA

	scaleBuffer *ebiten.Image
	frames      int64
}

// Update updates the state of the game.
func (g *Game) Update() error {
	g.frames++

	return nil
}

// Draw draws the game screen in ebiten.
func (g *Game) Draw(screen *ebiten.Image) {
	dc := g.ggCtx

	dc.SetRGBA(0, 0, 0, 1)
	dc.Clear()
	dc.SetRGBA(1, 1, 1, 0.1)

	const (
		waveFreq           = 0.1
		waveRotStepRadians = math.Pi / 8.0
		rotFreq            = 0.005
		hw                 = screenWidth / 2
		hh                 = screenHeight / 2
	)
	dc.Push()
	dc.RotateAbout(float64(g.frames)*rotFreq, hw, hh)
	waveAngle := 0.0
	for {
		dc.Push()
		dc.RotateAbout(float64(waveAngle), hw, hh)
		dc.DrawEllipse(hw, hh, screenWidth*7/16, screenHeight/8)
		dc.Fill()
		dc.Pop()
		waveAngle += waveRotStepRadians
		if math.Abs(math.Mod(waveAngle, 2*math.Pi)) < .1 {
			break
		}
	}
	dc.Pop()
	screen.ReplacePixels(g.frameBuffer.Pix)
	ebitenutil.DebugPrint(screen,
		fmt.Sprintf("TPS: %0.2f\nFPS: %0.2f",
			ebiten.CurrentTPS(),
			ebiten.CurrentFPS()))
}

// Layout has a	party with gnomes.
func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

// Run initializes and runs the game.
func Run() {
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Ellipses")
	img := image.NewRGBA(image.Rect(0, 0, screenWidth, screenHeight))
	g := &Game{
		ggCtx:       gg.NewContextForRGBA(img),
		frameBuffer: img,
		frames:      0,
	}
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
