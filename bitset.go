package bitset

import (
	"fmt"
	"math/big"
)

type BitSet struct {
	bits big.Int
}

func (b *BitSet) Set(bit uint) *BitSet {
	b.bits.SetBit(&b.bits, int(bit), 1)
	return b
}

func (b *BitSet) Clear(bit uint) *BitSet {
	b.bits.SetBit(&b.bits, int(bit), 0)
	return b
}

func (b *BitSet) Test(bit uint) bool {
	return b.bits.Bit(int(bit)) == 1
}

func (b *BitSet) Flip(bit uint) *BitSet {
	if !b.Test(bit) {
		return b.Set(bit)
	}
	return b.Clear(bit)
}

func (b1 *BitSet) Union(b2 *BitSet) (b3 *BitSet) {
	b3 = new(BitSet)
	b3.bits.Or(&b1.bits, &b2.bits)
	return
}

func (b1 *BitSet) Intersection(b2 *BitSet) (b3 *BitSet) {
	b3 = new(BitSet)
	b3.bits.And(&b1.bits, &b2.bits)
	return
}

func (b1 *BitSet) Difference(b2 *BitSet) (b3 *BitSet) {
	b3 = new(BitSet)
	b3.bits.AndNot(&b1.bits, &b2.bits)
	return
}

func Empty() *BitSet {
	return new(BitSet)
}

func (b BitSet) String() string {
	return fmt.Sprintf("%b", &b.bits)
}

func (b BitSet) Len() int {
	return b.bits.BitLen()
}

func (b BitSet) Bit2Ints() (ids []int32) {
	for i := 0; i < b.Len(); i++ {
		if b.Test(uint(i)) {
			ids = append(ids, int32(i))
		}
	}
	return ids
}
