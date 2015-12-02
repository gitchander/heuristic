package hexm

import (
	"math"
)

const (
	sqrtThree = 1.73205080757 // sqrt(3)

	factorX = 1.5
	factorY = sqrtThree * 0.5
)

type Vector struct {
	X, Y float64
}

func (a Vector) Add(b Vector) (c Vector) {
	c.X = a.X + b.X
	c.Y = a.Y + b.Y
	return
}

func (a Vector) Sub(b Vector) (c Vector) {
	c.X = a.X - b.X
	c.Y = a.Y - b.Y
	return
}

func (a Vector) MulScalar(t float64) (b Vector) {
	b.X = a.X * t
	b.Y = a.Y * t
	return
}

func (a Vector) DivScalar(t float64) (b Vector) {
	b.X = a.X / t
	b.Y = a.Y / t
	return
}

func VertexPolygon() []Vector {
	return []Vector{
		Vector{-1.0, 0.0},
		Vector{-0.5, sqrtThree * 0.5},
		Vector{0.5, sqrtThree * 0.5},
		Vector{1.0, 0.0},
		Vector{0.5, -sqrtThree * 0.5},
		Vector{-0.5, -sqrtThree * 0.5},
	}
}

func CoordToVector(c Coord) (v Vector, err error) {

	if err = c.getError(); err != nil {
		return
	}

	switch {

	case (c.X == 0):
		{
			v.X = -factorX * float64(c.Z)
			v.Y = factorY * float64(c.Z-2*c.Y)
		}

	case (c.Y == 0):
		{
			v.X = factorX * float64(c.X-c.Z)
			v.Y = factorY * float64(c.Z+c.X)
		}

	case (c.Z == 0):
		{
			v.X = factorX * float64(c.X)
			v.Y = factorY * float64(c.X-2*c.Y)
		}
	}

	return
}

func VectorToCoord(v Vector) (Coord, error) {
	if v.X < 0.0 {
		if v.Y < -0.5*v.X/factorY {
			return vectorToCoordYZ(v) // x = 0
		}
	} else {
		if v.Y < 0.5*v.X/factorY {
			return vectorToCoordXY(v) // z = 0
		}
	}
	return vectorToCoordZX(v) // y = 0
}

func vectorToCoordXY(v Vector) (c Coord, err error) {

	var (
		X = v.X / factorX
		Y = 0.5 * (v.X/factorX - v.Y/factorY)
	)

	var (
		fX = math.Floor(X)
		fY = math.Floor(Y)
	)

	dx, dy := hexIndexes(X-fX, Y-fY)

	c = Coord{
		X: int(fX) + dx,
		Y: int(fY) + dy,
		Z: 0,
	}

	err = c.getError()

	return
}

func vectorToCoordYZ(v Vector) (c Coord, err error) {

	var (
		Y = -0.5 * (v.X/factorX + v.Y/factorY)
		Z = -v.X / factorX
	)

	var (
		fY = math.Floor(Y)
		fZ = math.Floor(Z)
	)

	dy, dz := hexIndexes(Y-fY, Z-fZ)

	c = Coord{
		X: 0,
		Y: int(fY) + dy,
		Z: int(fZ) + dz,
	}

	err = c.getError()

	return
}

func vectorToCoordZX(v Vector) (c Coord, err error) {

	var (
		Z = 0.5 * (v.Y/factorY - v.X/factorX)
		X = 0.5 * (v.X/factorX + v.Y/factorY)
	)

	var (
		fZ = math.Floor(Z)
		fX = math.Floor(X)
	)

	dz, dx := hexIndexes(Z-fZ, X-fX)

	c = Coord{
		X: int(fX) + dx,
		Y: 0,
		Z: int(fZ) + dz,
	}

	err = c.getError()

	return
}

func vectorInCell(v Vector, c Coord) bool {

	pos, err := CoordToVector(c)
	if err != nil {
		return false
	}

	vs := VertexPolygon()
	for i, _ := range vs {
		vs[i] = vs[i].Add(pos)
	}

	return VectorInPolygon(v, vs)
}

func hexIndexes(x, y float64) (dx, dy int) {

	if y < 1.0-x {

		switch {
		case (y > x*0.5+0.5):
			dx, dy = 0, 1

		case (y < x*2.0-1.0):
			dx, dy = 1, 0

		default:
			dx, dy = 0, 0
		}

	} else {
		switch {
		case (y > x*2.0):
			dx, dy = 0, 1

		case (y < x*0.5):
			dx, dy = 1, 0

		default:
			dx, dy = 1, 1
		}
	}

	return
}

func VectorInPolygon(v Vector, poly []Vector) bool {

	ok := false
	if n := len(poly); n > 0 {
		b := poly[n-1]
		for _, a := range poly {
			if (a.X <= v.X) && (b.X > v.X) || (b.X <= v.X) && (a.X > v.X) {
				if v.Y < (v.X-a.X)*(b.Y-a.Y)/(b.X-a.X)+a.Y {
					ok = !ok
				}
			}
			b = a
		}
	}
	return ok
}
