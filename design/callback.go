package design

import "fmt"

type callback interface {
	callback()
}

type B struct {
}

func (b *B) process(call callback) {
	fmt.Println("callback start:")
	call.callback()
	fmt.Println("callback end;")
}

type A struct {
}

func (a *A) pay(isSuccess bool) {
	// doSomething...
	b := &B{}
	if isSuccess {
		b.process(&success{})
	} else {
		b.process(&failed{})
	}
}

type success struct {
}

func (s *success) callback() {
	fmt.Println("pay success!")
}

type failed struct {
}

func (f *failed) callback() {
	fmt.Println("pay failed!")
}
