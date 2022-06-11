package engine

import "github.com/hajimehoshi/ebiten/v2"

var WantArrows = true
var WantSpacebar = true
var WantEnter = true
var WantEscape = true
var WantMouseButtons = true
var WantMouseWheel = true

type InputState struct {
	Q                       bool
	Left, Right, Up, Down   bool
	Spacebar, Enter, Escape bool
	MouseLeft, MouseRight   bool
	MouseX, MouseY          int
	WheelX, WheelY          float64
}

func GetInputState() *InputState {
	state := &InputState{}
	state.Q = ebiten.IsKeyPressed(ebiten.KeyQ)
	if WantArrows {
		state.Left = ebiten.IsKeyPressed(ebiten.KeyLeft)
		state.Right = ebiten.IsKeyPressed(ebiten.KeyRight)
		state.Up = ebiten.IsKeyPressed(ebiten.KeyUp)
		state.Down = ebiten.IsKeyPressed(ebiten.KeyDown)
	}
	if WantSpacebar {
		state.Spacebar = ebiten.IsKeyPressed(ebiten.KeySpace)
	}
	if WantEnter {
		state.Enter = ebiten.IsKeyPressed(ebiten.KeyEnter)
	}
	if WantEscape {
		state.Escape = ebiten.IsKeyPressed(ebiten.KeyEscape)
	}
	if WantMouseButtons {
		state.MouseLeft = ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft)
		state.MouseRight = ebiten.IsMouseButtonPressed(ebiten.MouseButtonRight)
	}
	if WantMouseWheel {
		state.WheelX, state.WheelY = ebiten.Wheel()
	}
	state.MouseX, state.MouseY = ebiten.CursorPosition()
	return state
}
