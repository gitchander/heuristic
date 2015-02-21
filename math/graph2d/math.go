package graph2d

import (
	"errors"
	"math"
)

const (
	Epsilon = 1e-6 // for equal

	Pi    = float32(math.Pi)
	twoPi = 2 * Pi
)

var ErrorDivByZero = errors.New("DivByZero")

func Floor(x float32) float32 {
	return float32(math.Floor(float64(x)))
}

func Round(x float32) float32 {
	return (Floor(x) + 0.5)
}

func Sqrt(x float32) float32 {
	return float32(math.Sqrt(float64(x)))
}

func Abs(x float32) float32 {
	return float32(math.Abs(float64(x)))
}

func Equal(x, y float32) bool {
	return (Abs(x-y) < Epsilon)
}

func SinCos(angle float32) (sin, cos float32) {

	s, c := math.Sincos(float64(angle))

	sin = float32(s)
	cos = float32(c)

	return
}

func angleNormalize(angle float32) float32 {

	for angle > twoPi {
		angle -= twoPi
	}

	for angle < 0.0 {
		angle += twoPi
	}

	return angle
}
