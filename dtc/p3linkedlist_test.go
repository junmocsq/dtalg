package dtc

import "testing"

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
}
