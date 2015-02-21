package graph2d

type Matrix [9]float32

func Identity() Matrix {
	return Matrix{
		1, 0, 0,
		0, 1, 0,
		0, 0, 1,
	}
}

func NewIdentity() *Matrix {
	i := Identity()
	return &i
}

func (m *Matrix) Reset() *Matrix {

	*m = Identity()
	return m
}

//---------------------
//
//		| 1 0 0 |
//	m * 	| 0 1 0 |
//		| x y 1 |
//
//---------------------
func (m *Matrix) Move(x, y float32) *Matrix {

	temp := m[2]
	m[0] += x * temp
	m[1] += y * temp

	temp = m[5]
	m[3] += x * temp
	m[4] += y * temp

	temp = m[8]
	m[6] += x * temp
	m[7] += y * temp

	return m
}

//---------------------
//
//		| x 0 0 |
//	m * 	| 0 y 0 |
//		| 0 0 1 |
//
//---------------------
func (m *Matrix) Scale(x, y float32) *Matrix {

	m[0] *= x
	m[1] *= y

	m[3] *= x
	m[4] *= y

	m[6] *= x
	m[7] *= y

	return m
}

//---------------------------
//
//		|  cosA  sinA  0 |
//	m * 	| -sinA  cosA  0 |
//		|   0     0    1 |
//
//---------------------------
func (m *Matrix) Rotate(angle float32) *Matrix {

	sinA, cosA := SinCos(angle)

	r := Matrix{
		cosA, sinA, 0,
		-sinA, cosA, 0,
		0, 0, 1,
	}

	m.Mul(r)

	return m
}

func vector_mul_matrix(v []float32, m []float32, w []float32) {

	a, b, c := v[0], v[1], v[2]

	w[0] = a*m[0] + b*m[3] + c*m[6]
	w[1] = a*m[1] + b*m[4] + c*m[7]
	w[2] = a*m[2] + b*m[5] + c*m[8]
}

func matrix_mul_vector(m []float32, v []float32, w []float32) {

	a, b, c := v[0], v[1], v[2]

	w[0] = a*m[0] + b*m[1] + c*m[2]
	w[1] = a*m[3] + b*m[4] + c*m[5]
	w[2] = a*m[6] + b*m[7] + c*m[8]
}

func (m *Matrix) mulVector(v Vector) (w Vector) {

	var _v, _w vect3

	_v = v.toVector3()
	matrix_mul_vector(m[:], _v[:], _w[:])
	w.fromVector3(_w)

	return
}

func (x *Matrix) Mul(y Matrix) *Matrix {

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
