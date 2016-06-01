package ternary

import "bytes"

const (
	tritCount = 9
	base      = 3
)

const (
	minDigit = -(base - 1) / 2
	maxDigit = +(base - 1) / 2
)

const (
	MinValue = -9841 // -(3^9−1)/2
	MaxValue = +9841 // +(3^9−1)/2
)

const prefix = "0t"

var digitsSamples = [][]rune{
	[]rune{'-', '0', '+'},
	[]rune{'N', '0', 'P'},
	[]rune{'N', 'Z', 'P'},
	[]rune{'\u25cf', '\u25d2', '\u25cb'}, // []rune{ '●', '◒', '○' }
}

var digs = digitsSamples[2]

type Tryte struct {
	trits []int8
}

func NewTryte() *Tryte {
	return &Tryte{make([]int8, tritCount)}
}

func NewTryteInt(x int) *Tryte {
	return NewTryte().SetInt(x)
}

func (t *Tryte) Clone() *Tryte {
	c := NewTryte()
	copy(c.trits, t.trits)
	return c
}

func (t *Tryte) String() string {

	var buffer bytes.Buffer

	index := t.indexDigit()

	buffer.WriteString(prefix)

	for i := index; i >= 0; i-- {
		buffer.WriteRune(digs[t.trits[i]+maxDigit])
	}

	return buffer.String()
}

func (t *Tryte) SetInt(v int) *Tryte {

	ts := t.trits

	if v > 0 {
		for i := range ts {
			quo, rem := quoRem(v+maxDigit, base)
			ts[i] = int8(rem - maxDigit)
			v = quo
		}
	} else if v < 0 {
		for i := range ts {
			quo, rem := quoRem(v-maxDigit, base)
			ts[i] = int8(rem + maxDigit)
			v = quo
		}
	} else {
		for i := range ts {
			ts[i] = 0
		}
	}

	return t
}

func (t *Tryte) Int() int {

	var (
		v     = 0
		power = 1
	)

	for _, trit := range t.trits {
		v += power * int(trit)
		power *= base
	}

	return v
}

func (t *Tryte) indexDigit() (index int) {

	for i, trit := range t.trits {
		if trit != 0 {
			index = i
		}
	}

	return
}

// Sign return:
// 	-1 	if t < 0
// 	 0	if t == 0
// 	+1 	if t > 0
func (t *Tryte) Sign() int {

	index := t.indexDigit()

	switch {
	case t.trits[index] > 0:
		return +1

	case t.trits[index] < 0:
		return -1
	}

	return 0
}

func (c *Tryte) Add(a, b *Tryte) *Tryte {

	var (
		a_ts = a.trits
		b_ts = b.trits
		c_ts = c.trits
	)

	var carry int = 0

	for i := 0; i < tritCount; i++ {

		temp := int(a_ts[i]) + int(b_ts[i]) + carry

		if temp > maxDigit {
			temp -= base
			carry = maxDigit
		} else if temp < minDigit {
			temp += base
			carry = minDigit
		} else {
			carry = 0
		}

		c_ts[i] = int8(temp)
	}

	return c
}

func (c *Tryte) Sub(a, b *Tryte) *Tryte {

	var (
		a_ts = a.trits
		b_ts = b.trits
		c_ts = c.trits
	)

	var carry int = 0

	for i := 0; i < tritCount; i++ {

		temp := int(a_ts[i]) - int(b_ts[i]) + carry

		if temp > maxDigit {
			temp -= base
			carry = maxDigit
		} else if temp < minDigit {
			temp += base
			carry = minDigit
		} else {
			carry = 0
		}

		c_ts[i] = int8(temp)
	}

	return c
}

func (c *Tryte) Mul(a, b *Tryte) *Tryte {

	if c == a {
		a = a.Clone()
	}

	if c == b {
		b = b.Clone()
	}

	c.SetInt(0)

	var (
		a_ts = a.trits
		b_ts = b.trits
		c_ts = c.trits
	)

	for ia := 0; ia < tritCount; ia++ {
		carry := 0
		for ib := 0; ib < (tritCount - ia); ib++ {
			temp := int(a_ts[ia])*int(b_ts[ib]) + int(c_ts[ia+ib]) + carry

			if temp > maxDigit {
				temp -= base
				carry = maxDigit
			} else if temp < minDigit {
				temp += base
				carry = minDigit
			} else {
				carry = 0
			}

			c_ts[ia+ib] = int8(temp)
		}
	}

	return c
}

func (quo *Tryte) QuoRem(x, y, rem *Tryte) *Tryte {

	// In Work!

	return quo
}

func quoRem(a, b int) (quo, rem int) {

	quo = a / b
	rem = a - quo*b

	return
}
