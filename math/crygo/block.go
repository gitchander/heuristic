package crygo

import (
	"crypto/cipher"
)

type blockCipher struct {
	x []uint32
	r replacer
}

func NewBlock(key []byte) (cipher.Block, error) {

	return NewBlockTable(key, Table1)
}

func NewBlockTable(key []byte, table []byte) (cipher.Block, error) {

	block := new(blockCipher)
	var err error

	if block.r, err = newReplacer256x4(table); err != nil {
		return nil, err
	}

	if block.x, err = newKey(key); err != nil {
		return nil, err
	}

	return block, nil
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

	encryptBlock(this.x, this.r, dst, src)
}

func (this *blockCipher) Decrypt(dst, src []byte) {

	if len(src) < BlockSize {
		panic("crypto/crygo: input not full block")
	}

	if len(dst) < BlockSize {
		panic("crypto/crygo: output not full block")
	}

	decryptBlock(this.x, this.r, dst, src)
}
