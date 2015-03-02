package crygo

import (
	"crypto/cipher"
	"errors"
)

const (
	C0 uint32 = 0x01010101
	C1 uint32 = 0x01010104
)

type streamCipher struct {
	b        cipher.Block
	s        []uint32
	out      []byte
	outIndex int
}

func NewStream(block cipher.Block, syn []byte) (cipher.Stream, error) {

	size := block.BlockSize()

	if len(syn) != size {
		return nil, errors.New("wrong syn len")
	}

	synEnc := make([]byte, size)
	block.Encrypt(synEnc, syn)

	stream := &streamCipher{
		b: block,
		s: []uint32{
			byteOrder.Uint32(synEnc[0:4]),
			byteOrder.Uint32(synEnc[4:8]),
		},
		out:      make([]byte, size),
		outIndex: 0,
	}

	stream.refill()

	return stream, nil
}

func (this *streamCipher) refill() {

	s := this.s

	s[0] = s[0] + C0
	s[1] = sum_mod32m1(s[1], C1)

	byteOrder.PutUint32(this.out[0:4], s[0])
	byteOrder.PutUint32(this.out[4:8], s[1])

	this.b.Encrypt(this.out, this.out)

	this.outIndex = 0
}

func (this *streamCipher) XORKeyStream(dst, src []byte) {

	for len(src) > 0 {
		if this.outIndex >= len(this.out) {
			this.refill()
		}
		n := safeXORBytes(dst, src, this.out[this.outIndex:])
		src = src[n:]
		dst = dst[n:]
		this.outIndex += n
	}
}

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

/*
func duplicate(a []byte) []byte {
	b := make([]byte, len(a))
	copy(b, a)
	return b
}
*/
