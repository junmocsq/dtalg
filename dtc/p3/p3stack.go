package p3

import "errors"

type stacker interface {
	IsEmpty() bool
	MakeEmpty()
	Push(int2 int)
	Top() (int, error)
	Pop() (int, error)
}

type linkedStackNode struct {
	ele  int
	next *linkedStackNode
}
type linkedStack linkedStackNode

func NewLinkedStack() *linkedStack {
	return &linkedStack{}
}

func (ls *linkedStack) IsEmpty() bool {
	return ls.next == nil
}

func (ls *linkedStack) MakeEmpty() {
	ls.next = nil
}
func (ls *linkedStack) Push(ele int) {
	temp := ls.next
	ls.next = &linkedStackNode{ele: ele, next: temp}
}
func (ls *linkedStack) Top() (int, error) {
	if ls.IsEmpty() {
		return 0, errors.New("stack is empty")
	}
	return ls.next.ele, nil
}
func (ls *linkedStack) Pop() (int, error) {
	if ls.IsEmpty() {
		return 0, errors.New("stack is empty")
	}
	temp := ls.next
	ls.next = temp.next
	return temp.ele, nil
}

type arrayStack struct {
	arr []int
}

func NewArrayStack() *arrayStack {
	return &arrayStack{}
}

func (ls *arrayStack) IsEmpty() bool {
	return len(ls.arr) == 0
}

func (ls *arrayStack) MakeEmpty() {
	ls.arr = ls.arr[:0]
}
func (ls *arrayStack) Push(ele int) {
	ls.arr = append(ls.arr, ele)
}
func (ls *arrayStack) Top() (int, error) {
	if ls.IsEmpty() {
		return 0, errors.New("stack is empty")
	}
	return ls.arr[len(ls.arr)-1], nil
}
func (ls *arrayStack) Pop() (int, error) {
	if ls.IsEmpty() {
		return 0, errors.New("stack is empty")
	}
	temp := ls.arr[len(ls.arr)-1]
	ls.arr = ls.arr[:len(ls.arr)-1]
	return temp, nil
}
