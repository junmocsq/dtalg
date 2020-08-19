package p3

import (
	"fmt"
	"testing"
)

func TestCursor(t *testing.T) {
	initCursor()
	list := NewCursor()
	fmt.Println(list)
	list.add("a")
	list.add("b")
	list.add("c")
	list.add("d")
	list.add("e")
	list.add("f")
	list.print()
	list.delete(6)
	list.add("g")
	list.print()
}
