package crygo

import (
	"testing"
)

func TestSumMod32(t *testing.T) {

	r := newRand()

	for i := 0; i < 1000000; i++ {

		a := r.Uint32()
		b := r.Uint32()

		err := sumMod32Test(a, b)
		if err != nil {
			t.Error(err)
			return
		}
	}
}

func TestSumMod32M1(t *testing.T) {

	var compare = func(a, b uint32) {
		err := sumMod32M1Test(a, b)
		if err != nil {
			t.Error(err)
		}
	}

	r := newRand()

	for i := 0; i < 1000000; i++ {

		a := r.Uint32()
		b := r.Uint32()

		compare(a, b)
	}

	var rangeTest = func(min, max int64) {

		for ia := min; ia <= max; ia++ {
			for ib := min; ib <= max; ib++ {

				a := uint32(ia)
				b := uint32(ib)

				compare(a, b)
			}
		}
	}

	rangeTest(0, 1000)
	rangeTest(maxUint32-1000, maxUint32)

	compare(maxUint32, 1)
}
