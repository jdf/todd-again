package game

import (
	"github.com/jdf/todd-again/engine"
)

func Clamp(v, min, max float32) float32 {
	switch {
	case v < min:
		return min
	case v > max:
		return max
	default:
		return v
	}
}

/**
 * @param a value at t=0
 * @param b value at t=1
 * @param t [0, 1] parameter
 */
func Lerp(a, b, t float32) float32 {
	return a + (b-a)*t
}

/**
 * @param v value at t=0
 * @param w value at t=1
 * @param t [0, 1] parameter
 */
func LerpVec(v, w *engine.Vec2, t float32) *engine.Vec2 {
	return &engine.Vec2{X: Lerp(v.X, w.X, t), Y: Lerp(v.Y, w.Y, t)}
}

type Animation interface {
	IsDone(nowSeconds float32) bool
	Value(nowSeconds float32) float32
}

func NewAnimation(startValue, endValue, nowSeconds, durationSeconds float32) Animation {
	return &timeBasedAnimation{
		startValue:      startValue,
		endValue:        endValue,
		startSeconds:    nowSeconds,
		endSeconds:      nowSeconds + durationSeconds,
		durationSeconds: durationSeconds,
	}
}

type timeBasedAnimation struct {
	startSeconds float32
	endSeconds   float32
	startValue   float32
	endValue     float32

	durationSeconds float32
}

func (a *timeBasedAnimation) IsDone(nowSeconds float32) bool {
	return nowSeconds >= a.endSeconds
}

func (a *timeBasedAnimation) Value(nowSeconds float32) float32 {
	if a.IsDone(nowSeconds) {
		return a.endValue
	}
	return Lerp(a.startValue, a.endValue, (nowSeconds-a.startSeconds)/a.durationSeconds)
}
