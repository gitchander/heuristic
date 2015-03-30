package cubspl

type splineTuple struct {
	a       Point
	b, c, d float64
}

type CubicSpline struct {
	tuples []splineTuple
	alpha  []float64
	beta   []float64
}

func NewCubicSpline(ps []Point) (*CubicSpline, error) {

	if err := errorPoints(ps); err != nil {
		return nil, err
	}

	n := len(ps)

	var (
		tuples = make([]splineTuple, n)
		alpha  = make([]float64, n-1)
		beta   = make([]float64, n-1)
	)

	for i, p := range ps {
		tuples[i].a = p
	}

	s := &CubicSpline{tuples, alpha, beta}
	s.Recalculate()

	return s, nil
}

func (this *CubicSpline) Recalculate() {

	var (
		tuples = this.tuples
		alpha  = this.alpha
		beta   = this.beta
	)

	n := len(tuples)

	tuples[0].c = 0
	tuples[n-1].c = 0

	alpha[0] = 0
	beta[0] = 0

	for i := 1; i < (n - 1); i++ {

		var (
			prev = tuples[i-1].a
			curr = tuples[i].a
			next = tuples[i+1].a
		)

		var (
			h_i  = curr.X - prev.X
			h_i1 = next.X - curr.X
		)

		var (
			A = h_i
			C = 2.0 * (h_i + h_i1)
			B = h_i1
		)

		var (
			F = 6.0 * ((next.Y-curr.Y)/h_i1 - (curr.Y-prev.Y)/h_i)
			z = (A*alpha[i-1] + C)
		)

		alpha[i] = -B / z
		beta[i] = (F - A*beta[i-1]) / z
	}

	for i := n - 2; i > 0; i-- {

		tuples[i].c = alpha[i]*tuples[i+1].c + beta[i]
	}

	for i := n - 1; i > 0; i-- {

		var (
			prev = tuples[i-1].a
			curr = tuples[i].a
		)

		var (
			h_i = curr.X - prev.X
		)

		tuples[i].d = (tuples[i].c - tuples[i-1].c) / h_i
		tuples[i].b = h_i*(2.0*tuples[i].c+tuples[i-1].c)/6.0 + (curr.Y-prev.Y)/h_i
	}
}

func (this *CubicSpline) Interpolate(x float64) (y float64) {

	var (
		n      = len(this.tuples)
		tuple  *splineTuple
		tuples = this.tuples
	)

	if x < tuples[0].a.X {

		tuple = &(tuples[1])

	} else if x > tuples[n-1].a.X {

		tuple = &(tuples[n-1])

	} else {

		i, j := 0, n-1
		for i+1 < j {
			k := i + ((j - i) >> 1)
			if x <= tuples[k].a.X { // operator - '<='
				j = k
			} else {
				i = k
			}
		}
		tuple = &(tuples[j])
	}

	dx := (x - tuple.a.X)

	y = tuple.a.Y + (tuple.b+(tuple.c*0.5+tuple.d*dx/6.0)*dx)*dx

	return
}

func (this *CubicSpline) SetPointByIndex(index int, p Point) error {

	n := len(this.tuples)

	if (index < 0) || (index >= n) {
		return ErrorIndexOutOfRange
	}

	if index > 0 {
		prev := this.tuples[index-1].a
		if lessOrEqual(p.X, prev.X) { // a <= b
			return newError("X[i] <= X[i-1]")
		}
	}

	if index < n-1 {
		next := this.tuples[index+1].a
		if moreOrEqual(p.X, next.X) { // a >= b
			return newError("X[i] >= X[i+1]")
		}
	}

	tuple := &(this.tuples[index])

	tuple.a = p

	return nil
}

func (this *CubicSpline) GetPointByIndex(index int) (Point, error) {

	p := Point{}

	n := len(this.tuples)

	if (index < 0) || (index >= n) {
		return p, ErrorIndexOutOfRange
	}

	tuple := &(this.tuples[index])

	return tuple.a, nil
}
