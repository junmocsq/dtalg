package p6

import (
	"fmt"
	"testing"
)

func TestNewSkewHeap(t *testing.T) {
	heap := NewSkewHeap()
	heap.Insert(3)
	heap.Insert(10)
	heap.Insert(8)
	heap.Insert(21)
	heap.Insert(14)
	heap.Insert(23)
	heap.Insert(26)

	heap1 := NewSkewHeap()
	heap1.Insert(6)
	heap1.Insert(12)
	heap1.Insert(7)
	heap1.Insert(18)
	heap1.Insert(24)
	heap1.Insert(37)
	heap1.Insert(18)
	heap1.Insert(33)
	heap.Merge(heap1.queue)

	heap.inOrderPrint(heap.queue)
	fmt.Println()
	heap.midOrderPrint(heap.queue)
	fmt.Println()
	heap.postOrderPrint(heap.queue)
	fmt.Println()

	for {
		ele, e := heap.DeleteMin()
		if e != nil {
			break
		}
		fmt.Println(ele.ele)
	}
}
