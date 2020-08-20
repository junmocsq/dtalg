package iterator

import (
	"fmt"
	"testing"
)

func TestIterator(t *testing.T) {
	// 实现带快照的迭代器
	vector := NewVector()
	iter := NewIterator(vector)
	for i := 0; i < 20; i++ {
		vector.Add(i)
	}
	iter.Init()
	vector.Remove(0)
	vector.Remove(11)
	vector.Add(100)
	for iter.hasNext() {
		fmt.Printf("%v\t", iter.currentItem())
	}
	fmt.Println()
	fmt.Println("========================================")
	iter.Init()
	for iter.hasNext() {
		fmt.Printf("%v\t", iter.currentItem())
	}
	fmt.Println()
}
