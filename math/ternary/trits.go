package ternary

//import "bytes"

const (
	//	tritCount = 9
	base = 3
)

const (
	minDigit = -(base - 1) / 2
	maxDigit = +(base - 1) / 2
)

const prefix = "0t"

//var digitsSamples = [][]rune{
//	0: []rune{'-', '0', '+'},
//	1: []rune{'N', '0', 'P'},
//	2: []rune{'N', 'Z', 'P'},
//	3: []rune{'\u25cf', '\u25d2', '\u25cb'}, // []rune{ '●', '◒', '○' }
//}

//var digs = digitsSamples[2]

//func (t *Tryte) indexDigit() (index int) {

//	for i, trit := range t.trits {
//		if trit != 0 {
//			index = i
//		}
//	}

//	return
//}

//// Sign return:
//// 	-1 	if t < 0
//// 	 0	if t == 0
//// 	+1 	if t > 0
//func (t *Tryte) Sign() int {

//	index := t.indexDigit()

//	switch {
//	case t.trits[index] > 0:
//		return +1

//	case t.trits[index] < 0:
//		return -1
//	}

//	return 0
//}

//func (c *Tryte) Add(a, b *Tryte) *Tryte {

//	var (
//		a_ts = a.trits
//		b_ts = b.trits
//		c_ts = c.trits
//	)

//	var carry int = 0

//	for i := 0; i < tritCount; i++ {

//		temp := int(a_ts[i]) + int(b_ts[i]) + carry

//		if temp > maxDigit {
//			temp -= base
//			carry = maxDigit
//		} else if temp < minDigit {
//			temp += base
//			carry = minDigit
//		} else {
//			carry = 0
//		}

//		c_ts[i] = int8(temp)
//	}

//	return c
//}

//func (c *Tryte) Sub(a, b *Tryte) *Tryte {

//	var (
//		a_ts = a.trits
//		b_ts = b.trits
//		c_ts = c.trits
//	)

//	var carry int = 0

//	for i := 0; i < tritCount; i++ {

//		temp := int(a_ts[i]) - int(b_ts[i]) + carry

//		if temp > maxDigit {
//			temp -= base
//			carry = maxDigit
//		} else if temp < minDigit {
//			temp += base
//			carry = minDigit
//		} else {
//			carry = 0
//		}

//		c_ts[i] = int8(temp)
//	}

//	return c
//}

//func (c *Tryte) Mul(a, b *Tryte) *Tryte {

//	if c == a {
//		a = a.Clone()
//	}

//	if c == b {
//		b = b.Clone()
//	}

//	c.SetInt(0)

//	var (
//		a_ts = a.trits
//		b_ts = b.trits
//		c_ts = c.trits
//	)

//	for ia := 0; ia < tritCount; ia++ {
//		carry := 0
//		for ib := 0; ib < (tritCount - ia); ib++ {
//			temp := int(a_ts[ia])*int(b_ts[ib]) + int(c_ts[ia+ib]) + carry

//			if temp > maxDigit {
//				temp -= base
//				carry = maxDigit
//			} else if temp < minDigit {
//				temp += base
//				carry = minDigit
//			} else {
//				carry = 0
//			}

//			c_ts[ia+ib] = int8(temp)
//		}
//	}

//	return c
//}

func quoRem(a, b int) (quo, rem int) {
	quo = a / b
	rem = a - quo*b
	return
}

func toStringBytes(trits []int8) string {

	var digs = []byte{'A', 'B', 'C'} // A < B < C
	//var digs = []byte{'N', 'Z', 'P'} // N < Z < P
	//var digs = []byte{'-', '0', '+'} // '-' < 0 < '+'

	plen := len(prefix)
	n := len(trits)
	bs := make([]byte, plen+n)
	j := len(bs) - 1
	jN := j
	for _, trit := range trits {
		var b byte

		if trit < 0 {
			b = digs[0]
			jN = j
		} else if trit == 0 {
			b = digs[1]
		} else if trit > 0 {
			b = digs[2]
			jN = j
		}

		bs[j] = b
		j--
	}

	bs = bs[jN-plen:]
	copy(bs, prefix)

	return string(bs)
}

func toStringRunes(trits []int8) string {

	//var digs = []rune{'A', 'B', 'C'} // A < B < C
	//var digs = []byte{'N', 'Z', 'P'} // N < Z < P
	//var digs = []byte{'-', '0', '+'} // '-' < 0 < '+'
	var digs = []rune{'●', '◒', '○'}

	rsPrefix := []rune(prefix)

	plen := len(rsPrefix)
	n := len(trits)
	rs := make([]rune, plen+n)
	j := len(rs) - 1
	jN := j
	for i := 0; i < n; i++ {
		trit := trits[i]
		var r rune

		if trit < 0 {
			r = digs[0]
			jN = j
		} else if trit == 0 {
			r = digs[1]
		} else if trit > 0 {
			r = digs[2]
			jN = j
		}

		rs[j] = r
		j--
	}

	rs = rs[jN-plen:]
	copy(rs, rsPrefix)

	return string(rs)
}

func getInt(trits []int8) int {
	var (
		v     = 0
		power = 1
	)
	for _, trit := range trits {
		v += power * int(trit)
		power *= base
	}
	return v
}

func setInt(trits []int8, v int) {
	ts := trits
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
}

func add(a, b, c []int8, n int) {
	var carry int8
	for i := 0; i < n; i++ {
		temp := a[i] + b[i] + carry
		if temp > maxDigit {
			temp -= base
			carry = maxDigit
		} else if temp < minDigit {
			temp += base
			carry = minDigit
		} else {
			carry = 0
		}
		c[i] = temp
	}
}

func sub(a, b, c []int8, n int) {
	var carry int8
	for i := 0; i < n; i++ {
		temp := a[i] - b[i] + carry
		if temp > maxDigit {
			temp -= base
			carry = maxDigit
		} else if temp < minDigit {
			temp += base
			carry = minDigit
		} else {
			carry = 0
		}
		c[i] = temp
	}
}

func mul(a, b, c []int8, n int) {

	//	if c == a {
	//		a = a.Clone()
	//	}

	//	if c == b {
	//		b = b.Clone()
	//	}

	//	c.SetInt(0)

	//	var (
	//		a_ts = a.trits
	//		b_ts = b.trits
	//		c_ts = c.trits
	//	)

	for ia := 0; ia < n; ia++ {
		var carry int8
		for ib := 0; ib < (n - ia); ib++ {
			temp := a[ia]*b[ib] + c[ia+ib] + carry

			if temp > maxDigit {
				temp -= base
				carry = maxDigit
			} else if temp < minDigit {
				temp += base
				carry = minDigit
			} else {
				carry = 0
			}

			c[ia+ib] = temp
		}
	}
}

func setTrits(trits []int8, v int8) {
	for i := range trits {
		trits[i] = v
	}
}

// T << offset
func shl(trits []int8, offset int) {
	if offset < 0 {
		panic("negative offset")
	}
	n := len(trits)
	i := n - 1
	if offset < n {
		for ; i >= offset; i-- {
			trits[i] = trits[i-offset]
		}
	}
	for ; i >= 0; i-- {
		trits[i] = 0
	}
}

// T >> offset
func shr(trits []int8, offset int) {
	if offset < 0 {
		panic("negative offset")
	}
	n := len(trits)
	i := 0
	if offset < n {
		for ; i < n-offset; i++ {
			trits[i] = trits[i+offset]
		}
	}
	for ; i < n; i++ {
		trits[i] = 0
	}
}
