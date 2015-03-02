package crygo

import (
	"errors"
	"fmt"
)

const maxUint32 = (1 << 32) - 1 // 2 ^ 32 - 1

func mod(x, y int64) int64 {

	t := x % y

	if t < 0 {
		t += y
	}

	return t
}

// (a + b) mod (2 ^ 32)
func sum_mod32_v1(a, b uint32) uint32 {

	return a + b
}

func sum_mod32_v2(a, b uint32) uint32 {

	_a := maxUint32 - a
	if b > _a {
		return (b - _a - 1)
	}
	return (a + b)
}

func sum_mod32_v3(a, b uint32) uint32 {

	A := int64(a)
	B := int64(b)
	C := int64(1 << 32) // 2^32

	return uint32(mod(A+B, C))
}

func sumMod32Test(a, b uint32) error {

	s1 := sum_mod32_v1(a, b)
	s2 := sum_mod32_v2(a, b)
	s3 := sum_mod32_v3(a, b)

	const format = "wrong (%d + %d) mod (2^32)"

	if s1 != s2 {
		return errors.New(fmt.Sprintf(format, a, b))
	}

	if s1 != s3 {
		return errors.New(fmt.Sprintf(format, a, b))
	}

	return nil
}

// (a + b) mod (2^32 - 1)*
func sum_mod32m1(a, b uint32) uint32 {

	_a := maxUint32 - a
	if b > _a {
		return (b - _a)
	}
	return (a + b)
}

func sum_mod32m1_sample(a, b uint32) uint32 {

	s := int64(a) + int64(b)
	if s < (1 << 32) {
		return uint32(s)
	}

	return uint32(s - (1 << 32) + 1)
}

/*
func sum_mod32m1_v3(a, b uint32) uint32 {

	A := int64(a)
	B := int64(b)
	C := int64((1 << 32) - 1) // (2^32 - 1)

	return uint32(mod(A+B-1, C) + 1)
}
*/

func sumMod32M1Test(a, b uint32) error {

	s1 := sum_mod32m1(a, b)
	s2 := sum_mod32m1_sample(a, b)

	if s1 != s2 {
		return errors.New(fmt.Sprintf("wrong (%d + %d) mod (2^32 - 1)*", a, b))
	}

	return nil
}
