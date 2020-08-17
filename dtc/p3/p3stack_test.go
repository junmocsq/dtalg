package p3

import "testing"

func TestStack(t *testing.T) {
	linked := NewLinkedStack()
	stackTest(linked, t)

	arr := NewArrayStack()
	stackTest(arr, t)
}

func stackTest(stack stacker, t *testing.T) {
	for i := 100; i <= 1000; i++ {
		stack.Push(i)
	}

	for i := 1000; i >= 100; i-- {
		if res, _ := stack.Pop(); res != i {
			t.Fatal("linked push pop failed!")
		}
	}
}
