// zheng-ji.info

package bitset

import (
	"fmt"
	"testing"
	"time"
)

/*
* [20000]
* 393.504µs
*
 */
func TestBitInts(b *testing.T) {
	d := Empty()
	num := 20000
	d.Set(uint(num))

	start := time.Now()

	fmt.Println(d.Bit2Ints())

	t := time.Now()
	elapsed := t.Sub(start)
	fmt.Println(elapsed)
}
