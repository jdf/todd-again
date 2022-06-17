package engine

import "github.com/hajimehoshi/ebiten/v2"

type UpdateState struct {
	Input        *InputState
	NowSeconds   float64
	DeltaSeconds float64
}

type Game interface {
	Resize(w, h int)
	Update(*UpdateState)
	Draw(*ebiten.Image, *Graphics)
}
