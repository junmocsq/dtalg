package p3

import "math"

type radix struct {
	arr []int
}

func NewRadix(a []int) *radix {
	return &radix{arr: a}
}

// 基数排序
// 从最低位开始排
func (r *radix) radixSort() {

	var tempArr [10][]int
	max := 0

	for _, v := range r.arr {
		index := v % 10
		tempArr[index] = append(tempArr[index], v)
		if v > max {
			max = v
		}
	}
	i := 1
	for {
		base := int(math.Pow10(i))
		var newTempArr [10][]int
		for _, v := range tempArr {
			if v == nil {
				continue
			}
			for _, _v := range v {
				index := _v / base % 10
				newTempArr[index] = append(newTempArr[index], _v)
			}
		}
		i++
		tempArr = newTempArr
		if base > max {
			r.arr = newTempArr[0]
			break
		}
	}
}
