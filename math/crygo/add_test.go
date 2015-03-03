package crygo

import (
	"testing"
)

func TestAddMod32(t *testing.T) {

	r := newRand()

	for i := 0; i < 1000000; i++ {

		a := r.Uint32()
		b := r.Uint32()

		err := addMod32Test(a, b)
		if err != nil {
			t.Error(err)
			return
		}
	}
}

func TestAddMod32M1(t *testing.T) {

	r := newRand()

	for i := 0; i < 1000000; i++ {

		a := r.Uint32()
		b := r.Uint32()

		err := addMod32M1Test(a, b)
		if err != nil {
			t.Error(err)
		}
	}

	var rangeTest = func(min, max int64) {

		for ia := min; ia <= max; ia++ {
			for ib := min; ib <= max; ib++ {

				a := uint32(ia)
				b := uint32(ib)

				err := addMod32M1Test(a, b)
				if err != nil {
					t.Error(err)
				}
			}
		}
	}

	rangeTest(0, 1000)
	rangeTest(maxUint32-1000, maxUint32)
}
