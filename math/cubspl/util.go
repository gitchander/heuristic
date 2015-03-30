package cubspl

import "math"

const epsilon = 1e-6

func equal(a, b float64) bool {
	return math.Abs(a-b) < epsilon
}

// a <= b
func lessOrEqual(a, b float64) bool {

	if a < b {
		return true
	}

	if equal(a, b) {
		return true
	}

	return false
}

// a >= b
func moreOrEqual(a, b float64) bool {

	if a > b {
		return true
	}

	if equal(a, b) {
		return true
	}

	return false
}
