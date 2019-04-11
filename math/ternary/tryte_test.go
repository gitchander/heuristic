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

func TestTs9SetGet(t *testing.T) {

	for i := MinInt9; i <= MaxInt9; i++ {

		a := Ts9Int(i)
		j := a.Int()

		if i != j {
			t.Errorf("%d -> (set / get) -> %d", i, j)
			return
		}
	}
}

func TestTryteMulSimple(t *testing.T) {
	for x := MinInt9; x <= MaxInt9; x++ {
		for y := MinInt9; y <= MaxInt9; y++ {

			mulXY := x * y

			if (mulXY >= MinInt9) && (mulXY <= MaxInt9) {

				a := Ts9Int(x)
				b := Ts9Int(y)
				c := a.Mul(b)

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

	r := randomize()

	for i := 0; i < 1000000; i++ {

		x := intGivenRange(r, MinInt9, MaxInt9+1)
		y := intGivenRange(r, MinInt9, MaxInt9+1)

		mulXY := x * y

		mulXY = Ts9Int(mulXY).Int() // normalize

		a := Ts9Int(x)
		b := Ts9Int(y)
		c := a.Mul(b)

		tryteMulXY := c.Int()

		if tryteMulXY != mulXY {
			t.Errorf("(%d * %d = %d); tryte mul result: %d", x, y, mulXY, tryteMulXY)
			return
		}
	}
}
