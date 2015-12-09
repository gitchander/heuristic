package ternary

import (
	"math/rand"
	"testing"
	"time"
)

func randomize() *rand.Rand {
	return rand.New(rand.NewSource(time.Now().UnixNano()))
}

func intGivenRange(r *rand.Rand, min, max int) int {
	if min > max {
		min, max = max, min
	}
	return min + r.Intn(max-min)
}

func TestTryteSetGet(t *testing.T) {

	a := NewTryte()

	for i := MinValue; i <= MaxValue; i++ {

		a.SetInt(i)
		j := a.Int()

		if i != j {
			t.Errorf("%d -> (set / get) -> %d", i, j)
			return
		}
	}
}

func TestTryteMulSimple(t *testing.T) {

	a := NewTryte()
	b := NewTryte()
	c := NewTryte()

	for x := MinValue; x <= MaxValue; x++ {
		for y := MinValue; y <= MaxValue; y++ {

			mulXY := x * y

			if (mulXY >= MinValue) && (mulXY <= MaxValue) {

				a.SetInt(x)
				b.SetInt(y)
				c.Mul(a, b)

				tryteMulXY := c.Int()

				if tryteMulXY != mulXY {
					t.Errorf("(%d * %d = %d); tryte mul result: %d", x, y, mulXY, tryteMulXY)
					return
				}
			}
		}
	}
}

func TestTryteMulRand(t *testing.T) {

	a := NewTryte()
	b := NewTryte()
	c := NewTryte()

	r := randomize()

	for i := 0; i < 1000000; i++ {

		x := intGivenRange(r, MinValue, MaxValue+1)
		y := intGivenRange(r, MinValue, MaxValue+1)

		mulXY := x * y

		mulXY = a.SetInt(mulXY).Int() // normalize

		a.SetInt(x)
		b.SetInt(y)
		c.Mul(a, b)

		tryteMulXY := c.Int()

		if tryteMulXY != mulXY {
			t.Errorf("(%d * %d = %d); tryte mul result: %d", x, y, mulXY, tryteMulXY)
			return
		}
	}
}
