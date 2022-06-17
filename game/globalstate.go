package game

import (
	"github.com/jdf/todd-again/engine"
)

var World = GlobalStateType{}

type GlobalStateType struct {
	Controller   Controller
	TumbleLevels []float64
	JumpState    JumpStateType
}

type Controller interface {
	Left() bool
	Right() bool
	Jump() bool
}

type engineControllerWrapper struct {
	*engine.InputState
}

func (wrapper engineControllerWrapper) Left() bool {
	return wrapper.InputState.Left
}

func (wrapper engineControllerWrapper) Right() bool {
	return wrapper.InputState.Right
}

func (wrapper engineControllerWrapper) Jump() bool {
	return wrapper.InputState.Spacebar
}

type JumpStateType int

const (
	JumpStateIdle JumpStateType = iota
	JumpStateJumping
	JumpStateLanded
)

func SetControllerState(input *engine.InputState) {
	World.Controller = engineControllerWrapper{input}
}

var WorldBounds = engine.NewRect(-1000, 0, 1000, 200)
