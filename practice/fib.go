package practice

import (
	"fmt"
)

type fibNode struct {
	isBack int
	n      int
}

type fibStruct struct {
	arr []*fibNode
}

func (f *fibStruct) fib(n int) {
	f.push(n, 0)
	sum := 0
	for !f.isEmpty() {
		//time.Sleep(time.Second)
		//fmt.Println(f.arr)
		node := f.pop()
		if node.n <= 2 {
			sum += 1
			continue
		}
		f.push(node.n-1, 0)
		f.push(node.n-2, 0)
		//f.push(node.n-3,sum)
	}
	fmt.Println(sum)
}

func (f *fibStruct) push(n, isBack int) {
	f.arr = append(f.arr, &fibNode{isBack: isBack, n: n})
}

func (f *fibStruct) pop() *fibNode {
	length := len(f.arr)
	node := f.arr[length-1]
	f.arr = f.arr[:length-1]
	return node
}

func (f *fibStruct) isEmpty() bool {
	return len(f.arr) == 0
}
