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

func randIntRange(r *rand.Rand, min, max int) int {
	if min > max {
		min, max = max, min
	}
	return min + r.Intn(max-min)
}

func randomCoord(r *rand.Rand, min, max int) Coord {

	var x, y, z int

	switch n := r.Intn(3); n {

	case 0:
		{
			x = 0
			y = randIntRange(r, min, max)
			z = randIntRange(r, min, max)
		}

	case 1:
		{
			x = randIntRange(r, min, max)
			y = 0
			z = randIntRange(r, min, max)
		}

	case 2:
		{
			x = randIntRange(r, min, max)
			y = randIntRange(r, min, max)
			z = 0
		}
	}

	return Coord{x, y, z}
}

func neighborAxis(c Coord, axis int, nd NeighborDir) (n Coord, err error) {

	switch axis {
	case 0:
		n, err = NeighborX(c, nd)

	case 1:
		n, err = NeighborY(c, nd)

	case 2:
		n, err = NeighborZ(c, nd)

	default:
		err = errors.New("Wrong neighbor axis")
	}

	return
}

func TestCoordNeighbors(t *testing.T) {

	var (
		a, b, c Coord
		err     error
	)

	r := newRand()
	for i := 0; i < 1000000; i++ {

		axis := r.Intn(3)

		a = randomCoord(r, 0, 10000)

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

		if !a.Equal(c) {
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

		a = randomCoord(r, 0, 10000)

		v, _ := CoordToVector(a)
		b, err = VectorToCoord(v)
		if err != nil {
			t.Error(err)
			return
		}

		if !a.Equal(b) {
			t.Error(errors.New("not equal"))
			return
		}
	}
}
