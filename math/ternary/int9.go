package ternary

const (
	MinInt9 = -9841 // -(3^9−1)/2
	MaxInt9 = +9841 // +(3^9−1)/2
)

// synonym: Tryte9

type Ts9 [9]int8

func Ts9Int(v int) (t Ts9) {
	setInt(t[:], v)
	return
}

func (t Ts9) String() string {
	return toStringBytes(t[:])
	//return toStringRunes(t[:])
}

func (t Ts9) Int() int {
	return getInt(t[:])
}

func (a Ts9) Add(b Ts9) (c Ts9) {
	add(a[:], b[:], c[:], 9)
	return
}

func (a Ts9) Sub(b Ts9) (c Ts9) {
	sub(a[:], b[:], c[:], 9)
	return
}

func (a Ts9) Mul(b Ts9) (c Ts9) {
	mul(a[:], b[:], c[:], 9)
	return
}

func (a Ts9) Shl(n int) Ts9 {
	shl(a[:], n)
	return a
}

func (a Ts9) Shr(n int) Ts9 {
	shr(a[:], n)
	return a
}
