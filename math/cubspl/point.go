package cubspl

type Point struct {
	X, Y float64
}

func (this *Point) Set(x, y float64) {

	this.X = x
	this.Y = y
}

func (this *Point) Get() (x, y float64) {

	x = this.X
	y = this.Y

	return
}

func (this *Point) Equal(other *Point) bool {

	if !equal(this.X, other.X) {
		return false
	}

	if !equal(this.Y, other.Y) {
		return false
	}

	return true
}

func errorPoints(ps []Point) error {

	n := len(ps)

	if n < 2 {
		return ErrorInputSize
	}

	x := ps[0].X
	for i := 1; i < n; i++ {
		if x < ps[i].X {
			x = ps[i].X
		} else {
			return newErrorf("points error: x[%d] >= x[%d]", i-1, i)
		}
	}

	return nil
}
