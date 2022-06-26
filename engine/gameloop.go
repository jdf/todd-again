package engine

import (
	"fmt"
	"image/color"
	"log"
	"time"

	"github.com/golang/freetype/truetype"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/jdf/todd-again/engine/dbgassets"
	"golang.org/x/image/font"
)

const (
	// Fixed timestep for physics simulation, independent of ebiten update frequency.
	tick                        = 1 / 180.0
	maxOffscreenBufferDimension = 4096
)

const debug = false

type windowInfo struct {
	// width to height ratio of the game's desired viewport
	aspectRatio float64

	// Managing layout.
	// lastW and lastH are the last device pixel width and height requested
	// during a Layout call. Because we do expensive things when the size changes,
	// we fast-return if the size hasn't changed.
	lastW, lastH int
	// bufW and bufH are the actual width and height of the offscreen buffer. They will
	// have the right aspect ratio, but may be smaller than the lastW and lastH.
	bufW, bufH int
}

type ebitenGame struct {
	userGame GameModule

	window windowInfo

	// Graphics stuff.
	gfx       *Graphics
	debugFont *truetype.Font
	debugFace font.Face

	// We maintain an accumulator of excess time remaining after going through
	// physics ticks.
	accumulator          float64
	lastEbitenUpdate     time.Time
	lastWorldTimeSeconds float64
}

// Update is called periodically by ebiten. We use it to run physics
// simulation and update the game state, likely breaking each ebiten
// update into multiple physics ticks.
func (game *ebitenGame) Update() error {
	input := GetInputState()

	now := time.Now()
	dt := now.Sub(game.lastEbitenUpdate)
	game.lastEbitenUpdate = now

	us := &UpdateState{
		Input:        input,
		DeltaSeconds: tick,
	}

	game.userGame.UpdateInput(us)
	for game.accumulator += dt.Seconds(); game.accumulator >= tick; game.accumulator -= tick {
		us.NowSeconds = game.lastWorldTimeSeconds + tick
		game.userGame.UpdatePhysics(us)
		game.lastWorldTimeSeconds += tick
	}

	return nil
}

func drawDebugInfo(img *ebiten.Image, game *ebitenGame) {
	g := game.gfx
	g.SetFont(game.debugFace)
	g.SetColor(color.RGBA{0, 0, 0, 200})
	//g.FillRectScreen(NewRect(2, 2, 120, 24))

	g.SetColor(color.RGBA{128, 128, 128, 255})
	g.DrawTextScreen(img,
		fmt.Sprintf("FPS: %0.2f", ebiten.CurrentFPS()),
		4, 18)
}

// Draw draws the game screen in ebiten.
func (game *ebitenGame) Draw(screen *ebiten.Image) {
	screen.Fill(color.Black)
	game.userGame.Draw(screen, game.gfx)

	if debug {
		drawDebugInfo(screen, game)
	}
}

// Layout has a	party with gnomes.
func (game *ebitenGame) Layout(outsideWidth, outsideHeight int) (int, int) {
	win := &game.window
	if outsideWidth != win.lastW || outsideHeight != win.lastH {
		log.Printf("layout %dx%d", outsideWidth, outsideHeight)
		win.lastW = outsideWidth
		win.lastH = outsideHeight
		s := ebiten.DeviceScaleFactor()
		w, h := s*float64(outsideWidth), s*float64(outsideHeight)
		log.Printf("scaled %0.1fx%0.1f", w, h)
		if w/h > win.aspectRatio {
			w = h * win.aspectRatio
		} else {
			h = w / win.aspectRatio
		}
		for w > maxOffscreenBufferDimension || h > maxOffscreenBufferDimension {
			w *= .5
			h *= .5
		}
		win.bufW, win.bufH = int(w), int(h)
		game.gfx = NewGraphics()
		game.userGame.Resize(win.bufW, win.bufH)
		game.debugFace = truetype.NewFace(game.debugFont, &truetype.Options{
			Size:    9,
			DPI:     72 * ebiten.DeviceScaleFactor(),
			Hinting: font.HintingFull,
		})
		log.Printf("buffer size: %d, %d", win.bufW, win.bufH)
	}
	return win.bufW, win.bufH
}

// RunGameLoop initializes and runs the game.
func RunGameLoop(userGame GameModule, width, height int, title string) {
	ebiten.SetWindowSize(width, height)
	ebiten.SetWindowTitle(title)
	ebiten.SetScreenClearedEveryFrame(true)

	font := dbgassets.GetFontOrDie("InstructionBold.ttf")

	game := &ebitenGame{
		userGame:         userGame,
		window:           windowInfo{aspectRatio: float64(width) / float64(height)},
		debugFont:        font,
		debugFace:        truetype.NewFace(font, &truetype.Options{Size: 72}),
		lastEbitenUpdate: time.Now(),
	}
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
