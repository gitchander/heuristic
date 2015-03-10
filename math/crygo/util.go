package crygo

import (
	"encoding/binary"
)

var byteOrder = binary.LittleEndian

func safeXORBytes(dst, a, b []byte) int {

	n := len(a)
	if len(b) < n {
		n = len(b)
	}
	for i := 0; i < n; i++ {
		dst[i] = a[i] ^ b[i]
	}
	return n
}

func duplicateBytes(a []byte) []byte {
	b := make([]byte, len(a))
	copy(b, a)
	return b
}

func fillBytes(bs []byte, b byte) {
	for i, _ := range bs {
		bs[i] = b
	}
}

func quoRem(x, y int) (q, r int) {
	q = x / y
	r = x - q*y
	return
}

func mod(x, y int64) int64 {

	t := x % y

	if t < 0 {
		t += y
	}

	return t
}

func getTwoUint32(src []byte) []uint32 {

	return []uint32{
		byteOrder.Uint32(src[0:4]),
		byteOrder.Uint32(src[4:8]),
	}
}

func putTwoUint32(dst []byte, src []uint32) {

	byteOrder.PutUint32(dst[0:4], src[0])
	byteOrder.PutUint32(dst[4:8], src[1])
}
