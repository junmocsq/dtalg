package dtc

import (
	"fmt"
	"testing"
)

func TestNewLinkedList(t *testing.T) {
	list := NewLinkedList()
	list.Insert(1, list)
	list.Insert(2, list)
	list.Insert(3, list)
	list.Insert(4, list)
	list.Print()
	if list.FindPrevious(1).ele != 2 {
		t.Errorf("list.FindPrevious(1).ele is %d actual:%d ", list.FindPrevious(1).ele, 2)
	}
	if list.Find(2).ele != 2 {
		t.Errorf("list.Find(2) is %d actual:%d ", list.Find(2).ele, 2)
	}

	l := NewLinkedList()
	l.Add(1)
	l.Add(2)
	l.Add(3)
	l.Add(4)
	l.Add(5)
	l.Add(6)
	l.Print()
	l.Reverse()
	l.Print()
	l.Reverse()
	l.Print()
	if l.CheckRing()  {
		t.Log("l.CheckRing() is failed")
	}
	l.Find(6).next = l.next.next

	if !l.CheckRing()  {
		t.Log("l.CheckRing() is failed")
	}

	a := NewLinkedList()
	a.Add(10)
	a.Add(15)
	a.Add(20)
	a.Add(30)
	a.Add(40)
	a.Add(50)
	b := NewLinkedList()
	b.Add(1)
	b.Add(2)
	b.Add(13)
	b.Add(28)
	b.Add(31)
	b.Add(51)
	a.Merge(b)
	a.Print()
	fmt.Println(a.FindMiddleNode())
	a.DeleteNNodeFromLast(10)
	a.Print()
	fmt.Println(a.FindMiddleNode())

}
