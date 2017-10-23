## bitset

一个用 Go 实现的的 BitSet 

## How to Use

```go
import (
    "fmt"
    "github.com/zheng-ji/bitset"
)

func TestBitInts(b *testing.T) {
	d1 := bitset.Empty()
	d2 := bitset.Empty()

    // set content
	d1.Set(uint(20000))
	d2.Set(uint(10000))

    // output bitset content
	fmt.Println(d1.Bit2Ints())

    // intersetion
    b1.Intersection(b2)

    // Difference
    b1.Intersection(b2)

    // different
    b1.Difference(b2)

    //Len
    b1.Len()
}
```

License
-------

Copyright (c) 2017 by [zheng-ji](zheng-ji.info) released under a MIT style license.
