package p2

import "fmt"

// 逆序对
type inversion struct {
	num int
}

func NewInversion() *inversion {
	return &inversion{}
}

func (i *inversion) iv(arr []int) int {
	res := i.mergeSort(arr, 0, len(arr)-1)
	for k, v := range res {
		arr[k] = v
	}
	return i.num
}

func (i *inversion) mergeSort(arr []int, start, end int) []int {
	if end < start {
		return []int{}
	}
	if end == start {
		return arr[start : end+1]
	}
	mid := (start + end) / 2
	res := i.merge(i.mergeSort(arr, start, mid), i.mergeSort(arr, mid+1, end))

	return res
}

func (i *inversion) merge(arr1, arr2 []int) []int {
	fmt.Println(arr1, arr2)
	res := make([]int, 0)

	l1 := len(arr1)
	l2 := len(arr2)
	if l1 == 0 {
		return arr2
	}
	if l2 == 0 {
		return arr1
	}
	var m, n int
	for {

		if arr1[m] > arr2[n] {
			res = append(res, arr2[n])
			// 此种情况 arr1[m]之后的元素都大于arr2[n],arr1[m]和之后的元素的个数就是本次需要添加的逆序对
			i.num += l1 - m
			n++
			if n == l2 {
				res = append(res, arr1[m:]...)
				break
			}
		} else {
			res = append(res, arr1[m])
			m++
			if m == l1 {
				res = append(res, arr2[n:]...)
				break
			}
		}
	}
	fmt.Println("==", res)
	return res
}
