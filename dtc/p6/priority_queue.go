package p6

import (
	"errors"
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
	index := h.size
	for index != 0 {
		if h.compare(h.eles[h.parent(index)], ele) {
			h.eles[index] = h.eles[h.parent(index)]
			index = h.parent(index)
		} else {
			break
		}
	}
	h.eles[index] = ele
	h.size++
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
	index := 0
	for child := h.childLeft(index); child < h.size; child = h.childLeft(index) {
		if child+1 != h.size && h.compare(h.eles[child], h.eles[child+1]) {
			child++
		}
		// 下滤
		if h.compare(lastEle, h.eles[child]) {
			h.eles[index] = h.eles[child]
			index = child
		} else {
			break
		}

	}
	h.eles[index] = lastEle
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
		child := h.childLeft(i)
		if child+1 <= h.size-1 && h.compare(h.eles[child], h.eles[child+1]) {
			child++
		}
		if h.compare(h.eles[i], h.eles[child]) {
			// 下滤
			h.eles[i], h.eles[child] = h.eles[child], h.eles[i]
		}
	}
	return nil
}
