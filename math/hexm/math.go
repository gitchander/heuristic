package hexm

import (
	"math"
)

const (
	sqrtThree = 1.73205080757
	twoPi     = 2 * math.Pi
)

type Point struct {
	X, Y float32
}

func PointInPolygon(m Point, poly []Point) bool {

	ok := false
	if n := len(poly); n > 0 {
		b := poly[n-1]
		for _, a := range poly {
			if ((a.X <= m.X) && (b.X > m.X) || (b.X <= m.X) && (a.X > m.X)) && (m.Y < (m.X-a.X)*(b.Y-a.Y)/(b.X-a.X)+a.Y) {
				ok = !ok
			}
			b = a
		}
	}
	return ok
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
func CoordToPosition(c Coord) (p Point) {

	const (
		dX = 1.5
		dY = sqrtThree * 0.5
	)

	x, y, z := c.GetCoord()

	switch {

	case (x == 0):
		{
			p.X = -dX * float32(z)
			p.Y = dY * float32(z-2*y)
		}

	case (y == 0):
		{
			p.X = dX * float32(x-z)
			p.Y = dY * float32(z+x)
		}

	case (z == 0):
		{
			p.X = dX * float32(x)
			p.Y = dY * float32(x-2*y)
		}
	}

	return
}

func HexagonVertexes() []Point {

	p := make([]Point, 6)

	p[0] = Point{-1.0, 0.0}
	p[1] = Point{-0.5, sqrtThree * 0.5}
	p[2] = Point{0.5, sqrtThree * 0.5}
	p[3] = Point{1.0, 0.0}
	p[4] = Point{0.5, -sqrtThree * 0.5}
	p[5] = Point{-0.5, -sqrtThree * 0.5}

	return p
}

//--------------------------------------------------------------
