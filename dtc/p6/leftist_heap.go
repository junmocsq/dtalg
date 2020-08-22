package p6

import (
	"errors"
	"fmt"
)

// 零路径长(null path length) npl：节点到一个没有两个儿子的节点的路径长
// 左式堆:对于堆中的每一个节点，左儿子的零路径长至少与右儿子的零路径长一样大
// 作用是为了高速合并

type leftistNode struct {
	ele   int
	left  *PriorityQueue
	right *PriorityQueue
	npl   int
}

type PriorityQueue leftistNode
type LeftistHeap struct {
	queue *PriorityQueue
}

func NewLeftistHeap() *LeftistHeap {
	return &LeftistHeap{}
}

// 左式堆合并
func (p *LeftistHeap) Merge(queue *PriorityQueue) {
	p.queue = p.merge1(p.queue, queue)
}

// 判断是否有nil，存在直接返回另一个
//
func (p *LeftistHeap) merge1(h1, h2 *PriorityQueue) *PriorityQueue {
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
func (p *LeftistHeap) merge2(h1, h2 *PriorityQueue) *PriorityQueue {
	if h1.left == nil {
		h1.left = h2
	} else {
		h1.right = p.merge1(h1.right, h2)
		if h1.left.npl < h1.right.npl {
			left := h1.left
			h1.left = h1.right
			h1.right = left
		}
		h1.npl = h1.right.npl + 1
	}
	return h1
}

func (p *LeftistHeap) DeleteMin() (*leftistNode, error) {
	if p.queue == nil {
		return nil, errors.New("leftist heap is empty")
	}
	delNode := p.queue
	p.queue = p.merge1(delNode.left, delNode.right)
	return (*leftistNode)(delNode), nil
}

func (p *LeftistHeap) Insert(ele int) {
	node := &leftistNode{ele: ele}
	p.queue = p.merge1(p.queue, (*PriorityQueue)(node))
}

func (p *LeftistHeap) inOrderPrint(queue *PriorityQueue) {
	if queue == nil {
		return
	}
	fmt.Printf("npl:%d ele:%d\t", queue.npl, queue.ele)
	p.inOrderPrint(queue.left)
	p.inOrderPrint(queue.right)
}

func (p *LeftistHeap) midOrderPrint(queue *PriorityQueue) {
	if queue == nil {
		return
	}
	p.inOrderPrint(queue.left)
	fmt.Printf("npl:%d ele:%d\t", queue.npl, queue.ele)
	p.inOrderPrint(queue.right)
}

func (p *LeftistHeap) postOrderPrint(queue *PriorityQueue) {
	if queue == nil {
		return
	}
	p.inOrderPrint(queue.left)
	p.inOrderPrint(queue.right)
	fmt.Printf("npl:%d ele:%d\t", queue.npl, queue.ele)
}
