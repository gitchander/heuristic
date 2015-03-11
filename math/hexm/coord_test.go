package hexm

import (
	"errors"
	"math/rand"
	"testing"
	"time"
)

func newRand() *rand.Rand {
	return rand.New(rand.NewSource(time.Now().UnixNano()))
}

func intGivenRange(r *rand.Rand, min, max int) int {

	if min > max {
		min, max = max, min
	}
	n := max - min + 1
	return min + r.Intn(n)
}

func randomCoord(r *rand.Rand, min, max int) (Coord, error) {

	var x, y, z int

	switch n := r.Intn(3); n {

	case 0:
		{
			x = 0
			y = intGivenRange(r, min, max)
			z = intGivenRange(r, min, max)
		}

	case 1:
		{
			x = intGivenRange(r, min, max)
			y = 0
			z = intGivenRange(r, min, max)
		}

	case 2:
		{
			x = intGivenRange(r, min, max)
			y = intGivenRange(r, min, max)
			z = 0
		}
	}

	return NewCoord(x, y, z)
}

func neighborAxis(c Coord, axis int, nd NeighborDir) (Coord, error) {

	switch axis {
	case 0:
		return NeighborX(c, nd)

	case 1:
		return NeighborY(c, nd)

	case 2:
		return NeighborZ(c, nd)
	}

	return nil, errors.New("Wrong neighbor axis")
}

func TestCoordNeighbors(t *testing.T) {

	var (
		a, b, c Coord
		err     error
	)

	r := newRand()
	for i := 0; i < 1000000; i++ {

		axis := r.Intn(3)

		a, err = randomCoord(r, 0, 10000)
		if err != nil {
			t.Error(err)
			return
		}

		b, err = neighborAxis(a, axis, ND_NEGATIVE)
		if err != nil {
			t.Error(err)
			return
		}

		c, err = neighborAxis(b, axis, ND_POSITIVE)
		if err != nil {
			t.Error(err)
			return
		}

		if !(a.Equal(c)) {
			t.Error(errors.New("not equal"))
			return
		}
	}
}

func TestCoordToVector(t *testing.T) {

	var (
		a, b Coord
		err  error
	)

	r := newRand()
	for i := 0; i < 1000000; i++ {

		a, err = randomCoord(r, 0, 10000)
		if err != nil {
			t.Error(err)
			return
		}

		v := CoordToVector(a)
		b, err = VectorToCoord(v)
		if err != nil {
			t.Error(err)
			return
		}

		if !(a.Equal(b)) {
			t.Error(errors.New("not equal"))
			return
		}
	}
}
