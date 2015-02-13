package ternary

import (
	"bytes"
)

const (
	count     = 9
	base      = 3
	min_digit = -(base - 1) / 2
	max_digit = (base - 1) / 2
)

const (
	MinValue = -9841 // -(3^9−1)/2
	MaxValue = +9841 // +(3^9−1)/2
)

//var digs = []rune{'-', '0', '+'}

var digs = []rune{'N', '0', 'P'}

/*
	'\u25cf' - ●		-1
	'\u25d2' - ◒		0
	'\u25cb' - ○		+1

var digs = []rune{'\u25cf', '\u25d2', '\u25cb'}
*/

type trit int8

/*
type ITryte interface {
	SetInt(int) ITryte
	GetInt() int
	Add(x, y ITryte) ITryte
	Sub(x, y ITryte) ITryte
	Mul(x, y ITryte) ITryte
	QuoRem(x, y, rem ITryte) (q ITryte, r ITryte)
	String() string
	Sign() int
}
*/

type Tryte [count]trit

func NewTryte(x int) *Tryte {
	return new(Tryte).SetInt(x)
}

func (this *Tryte) сlone() *Tryte {

	q := new(Tryte)

	copy(q[:], this[:])

	/*
		for i := 0; i < count; i++ {
			q[i] = this[i]
		}
	*/

	return q
}

func (t *Tryte) String() string {

	var buffer bytes.Buffer

	index := t.indexMostSignificantDigit()

	//buffer.WriteString("0t")

	for i := index; i >= 0; i-- {
		buffer.WriteRune(digs[t[i]+1])
	}

	return buffer.String()
}

func divmod(a, b int) (quotient, remainder int) {

	quotient = a / b
	remainder = a - quotient*b

	return
}

func (t *Tryte) SetZero() {
	for i := 0; i < count; i++ {
		t[i] = 0
	}
}

func (t *Tryte) SetInt(v int) *Tryte {

	var q, r int

	switch {
	case v > 0:

		for i := 0; i < count; i++ {
			q, r = divmod(v+max_digit, base)
			t[i] = trit(r - max_digit)
			v = q
		}

	case v < 0:

		for i := 0; i < count; i++ {
			q, r = divmod(v-max_digit, base)
			t[i] = trit(r + max_digit)
			v = q
		}

	default:

		t.SetZero()
	}

	return t
}

func (t *Tryte) Int() int {

	var v int = 0
	var power int = 1

	for i := 0; i < count; i++ {
		v += power * int(t[i])
		power *= base
	}

	return v
}

func (t *Tryte) indexMostSignificantDigit() (index int) {

	for i := 1; i < count; i++ {
		if t[i] != 0 {
			index = i
		}
	}

	return
}

// Sign returns:
// 	-1 	if t < 0
// 	 0	if t == 0
// 	+1 	if t > 0
func (t *Tryte) Sign() int {

	index := t.indexMostSignificantDigit()

	switch {
	case t[index] > 0:
		return +1

	case t[index] < 0:
		return -1
	}

	return 0
}

func (c *Tryte) Add(a, b *Tryte) *Tryte {

	var carry int = 0
	for i := 0; i < count; i++ {

		temp := int(a[i]+b[i]) + carry

		if temp > max_digit {
			temp -= base
			carry = max_digit
		} else if temp < min_digit {
			temp += base
			carry = min_digit
		} else {
			carry = 0
		}

		c[i] = trit(temp)
	}
	return c
}

func (c *Tryte) Sub(a, b *Tryte) *Tryte {

	var carry int = 0
	for i := 0; i < count; i++ {

		temp := int(a[i]-b[i]) + carry

		if temp > max_digit {
			temp -= base
			carry = max_digit
		} else if temp < min_digit {
			temp += base
			carry = min_digit
		} else {
			carry = 0
		}

		c[i] = trit(temp)
	}
	return c
}

func (c *Tryte) Mul(a, b *Tryte) *Tryte {

	if c == a {
		a = a.сlone()
	}

	if c == b {
		b = b.сlone()
	}

	c.SetZero()
	for ia := 0; ia < count; ia++ {
		carry := 0
		for ib := 0; ib < (count - ia); ib++ {
			temp := int(a[ia]*b[ib]+c[ia+ib]) + carry

			if temp > max_digit {
				temp -= base
				carry = max_digit
			} else if temp < min_digit {
				temp += base
				carry = min_digit
			} else {
				carry = 0
			}

			c[ia+ib] = trit(temp)
		}
	}

	return c
}
