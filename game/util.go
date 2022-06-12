package game

import (
	"github.com/jdf/todd-again/engine"
)

func clamp(v, min, max) {
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
func Lerp(a, b, t float64) float64 {
	return a + (b-a)*t
}

/**
 * @param v value at t=0
 * @param w value at t=1
 * @param t [0, 1] parameter
 */
func LerpVec(v, w engine.Vec2, t float64) engine.Vec2 {
	return engine.Vec2{X: Lerp(v.X, w.X, t), Y: Lerp(v.Y, w.Y, t)}
}

type Animation interface {
	IsDone(nowSeconds float64) bool
	Value(nowSeconds float64) float64
}

func NewAnimation(startValue, endValue, nowSeconds, durationSeconds float64) Animation {
	return &timeBasedAnimation{
		startValue:      startValue,
		endValue:        endValue,
		startSeconds:    nowSeconds,
		endSeconds:      nowSeconds + durationSeconds,
		durationSeconds: durationSeconds,
	}
}

type timeBasedAnimation struct {
	startSeconds float64
	endSeconds   float64
	startValue   float64
	endValue     float64

	durationSeconds float64
}

func (a *timeBasedAnimation) IsDone(nowSeconds float64) bool {
	return nowSeconds >= a.endSeconds
}

func (a *timeBasedAnimation) Value(nowSeconds float64) float64 {
	if a.IsDone(nowSeconds) {
		return a.endValue
	}
	return Lerp(a.startValue, a.endValue, (nowSeconds-a.startSeconds)/a.durationSeconds)
}
