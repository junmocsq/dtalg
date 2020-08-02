package dtc

import "fmt"

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
