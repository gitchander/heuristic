package hexm

import (
	"github.com/gitchander/heuristic/math/graph2d"
)

const sqrtThree = 1.73205080757 // sqrt(3)

const (
	factorX = 1.5
	factorY = sqrtThree * 0.5
)

// quo = x / y
// rem = x % y
func quoRem(x, y int) (quo, rem int) {

	quo = x / y
	rem = x - quo*y

	return
}

func HexPolygon() []graph2d.Vector {
	return []graph2d.Vector{
		graph2d.Vector{-1.0, 0.0},
		graph2d.Vector{-0.5, sqrtThree * 0.5},
		graph2d.Vector{0.5, sqrtThree * 0.5},
		graph2d.Vector{1.0, 0.0},
		graph2d.Vector{0.5, -sqrtThree * 0.5},
		graph2d.Vector{-0.5, -sqrtThree * 0.5},
	}
}

func CoordToVector(c Coord) (v graph2d.Vector) {

	x, y, z := c.GetCoord()

	switch {

	case (x == 0):
		{
			v.X = -factorX * float32(z)
			v.Y = factorY * float32(z-2*y)
		}

	case (y == 0):
		{
			v.X = factorX * float32(x-z)
			v.Y = factorY * float32(z+x)
		}

	case (z == 0):
		{
			v.X = factorX * float32(x)
			v.Y = factorY * float32(x-2*y)
		}
	}

	return
}

func VectorToCoord(v graph2d.Vector) (Coord, error) {

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

func vectorInCell(v graph2d.Vector, c Coord) bool {

	pos := CoordToVector(c)
	vs := HexPolygon()
	for i, _ := range vs {
		vs[i] = vs[i].Add(pos)
	}

	return graph2d.VectorInPolygon(v, vs)
}

func hexIndexes(x, y float32) (dx, dy int) {

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

func vectorToCoordXY(v graph2d.Vector) (Coord, error) {

	var (
		X = v.X / factorX
		Y = 0.5 * (v.X/factorX - v.Y/factorY)
	)

	var (
		fX = graph2d.Floor(X)
		fY = graph2d.Floor(Y)
	)

	dx, dy := hexIndexes(X-fX, Y-fY)
	c, err := NewCoord(int(fX)+dx, int(fY)+dy, 0)
	if err != nil {
		return nil, err
	}

	return c, nil
}

func vectorToCoordYZ(v graph2d.Vector) (Coord, error) {

	var (
		Y = -0.5 * (v.X/factorX + v.Y/factorY)
		Z = -v.X / factorX
	)

	var (
		fY = graph2d.Floor(Y)
		fZ = graph2d.Floor(Z)
	)

	dy, dz := hexIndexes(Y-fY, Z-fZ)
	c, err := NewCoord(0, int(fY)+dy, int(fZ)+dz)
	if err != nil {
		return nil, err
	}

	return c, nil
}

func vectorToCoordZX(v graph2d.Vector) (Coord, error) {

	var (
		Z = 0.5 * (v.Y/factorY - v.X/factorX)
		X = 0.5 * (v.X/factorX + v.Y/factorY)
	)

	var (
		fZ = graph2d.Floor(Z)
		fX = graph2d.Floor(X)
	)

	dz, dx := hexIndexes(Z-fZ, X-fX)
	c, err := NewCoord(int(fX)+dx, 0, int(fZ)+dz)
	if err != nil {
		return nil, err
	}

	return c, nil
}

/*
func vectorToCoordXY(v graph2d.Vector) (Coord, error) {

	var (
		fX = v.X / factorX
		fY = 0.5 * (v.X/factorX - v.Y/factorY)
	)

	var (
		x0 = int(graph2d.Floor(fX))
		y0 = int(graph2d.Floor(fY))
	)

	for dx := 0; dx < 2; dx++ {
		for dy := 0; dy < 2; dy++ {

			c, err := NewCoord(x0+dx, y0+dy, 0)
			if err != nil {
				return nil, err
			}

			if vectorInCell(v, c) {
				return c, nil
			}
		}
	}

	return nil, ErrorVectorToCoord
}

func vectorToCoordYZ(v graph2d.Vector) (Coord, error) {

	var (
		fY = -0.5 * (v.X/factorX + v.Y/factorY)
		fZ = -v.X / factorX
	)

	var (
		y0 = int(graph2d.Floor(fY))
		z0 = int(graph2d.Floor(fZ))
	)

	for dy := 0; dy < 2; dy++ {
		for dz := 0; dz < 2; dz++ {

			c, err := NewCoord(0, y0+dy, z0+dz)
			if err != nil {
				return nil, err
			}

			if vectorInCell(v, c) {
				return c, nil
			}
		}
	}

	return nil, ErrorVectorToCoord
}

func vectorToCoordZX(v graph2d.Vector) (Coord, error) {

	var (
		fZ = 0.5 * (v.Y/factorY - v.X/factorX)
		fX = 0.5 * (v.X/factorX + v.Y/factorY)
	)

	var (
		z0 = int(graph2d.Floor(fZ))
		x0 = int(graph2d.Floor(fX))
	)

	for dz := 0; dz < 2; dz++ {
		for dx := 0; dx < 2; dx++ {

			c, err := NewCoord(x0+dx, 0, z0+dz)
			if err != nil {
				return nil, err
			}

			if vectorInCell(v, c) {
				return c, nil
			}
		}
	}

	return nil, ErrorVectorToCoord
}
*/
