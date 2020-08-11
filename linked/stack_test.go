package linked

import (
	"fmt"
	"testing"
)

func TestArrayStack(t *testing.T) {
	s := NewArrayStack()
	for i := 1; i < 100; i++ {
		s.Push(i * i)
	}
	for {
		a := s.Pop()
		if a == 0 {
			break
		}
		fmt.Printf("%d ", a)
	}
	fmt.Println()
}
