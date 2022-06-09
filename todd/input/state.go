package input

import "github.com/hajimehoshi/ebiten/v2"

type State struct {
	Left, Right, Up, Down   bool
	Spacebar, Enter, Escape bool
	MouseLeft, MouseRight   bool
	MouseX, MouseY          int
}

func GetState() *State {
	mouseX, mouseY := ebiten.CursorPosition()
	return &State{
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
