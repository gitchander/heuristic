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

func TestCoordToVector(t *testing.T) {
	r := newRand()
	for i := 0; i < 1000000; i++ {
		var (
			a = randomCoord(r, 0, 10000)
			v = CoordToVector(a)
			b = VectorToCoord(v)
		)
		if !a.Equal(b) {
			t.Error(errors.New("not equal"))
			return
		}
	}
}
