package hexm

import (
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

func randCoord(r *rand.Rand, min, max int) Coord {
	return Coord{
		X: randIntRange(r, min, max),
		Y: randIntRange(r, min, max),
		Z: randIntRange(r, min, max),
	}
}

func TestCoordToVector(t *testing.T) {
	r := newRand()
	or := Flat
	for i := 0; i < 1000000; i++ {
		var (
			a = randCoord(r, -1000, 1000).Norm()
			v = CoordToVector(or, a)
			b = VectorToCoord(or, v)
		)
		if !a.Equal(b) {
			t.Fatalf("%v != %v", a, b)
		}
	}
}
