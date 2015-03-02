package crygo

import (
	"encoding/binary"
	"errors"
)

// http://www.kzi-manual.ru/standarts/?id=2

const (
	BlockSize = 8
	KeySize   = 32
)

var byteOrder = binary.LittleEndian

func replace8_to_replace4(t8, t4 []byte) {

	// x * 256 = x << 8
	// x * 32 = x << 5
	// x * 16 = x << 4
	// x * 2^n = x << n

	for i := 0; i < 4; i++ {
		for j := 0; j < 16; j++ {
			for k := 0; k < 16; k++ {
				t4[(i<<8)+(j<<4)+k] = t8[(i<<5)+k] | (t8[(i<<5)+16+j] << 4)
			}
		}
	}
}

func replace8(t []byte, s0 uint32) (s1 uint32) {

	s1 |= uint32(t[((s0>>0x00)&0x0F)+0x00]) << 0x00
	s1 |= uint32(t[((s0>>0x04)&0x0F)+0x10]) << 0x04
	s1 |= uint32(t[((s0>>0x08)&0x0F)+0x20]) << 0x08
	s1 |= uint32(t[((s0>>0x0C)&0x0F)+0x30]) << 0x0C
	s1 |= uint32(t[((s0>>0x10)&0x0F)+0x40]) << 0x10
	s1 |= uint32(t[((s0>>0x14)&0x0F)+0x50]) << 0x14
	s1 |= uint32(t[((s0>>0x18)&0x0F)+0x60]) << 0x18
	s1 |= uint32(t[((s0>>0x1C)&0x0F)+0x70]) << 0x1C

	return
}

func replace4(t []byte, s0 uint32) (s1 uint32) {

	s1 |= uint32(t[((s0>>0x00)&0xFF)+0x0000]) << 0x00
	s1 |= uint32(t[((s0>>0x08)&0xFF)+0x0100]) << 0x08
	s1 |= uint32(t[((s0>>0x10)&0xFF)+0x0200]) << 0x10
	s1 |= uint32(t[((s0>>0x18)&0xFF)+0x0300]) << 0x18

	return
}

func baseStep(n []uint32, x uint32, r replacer) {

	s := n[0] + x // (n + x) mod 32
	s = r.replace(s)
	s = (s << 11) | (s >> 21)
	s ^= n[1]

	n[1] = n[0]
	n[0] = s
}

func encrypt(xs []uint32, r replacer, n []uint32) {

	for j := 0; j < 3; j++ {
		for i := 0; i < 8; i++ {
			baseStep(n, xs[i], r)
		}
	}
	for i := 8; i > 0; i-- {
		baseStep(n, xs[i-1], r)
	}

	n[0], n[1] = n[1], n[0]
}

func decrypt(xs []uint32, r replacer, n []uint32) {

	for i := 0; i < 8; i++ {
		baseStep(n, xs[i], r)
	}
	for j := 0; j < 3; j++ {
		for i := 8; i > 0; i-- {
			baseStep(n, xs[i-1], r)
		}
	}

	n[0], n[1] = n[1], n[0]
}

func encryptBlock(xs []uint32, r replacer, dst, src []byte) {

	n := []uint32{
		byteOrder.Uint32(src[0:4]),
		byteOrder.Uint32(src[4:8]),
	}

	encrypt(xs, r, n)

	byteOrder.PutUint32(dst[0:4], n[0])
	byteOrder.PutUint32(dst[4:8], n[1])
}

func decryptBlock(xs []uint32, r replacer, dst, src []byte) {

	n := []uint32{
		byteOrder.Uint32(src[0:4]),
		byteOrder.Uint32(src[4:8]),
	}

	decrypt(xs, r, n)

	byteOrder.PutUint32(dst[0:4], n[0])
	byteOrder.PutUint32(dst[4:8], n[1])
}

func newKey(key []byte) ([]uint32, error) {

	if len(key) != KeySize {
		return nil, errors.New("wrong key size")
	}

	xs := make([]uint32, 8)
	for i, _ := range xs {
		xs[i] = byteOrder.Uint32(key[i*4:])
	}

	return xs, nil
}
