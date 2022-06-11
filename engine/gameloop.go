package engine

import (
	"errors"
	"fmt"
	"image"
	"image/color"
	"log"
	"time"

	"github.com/fogleman/gg"
	"github.com/golang/freetype/truetype"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/jdf/todd-again/engine/dbgassets"
	"golang.org/x/image/font"
)

const (
	// Fixed timestep for physics simulation, independent of ebiten update frequency.
	tick                        = 1 / 180.0
	maxOffscreenBufferDimension = 2048
)

const debug = true

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
	userGame Game

	window windowInfo

	// Graphics stuff.
	gfx         *Graphics
	frameBuffer *image.RGBA
	debugFont   *truetype.Font
	debugFace   font.Face

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

	if input.Q {
		return errors.New("Quit")
	}

	now := time.Now()
	dt := now.Sub(game.lastEbitenUpdate)
	game.lastEbitenUpdate = now

	us := &UpdateState{
		Input:        input,
		DeltaSeconds: tick,
	}

	for game.accumulator += dt.Seconds(); game.accumulator >= tick; game.accumulator -= tick {
		us.NowSeconds = game.lastWorldTimeSeconds + tick
		game.userGame.Update(us)
		game.lastWorldTimeSeconds += tick
	}

	return nil
}

func drawDebugInfo(game *ebitenGame) {
	g := game.gfx
	g.SetFontFace(game.debugFace)
	g.SetColor(color.RGBA{0, 0, 0, 200})
	g.FillRectScreen(NewRect(2, 2, 120, 24))

	g.SetColor(color.RGBA{128, 128, 128, 255})
	g.DrawTextScreen(
		fmt.Sprintf("FPS: %0.2f", ebiten.CurrentFPS()),
		4, 18)
}

// Draw draws the game screen in ebiten.
func (game *ebitenGame) Draw(screen *ebiten.Image) {
	g := game.gfx

	g.SetRGB(0, 0, 0)
	g.Clear()
	game.userGame.Draw(g)

	if debug {
		drawDebugInfo(game)
	}
	screen.ReplacePixels(game.frameBuffer.Pix)
}

// Layout has a	party with gnomes.
func (game *ebitenGame) Layout(outsideWidth, outsideHeight int) (int, int) {
	win := &game.window
	if outsideWidth != win.lastW || outsideHeight != win.lastH {
		log.Printf("layout %dx%d", outsideWidth, outsideHeight)
		win.lastW = outsideWidth
		win.lastH = outsideHeight
		s := 1.0 // ebiten.DeviceScaleFactor()
		w, h := s*float64(outsideWidth), s*float64(outsideHeight)
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
		img := image.NewRGBA(image.Rect(0, 0, win.bufW, win.bufH))
		game.frameBuffer = img
		game.gfx = &Graphics{Context: *gg.NewContextForRGBA(img)}
		game.userGame.Resize(win.bufW, win.bufH)
		game.debugFace = truetype.NewFace(game.debugFont, &truetype.Options{
			Size: 9,
			DPI:  72 * ebiten.DeviceScaleFactor(),
		})
		log.Printf("buffer size: %d, %d", win.bufW, win.bufH)
	}
	return win.bufW, win.bufH
}

// RunGameLoop initializes and runs the game.
func RunGameLoop(userGame Game, width, height int, title string) {
	ebiten.SetWindowSize(width, height)
	ebiten.SetWindowTitle(title)
	ebiten.SetScreenClearedEveryFrame(false) // we blit the whole frame anyway

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
