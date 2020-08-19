package linked

import (
	"fmt"
	"testing"
)

func TestSkipList(t *testing.T) {
	s := NewSkipList()
	s.Insert(1, 11)
	s.Insert(1, 12)
	s.Insert(2, 12)
	s.Insert(3, 13)
	s.Insert(4, 14)
	s.Insert(5, 15)
	s.Insert(6, 1)
	s.Insert(7, 2)
	s.Insert(8, 2)
	s.print()
	t.Log(s.FindByScore(5, 15))
	t.Log(s.FindByScore(14, 20))
	fmt.Println(s)
	s.DeleteByScore(4, 15)
	fmt.Println(s)
	s.print()
}
