package game

import (
	"fmt"
	"math"
)

func nextTumbleHeight(height float64) float64 {
	for _, h := range TumbleLevels {
		if h < height {
			return h
		}
	}
	panic(fmt.Errorf("nowhere to tumble to from %v", height))
}

type TumbleAnimation struct {
	sign         float64
	startHeight  float64
	startAngle   float64
	targetHeight float64
}

type RotationDirection int

const (
	CounterClockwise = RotationDirection(1)
	Clockwise        = RotationDirection(-1)
)

func NewTumbleAnimation(direction RotationDirection, startHeight float64) *TumbleAnimation {
	return &TumbleAnimation{
		sign:         float64(direction),
		startHeight:  startHeight,
		startAngle:   0,
		targetHeight: nextTumbleHeight(startHeight),
	}
}

func (t *TumbleAnimation) AngleFor(height float64) float64 {
	if height >= t.startHeight {
		return t.startAngle
	}

	if height <= t.targetHeight {
		t.startHeight = t.targetHeight
		t.targetHeight = nextTumbleHeight(t.startHeight)
		t.startAngle += (math.Pi / 2.0) * float64(t.sign)
	}

	totalDelta := t.targetHeight - t.startHeight
	delta := height - t.startHeight
	ratio := delta / totalDelta

	return t.startAngle + ratio*(math.Pi/2.0)*t.sign
}
