package graph2d

type Transform interface {
	Move(x, y float32)
	Scale(x, y float32)
	Rotate(angle float32)
	Apply(v Vector) (w Vector)
	Invert() Transform
}

func NewTransform() Transform {
	return newMatrixIdentity()
}

type matrix [9]float32

/*
   0 1 2
   3 4 5
   6 7 8
*/

func newMatrixIdentity() *matrix {
	return &matrix{
		1, 0, 0,
		0, 1, 0,
		0, 0, 1,
	}
}

func newMatrixMove(x, y float32) *matrix {
	return &matrix{
		1, 0, 0,
		0, 1, 0,
		x, y, 1,
	}
}

func newMatrixScale(x, y float32) *matrix {
	return &matrix{
		x, 0, 0,
		0, y, 0,
		0, 0, 1,
	}
}

func newMatrixRotate(angle float32) *matrix {
	s, c := SinCos(angle)
	return &matrix{
		c, s, 0,
		-s, c, 0,
		0, 0, 1,
	}
}

func (m *matrix) Move(x, y float32) {

	temp := m[2]
	m[0] += x * temp
	m[1] += y * temp

	temp = m[5]
	m[3] += x * temp
	m[4] += y * temp

	temp = m[8]
	m[6] += x * temp
	m[7] += y * temp
}

func (m *matrix) Scale(x, y float32) {

	m[0] *= x
	m[1] *= y

	m[3] *= x
	m[4] *= y

	m[6] *= x
	m[7] *= y
}

/*

func (m *matrix) Move(x, y float32) {
	m.mul(newMatrixMove(x, y))
}

func (m *matrix) Scale(x, y float32) {
	m.mul(newMatrixScale(x, y))
}

*/

func (m *matrix) Rotate(angle float32) {
	m.mul(newMatrixRotate(angle))
}

func (m *matrix) Apply(v Vector) (w Vector) {
	return m.mulVectorR(v)
}

func (m *matrix) Invert() Transform {
	n := new(matrix)
	invert(m, n)
	return n
}

// matrix * vector
func (m *matrix) mulVectorL(v Vector) (w Vector) {

	var _v, _w vect3
	_v.fromVector(v)
	matrix_mul_vector(m, &_v, &_w)
	w = _w.toVector()

	return
}

// vector * matrix
func (m *matrix) mulVectorR(v Vector) (w Vector) {

	var _v, _w vect3
	_v.fromVector(v)
	vector_mul_matrix(&_v, m, &_w)
	w = _w.toVector()

	return
}

func (x *matrix) mul(y *matrix) *matrix {

	a, b, c := x[0], x[1], x[2]

	x[0] = a*y[0] + b*y[3] + c*y[6]
	x[1] = a*y[1] + b*y[4] + c*y[7]
	x[2] = a*y[2] + b*y[5] + c*y[8]

	a, b, c = x[3], x[4], x[5]

	x[3] = a*y[0] + b*y[3] + c*y[6]
	x[4] = a*y[1] + b*y[4] + c*y[7]
	x[5] = a*y[2] + b*y[5] + c*y[8]

	a, b, c = x[6], x[7], x[8]

	x[6] = a*y[0] + b*y[3] + c*y[6]
	x[7] = a*y[1] + b*y[4] + c*y[7]
	x[8] = a*y[2] + b*y[5] + c*y[8]

	return x
}

func det(m *matrix) (d float32) {

	d += m[0] * (m[4]*m[8] - m[5]*m[7])
	d -= m[1] * (m[3]*m[8] - m[5]*m[6])
	d += m[2] * (m[3]*m[7] - m[4]*m[6])

	return
}

func transpose(m *matrix) {

	m[1], m[3] = m[3], m[1]
	m[2], m[6] = m[6], m[2]
	m[5], m[7] = m[7], m[5]
}

func invert(n, m *matrix) {

	invDet := 1.0 / det(n)

	m[0] = +invDet * (n[4]*n[8] - n[5]*n[7])
	m[1] = -invDet * (n[3]*n[8] - n[5]*n[6])
	m[2] = +invDet * (n[3]*n[7] - n[4]*n[6])

	m[3] = -invDet * (n[1]*n[8] - n[2]*n[7])
	m[4] = +invDet * (n[0]*n[8] - n[2]*n[6])
	m[5] = -invDet * (n[0]*n[7] - n[1]*n[6])

	m[6] = +invDet * (n[1]*n[5] - n[2]*n[4])
	m[7] = -invDet * (n[0]*n[5] - n[2]*n[3])
	m[8] = +invDet * (n[0]*n[4] - n[1]*n[3])

	transpose(m)
}

func vector_mul_matrix(v *vect3, m *matrix, w *vect3) {

	a, b, c := v[0], v[1], v[2]

	w[0] = a*m[0] + b*m[3] + c*m[6]
	w[1] = a*m[1] + b*m[4] + c*m[7]
	w[2] = a*m[2] + b*m[5] + c*m[8]
}

func matrix_mul_vector(m *matrix, v *vect3, w *vect3) {

	a, b, c := v[0], v[1], v[2]

	w[0] = a*m[0] + b*m[1] + c*m[2]
	w[1] = a*m[3] + b*m[4] + c*m[5]
	w[2] = a*m[6] + b*m[7] + c*m[8]
}

//----------------------------------------
type vect3 [3]float32

func (v *vect3) normalize() error {

	if m := v[2]; !Equal(m, 1) {

		if Equal(m, 0) {
			return ErrorDivByZero
		}

		inv_m := 1 / m
		v[0] *= inv_m
		v[1] *= inv_m
		v[2] = 1
	}

	return nil
}

func (v *vect3) toVector() Vector {

	v.normalize()
	return Vector{v[0], v[1]}
}

func (v *vect3) fromVector(w Vector) {
	v[0] = w.X
	v[1] = w.Y
	v[2] = 1
}
