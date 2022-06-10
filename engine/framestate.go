package engine

import (
	"time"
)

type FrameState struct {
	Camera *Camera
	Input  *InputState
	Now    time.Time
	DeltaT float64
}
