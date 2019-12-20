package ring

import (
	"container/ring"
	"fmt"
	"testing"
)

func Test_ring(t *testing.T) {
	r := ring.New(10)
	r.Value = 1
	r.Next().Value = 2
	r.Next().Next().Next().Value = 3

	r2 := ring.New(20)
	r2.Value = 11
	r.Next().Value = 12
	r.Next().Next().Value = 13
	r.Link(r2)

	i, n := 0, r.Len()
	for p := r; i < n; p = p.Next() {
		fmt.Println(p.Value)
		i++
	}

}
