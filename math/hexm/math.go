package hexm

import (
	"math"
)

const (
	sqrtThree = 1.73205080757
	twoPi     = 2 * math.Pi
)

const (
	factorX = 1.5
	factorY = sqrtThree * 0.5
)

func floor(x float32) float32 {
	return float32(math.Floor(float64(x)))
}

func round(x float32) float32 {
	return (floor(x) + 0.5)
}

func HexPolygon() []Vector {
	return []Vector{
		Vector{-1.0, 0.0},
		Vector{-0.5, sqrtThree * 0.5},
		Vector{0.5, sqrtThree * 0.5},
		Vector{1.0, 0.0},
		Vector{0.5, -sqrtThree * 0.5},
		Vector{-0.5, -sqrtThree * 0.5},
	}
}

//--------------------------------------------------------------
// dividend / divisor= quotient
// dividend % divisor= remainder

// dividend = quotient * divisor + remainder

func divmod(dividend, divisor int) (quotient, remainder int) {

	quotient = dividend / divisor
	remainder = dividend - quotient*divisor

	return
}

//--------------------------------------------------------------
func AngleNormalize(angle float32) float32 {

	for angle > twoPi {
		angle -= twoPi
	}

	for angle < 0.0 {
		angle += twoPi
	}

	//double norm_angle= fMath::Mod(Angle, Math::TwoPi);
	//if (norm_angle < 0.0f) norm_angle+= Math::TwoPi;

	// result -> [ 0 ... 2Pi ]

	return angle
}

//--------------------------------------------------------------
func CoordToPosition(c Coord) (v Vector) {

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

func vectorInCell(v Vector, c Coord) bool {

	pos := CoordToPosition(c)
	vs := HexPolygon()
	for i, _ := range vs {
		vs[i] = vs[i].Add(pos)
	}

	return VectorInPolygon(v, vs)
}

func vectorToCoordXY(v Vector) (Coord, error) {

	var (
		fX = v.X / factorX
		fY = 0.5 * (v.X/factorX - v.Y/factorY)
	)

	var (
		x0 = int(floor(fX))
		y0 = int(floor(fY))
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

func vectorToCoordYZ(v Vector) (Coord, error) {

	var (
		fY = -0.5 * (v.X/factorX + v.Y/factorY)
		fZ = -v.X / factorX
	)

	var (
		y0 = int(floor(fY))
		z0 = int(floor(fZ))
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

func vectorToCoordZX(v Vector) (Coord, error) {

	var (
		fZ = 0.5 * (v.Y/factorY - v.X/factorX)
		fX = 0.5 * (v.X/factorX + v.Y/factorY)
	)

	var (
		z0 = int(floor(fZ))
		x0 = int(floor(fX))
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

//--------------------------------------------------------------
