package hashxy

import (
	"testing"
)

func TestRemainder(t *testing.T) {
	const minX, maxX = -10000, +10000
	for y := 1; y < 100; y++ {
		for x := minX; x <= maxX; x++ {
			sample := modBase(x, y)
			r := mod(x, y)
			if r != sample {
				t.Errorf("wrong mod(%d, %d)", x, y)
				return
			}
		}
	}
}
