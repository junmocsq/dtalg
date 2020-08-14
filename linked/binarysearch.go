package linked

import (
	"math"
)

type binarySearchSquare struct {
}

func (b *binarySearchSquare) square(n int) float64 {
	if n < 0 {
		return -1
	}
	start := float64(0)
	end := float64(n)
	for {
		mid := (start + end) / 2
		temp := mid * mid
		if math.Abs(temp-float64(n)) < 1e-7 {
			return float64(mid)
		} else if temp > float64(n) {
			end = mid
		} else {
			start = mid
		}
	}
}

/*
flag
	1 查找第一个等于给定值的元素
	2 查找最后一个等于给定值的元素
	3 查找第一个大于等于给定值的元素
	4 查找最后一个小于等于给定值的元素
*/
func (b *binarySearchSquare) search(arr []int, val, flag int) int {
	length := len(arr)
	start := 0
	end := length - 1
	index := -1
	for start <= end {
		mid := (start + end) / 2
		index = mid
		if arr[mid] == val {
			break
		} else if arr[mid] > val {
			end = mid - 1
		} else {
			start = mid + 1
		}
	}
	switch flag {
	case 1:
		if arr[index] != val {
			return -1
		}
		for index > 0 {
			index--
			if arr[index] != val {
				return index + 1
			}
		}
		return 0
	case 2:
		if arr[index] != val {
			return -1
		}
		for index < length-1 {
			index++
			if arr[index] != val {
				return index - 1
			}
		}
		return length - 1
	case 3:
		if arr[index] == val {
			for index > 0 {
				index--
				if arr[index] != val {
					return index + 1
				}
			}
			return 0
		} else {
			return index
		}
	case 4:
		// 查找最后一个小于等于给定值的元素
		if arr[index] == val {
			for index < length-1 {
				index++
				if arr[index] != val {
					return index - 1
				}
			}
			return length - 1
		} else {
			return index - 1
		}
	}
	return 0
}
