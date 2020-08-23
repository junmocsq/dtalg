package p6

import (
	"errors"
	"fmt"
)

// 优先队列

type binaryHeap struct {
	eles     []int
	size     int
	capacity int
	compare  func(a, b int) bool
}

func compareMin(a, b int) bool {
	return a > b
}
func compareMax(a, b int) bool {
	return !compareMin(a, b)
}

// 最小优先队列
func NewBinaryMinHeap(cap int) *binaryHeap {
	return &binaryHeap{
		eles:     make([]int, cap, cap),
		capacity: cap,
		compare:  compareMin,
	}
}

// 最大优先队列
func NewBinaryMaxHeap(cap int) *binaryHeap {
	return &binaryHeap{
		eles:     make([]int, cap, cap),
		capacity: cap,
		compare:  compareMax,
	}
}

func (h *binaryHeap) Insert(ele int) error {
	if h.IsFull() {
		return errors.New("heap is full")
	}
	h.size++
	h.eles[h.size-1] = ele
	h.percolateUp(h.size - 1)
	return nil
}

func (h *binaryHeap) parent(index int) int {
	return (index - 1) / 2
}

func (h *binaryHeap) childLeft(index int) int {
	return (index+1)*2 - 1
}

func (h *binaryHeap) DeleteFirst() (int, error) {
	if h.IsEmpty() {
		return 0, errors.New("heap is empty")
	}
	del := h.eles[0]
	h.size--
	lastEle := h.eles[h.size]
	h.eles[0] = lastEle
	h.percolateDown(0)
	return del, nil
}

func (h *binaryHeap) FindFirst() (int, error) {
	if h.IsEmpty() {
		return 0, errors.New("heap is empty")
	}
	return h.eles[0], nil
}

func (h *binaryHeap) IsEmpty() bool {
	return h.size == 0
}

func (h *binaryHeap) IsFull() bool {
	return h.size == h.capacity
}

// 减少指定节点的值 val为正值
func (h *binaryHeap) DecreaseKey(index, val int) bool {
	if index >= h.size || index < 0 {
		return false
	}
	if val < 0 {
		val = -val
	}
	h.eles[index] -= val
	if h.compare(2, 1) {
		h.percolateUp(index)
	} else {
		h.percolateDown(index)
	}
	return true
}

// 增加指定节点的值 val为正值
func (h *binaryHeap) IncreaseKey(index, val int) bool {
	if index >= h.size || index < 0 {
		return false
	}
	if val < 0 {
		val = -val
	}
	h.eles[index] += val
	if h.compare(2, 1) {
		h.percolateDown(index)
	} else {
		h.percolateUp(index)
	}
	return true
}

func (h *binaryHeap) Delete(index int) (int, error) {
	if h.size == 0 {
		return 0, errors.New("priority queue is empty!")
	}
	h.size--
	delNode := h.eles[index]
	h.eles[index] = h.eles[h.size]
	h.percolateDown(index)
	h.percolateUp(index)
	return delNode, nil
}

// 构建堆
func (h *binaryHeap) BuildHeap(arr []int) error {
	length := len(arr)
	if h.capacity < length {
		return errors.New("数组过小")
	}
	h.size = length
	for k, v := range arr {
		h.eles[k] = v
	}

	for i := h.parent(h.size); i >= 0; i-- {
		h.percolateDown(i)
	}
	return nil
}

func (h *binaryHeap) FindKMinOrMax(arr []int, k int) error {
	h.BuildHeap(arr)
	for i := 0; i < k; i++ {
		if i == k-1 {
			fmt.Println(h.DeleteFirst())
		} else {
			h.DeleteFirst()
		}
	}
	return nil
}

// 上滤
func (h *binaryHeap) percolateUp(index int) {
	if index <= 0 {
		return
	}
	for index != 0 {
		parent := (index - 1) / 2
		if h.compare(h.eles[parent], h.eles[index]) {
			h.eles[parent], h.eles[index] = h.eles[index], h.eles[parent]
			index = parent
		} else {
			break
		}
	}
	return
}

// 下滤
func (h *binaryHeap) percolateDown(index int) {
	if index < 0 {
		return
	}
	for {
		child := index*2 + 1
		if child+1 < h.size && h.compare(h.eles[child], h.eles[child+1]) {
			child++
		}
		if child >= h.size {
			break
		}
		if h.compare(h.eles[index], h.eles[child]) {
			h.eles[index], h.eles[child] = h.eles[child], h.eles[index]
			index = child
		} else {
			break
		}
	}
	return
}

// 堆排序
func (h *binaryHeap) HeapSort(arr []int) []int {
	h.BuildHeap(arr)
	for i := h.size - 1; i > 0; i-- {
		del, err := h.DeleteFirst()
		if err != nil {
			panic("heap sort failed")
		}
		h.eles[i] = del
	}
	length := len(arr)
	for i := length - 1; i >= 0; i-- {
		arr[length-i-1] = h.eles[i]
	}
	return arr
}
