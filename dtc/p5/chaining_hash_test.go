package p5

import "testing"

func TestNewChainingHashTable(t *testing.T) {
	h := NewChainingHashTable(10007, NewHorner())
	h.Insert("lxe", "csq")
	h.Insert("zsf", "lxq")
	h.Insert("zhangsan", "mm")
	h.print()
}
