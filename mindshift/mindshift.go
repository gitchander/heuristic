package mindshift

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

type Location struct {
	Start   Point
	Finish  Point
	Current Point
}

type Primitive struct {
	Location Location
	Points   []Point
}

func collisionShiftPrimitive(ps []Primitive, index int, shift Point) bool {

	var (
		primIndex  = ps[index]
		shiftIndex = primIndex.Location.Current.Add(shift)
	)

	for i, p := range ps {
		if i != index {
			if checkCollision(primIndex.Points, shiftIndex,
				p.Points, p.Location.Current) {
				return true
			}
		}
	}

	return false
}

func checkSolved(ps []Primitive) bool {

	for _, p := range ps {
		if !p.Location.Current.Equal(p.Location.Finish) {
			return false
		}
	}

	return true
}
