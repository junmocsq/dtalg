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

	for i:=len(arr)-1;i>=0;i-- {
		index := temp[arr[i]-min]-1
		r[index] = arr[i]
		temp[arr[i]-min]--
	}
	fmt.Println(temp)
	for k,_ := range r{
		arr[k] = r[k]
	}
}
