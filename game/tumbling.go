package game

import (
	"context"
	"fmt"
	"math"
)

func nextTumbleHeight(height float32) float32 {
	for _, h := range TumbleLevels {
		if h < height {
			return h
		}
	}
	panic(fmt.Errorf("nowhere to tumble to from %v", height))
}

type TumbleAnimation struct {
	sign         float32
	startHeight  float32
	startAngle   float32
	targetHeight float32
}

type RotationDirection int

const (
	CounterClockwise = RotationDirection(1)
	Clockwise        = RotationDirection(-1)
)

func NewTumbleAnimation(direction RotationDirection, startHeight float32) *TumbleAnimation {
	return &TumbleAnimation{
		sign:         float32(direction),
		startHeight:  startHeight,
		startAngle:   0,
		targetHeight: nextTumbleHeight(startHeight),
	}
}

func (t *TumbleAnimation) AngleFor(height float32) float32 {
	if height >= t.startHeight {
		return t.startAngle
	}

	if height <= t.targetHeight {
		t.startHeight = t.targetHeight
		t.targetHeight = nextTumbleHeight(t.startHeight)
		t.startAngle += (math.Pi / 2.0) * float32(t.sign)
		Bus.Emit(context.Background(), ToddVerticalLevelChanged, t.targetHeight)
	}

	totalDelta := t.targetHeight - t.startHeight
	delta := height - t.startHeight
	ratio := delta / totalDelta

	return t.startAngle + ratio*(math.Pi/2.0)*t.sign
}
