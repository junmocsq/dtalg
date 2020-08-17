package p3

import "fmt"

type cursor interface {
	IsEmpty()
	IsLast()
}

var cursorSpace [100]cursorNode

type cursorNode struct {
	ele  string
	next int
}

func initCursor() {
	for i := 0; i < 100; i++ {
		cursorSpace[i] = cursorNode{next: (i + 1) % 100}
	}
}

func cursorAlloc() int {
	position := cursorSpace[0].next
	cursorSpace[0].next = cursorSpace[position].next
	cursorSpace[position].next = 0
	return position
}

func cursorFree(position int) {
	cursorSpace[position].next = cursorSpace[0].next
	cursorSpace[0].next = position
}

type cursorList int

// 带哨兵的游标实现
func NewCursor() cursorList {
	p := cursorAlloc()
	return cursorList(p)
}
func (list cursorList) isEmpty() bool {
	return cursorSpace[list].next == 0
}

func (list cursorList) add(ele string) {
	p := cursorAlloc()
	cursorSpace[p].ele = ele
	temp := cursorSpace[list].next
	if temp == 0 {
		cursorSpace[list].next = p
		return
	}
	for cursorSpace[temp].next != 0 {
		temp = cursorSpace[temp].next
	}
	cursorSpace[temp].next = p
}

func (list cursorList) delete(position int) {
	p := int(list)
	for cursorSpace[p].next != position {
		p = cursorSpace[p].next
	}
	cursorSpace[p].next = cursorSpace[position].next
	cursorFree(position)
}

func (list cursorList) print() {
	position := cursorSpace[list].next
	for position != 0 {
		fmt.Printf("position:%d ele:%s ", position, cursorSpace[position].ele)
		position = cursorSpace[position].next
	}
	fmt.Println()
}
