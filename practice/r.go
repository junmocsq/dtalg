package practice

import "fmt"

var index = 0

func R() {
	arr := []int{1, 2, 5, 10, 20, 50, 100}
	total := 100
	res := make([]int, len(arr))
	foreach(arr, res, total, 0)
	//fmt.Println(index)
}

func r(arr, res []int, total int, index int) {
	if index == len(arr) {
		if total == 0 {
			fmt.Println(res)
		}
		return
	}
	val := arr[index]
	for m := 0; m*val <= total; m++ {
		temp := make([]int, len(arr))
		copy(temp, res)
		temp[index] = m
		r(arr, temp, total-m*val, index+1)
	}
}

func foreach(arr, res []int, total int, index int) {
	s := &stack{}
	s.push(res, total, index)
	for !s.isEmpty() {
		node := s.pop()
		index = node.index
		total = node.total
		res = node.prefixArr
		if index == len(arr) {
			if total == 0 {
				fmt.Println(node.prefixArr)
			}
			continue
		}
		val := arr[index]
		for m := 0; m*val <= total; m++ {
			temp := make([]int, len(arr))
			copy(temp, res)
			temp[index] = m
			s.push(temp, total-m*val, index+1)
		}
	}
}

type node struct {
	total     int
	index     int
	prefixArr []int
}
type stack struct {
	arr []*node
}

func (n *stack) isEmpty() bool {
	return len(n.arr) == 0
}
func (n *stack) push(prefixArr []int, total, index int) {
	n.arr = append(n.arr, &node{total: total, prefixArr: prefixArr, index: index})
}

func (n *stack) pop() *node {
	length := len(n.arr)
	node := n.arr[length-1]
	n.arr = n.arr[:length-1]
	return node
}

func fib(n int) {

}
