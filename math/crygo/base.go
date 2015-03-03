package crygo

import (
	"encoding/binary"
)

const (
	BlockSize = 8
	KeySize   = 32
)

var byteOrder = binary.LittleEndian

func baseStep(n []uint32, x uint32, t Table) {

	s := n[0] + x // (n + x) mod 32
	s = t.Replace(s)
	s = (s << 11) | (s >> 21)
	s ^= n[1]

	n[1] = n[0]
	n[0] = s
}

func encrypt(xs []uint32, t Table, n []uint32) {

	for j := 0; j < 3; j++ {
		for i := 0; i < 8; i++ {
			baseStep(n, xs[i], t)
		}
	}
	for i := 8; i > 0; i-- {
		baseStep(n, xs[i-1], t)
	}

	n[0], n[1] = n[1], n[0]
}

func decrypt(xs []uint32, t Table, n []uint32) {

	for i := 0; i < 8; i++ {
		baseStep(n, xs[i], t)
	}
	for j := 0; j < 3; j++ {
		for i := 8; i > 0; i-- {
			baseStep(n, xs[i-1], t)
		}
	}

	n[0], n[1] = n[1], n[0]
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

func encryptBlock(xs []uint32, t Table, dst, src []byte) {

	n := getTwoUint32(src)
	encrypt(xs, t, n)
	putTwoUint32(dst, n)
}

func decryptBlock(xs []uint32, t Table, dst, src []byte) {

	n := getTwoUint32(src)
	decrypt(xs, t, n)
	putTwoUint32(dst, n)
}
