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

	h = NewBinaryMaxHeap(20)
	for i := 30; i > 10; i-- {
		h.Insert(i)
	}
	t.Log(h.eles)
	for i := 30; i > 10; i-- {
		fmt.Println(h.DeleteFirst())
	}
	t.Log(h.eles)

	h = NewBinaryMinHeap(100)
	var arr []int
	arr = append(arr, 10)
	arr = append(arr, 12)
	arr = append(arr, 1)
	arr = append(arr, 14)
	arr = append(arr, 6)
	arr = append(arr, 5)
	arr = append(arr, 8)
	arr = append(arr, 15)
	arr = append(arr, 3)
	arr = append(arr, 9)
	arr = append(arr, 7)
	arr = append(arr, 4)
	arr = append(arr, 11)
	arr = append(arr, 13)
	arr = append(arr, 2)

	fmt.Println(h.HeapSort(arr))
	//h.size = 0
	//for _, v := range arr {
	//	h.Insert(v)
	//}
	//fmt.Println(h.eles)
	//h.IncreaseKey(0, 15)
	//fmt.Println(h.eles)

}
