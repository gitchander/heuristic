package graph2d

type Matrix [4]float32

func (this *Matrix) Indent() *Matrix {

	*this = [4]float32{
		1.0, 0.0,
		0.0, 1.0,
	}

	return this
}

func (this *Matrix) Move(dx, dy float32) *Matrix {

	return this
}

func (this *Matrix) Scale(dx, dy float32) *Matrix {

	return this
}

func (this *Matrix) Rotate(dx, dy float32) *Matrix {

	return this
}

func (m Matrix) MulVector(v Vector) (w Vector) {

	w.X = m[0]*v.X + m[1]*v.Y
	w.Y = m[2]*v.X + m[3]*v.Y

	return
}

func (a Matrix) MulMatrix(b Matrix) (c Matrix) {

	c[0] = a[0]*b[0] + a[1]*b[2]
	c[1] = a[0]*b[1] + a[1]*b[3]
	c[2] = a[2]*b[0] + a[3]*b[2]
	c[3] = a[2]*b[1] + a[3]*b[3]

	return
}
