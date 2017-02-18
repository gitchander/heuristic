package hexm

// quo = x / y
// rem = x % y
func quoRem(x, y int) (quo, rem int) {

	quo = x / y
	rem = x - quo*y

	return
}

func min3(x, y, z int) int {
	if x < y {
		if x < z {
			return x
		} else {
			return z
		}
	} else {
		if y < z {
			return y
		} else {
			return z
		}
	}
}
