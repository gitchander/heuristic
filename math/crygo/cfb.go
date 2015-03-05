package crygo

import (
	"crypto/cipher"
)

// Cipher Feedback Mode - CFB

func NewCFBEncrypter(block cipher.Block, syn []byte) (cipher.Stream, error) {
	return newCFBCipher(block, syn, false)
}

func NewCFBDecrypter(block cipher.Block, syn []byte) (cipher.Stream, error) {
	return newCFBCipher(block, syn, true)
}

type cfb struct {
	b        cipher.Block
	out      []byte
	outIndex int

	decrypt bool
}

func newCFBCipher(block cipher.Block, syn []byte, decrypt bool) (*cfb, error) {

	blockSize := block.BlockSize()

	if len(syn) != blockSize {
		return nil, ErrorSynLen
	}

	out := make([]byte, blockSize)
	copy(out, syn)

	return &cfb{
		b:        block,
		out:      out,
		outIndex: blockSize,
		decrypt:  decrypt,
	}, nil
}

func (this *cfb) XORKeyStream(dst, src []byte) {

	for len(src) > 0 {

		if this.outIndex >= len(this.out) {
			this.b.Encrypt(this.out, this.out)
			this.outIndex = 0
		}

		n := safeXORBytes(dst, src, this.out[this.outIndex:])

		if this.decrypt {
			copy(this.out[this.outIndex:this.outIndex+n], src[:n])
		} else {
			copy(this.out[this.outIndex:this.outIndex+n], dst[:n])
		}

		src = src[n:]
		dst = dst[n:]
		this.outIndex += n
	}
}
