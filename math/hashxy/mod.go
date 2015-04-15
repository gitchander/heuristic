package hashxy

func mod(x, y int) int {

	t := x % y

	if t < 0 {
		t += y
	}

	return t
}
