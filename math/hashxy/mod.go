package hashxy

func modBase(x, y int) int {

	for x >= y {
		x -= y
	}

	for x < 0 {
		x += y
	}

	return x
}

func mod(x, y int) int {

	t := x % y

	if t < 0 {
		t += y
	}

	return t
}
