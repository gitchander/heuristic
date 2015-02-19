package graph2d

import "math"

const twoPi = 2 * math.Pi

func Floor(x float32) float32 {
	return float32(math.Floor(float64(x)))
}

func Round(x float32) float32 {
	return (Floor(x) + 0.5)
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
