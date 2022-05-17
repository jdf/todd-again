package main

import (
	"flag"
	"fmt"
	"image"
	"log"
	"math"
	"os"
	"runtime/pprof"

	"github.com/fogleman/gg"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

const (
	screenWidth  = 640
	screenHeight = 640
)

// Game is a game state.
type Game struct {
	ctx   *gg.Context
	img   *image.RGBA
	w, h  int
	frame int64
}

// Update updates the state of the game.
func (g *Game) Update() error {
	g.frame++

	dc := g.ctx

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
	dc.RotateAbout(float64(g.frame)*rotFreq, hw, hh)
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

	return nil
}

// Draw draws the game screen in ebiten.
func (g *Game) Draw(screen *ebiten.Image) {
	screen.ReplacePixels(g.img.Pix)
	ebitenutil.DebugPrint(screen,
		fmt.Sprintf("TPS: %0.2f\nFPS: %0.2f",
			ebiten.CurrentTPS(),
			ebiten.CurrentFPS()))
}

// Layout has a	party with gnomes.
func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func main() {
	var cpuprofile = flag.String("cpuprofile", "", "write cpu profile to `file`")
	flag.Parse()
	if *cpuprofile != "" {
		f, err := os.Create(*cpuprofile)
		if err != nil {
			log.Fatal("could not create CPU profile: ", err)
		}
		defer f.Close() // error handling omitted for example
		if err := pprof.StartCPUProfile(f); err != nil {
			log.Fatal("could not start CPU profile: ", err)
		}
		defer pprof.StopCPUProfile()
	}

	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Ellipses")
	img := image.NewRGBA(image.Rect(0, 0, screenWidth, screenHeight))
	g := &Game{
		ctx:   gg.NewContextForRGBA(img),
		img:   img,
		w:     screenWidth,
		h:     screenHeight,
		frame: 0,
	}
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
