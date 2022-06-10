package engine

type UpdateState struct {
	Input        *InputState
	NowSeconds   float64
	DeltaSeconds float64
}

type Game interface {
	Resize(w, h int)
	Update(*UpdateState)
	Draw(*Graphics)
}
