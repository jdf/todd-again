package todd

import (
	"fmt"
	"image"
	"image/color"
	"log"
	"math/rand"

	"github.com/fogleman/gg"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/lucasb-eyer/go-colorful"
)

const (
	screenWidth  = 1200
	screenHeight = 600

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

	if ebiten.IsKeyPressed(ebiten.KeyArrowLeft) && g.camera.Left() > worldLeft-5 {
		g.camera.Pan(&Vec2{-2, 0})
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowRight) && g.camera.Right() < worldRight+5 {
		g.camera.Pan(&Vec2{2, 0})
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

	dc.SetRGB(0, 0, 0)
	dc.Clear()

	for _, x := range []int{-100, -50, 0, 50, 100} {
		for _, y := range []int{0, 25, 50, 75} {
			dc.SetColor(color.RGBA{00, 128, 00, 255})
			dc.DrawText(g.camera, fmt.Sprintf("%d,%d", x, y), float64(x), float64(y))
		}
	}

	dc.SetColor(color.White)
	dc.SetLineWidth(4)
	dc.DrawLine(g.camera, worldLeft, 0, worldRight, 50)
	dc.Stroke()

	for _, box := range g.boxes {
		if !g.camera.CanSee(box.bounds) {
			continue
		}
		dc.SetColor(g.throbColors[box.colorIndex])
		dc.FillRect(g.camera, box.bounds)
	}

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
	g := &Game{
		gfx: &Graphics{Context: *gg.NewContextForRGBA(img)},
		camera: NewCamera(
			NewRect(worldLeft, -1, worldLeft+100, 51),
			NewRect(0, 0, screenWidth, screenHeight)),
		frameBuffer: img,
		frames:      0,
		boxes:       initBoxes(),
		throbColors: initColors(),
	}
	g.camera.SetInvertY(true)
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
