package crygo

import (
	"bytes"
	"math/rand"
	"testing"
	"time"
)

func newRand() *rand.Rand {
	return rand.New(rand.NewSource(time.Now().UnixNano()))
}

func fillBytes(r *rand.Rand, bs []byte) {
	for i, _ := range bs {
		bs[i] = byte(r.Intn(256))
	}
}

func TestBlockCrypt(t *testing.T) {

	r := newRand()

	key := make([]byte, 32)

	s1 := make([]byte, BlockSize)
	s2 := make([]byte, BlockSize)
	s3 := make([]byte, BlockSize)

	for i := 0; i < 100; i++ {

		fillBytes(r, key)

		c, err := NewBlock(key)
		if err != nil {
			t.Error(err)
			return
		}

		for j := 0; j < 1000; j++ {

			fillBytes(r, s1)

			c.Encrypt(s2, s1)
			if bytes.Compare(s1, s2) == 0 {
				t.Error("Encrypt compare true")
				return
			}

			c.Decrypt(s3, s2)
			if bytes.Compare(s1, s3) != 0 {
				t.Error("Decrypt compare false")
				return
			}
		}
	}
}

func TestSamples(t *testing.T) {

	type Sample struct {
		Input  []byte
		Output []byte
	}

	type TestSamples struct {
		Table   []byte
		Key     []byte
		Samples []Sample
	}

	var ts = TestSamples{
		Table: Table2,
		Key: []byte{
			0x81, 0x82, 0x83, 0x84, 0x85, 0x86, 0x87, 0x88,
			0x89, 0x8a, 0x8b, 0x8c, 0x8d, 0x8e, 0x8f, 0x80,
			0xd1, 0xd2, 0xd3, 0xd4, 0xd5, 0xd6, 0xd7, 0xd8,
			0xd9, 0xda, 0xdb, 0xdc, 0xdd, 0xde, 0xdf, 0xd0,
		},
		Samples: []Sample{

			Sample{
				Input:  []byte{0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08},
				Output: []byte{0xce, 0x5a, 0x5e, 0xd7, 0xe0, 0x57, 0x7a, 0x5f},
			},

			Sample{
				Input:  []byte{0xf1, 0xf2, 0xf3, 0xf4, 0xf5, 0xf6, 0xf7, 0xf8},
				Output: []byte{0xd0, 0xcc, 0x85, 0xce, 0x31, 0x63, 0x5b, 0x8b},
			},
		},
	}

	c, err := NewBlockTable(ts.Key, ts.Table)
	if err != nil {
		t.Error(err)
		return
	}

	dst := make([]byte, BlockSize)

	for _, sample := range ts.Samples {
		c.Encrypt(dst, sample.Input)
		if bytes.Compare(dst, sample.Output) != 0 {
			t.Error("not compare")
		}
	}
}

func TestFastTable(t *testing.T) {

	r := newRand()

	t11, err := newReplacer256x4(Table1)
	if err != nil {
		t.Error(err)
		return
	}

	t12, err := newReplacer16x8(Table1)
	if err != nil {
		t.Error(err)
		return
	}

	t21, err := newReplacer256x4(Table2)
	if err != nil {
		t.Error(err)
		return
	}

	t22, err := newReplacer16x8(Table2)
	if err != nil {
		t.Error(err)
		return
	}

	type Tables struct {
		t1 replacer
		t2 replacer
	}

	ts := []Tables{
		Tables{
			t11,
			t12,
		},
		Tables{
			t21,
			t22,
		},
	}

	for _, y := range ts {
		for i := 0; i < 1000000; i++ {
			u := r.Uint32()
			if y.t1.replace(u) != y.t2.replace(u) {
				t.Error("wrong mix")
				return
			}
		}
	}
}
