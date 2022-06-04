package todd

import (
	"fmt"
	"image"
	"image/color"
	"log"
	"time"

	"github.com/fogleman/gg"
	"github.com/golang/freetype/truetype"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/jdf/todd-again/todd/assets"
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
	gfx         *Graphics
	frameBuffer *image.RGBA
	font        *truetype.Font
	debugFace   font.Face

	lastUpdate         time.Time
	lastUpdateDebugLog time.Time

	// Entities
	camera *Camera
}

// Update updates the state of the game.
func (game *Game) Update() error {
	now := time.Now()
	//dt := now.Sub(game.lastUpdate)

	game.lastUpdate = now
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
		game.gfx = &Graphics{Context: *gg.NewContextForRGBA(img)}
		game.camera.SetScreenRect(NewRect(0, 0, float64(calculatedOw), float64(calculatedOh)))
		game.debugFace = truetype.NewFace(game.font, &truetype.Options{
			Size: 9,
			DPI:  72 * ebiten.DeviceScaleFactor(),
		})
		log.Printf("buffer size: %d, %d", calculatedOw, calculatedOh)
	}
	return calculatedOw, calculatedOh
}

// Run initializes and runs the game.
func Run() {
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Todd")
	ebiten.SetScreenClearedEveryFrame(false) // we blit the whole frame anyway
	img := image.NewRGBA(image.Rect(0, 0, screenWidth, screenHeight))

	font := assets.GetFontOrDie("InstructionBold.ttf")

	game := &Game{
		gfx: &Graphics{Context: *gg.NewContextForRGBA(img)},
		camera: NewCamera(
			NewRect(worldLeft, -1, worldLeft+100, 51),
			NewRect(0, 0, screenWidth, screenHeight)),
		frameBuffer:        img,
		font:               font,
		debugFace:          truetype.NewFace(font, &truetype.Options{Size: 72}),
		lastUpdate:         time.Now(),
		lastUpdateDebugLog: time.Now(),
	}
	game.camera.SetInvertY(true)
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
