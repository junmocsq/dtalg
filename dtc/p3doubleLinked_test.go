package dtc

import "testing"

func TestNewDoubleLinkedList(t *testing.T) {
	doubleLinked := NewDoubleLinkedList()
	doubleLinked.Add(1)
	doubleLinked.Add(2)
	doubleLinked.Add(3)
	doubleLinked.Add(4)
	doubleLinked.Add(5)
	if doubleLinked.IsEmpty() != false {
		t.Error("doubleLinkedList IsEmpty() failed")
	}
	end := doubleLinked.Find(5)
	if end.ele != 5 {
		t.Error("doubleLinkedList Find(5) failed")
	}
	if doubleLinked.IsLast(end) == false {
		t.Error("doubleLinkedList IsLast(end) failed")
	}

	if !doubleLinked.Delete(5) {
		t.Error("doubleLinkedList Delete(5) failed")
	}

	if !doubleLinked.Delete(2) {
		t.Error("doubleLinkedList Delete(2) failed")
	}
	if doubleLinked.IsLast(doubleLinked.Find(4)) == false {
		t.Error("doubleLinkedList IsLast(end) failed")
	}
}
