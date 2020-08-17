package p3

import (
	"fmt"
)

type linkedListNode struct {
	ele  int
	next *linkedListNode
}

type LinkedLister interface {
	MakeEmpty() *linkedListNode
	IsEmpty() bool
	IsLast(position *linkedListNode) bool
	Find(ele int) *linkedListNode
	Delete(ele int) bool
	FindPrevious(ele int) *linkedListNode
	Insert(ele int, position *linkedListNode) bool
	DeleteList()
	Header() *linkedListNode
	First() *linkedListNode
	Advance(position *linkedListNode) *linkedListNode
	Retrieve(position *linkedListNode) int
}

func NewLinkedList() *linkedListNode {
	return &linkedListNode{}
}

func (l *linkedListNode) MakeEmpty() *linkedListNode {
	l.next = nil
	return l
}

func (l *linkedListNode) IsEmpty() bool {
	return l.next == nil
}
func (l *linkedListNode) IsLast(position *linkedListNode) bool {
	return position.next == nil
}
func (l *linkedListNode) Find(ele int) *linkedListNode {
	temp := l.next
	for temp != nil && temp.ele != ele {
		temp = temp.next
	}
	return temp
}
func (l *linkedListNode) Delete(ele int) bool {
	node := l.FindPrevious(ele)
	if l.IsLast(node) {
		return false
	} else {
		node.next = node.next.next
		return true
	}
}
func (l *linkedListNode) FindPrevious(ele int) *linkedListNode {
	temp := l
	for temp.next != nil && temp.next.ele != ele {
		temp = temp.next
	}
	return temp
}
func (l *linkedListNode) Insert(ele int, position *linkedListNode) bool {
	node := &linkedListNode{ele: ele}
	next := position.next
	position.next = node
	node.next = next
	return true
}

func (l *linkedListNode) DeleteList() {
	p := l.next
	l.next = nil
	for p != nil {
		tmp := p.next
		p = nil
		p = tmp

	}

}
func (l *linkedListNode) Header() *linkedListNode {
	return l
}
func (l *linkedListNode) First() *linkedListNode {
	return l.next
}
func (l *linkedListNode) Advance(position *linkedListNode) *linkedListNode {
	temp := l
	for temp.next != position {
		temp = temp.next
	}
	return temp
}
func (l *linkedListNode) Retrieve(position *linkedListNode) int {
	return position.ele
}

func (l *linkedListNode) Print() {
	temp := l.next
	for temp != nil {
		fmt.Printf("%d ", temp.ele)
		temp = temp.next
	}
	fmt.Println()
}

func (l *linkedListNode) Add(ele int) bool {
	node := &linkedListNode{ele: ele}
	tmp := l
	for tmp.next != nil {
		tmp = tmp.next
	}
	tmp.next = node
	return true
}

// 单链表反转
// 三个变量交替
func (l *linkedListNode) Reverse() {
	firstNode := l.next
	if firstNode == nil || firstNode.next == nil {
		return
	}
	// 当前节点 下一节点 下下节点
	cur := firstNode
	next := firstNode.next
	firstNode.next = nil
	nextNext := next.next
	for nextNext != nil {
		next.next = cur
		cur = next
		next = nextNext
		nextNext = nextNext.next
	}
	next.next = cur
	l.next = next
}

// 环检测
// 双指针
func (l *linkedListNode) CheckRing() bool {
	// 双指针，一个走一步 一个走两步 如果存在环，则一定会相遇
	a := l.next
	b := l.next
	for a.next != nil && b.next != nil && b.next.next != nil {
		a = a.next
		b = b.next.next
		if a == b {
			return true
		}
	}
	return false
}

// 和并两个有序链表 假设都为升序
func (l *linkedListNode) Merge(ll *linkedListNode) {
	tmp := l
	ll = ll.next
	for tmp.next != nil && ll != nil {
		if ll.ele < tmp.next.ele {
			t := ll
			ll = ll.next
			t.next = tmp.next
			tmp.next = t
		}
		tmp = tmp.next
	}
	if ll != nil {
		tmp.next = ll
	}
}

/**
删除倒数第N个节点 N从1开始
双指针
*/
func (l *linkedListNode) DeleteNNodeFromLast(n int) bool {
	n = n - 1 // 转换为从0开始

	firstNode := l.next
	if firstNode == nil {
		// 空节点不能删除
		return false
	}
	secondNode := l // 删除节点的上一个节点
	i := 0
	for firstNode.next != nil {
		if i >= n {
			secondNode = secondNode.next
		}
		firstNode = firstNode.next
		i++
	}
	// 在最后一个i++之前，i必须满足i=n-1,否则就是不足n+1个元素，没有办法作删除
	if i < n {
		return false
	}
	// 删除
	secondNode.next = secondNode.next.next
	return true
}

/**
查找中间节点
双指针
*/
func (l *linkedListNode) FindMiddleNode() (a, b *linkedListNode) {
	firstNode := l
	secondnode := l

	for firstNode != nil && firstNode.next != nil {
		firstNode = firstNode.next.next
		secondnode = secondnode.next
	}
	if firstNode != nil { // 偶数数个节点
		return secondnode, secondnode.next
	} else {
		return secondnode, nil
	}

}
