package p3

import "fmt"

type p3practice struct {
}

func NewP3Practice() *p3practice {
	return &p3practice{}
}

func (p *p3practice) p10_1(m, n int) {
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		arr[i] = 1
	}
	index := 0
	M := 0
	size := n
	for {
		if arr[index] == 1 {
			if M == m {
				arr[index] = 0
				M = 0
				fmt.Println(size, index+1)
				size--
				if size == 0 {
					break
				}
			} else {
				M++
			}
		}
		index++
		if index == n {
			index = 0
		}
	}
}

// 一个数组实现里欧昂个栈，可以从一个数组的两端向中间生长，
