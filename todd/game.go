package todd

import (
	"fmt"
	"image"
	"image/color"
	"log"
	"math"
	"math/rand"

	"github.com/fogleman/gg"
	"github.com/golang/freetype/truetype"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/jdf/todd-again/todd/assets"
	"github.com/lucasb-eyer/go-colorful"
	"golang.org/x/image/font"
)

const (
	screenWidth  = 1200
	screenHeight = 600

	aspectRatio = float64(screenWidth) / float64(screenHeight)

	worldLeft  = -100.0
	worldRight = 100.0
)

// Box is a box is a box is a box. Loveliness extreme.
type Box struct {
	bounds         *Rect
	colorIndex     int
	colorDirection int
}

// Game is a game state.
type Game struct {
	// Graphics stuff.
	gfx         *Graphics
	frameBuffer *image.RGBA
	font        *truetype.Font
	debugFace   font.Face

	// Game state.
	frames int64

	// Entities
	camera      *Camera
	boxes       []*Box
	throbColors []colorful.Color
}

// Update updates the state of the game.
func (g *Game) Update() error {
	g.frames++

	if ebiten.IsKeyPressed(ebiten.KeyArrowLeft) {
		g.camera.Pan(&Vec2{-2, 0})
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowRight) {
		g.camera.Pan(&Vec2{2, 0})
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowUp) {
		g.camera.Pan(&Vec2{0, 2})
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowDown) {
		g.camera.Pan(&Vec2{0, -2})
	}

	_, wheelY := ebiten.Wheel()
	if math.Abs(wheelY) > 0.0 {
		g.camera.ZoomInto(
			1+(wheelY*.005),
			g.camera.ToWorldVec2(Vec(ebiten.CursorPosition())))
	}

	for _, b := range g.boxes {
		b.colorIndex += b.colorDirection
		if b.colorIndex >= len(g.throbColors) {
			b.colorIndex = len(g.throbColors) - 1
			b.colorDirection = -1
		}
		if b.colorIndex < 0 {
			b.colorIndex = 0
			b.colorDirection = 1
		}
	}

	return nil
}

// Draw draws the game screen in ebiten.
func (g *Game) Draw(screen *ebiten.Image) {
	dc := g.gfx

	dc.SetRGB(230, 230, 124)
	dc.Clear()

	dc.SetFontFace(g.debugFace)

	dc.SetColor(color.White)
	dc.SetLineWidth(4)
	dc.DrawLine(g.camera, worldLeft, 0, worldRight, 50)
	dc.Stroke()

	for _, x := range []int{-100, -50, 0, 50, 100} {
		for _, y := range []int{0, 25, 50, 75} {
			dc.SetColor(color.RGBA{00, 170, 00, 255})
			dc.DrawText(g.camera, fmt.Sprintf("%d,%d", x, y), float64(x), float64(y))
		}
	}

	for _, box := range g.boxes {
		if !g.camera.CanSee(box.bounds) {
			continue
		}
		dc.SetColor(g.throbColors[box.colorIndex])
		dc.FillRect(g.camera, box.bounds)
	}

	dc.SetColor(color.RGBA{0, 0, 0, 200})
	dc.FillRectScreen(g.camera, NewRect(2, 2, 120, 24))

	dc.SetColor(color.RGBA{128, 128, 128, 255})
	dc.DrawTextScreen(g.camera,
		fmt.Sprintf("FPS: %0.2f", ebiten.CurrentFPS()),
		4, 18)

	screen.ReplacePixels(g.frameBuffer.Pix)
}

var (
	lastW, lastH               int
	calculatedOw, calculatedOh int
)

// Layout has a	party with gnomes.
func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
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
		for w > 2048 || h > 2048 {
			w *= .5
			h *= .5
		}
		calculatedOw, calculatedOh = int(w), int(h)
		img := image.NewRGBA(image.Rect(0, 0, calculatedOw, calculatedOh))
		g.frameBuffer = img
		g.gfx = &Graphics{Context: *gg.NewContextForRGBA(img)}
		g.camera.SetScreenRect(NewRect(0, 0, float64(calculatedOw), float64(calculatedOh)))
		g.debugFace = truetype.NewFace(g.font, &truetype.Options{
			Size: 9,
			DPI:  72 * ebiten.DeviceScaleFactor(),
		})
		log.Printf("buffer size: %d, %d", calculatedOw, calculatedOh)
	}
	return calculatedOw, calculatedOh
}

func initColors() []colorful.Color {
	throbColors := []colorful.Color{}
	a, err := colorful.Hex("#6932a8")
	if err != nil {
		panic("failed to parse color")
	}
	b, err := colorful.Hex("#e1e823")
	if err != nil {
		panic("failed to parse color")
	}
	const colorTableSize = 60
	for i := 0; i < colorTableSize; i++ {
		t := float64(i) / colorTableSize
		throbColors = append(throbColors, a.BlendHcl(b, t).Clamped())
	}
	return throbColors
}

func initBoxes() []*Box {
	boxes := []*Box{}
	const boxCount = 20
	const boxWidth = 5
	boxGap := (worldRight - worldLeft) / (boxCount + 1.0)
	for i := 0; i < boxCount; i++ {
		x := worldLeft + boxGap*((float64)(i+1))
		boxes = append(boxes, &Box{
			bounds:         NewRect(x, 10, x+boxWidth, 10+boxWidth),
			colorIndex:     rand.Intn(60),
			colorDirection: 1,
		})
	}
	return boxes
}

// Run initializes and runs the game.
func Run() {
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Todd")
	ebiten.SetScreenClearedEveryFrame(false) // we blit the whole frame anyway
	img := image.NewRGBA(image.Rect(0, 0, screenWidth, screenHeight))

	font := assets.GetFontOrDie("InstructionBold.ttf")

	g := &Game{
		gfx: &Graphics{Context: *gg.NewContextForRGBA(img)},
		camera: NewCamera(
			NewRect(worldLeft, -1, worldLeft+100, 51),
			NewRect(0, 0, screenWidth, screenHeight)),
		frameBuffer: img,
		frames:      0,
		boxes:       initBoxes(),
		throbColors: initColors(),
		font:        font,
		debugFace:   truetype.NewFace(font, &truetype.Options{Size: 72}),
	}
	g.camera.SetInvertY(true)
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
