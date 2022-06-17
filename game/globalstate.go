package game

import (
	"context"
	"fmt"
	"time"

	"github.com/jdf/todd-again/engine"
	"github.com/mustafaturan/bus/v3"
	"github.com/mustafaturan/monoton/v3"
	"github.com/mustafaturan/monoton/v3/sequencer"
)

const ToddMovementTopic = "todd.movement"
const ToddVerticalLevelChanged = "todd.vertical.level.changed"
const CameraVerticalLevelChangedHandlerKey = "camera.vertical.level.changed"

var (
	Bus    *bus.Bus
	World  = GlobalStateType{}
	Camera *engine.Camera
)

func init() {
	m, err := monoton.New(sequencer.NewMillisecond(), uint64(1), uint64(time.Now().UnixMilli()))
	if err != nil {
		panic(err)
	}
	var idGenerator bus.Next = m.Next
	b, err := bus.NewBus(idGenerator)
	if err != nil {
		panic(err)
	}
	Bus = b
	Bus.RegisterTopics(ToddMovementTopic, ToddVerticalLevelChanged)

	Bus.RegisterHandler(CameraVerticalLevelChangedHandlerKey, bus.Handler{
		Handle: func(ctx context.Context, e bus.Event) {
			fmt.Printf("boop %v", e.Data)
		},
		Matcher: ToddVerticalLevelChanged,
	})
}

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
