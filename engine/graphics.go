package engine

import (
	"image/color"
	"math"

	"github.com/fogleman/gg"
)

// Graphics is a graphics context that knows how to manipulate
// coordinates with a camera.
type Graphics struct {
	context gg.Context

	worldToScreen *Affine

	objectToWorld Affine
	stack         *Stack[Affine]

	cachedObjectToScreen *Affine
}

func NewGraphics(wrapped *gg.Context) *Graphics {
	return &Graphics{
		context:              *wrapped,
		objectToWorld:        *Identity(),
		stack:                &Stack[Affine]{},
		worldToScreen:        Identity(),
		cachedObjectToScreen: nil,
	}
}

func (g *Graphics) SetColor(color color.Color) {
	g.context.SetColor(color)
}

func (g *Graphics) Fill() {
	g.context.Fill()
}

func (g *Graphics) Stroke() {
	g.context.Stroke()
}

func (g *Graphics) GetScreenContext() *gg.Context {
	return &g.context
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
func (g *Graphics) DrawTextScreen(s string, x, y float64) {
	g.context.DrawString(s, x, y)
}

// DrawLine draws a line from x1,y1 to x2,y2 in object space.
func (g *Graphics) DrawLine(x1, y1, x2, y2 float64) {
	g.context.MoveTo(g.ObjectToScreen().Transform(x1, y1))
	g.context.LineTo(g.ObjectToScreen().Transform(x2, y2))
}

func (g *Graphics) DrawRect(r *Rect) {
	sr := g.ObjectToScreen().TransformRect(r)
	g.context.MoveTo(sr.Min.X, sr.Min.Y)
	g.context.LineTo(sr.Max.X, sr.Min.Y)
	g.context.LineTo(sr.Max.X, sr.Max.Y)
	g.context.LineTo(sr.Min.X, sr.Max.Y)
	g.context.LineTo(sr.Min.X, sr.Min.Y)
}

type PathMode int

const (
	PathModeContinue PathMode = iota
	PathModeNewShape
)

func (g *Graphics) drawEllipticalArc(center, radii *Vec2, startAngle, endAngle float64, mode PathMode) {
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

		x0, y0 = g.ObjectToScreen().Transform(x0, y0)
		cx, cy = g.ObjectToScreen().Transform(cx, cy)
		x2, y2 = g.ObjectToScreen().Transform(x2, y2)

		if i == 0 {
			if mode == PathModeContinue {
				//fmt.Printf("LineTo(%v, %v)\n", x0, y0)
				g.context.LineTo(x0, y0)
			} else {
				//fmt.Printf("MoveTo(%v, %v)\n", x0, y0)
				g.context.MoveTo(x0, y0)
			}
		}
		//fmt.Printf("QuadraticTo(%v, %v)\n", x2, y2)
		g.context.QuadraticTo(cx, cy, x2, y2)
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
	g.context.ClosePath()
}

func (g *Graphics) DrawEllipse(r *Rect) {
	g.drawEllipticalArc(r.Center(), r.Size().Mul(.5), 0, 2*math.Pi, PathModeNewShape)
	g.context.ClosePath()
}
