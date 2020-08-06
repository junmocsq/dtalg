package linked

import "sync"

type ArrayStacker interface {
	Push()
	Pop()
}

type arrayStack struct {
	stack []interface{}
	top   int
	cap   int
	lock  sync.Mutex
}

func NewArrayStack() *arrayStack {
	s := &arrayStack{
		stack: make([]interface{}, 8),
		cap:   0,
		top:   -1,
	}
	s.cap = len(s.stack)
	return s
}

func (s *arrayStack) Push(ele interface{}) {
	s.lock.Lock()
	defer s.lock.Unlock()
	s.top++
	if s.top >= s.cap {
		// 扩容
		newStack := make([]interface{}, s.top*2)
		s.cap = s.top * 2
		for i := 0; i < s.top; i++ {
			newStack[i] = s.stack[i]
		}
		s.stack = newStack
	}
	s.stack[s.top] = ele
}

func (s *arrayStack) Pop() interface{} {
	s.lock.Lock()
	defer s.lock.Unlock()
	if s.top == -1 {
		return 0
	}
	ele := s.stack[s.top]
	s.top--
	return ele
}

