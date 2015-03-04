package crygo

import (
	"crypto/cipher"
)

// Electronic Codebook - ECB

type blockCipher struct {
	x []uint32
	t Table
}

func NewBlockCipher(t Table, key []byte) (cipher.Block, error) {

	if len(key) != KeySize {
		return nil, ErrorKeyLen
	}

	xs := make([]uint32, 8)
	for i, _ := range xs {
		xs[i] = byteOrder.Uint32(key[i*4:])
	}

	return &blockCipher{
		x: xs,
		t: t,
	}, nil
}

func (this *blockCipher) BlockSize() int {
	return BlockSize
}

func (this *blockCipher) Encrypt(dst, src []byte) {

	if len(src) < BlockSize {
		panic("crypto/crygo: input not full block")
	}

	if len(dst) < BlockSize {
		panic("crypto/crygo: output not full block")
	}

	encryptBlock(this.x, this.t, dst, src)
}

func (this *blockCipher) Decrypt(dst, src []byte) {

	if len(src) < BlockSize {
		panic("crypto/crygo: input not full block")
	}

	if len(dst) < BlockSize {
		panic("crypto/crygo: output not full block")
	}

	decryptBlock(this.x, this.t, dst, src)
}
