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
	n := max - min + 1
	return min + r.Intn(n)
}

func TestTryteSetGet(t *testing.T) {

	a := new(Tryte)

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

	a := new(Tryte)
	b := new(Tryte)
	c := new(Tryte)

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

	a := new(Tryte)
	b := new(Tryte)
	c := new(Tryte)

	r := randomize()

	for i := 0; i < 1000000; i++ {

		x := intGivenRange(r, MinValue, MaxValue)
		y := intGivenRange(r, MinValue, MaxValue)

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
