package game

import (
	"context"
	"math/rand"
	"time"

	"github.com/jdf/todd-again/engine"
	"github.com/jdf/todd-again/game/tuning"
	"github.com/mustafaturan/bus/v3"
	"github.com/mustafaturan/monoton/v3"
	"github.com/mustafaturan/monoton/v3/sequencer"
)

const ToddMovementTopic = "todd.movement"
const ToddVerticalLevelChanged = "todd.vertical.level.changed"
const CameraVerticalLevelChangedHandlerKey = "camera.vertical.level.changed"

var (
	Bus         *bus.Bus
	Controller  engine.Controller
	Camera      *engine.Camera
	Todd        *Dude
	WorldBounds = engine.NewRect(-1000, 0, 1000, 200)
	JumpState   JumpStateType
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
			AnimateCameraVertical()
		},
		Matcher: ToddVerticalLevelChanged,
	})

	Controller = engine.EbitenController{}
	{
		Todd = &Dude{
			sideLength: tuning.Instance.Todd.GetSideLength(),
			fillColor:  tuning.RGBA(tuning.Instance.Todd.Color),
			pos:        engine.Vec2{X: 0, Y: 0},
			rnd:        rand.New(rand.NewSource(time.Now().UnixNano())),
		}
	}
}

type JumpStateType int

const (
	JumpStateIdle JumpStateType = iota
	JumpStateJumping
	JumpStateLanded
)
