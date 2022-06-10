package engine

import (
	"fmt"
	"image"
	"image/color"
	"log"
	"math"
	"time"

	"github.com/fogleman/gg"
	"github.com/golang/freetype/truetype"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/jdf/todd-again/engine/dbgassets"
	"golang.org/x/image/font"
)

const (
	screenWidth  = 1600
	screenHeight = 900

	aspectRatio = float64(screenWidth) / float64(screenHeight)

	worldLeft  = -100.0
	worldRight = 100.0

	debug = true
)

// Game is a game state.
type Game struct {
	// Graphics stuff.
	gfx         *Context
	frameBuffer *image.RGBA
	debugFont   *truetype.Font
	debugFace   font.Face

	lastUpdate         time.Time
	lastUpdateDebugLog time.Time

	camera *Camera
	level  *Level
}

// Update updates the state of the game.
func (game *Game) Update() error {
	now := time.Now()
	dt := now.Sub(game.lastUpdate)

	frameState := &FrameState{
		Camera: game.camera,
		Input:  GetInputState(),
		Now:    now,
		DeltaT: dt.Seconds(),
	}
	game.level.Update(frameState)

	game.lastUpdate = now

	if frameState.Input.Left {
		game.camera.Pan(-2, 0)
	}
	if frameState.Input.Right {
		game.camera.Pan(2, 0)
	}
	if frameState.Input.Up {
		game.camera.Pan(0, 2)
	}
	if frameState.Input.Down {
		game.camera.Pan(0, -2)
	}

	_, wheelY := ebiten.Wheel()
	if math.Abs(wheelY) > 0.0 {
		game.camera.ZoomInto(
			1+(wheelY*.005),
			game.camera.ToWorldVec2(Vec(ebiten.CursorPosition())))
	}

	return nil
}

func drawDebugInfo(game *Game, camera *Camera) {
	g := game.gfx
	g.SetFontFace(game.debugFace)
	g.SetColor(color.RGBA{0, 0, 0, 200})
	g.FillRectScreen(camera, NewRect(2, 2, 120, 24))

	g.SetColor(color.RGBA{128, 128, 128, 255})
	g.DrawTextScreen(camera,
		fmt.Sprintf("FPS: %0.2f", ebiten.CurrentFPS()),
		4, 18)
}

// Draw draws the game screen in ebiten.
func (game *Game) Draw(screen *ebiten.Image) {
	g := game.gfx

	g.SetRGB(0, 0, 0)
	g.Clear()

	game.level.Draw(g, game.camera)
	if debug {
		drawDebugInfo(game, game.camera)
	}

	screen.ReplacePixels(game.frameBuffer.Pix)
}

const (
	maxOffscreenBufferDimension = 2048
)

var (
	lastW, lastH               int
	calculatedOw, calculatedOh int
)

// Layout has a	party with gnomes.
func (game *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	if outsideWidth != lastW || outsideHeight != lastH {
		log.Printf("layout %dx%d", outsideWidth, outsideHeight)
		lastW = outsideWidth
		lastH = outsideHeight
		s := ebiten.DeviceScaleFactor()
		w, h := s*float64(outsideWidth), s*float64(outsideHeight)
		if w/h > aspectRatio {
			w = h * aspectRatio
		} else {
			h = w / aspectRatio
		}
		for w > maxOffscreenBufferDimension || h > maxOffscreenBufferDimension {
			w *= .5
			h *= .5
		}
		calculatedOw, calculatedOh = int(w), int(h)
		img := image.NewRGBA(image.Rect(0, 0, calculatedOw, calculatedOh))
		game.frameBuffer = img
		game.gfx = &Context{Context: *gg.NewContextForRGBA(img)}
		game.camera.SetScreenRect(NewRect(0, 0, float64(calculatedOw), float64(calculatedOh)))
		game.debugFace = truetype.NewFace(game.debugFont, &truetype.Options{
			Size: 9,
			DPI:  72 * ebiten.DeviceScaleFactor(),
		})
		log.Printf("buffer size: %d, %d", calculatedOw, calculatedOh)
	}
	return calculatedOw, calculatedOh
}

// RunGameLoop initializes and runs the game.
func RunGameLoop() {
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Todd")
	ebiten.SetScreenClearedEveryFrame(false) // we blit the whole frame anyway
	img := image.NewRGBA(image.Rect(0, 0, screenWidth, screenHeight))

	font := dbgassets.GetFontOrDie("InstructionBold.ttf")

	game := &Game{
		gfx: &Context{Context: *gg.NewContextForRGBA(img)},
		camera: NewCamera(
			NewRect(worldLeft, -1, worldLeft+100, 51),
			NewRect(0, 0, screenWidth, screenHeight),
			FlipYAxis),
		frameBuffer:        img,
		debugFont:          font,
		debugFace:          truetype.NewFace(font, &truetype.Options{Size: 72}),
		lastUpdate:         time.Now(),
		lastUpdateDebugLog: time.Now(),
		level:              Level1(),
	}
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
