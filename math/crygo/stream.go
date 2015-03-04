package crygo

import (
	"crypto/cipher"
)

const (
	c0 = 0x01010101
	c1 = 0x01010104
)

type streamCipher struct {
	block    cipher.Block
	s        []uint32
	out      []byte
	outIndex int
}

func NewStreamCipher(block cipher.Block, syn []byte) (cipher.Stream, error) {

	size := block.BlockSize()

	if len(syn) != size {
		return nil, ErrorSynLen
	}

	sc := &streamCipher{
		block:    block,
		out:      make([]byte, size),
		outIndex: 0,
	}

	synEnc := make([]byte, size)
	block.Encrypt(synEnc, syn)
	sc.s = getTwoUint32(synEnc)

	sc.nextFill()

	return sc, nil
}

func (this *streamCipher) nextFill() {

	s := this.s

	s[0] = s[0] + c0
	s[1] = add_mod32m1(s[1], c1)

	putTwoUint32(this.out, s)

	this.block.Encrypt(this.out, this.out)
	this.outIndex = 0
}

func (this *streamCipher) XORKeyStream(dst, src []byte) {

	for len(src) > 0 {

		if this.outIndex >= len(this.out) {
			this.nextFill()
		}

		n := safeXORBytes(dst, src, this.out[this.outIndex:])

		src = src[n:]
		dst = dst[n:]
		this.outIndex += n
	}
}
