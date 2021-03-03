package ternary

import (
	"math/rand"
	"testing"
	"time"
)

func newRandNow() *rand.Rand {
	return rand.New(rand.NewSource(time.Now().UnixNano()))
}

func randByInterval(r *rand.Rand, min, max int) int {
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

	r := newRandNow()

	var (
		min = MinInt9
		max = MaxInt9 + 1
	)

	for i := 0; i < 1000000; i++ {

		x := randByInterval(r, min, max)
		y := randByInterval(r, min, max)

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
