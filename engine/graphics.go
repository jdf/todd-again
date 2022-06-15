package engine

import (
	"math"

	"github.com/fogleman/gg"
)

// Graphics is a graphics context that knows how to manipulate
// coordinates with a camera.
type Graphics struct {
	gg.Context

	worldToScreen *Affine

	objectToWorld Affine
	stack         *Stack[Affine]

	cachedObjectToScreen *Affine
}

func NewGraphics(wrapped *gg.Context) *Graphics {
	return &Graphics{
		Context:              *wrapped,
		objectToWorld:        *Identity(),
		stack:                &Stack[Affine]{},
		worldToScreen:        Identity(),
		cachedObjectToScreen: nil,
	}
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
	g.stack.Push(g.objectToWorld)
}

func (g *Graphics) Pop() {
	g.objectToWorld = g.stack.Top()
	g.stack.Pop()
	g.cachedObjectToScreen = nil
}

// DrawTextScreen draws text at the given position in screen space.
func (g *Graphics) DrawTextScreen(s string, x, y float64) {
	g.DrawString(s, x, y)
}

// DrawLine draws a line from x1,y1 to x2,y2 in object space.
func (g *Graphics) DrawLine(x1, y1, x2, y2 float64) {
	g.MoveTo(g.ObjectToScreen().Transform(x1, y1))
	g.LineTo(g.ObjectToScreen().Transform(x2, y2))
}

func (g *Graphics) DrawRect(r *Rect) {
	sr := g.ObjectToScreen().TransformRect(r)
	g.MoveTo(sr.Min.X, sr.Min.Y)
	g.LineTo(sr.Max.X, sr.Min.Y)
	g.LineTo(sr.Max.X, sr.Max.Y)
	g.LineTo(sr.Min.X, sr.Max.Y)
	g.LineTo(sr.Min.X, sr.Min.Y)
}

type PathMode int

const (
	PathModeContinue PathMode = iota
	PathModeNewShape
)

func (g *Graphics) drawEllipticalArc(center, radii *Vec2, startAngle, endAngle float64, mode PathMode) {
	const n = 4.0

	x, y := center.X, center.Y
	rx, ry := radii.X, radii.Y
	for i := 0; i < n; i++ {
		p1 := float64(i+0) / n
		p2 := float64(i+1) / n
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

		x0, y0 = g.ObjectToScreen().Transform(x0, y0)
		cx, cy = g.ObjectToScreen().Transform(cx, cy)
		x2, y2 = g.ObjectToScreen().Transform(x2, y2)

		if i == 0 {
			if mode == PathModeContinue {
				//fmt.Printf("LineTo(%v, %v)\n", x0, y0)
				g.LineTo(x0, y0)
			} else {
				//fmt.Printf("MoveTo(%v, %v)\n", x0, y0)
				g.MoveTo(x0, y0)
			}
		}
		//fmt.Printf("QuadraticTo(%v, %v)\n", x2, y2)
		g.QuadraticTo(cx, cy, x2, y2)
	}
}

var kCornerAngles = []float64{1.5 * math.Pi, math.Pi, .5 * math.Pi, 0, -.5 * math.Pi}

func (g *Graphics) DrawRoundedRect(r *Rect, radius float64) {
	maxRadius := math.Min(.5*r.Width(), .5*r.Height())
	actualRadius := math.Min(radius, maxRadius)
	radii := Vec(actualRadius, actualRadius)
	arcCenters := r.Inset(radii).Corners()
	for i := 0; i < 4; i++ {
		pathMode := PathModeContinue
		if i == 0 {
			pathMode = PathModeNewShape
		}
		g.drawEllipticalArc(&arcCenters[i], radii, kCornerAngles[i], kCornerAngles[i+1], pathMode)
	}
	g.ClosePath()
}
