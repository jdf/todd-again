package game

import (
	"image/color"
	"sort"

	"github.com/jdf/todd-again/engine"
)

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

var ToddController Controller

const (
	JumpStateIdle = iota
	JumpStateJumping
	JumpStateLanded
)

var JumpState = JumpStateIdle

var TumbleLevels []float64

var Platforms = []Platform{
	{engine.NewRect(100, 110, 250, 130), color.RGBA{190, 190, 255, 255}},
	{engine.NewRect(300, 210, 500, 230), color.RGBA{190, 255, 190, 255}},
}

func InitPlatforms() {
	sort.Slice(Platforms, func(i, j int) bool {
		return Platforms[i].bounds.Top() < Platforms[j].bounds.Top()
	})
	for _, plat := range Platforms {
		n := len(TumbleLevels)
		if n > 0 && TumbleLevels[n-1] == plat.bounds.Top() {
			continue
		}
		TumbleLevels = append(TumbleLevels, plat.bounds.Top())
	}
	TumbleLevels = append(TumbleLevels, 0)
}

func SetControllerState(input *engine.InputState) {
	ToddController = engineControllerWrapper{input}
}
