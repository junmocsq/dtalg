package practice

import "testing"

func TestFib(t *testing.T) {
	f := &fibStruct{}
	f.fib(10)

	fact := &factStruct{}
	fact.fact(3)
}
