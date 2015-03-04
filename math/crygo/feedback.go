package crygo

import (
	"crypto/cipher"
)

// Cipher Feedback Mode - CFB

type FeedbackEncrypter interface {
	Encrypt(dst, src []byte)
}

type FeedbackDecrypter interface {
	Decrypt(dst, src []byte)
}

func NewFeedbackEncrypter(block cipher.Block, syn []byte) (FeedbackEncrypter, error) {
	return newFeedbackCipher(block, syn)
}

func NewFeedbackDecrypter(block cipher.Block, syn []byte) (FeedbackDecrypter, error) {
	return newFeedbackCipher(block, syn)
}

type feedbackCipher struct {
	block    cipher.Block
	out      []byte
	outIndex int
}

func newFeedbackCipher(block cipher.Block, syn []byte) (*feedbackCipher, error) {

	size := block.BlockSize()

	if len(syn) != size {
		return nil, ErrorSynLen
	}

	out := make([]byte, size)
	block.Encrypt(out, syn)

	return &feedbackCipher{
		block:    block,
		out:      out,
		outIndex: 0,
	}, nil
}

func (this *feedbackCipher) nextFill() {

	this.block.Encrypt(this.out, this.out)
	this.outIndex = 0
}

func (this *feedbackCipher) Encrypt(dst, src []byte) {

	for len(src) > 0 {

		if this.outIndex >= len(this.out) {
			this.nextFill()
		}

		n := safeXORBytes(dst, src, this.out[this.outIndex:])
		copy(this.out[this.outIndex:this.outIndex+n], dst[:n])

		src = src[n:]
		dst = dst[n:]
		this.outIndex += n
	}
}

func (this *feedbackCipher) Decrypt(dst, src []byte) {

	for len(src) > 0 {

		if this.outIndex >= len(this.out) {
			this.nextFill()
		}

		n := safeXORBytes(dst, src, this.out[this.outIndex:])
		copy(this.out[this.outIndex:this.outIndex+n], src[:n])

		src = src[n:]
		dst = dst[n:]
		this.outIndex += n
	}
}
