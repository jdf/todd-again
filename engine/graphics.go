package engine

import (
	"fmt"
	"image"
	"image/color"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
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

func (g *Graphics) Translate(x, y float64) {
	g.objectToWorld.Append(Translation(x, y))
	g.cachedObjectToScreen = nil
}

func (g *Graphics) Rotate(angle float64) {
	g.objectToWorld.Append(Rotation(angle))
	g.cachedObjectToScreen = nil
}

func (g *Graphics) RotateAround(angle float64, center *Vec2) {
	g.objectToWorld.Append(RotationAround(angle, center))
	g.cachedObjectToScreen = nil
}

// ToPixel transforms a point in object space to a point in pixel space,
// rounding to the nearest integer.
func (g *Graphics) ToPixel(xWorld, yWorld float64) (float64, float64) {
	xScreen, yScreen := g.ObjectToScreen().Transform(xWorld, yWorld)
	return math.Round(xScreen), math.Round(yScreen)
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
func (g *Graphics) DrawLine(img *ebiten.Image, x1, y1, x2, y2 float64) {
	path := &vector.Path{}
	x1, y1 = g.ToPixel(x1, y1)
	x2, y2 = g.ToPixel(x2, y2)
	path.MoveTo(float32(x1), float32(y1))
	path.LineTo(float32(x2), float32(y2))
	g.fillPath(img, path)
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

func (g *Graphics) appendEllipticalArc(p *vector.Path, center, radii *Vec2, startAngle, endAngle float64, mode PathMode) {
	angleDelta := math.Abs(math.Atan2(math.Sin(startAngle-endAngle), math.Cos(startAngle-endAngle)))
	n := int(math.Round(4.0 * angleDelta / (.5 * math.Pi)))
	if n == 0 {
		n = 16
	}
	x, y := center.X, center.Y
	rx, ry := radii.X, radii.Y
	for i := 0; i < n; i++ {
		p1 := float64(i+0) / float64(n)
		p2 := float64(i+1) / float64(n)
		a1 := startAngle + (endAngle-startAngle)*p1
		a2 := startAngle + (endAngle-startAngle)*p2
		x0 := x + rx*math.Cos(a1)
		y0 := y + ry*math.Sin(a1)
		x1 := x + rx*math.Cos((a1+a2)/2)
		y1 := y + ry*math.Sin((a1+a2)/2)
		x2 := x + rx*math.Cos(a2)
		y2 := y + ry*math.Sin(a2)
		cx := 2*x1 - x0/2 - x2/2
		cy := 2*y1 - y0/2 - y2/2

		x0, y0 = g.ToPixel(x0, y0)
		cx, cy = g.ToPixel(cx, cy)
		x2, y2 = g.ToPixel(x2, y2)

		if i == 0 {
			if mode == PathModeContinue {
				p.LineTo(float32(x0), float32(y0))
			} else {
				p.MoveTo(float32(x0), float32(y0))
			}
		}
		p.QuadTo(float32(cx), float32(cy), float32(x2), float32(y2))
	}
}

var kCornerAngles = []float64{1.5 * math.Pi, math.Pi, .5 * math.Pi, 0, -.5 * math.Pi}

var DebugRoundRect = false

func (g *Graphics) DrawRoundedRect(img *ebiten.Image, r *Rect, radius float64) {
	if DebugRoundRect {
		fmt.Println(g.ObjectToScreen().TransformRect(r).Max.X)
	}
	path := &vector.Path{}
	maxRadius := math.Min(.5*r.Width(), .5*r.Height())
	actualRadius := math.Min(radius, maxRadius)
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
