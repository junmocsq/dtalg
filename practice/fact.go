package practice

import "fmt"

func fact(n int) int {
	if n <= 1 {
		return 1
	}
	return n * fact(n-1)
}

type factNode struct {
	index  int
	sum    *int
	isBack int
}

type factStruct struct {
	arr []*factNode
}

func (f *factStruct) fact(n int) {
	var sum int
	f.push(&factNode{index: n, sum: &sum})

	for !f.isEmpty() {

		node := f.pop()
		fmt.Println(node.index, *node.sum)
		if node.index == 1 {
			sum = 1
		} else if node.isBack == 0 {
			f.push(&factNode{index: node.index, sum: node.sum, isBack: 1})
			f.push(&factNode{index: node.index - 1, sum: node.sum, isBack: 0})
		} else {
			sum *= node.index
		}
	}
	fmt.Println(sum)
}

func (f *factStruct) isEmpty() bool {
	return len(f.arr) == 0
}
func (f *factStruct) push(node2 *factNode) {
	f.arr = append(f.arr, node2)
}
func (f *factStruct) pop() *factNode {
	length := len(f.arr)
	node := f.arr[length-1]
	f.arr = f.arr[:length-1]
	return node
}
