package crygo

import (
	"encoding/binary"
	"math/rand"
	"time"
)

type randomer interface {
	Intn(n int) int
	Uint32() uint32
	FillBytes(data []byte)
}

type privRandomer struct {
	r *rand.Rand
}

func newRand() randomer {

	return &privRandomer{
		r: rand.New(rand.NewSource(time.Now().UnixNano())),
	}
}

func (this *privRandomer) Intn(n int) int {
	return this.r.Intn(n)
}

func (this *privRandomer) Uint32() uint32 {
	return this.r.Uint32()
}

func quo_rem(x, y int) (q, r int) {
	q = x / y
	r = x - q*y
	return
}

func (this *privRandomer) FillBytes(data []byte) {

	const sizeOfUint32 = 4
	quo, rem := quo_rem(len(data), sizeOfUint32)

	if quo > 0 {
		bo := binary.BigEndian
		for i := 0; i < quo; i++ {
			bo.PutUint32(data, this.r.Uint32())
			data = data[sizeOfUint32:]
		}
	}

	if rem > 0 {
		u := this.r.Uint32()
		for i := 0; i < rem; i++ {
			data[i] = byte(u)
			u >>= 8
		}
	}
}
