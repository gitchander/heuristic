package hexm

import "math"

const (
	sqrt3     = 1.73205080757 // sqrt(3)
	halfSqrt3 = 0.5 * sqrt3
)

const (
	factorX_Flat = 1.5
	factorY_Flat = halfSqrt3

	factorX_Angled = halfSqrt3
	factorY_Angled = 1.5
)

func HexVertexes(or Orientation) (vs []Vector) {
	switch or {
	case Angled:
		vs = []Vector{
			{0, 1},
			{halfSqrt3, 0.5},
			{halfSqrt3, -0.5},
			{0, -1},
			{-halfSqrt3, -0.5},
			{-halfSqrt3, 0.5},
		}
	case Flat:
		vs = []Vector{
			{-1.0, 0.0},
			{-0.5, halfSqrt3},
			{0.5, halfSqrt3},
			{1.0, 0.0},
			{0.5, -halfSqrt3},
			{-0.5, -halfSqrt3},
		}
	}
	return
}

func CoordToVector(or Orientation, c Coord) (v Vector) {
	switch or {
	case Angled:
		v = coordToVectorAngled(c)
	case Flat:
		v = coordToVectorFlat(c)
	}
	return
}

func coordToVectorAngled(c Coord) (v Vector) {

	c = c.Norm()

	switch {

	case (c.X == 0):
		{
			v.X = -factorX_Angled * float64(c.Y+c.Z)
			v.Y = factorY_Angled * float64(c.Z-c.Y)
		}

	case (c.Y == 0):
		{
			v.X = -factorX_Angled * float64(c.Z-2*c.X)
			v.Y = factorY_Angled * float64(c.Z)
		}

	case (c.Z == 0):
		{
			v.X = -factorX_Angled * float64(c.Y-2*c.X)
			v.Y = -factorY_Angled * float64(c.Y)
		}
	}

	return
}

func coordToVectorFlat(c Coord) (v Vector) {

	c = c.Norm()

	switch {

	case (c.X == 0):
		{
			v.X = -factorX_Flat * float64(c.Z)
			v.Y = factorY_Flat * float64(c.Z-2*c.Y)
		}

	case (c.Y == 0):
		{
			v.X = factorX_Flat * float64(c.X-c.Z)
			v.Y = factorY_Flat * float64(c.Z+c.X)
		}

	case (c.Z == 0):
		{
			v.X = factorX_Flat * float64(c.X)
			v.Y = factorY_Flat * float64(c.X-2*c.Y)
		}
	}

	return
}

func VectorToCoord(or Orientation, v Vector) (c Coord) {
	switch or {
	case Flat:
		c = vectorToCoordFlat(v)
	case Angled:
		// In work
	}
	return
}

func vectorToCoordFlat(v Vector) Coord {
	if v.X < 0.0 {
		if v.Y < -0.5*v.X/factorY_Flat {
			return vectorToCoordYZ_Flat(v) // x = 0
		}
	} else {
		if v.Y < 0.5*v.X/factorY_Flat {
			return vectorToCoordXY_Flat(v) // z = 0
		}
	}
	return vectorToCoordZX_Flat(v) // y = 0
}

func vectorToCoordXY_Flat(v Vector) Coord {

	var (
		X = v.X / factorX_Flat
		Y = 0.5 * (v.X/factorX_Flat - v.Y/factorY_Flat)
	)

	var (
		fX = math.Floor(X)
		fY = math.Floor(Y)
	)

	dx, dy := hexIndexes(X-fX, Y-fY)

	return Coord{
		X: int(fX) + dx,
		Y: int(fY) + dy,
		Z: 0,
	}
}

func vectorToCoordYZ_Flat(v Vector) Coord {

	var (
		Y = -0.5 * (v.X/factorX_Flat + v.Y/factorY_Flat)
		Z = -v.X / factorX_Flat
	)

	var (
		fY = math.Floor(Y)
		fZ = math.Floor(Z)
	)

	dy, dz := hexIndexes(Y-fY, Z-fZ)

	return Coord{
		X: 0,
		Y: int(fY) + dy,
		Z: int(fZ) + dz,
	}
}

func vectorToCoordZX_Flat(v Vector) Coord {

	var (
		Z = 0.5 * (v.Y/factorY_Flat - v.X/factorX_Flat)
		X = 0.5 * (v.X/factorX_Flat + v.Y/factorY_Flat)
	)

	var (
		fZ = math.Floor(Z)
		fX = math.Floor(X)
	)

	dz, dx := hexIndexes(Z-fZ, X-fX)

	return Coord{
		X: int(fX) + dx,
		Y: 0,
		Z: int(fZ) + dz,
	}
}

func vectorInCell(or Orientation, v Vector, c Coord) bool {

	pos := CoordToVector(or, c)

	vs := HexVertexes(or)
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
