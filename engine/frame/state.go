package frame

import (
	"time"

	"github.com/jdf/todd-again/engine/camera"
	"github.com/jdf/todd-again/engine/input"
)

type State struct {
	Camera *camera.Camera
	Input  *input.State
	Now    time.Time
	DeltaT float64
}
