package mindshift

import "fmt"

type Point struct {
	X, Y int
}

func (p Point) Add(q Point) Point {
	return Point{p.X + q.X, p.Y + q.Y}
}

func (p Point) Sub(q Point) Point {
	return Point{p.X - q.X, p.Y - q.Y}
}

func (p Point) Equal(q Point) bool {

	if p.X != q.X {
		return false
	}

	if p.Y != q.Y {
		return false
	}

	return true
}

func (p Point) String() string {
	return fmt.Sprintf("(%d,%d)", p.X, p.Y)
}

func errorDuplicatePoints(ps []Point) error {

	const m = "coincidence points"
	//const m= "duplicate point"

	n := len(ps)
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if ps[i].Equal(ps[j]) {
				return newErrorf("%s: p[%d] = p[%d] = %s", m, i, j, ps[i])
			}
		}
	}

	return nil
}

func checkCollision(sliceA []Point, shiftA Point, sliceB []Point, shiftB Point) bool {

	for _, localA := range sliceA {
		worldA := localA.Add(shiftA)
		for _, localB := range sliceB {
			worldB := localB.Add(shiftB)
			if worldA.Equal(worldB) {
				return true
			}
		}
	}
	return false
}
