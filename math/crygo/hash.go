package crygo

import (
	"crypto/cipher"
	"hash"
)

type digest struct {
	block    cipher.Block
	out      []byte
	src      []byte
	srcIndex int
}

func NewHash(block cipher.Block) hash.Hash {
	return &digest{
		block: block,
		out:   make([]byte, block.BlockSize()),
		src:   make([]byte, block.BlockSize()),
	}
}

func (this *digest) nextFill() {

	safeXORBytes(this.out, this.out, this.src)

	for i := 0; i < 16; i++ {
		this.block.Encrypt(this.out, this.out)
	}

	fillBytes(this.src, 0)
	this.srcIndex = 0
}

func (this *digest) Write(src []byte) (count int, err error) {

	for len(src) > 0 {

		if this.srcIndex >= len(this.src) {
			this.nextFill()
		}

		n := len(src)
		if m := (len(this.src) - this.srcIndex); n > m {
			n = m
		}

		copy(this.src[this.srcIndex:this.srcIndex+n], src[:n])

		src = src[n:]
		this.srcIndex += n
		count += n
	}

	return
}

func (this *digest) checkSum() []byte {

	hash := make([]byte, this.block.BlockSize())

	safeXORBytes(hash, this.out, this.src)

	for i := 0; i < 16; i++ {
		this.block.Encrypt(hash, hash)
	}

	return hash
}

func (this *digest) Sum(in []byte) []byte {

	hash := this.checkSum()
	return append(in, hash...)
}

func (this *digest) Reset() {

	fillBytes(this.out, 0)
	fillBytes(this.src, 0)
	this.srcIndex = 0
}

func (this *digest) Size() int {

	return BlockSize
}

func (this *digest) BlockSize() int {

	return BlockSize
}
