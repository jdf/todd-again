package frame

import (
	"time"

	"github.com/jdf/todd-again/todd/camera"
	"github.com/jdf/todd-again/todd/input"
)

type State struct {
	Camera *camera.Camera
	Input  *input.State
	Now    time.Time
	DeltaT float64
}