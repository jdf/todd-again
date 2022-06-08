package todd

import (
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/jakecoffman/cp"
	"golang.org/x/exp/constraints"
)

// Numeric types can be used as coordinate components.
type Numeric interface {
	constraints.Float | constraints.Integer
}

// Entity is a thing in the game that can be updated and drawn.
type Entity interface {
	// Absolute time for animations; dt for physics.
	Update(frameState *FrameState, dt float64)
	Draw(*Graphics, *Camera)
	Bounds() *Rect

	RemoveFromSpace(space *cp.Space)
}

type InputState struct {
	Left, Right, Up, Down   bool
	Spacebar, Enter, Escape bool
	MouseLeft, MouseRight   bool
	MouseX, MouseY          int
}

func GetInputState() *InputState {
	mouseX, mouseY := ebiten.CursorPosition()
	return &InputState{
		Left:       ebiten.IsKeyPressed(ebiten.KeyLeft),
		Right:      ebiten.IsKeyPressed(ebiten.KeyRight),
		Up:         ebiten.IsKeyPressed(ebiten.KeyUp),
		Down:       ebiten.IsKeyPressed(ebiten.KeyDown),
		Spacebar:   ebiten.IsKeyPressed(ebiten.KeySpace),
		Enter:      ebiten.IsKeyPressed(ebiten.KeyEnter),
		Escape:     ebiten.IsKeyPressed(ebiten.KeyEscape),
		MouseLeft:  ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft),
		MouseRight: ebiten.IsMouseButtonPressed(ebiten.MouseButtonRight),
		MouseX:     mouseX,
		MouseY:     mouseY,
	}
}

type FrameState struct {
	Camera *Camera
	Input  *InputState
	Now    time.Time
	DeltaT float64
}
