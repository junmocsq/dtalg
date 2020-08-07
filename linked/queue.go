package linked

import (
	"errors"
)

var ERROR_QUEUE_EMPTY = errors.New("队列为空")

type Queue interface {
	Enqueue(ele int) error
	Dequeue() (int, error)
}

type arrayQueue struct {
	arr []int
}

func (aq *arrayQueue) Enqueue(ele int) error {
	aq.arr = append(aq.arr, ele)
	return nil
}

func (aq *arrayQueue) Dequeue() (int, error) {
	length := len(aq.arr)
	if length == 0 {
		return 0, ERROR_QUEUE_EMPTY
	}
	ele := aq.arr[0]
	aq.arr = aq.arr[1:length]
	return ele, nil
}

type linkedQueue struct {
	ele  int
	next *linkedQueue
}

func newLinkedQueue() *linkedQueue {
	return &linkedQueue{}
}

func (l *linkedQueue) Enqueue(ele int) error {

	node := &linkedQueue{ele: ele}
	next := l.next
	l.next = node
	node.next = next
	return nil
}

func (l *linkedQueue) Dequeue() (int, error) {

	if l.next == nil {
		return 0, ERROR_QUEUE_EMPTY
	}
	for l.next != nil && l.next.next != nil {
		l = l.next
	}
	ele := l.next.ele
	l.next = nil
	return ele, nil
}

type loopArrayQueue struct {
	eles []int
	cap  int
	head int
	tail int
}

func newLoopArrayQueue() *loopArrayQueue {
	return &loopArrayQueue{
		eles: make([]int, 10),
		cap:  10,
		head: 0,
		tail: 0,
	}
}

func (l *loopArrayQueue) Enqueue(ele int) error {
	l.expand()
	l.head--
	if l.head < 0 {
		l.head = l.cap - 1
	}
	l.eles[l.head] = ele
	return nil
}

func (l *loopArrayQueue) Dequeue() (int, error) {
	if l.isEmpty() {
		return 0, ERROR_QUEUE_EMPTY
	}
	l.tail--
	if l.tail < 0 {
		l.tail = l.cap - 1
	}
	ele := l.eles[l.tail]
	return ele, nil
}

func (l *loopArrayQueue) isFull() bool {
	return (l.head+l.cap-1)%l.cap == l.tail
}

func (l *loopArrayQueue) expand() {
	if !l.isFull() {
		return
	}
	newCap := l.cap * 2
	if l.cap > 1024 {
		newCap = l.cap + 1024
	}
	newEles := make([]int, newCap)
	start := l.head
	end := l.tail
	if start > end {
		end += l.cap
	}
	i := 0
	for ; start <= end; start++ {
		newEles[i] = l.eles[(start+l.cap)%l.cap]
		i++
	}
	l.cap = newCap
	l.eles = newEles
	l.head = 0
	l.tail = i - 1
}

func (l *loopArrayQueue) isEmpty() bool {
	return l.head == l.tail
}
