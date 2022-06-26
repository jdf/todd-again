package engine

import "github.com/hajimehoshi/ebiten/v2"

type UpdateState struct {
	Input        *InputState
	NowSeconds   float64
	DeltaSeconds float64
}

type GameModule interface {
	Resize(w, h int)
	// Called potentially multiple times per frame.
	UpdatePhysics(*UpdateState)
	// Called once per frame.
	UpdateInput(*UpdateState)
	// Called once per frame.
	Draw(*ebiten.Image, *Graphics)
}
