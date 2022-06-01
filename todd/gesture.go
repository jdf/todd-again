package todd

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

// InputModality represents a input device to provide strokes.
type InputModality interface {
	Position() (int, int)
	IsUp() bool
}

// MouseGesture is a mouse gesture.
type MouseGesture struct{}

func (m *MouseGesture) Position() (int, int) {
	return ebiten.CursorPosition()
}

func (m *MouseGesture) IsJustReleased() bool {
	return inpututil.IsMouseButtonJustReleased(ebiten.MouseButtonLeft)
}

// TouchGesture is a touch gesture.
type TouchGesture struct {
	ID ebiten.TouchID
}

func (t *TouchGesture) Position() (int, int) {
	return ebiten.TouchPosition(t.ID)
}

func (t *TouchGesture) IsJustReleased() bool {
	return inpututil.IsTouchJustReleased(t.ID)
}

// Gesture represents an in-progress gesture.
type Gesture struct {
	source InputModality

	start   *Vec2
	current *Vec2

	up bool
}

func NewGesture(source InputModality) *Gesture {
	cx, cy := source.Position()
	return &Gesture{
		source:  source,
		start:   Vec(cx, cy),
		current: Vec(cx, cy),
	}
}

func (s *Gesture) Update() {
	if s.up {
		return
	}
	if s.source.IsUp() {
		s.up = true
		return
	}
	s.current = Vec(s.source.Position())
}

func (s *Gesture) IsReleased() bool {
	return s.up
}

func (s *Gesture) Position() *Vec2 {
	return s.current
}

func (s *Gesture) Delta() *Vec2 {
	return s.current.Minus(s.start)
}
