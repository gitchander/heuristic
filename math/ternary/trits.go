package ternary

const (
	base = 3
)

const (
	minDigit = -(base - 1) / 2
	maxDigit = +(base - 1) / 2
)

const prefix = "0t"

var digsBytes = [][3]byte{
	[3]byte{'-', '0', '+'}, // '-' < '0' < '+'
	[3]byte{'A', 'B', 'C'}, // 'A' < 'B' < 'C'
	[3]byte{'A', 'B', 'C'}, // 'A' < 'B' < 'C'
	[3]byte{'N', 'Z', 'P'}, // 'N' < 'Z' < 'P'
	[3]byte{'N', '0', '1'}, // 'N' < '0' < '1'
	[3]byte{'T', '0', '1'}, // 'T' < '0' < '1'
}

var digsRunes = [][3]rune{
	[3]rune{'●', '◒', '○'}, // '●' < '◒' < '○'
}

func tritsToString(trits []int8) string {
	return toStringBytes(trits)
	//return toStringRunes(trits)
}

func toStringBytes(trits []int8) string {

	var digs = digsBytes[5]

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

	var digs = digsRunes[0]

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

func quoRem(a, b int) (quo, rem int) {
	quo = a / b
	rem = a % b
	return
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
