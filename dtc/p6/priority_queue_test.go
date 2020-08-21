package p6

import (
	"fmt"
	"testing"
)

func TestBinaryHeap(t *testing.T) {
	h := NewBinaryMinHeap(20)
	for i := 30; i > 10; i-- {
		h.Insert(i)
	}
	t.Log(h.eles)
	for i := 30; i > 10; i-- {
		fmt.Println(h.DeleteFirst())
	}
	t.Log(h.eles)

	var arr []int
	for i := 10; i < 30; i++ {
		arr = append(arr, i)
	}
	h.BuildHeap(arr)
	t.Log(h.eles)

	h = NewBinaryMaxHeap(20)
	for i := 30; i > 10; i-- {
		h.Insert(i)
	}
	t.Log(h.eles)
	for i := 30; i > 10; i-- {
		fmt.Println(h.DeleteFirst())
	}
	t.Log(h.eles)

}
