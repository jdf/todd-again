package todd

import (
	"fmt"
	"image"
	"image/color"
	"log"
	"math/rand"

	"github.com/fogleman/gg"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/lucasb-eyer/go-colorful"
)

const (
	screenWidth  = 600
	screenHeight = 300

	worldLeft  = -100.0
	worldRight = 100.0
)

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

	if ebiten.IsKeyPressed(ebiten.KeyArrowLeft) && g.camera.Left() > worldLeft {
		g.camera.AddToSelf(&Vec2{-2, 0})
		log.Printf("camera: %v<->%v", g.camera.Min.X, g.camera.Max.X)
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowRight) && g.camera.Right() < worldRight {
		g.camera.AddToSelf(&Vec2{2, 0})
		log.Printf("camera: %v<->%v", g.camera.Min.X, g.camera.Max.X)
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

	// Draw the world.
	dc.Push()

	dr := func(c color.Color, s string) {
		dc.SetColor(c)
		dc.DrawRectangle(0, 0, 10, 10)
		dc.Fill()
		dc.SetRGB(1, 0, 0)
		dc.DrawRectangle(0, 0, 2, 2)
		dc.Fill()
		dc.SetColor(color.RGBA{0, 255, 0, 255})
		dc.DrawString(s, 10, 22)
	}
	_ = dr

	//dr(color.White, "before invert")
	//dc.InvertY()
	// dr(color.White, "before xlate")
	// dr(color.White, "after xlate")
	//dc.InvertY()
	//dc.Scale(screenWidth/g.camera.Size().Width, screenHeight/g.camera.Size().Height)
	//dc.Translate(-g.camera.Center().X, -g.camera.Center().Y)

	// dr(color.White, "after scale")
	//dc.Translate(50, 25)
	//dr(color.White, "FINAL")

	for _, x := range []int{-100, 0, 100} {
		for _, y := range []int{0, 25, 50, 75} {
			dc.SetColor(color.White)
			dc.DrawString(fmt.Sprintf("%d,%d", x, y), float64(x), float64(y))
		}
	}

	dc.SetColor(color.White)
	dc.SetLineWidth(4)
	dc.DrawLine(worldLeft, 0, worldRight, 50)
	dc.Stroke()

	for _, box := range g.boxes {
		// if !box.bounds.Intersects(g.camera) {
		// 	continue
		// }
		dc.SetColor(g.throbColors[box.colorIndex])
		//dc.SetColor(color.White)
		r := box.bounds
		dc.DrawRectangle(r.Min.X, r.Min.Y, r.Size().X, r.Size().Y)
		dc.Fill()
	}

	dc.Pop()

	//dc.SetRGB(1, 1, 1)
	//dc.DrawRectangle(20, 20, 100, 100)
	//dc.Fill()

	screen.ReplacePixels(g.frameBuffer.Pix)
	// ebitenutil.DebugPrint(screen,
	// 	fmt.Sprintf("TPS: %0.2f\nFPS: %0.2f",
	// 		ebiten.CurrentTPS(),
	// 		ebiten.CurrentFPS()))
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
		gfx:         &Graphics{Context: *gg.NewContextForRGBA(img)},
		camera:      NewRect(worldLeft, 0, worldLeft+100, 50),
		frameBuffer: img,
		frames:      0,
		boxes:       initBoxes(),
		throbColors: initColors(),
	}
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
