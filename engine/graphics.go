package engine

import (
	"fmt"
	"image"
	"image/color"
	"math"

	"git.maze.io/go/math32"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/text"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"golang.org/x/image/font"
)

var (
	emptyImage = ebiten.NewImage(3, 3)

	// emptySubImage is an internal sub image of emptyImage.
	// Use emptySubImage at DrawTriangles instead of emptyImage in order to avoid bleeding edges.
	emptySubImage = emptyImage.SubImage(image.Rect(1, 1, 2, 2)).(*ebiten.Image)
)

func init() {
	emptyImage.Fill(color.White)
}

// Graphics is a graphics context that knows how to manipulate
// coordinates with a camera.
type Graphics struct {
	color color.Color
	font  font.Face

	worldToScreen *Affine

	objectToWorld Affine
	stack         *Stack[Affine]

	cachedObjectToScreen *Affine
}

func NewGraphics() *Graphics {
	return &Graphics{
		color:                color.White,
		objectToWorld:        *Identity(),
		stack:                &Stack[Affine]{},
		worldToScreen:        Identity(),
		cachedObjectToScreen: nil,
	}
}

func (g *Graphics) SetFont(font font.Face) {
	g.font = font
}

func (g *Graphics) SetColor(color color.Color) {
	g.color = color
}

func (g *Graphics) Translate(x, y float32) {
	g.objectToWorld.Append(Translation(x, y))
	g.cachedObjectToScreen = nil
}

func (g *Graphics) Rotate(angle float32) {
	g.objectToWorld.Append(Rotation(angle))
	g.cachedObjectToScreen = nil
}

func (g *Graphics) RotateAround(angle float32, center *Vec2) {
	g.objectToWorld.Append(RotationAround(angle, center))
	g.cachedObjectToScreen = nil
}

// ToPixel transforms a point in object space to a point in pixel space,
// rounding to the nearest integer.
func (g *Graphics) ToPixel(xWorld, yWorld float32) (float64, float64) {
	xScreen, yScreen := g.ObjectToScreen().Transform(xWorld, yWorld)
	return math.Round(float64(xScreen)), math.Round(float64(yScreen))
}

// ToPixel transforms a point in object space to a point in pixel space,
// rounding to the nearest integer.
func (g *Graphics) ToPixelRect(r *Rect) *Rect {
	minX, minY := g.ToPixel(r.Left(), r.Bottom())
	maxX, maxY := g.ToPixel(r.Right(), r.Top())
	return NewRect(minX, minY, maxX, maxY)
}

func (g *Graphics) ObjectToScreen() *Affine {
	if g.cachedObjectToScreen == nil {
		g.cachedObjectToScreen = Compose(&g.objectToWorld, g.worldToScreen)
	}
	return g.cachedObjectToScreen
}

func (g *Graphics) SetWorldToScreen(worldToScreen *Affine) {
	g.worldToScreen = worldToScreen
	g.cachedObjectToScreen = nil
}

func (g *Graphics) Push() {
	g.stack.Push(*g.objectToWorld.Copy())
}

func (g *Graphics) Pop() {
	g.objectToWorld = g.stack.Top()
	g.stack.Pop()
	g.cachedObjectToScreen = nil
}

// DrawTextScreen draws text at the given position in screen space.
func (g *Graphics) DrawTextScreen(img *ebiten.Image, s string, x, y int) {
	text.Draw(img, s, g.font, x, y, g.color)
}

func (g *Graphics) fillPath(img *ebiten.Image, p *vector.Path) {
	red, green, blue, _ := g.color.RGBA()
	vs, i := p.AppendVerticesAndIndicesForFilling(nil, nil)
	for i := range vs {
		vs[i].SrcX = 1
		vs[i].SrcY = 1
		vs[i].ColorR = float32(red>>8) / float32(0xff)
		vs[i].ColorG = float32(green>>8) / float32(0xff)
		vs[i].ColorB = float32(blue>>8) / float32(0xff)
	}
	img.DrawTriangles(vs, i, emptySubImage, nil)
}

// DrawLine draws a line from x1,y1 to x2,y2 in object space.
func (g *Graphics) DrawLine(img *ebiten.Image, x1, y1, x2, y2 float32) {
	x16, y16 := g.ToPixel(x1, y1)
	x26, y26 := g.ToPixel(x2, y2)
	ebitenutil.DrawLine(img, x16, y16, x26, y26, g.color)
}

func (g *Graphics) DrawRectScreen(img *ebiten.Image, left, top, right, bottom int) {
	img.SubImage(
		image.Rectangle{
			image.Point{left, top},
			image.Point{right, bottom},
		}).(*ebiten.Image).Fill(g.color)
}

func (g *Graphics) DrawRect(img *ebiten.Image, r *Rect) {
	sr := g.ToPixelRect(r)
	path := &vector.Path{}
	minx, miny, maxx, maxy := float32(sr.Min.X), float32(sr.Min.Y), float32(sr.Max.X), float32(sr.Max.Y)
	path.MoveTo(minx, miny)
	path.LineTo(maxx, miny)
	path.LineTo(maxx, maxy)
	path.LineTo(minx, maxy)
	path.LineTo(minx, miny)
	g.fillPath(img, path)
}

type PathMode int

const (
	PathModeContinue PathMode = iota
	PathModeNewShape
)

func (g *Graphics) appendEllipticalArc(p *vector.Path, center, radii *Vec2, startAngle, endAngle float32, mode PathMode) {
	angleDelta := math32.Abs(math32.Atan2(math32.Sin(startAngle-endAngle), math32.Cos(startAngle-endAngle)))
	n := int(math.Round(float64(4.0 * angleDelta / (.5 * math.Pi))))
	if n == 0 {
		n = 16
	}
	x, y := center.X, center.Y
	rx, ry := radii.X, radii.Y
	for i := 0; i < n; i++ {
		p1 := float32(i+0) / float32(n)
		p2 := float32(i+1) / float32(n)
		a1 := startAngle + (endAngle-startAngle)*p1
		a2 := startAngle + (endAngle-startAngle)*p2
		x0 := x + rx*math32.Cos(a1)
		y0 := y + ry*math32.Sin(a1)
		x1 := x + rx*math32.Cos((a1+a2)/2)
		y1 := y + ry*math32.Sin((a1+a2)/2)
		x2 := x + rx*math32.Cos(a2)
		y2 := y + ry*math32.Sin(a2)
		cx := 2*x1 - x0/2 - x2/2
		cy := 2*y1 - y0/2 - y2/2

		x0d, y0d := g.ToPixel(x0, y0)
		cxd, cyd := g.ToPixel(cx, cy)
		x2d, y2d := g.ToPixel(x2, y2)

		if i == 0 {
			if mode == PathModeContinue {
				p.LineTo(float32(x0d), float32(y0d))
			} else {
				p.MoveTo(float32(x0d), float32(y0d))
			}
		}
		p.QuadTo(float32(cxd), float32(cyd), float32(x2d), float32(y2d))
	}
}

var kCornerAngles = []float32{1.5 * math.Pi, math.Pi, .5 * math.Pi, 0, -.5 * math.Pi}

var DebugRoundRect = false

func (g *Graphics) DrawRoundedRect(img *ebiten.Image, r *Rect, radius float32) {
	if DebugRoundRect {
		fmt.Println(g.ObjectToScreen().TransformRect(r).Max.X)
	}
	path := &vector.Path{}
	maxRadius := math32.Min(.5*r.Width(), .5*r.Height())
	actualRadius := math32.Min(radius, maxRadius)
	radii := Vec(actualRadius, actualRadius)
	arcCenters := r.Inset(radii).Corners()
	for i := 0; i < 4; i++ {
		pathMode := PathModeContinue
		if i == 0 {
			pathMode = PathModeNewShape
		}
		g.appendEllipticalArc(path, &arcCenters[i], radii, kCornerAngles[i], kCornerAngles[i+1], pathMode)
	}
	g.fillPath(img, path)
}

func (g *Graphics) DrawEllipse(img *ebiten.Image, r *Rect) {
	path := &vector.Path{}
	g.appendEllipticalArc(path, r.Center(), r.Size().Mul(.5), 0, 2*math.Pi, PathModeNewShape)
	g.fillPath(img, path)
}
