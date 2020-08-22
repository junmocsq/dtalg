package p6

import (
	"errors"
	"fmt"
)

// 斜堆是左式堆的自调节形式
type skewNode struct {
	ele   int
	left  *PrioritySkewQueue
	right *PrioritySkewQueue
}

type PrioritySkewQueue skewNode
type skewHeap struct {
	queue *PrioritySkewQueue
}

func NewSkewHeap() *skewHeap {
	return &skewHeap{}
}

// 左式堆合并
func (p *skewHeap) Merge(queue *PrioritySkewQueue) {
	p.queue = p.merge1(p.queue, queue)
}

// 判断是否有nil，存在直接返回另一个
//
func (p *skewHeap) merge1(h1, h2 *PrioritySkewQueue) *PrioritySkewQueue {
	if h1 == nil {
		return h2
	}
	if h2 == nil {
		return h1
	}
	if h1.ele < h2.ele {
		return p.merge2(h1, h2)
	} else {
		return p.merge2(h2, h1)
	}
}

// h1的左节点为空，直接把h2转到h1左节点即可
// 先合并h1的右节点和h2合并，再把合并后的节点转到h1的右节点
func (p *skewHeap) merge2(h1, h2 *PrioritySkewQueue) *PrioritySkewQueue {
	if h1.left == nil {
		h1.left = h2
	} else {
		h1.right = p.merge1(h1.right, h2)
		if h1.left != nil && h1.right != nil {
			left := h1.left
			h1.left = h1.right
			h1.right = left
		}
	}
	return h1
}

func (p *skewHeap) DeleteMin() (*skewNode, error) {
	if p.queue == nil {
		return nil, errors.New("leftist heap is empty")
	}
	delNode := p.queue
	p.queue = p.merge1(delNode.left, delNode.right)
	return (*skewNode)(delNode), nil
}

func (p *skewHeap) Insert(ele int) {
	node := &skewNode{ele: ele}
	p.queue = p.merge1(p.queue, (*PrioritySkewQueue)(node))
}

func (p *skewHeap) inOrderPrint(queue *PrioritySkewQueue) {
	if queue == nil {
		return
	}
	fmt.Printf("ele:%d\t", queue.ele)
	p.inOrderPrint(queue.left)
	p.inOrderPrint(queue.right)
}

func (p *skewHeap) midOrderPrint(queue *PrioritySkewQueue) {
	if queue == nil {
		return
	}
	p.inOrderPrint(queue.left)
	fmt.Printf("ele:%d\t", queue.ele)
	p.inOrderPrint(queue.right)
}

func (p *skewHeap) postOrderPrint(queue *PrioritySkewQueue) {
	if queue == nil {
		return
	}
	p.inOrderPrint(queue.left)
	p.inOrderPrint(queue.right)
	fmt.Printf("ele:%d\t", queue.ele)
}
