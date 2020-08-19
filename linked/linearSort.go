package linked

import "fmt"

type BucketSort struct {
}

type CountingSort struct {
}

func (c *CountingSort) Sort(arr []int) {
	min := 0
	max := 0

	for _, v := range arr {
		if v > max {
			max = v
		}
		if v < min {
			min = v
		}
	}

	temp := make([]int, max-min+1)

	for _, v := range arr {
		temp[v-min]++
	}

	for k, _ := range temp {
		if k == 0 {
			continue
		}
		temp[k] = temp[k] + temp[k-1]
	}
	fmt.Println(temp)

	r := make([]int, len(arr))

	for i := len(arr) - 1; i >= 0; i-- {
		index := temp[arr[i]-min] - 1
		r[index] = arr[i]
		temp[arr[i]-min]--
	}
	fmt.Println(temp)
	for k, _ := range r {
		arr[k] = r[k]
	}
}

type NumberCharacterCategory struct {
}

// 按数字小写字母大写字母分开
func (c *NumberCharacterCategory) Sort(arr []byte) {
	fmt.Println(string(arr))

	index := 0
	for k, v := range arr {
		if v >= '0' && v <= '9' {
			arr[index], arr[k] = arr[k], arr[index]
			index++
		}
	}

	for i := index; i < len(arr); i++ {
		if arr[i] >= 'a' && arr[i] <= 'z' {
			arr[index], arr[i] = arr[i], arr[index]
			index++
		}
	}
	fmt.Println(string(arr))
}
