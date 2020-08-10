package linked

type Sorter interface {
	Sort([]int)
}

type BubbleSort struct {
}
// 冒泡排序
// 对比相邻两个元素交换数据，稳定排序
func (b *BubbleSort) Sort(arr []int) {
	length := len(arr)
	for i := length; i > 0; i-- {
		for j := 1; j < i; j++ {
			if arr[j] < arr[j-1] {
				arr[j], arr[j-1] = arr[j-1], arr[j]
			}
		}
	}
}

type InsertionSort struct {
}
// 插入排序 性能优于冒泡排序
// 稳定算法，不会改变相等元素的相对顺序
func (i *InsertionSort) Sort(arr []int) {
	length := len(arr)
	for i := 1; i < length; i++ {
		temp := arr[i]
		// 需要插入的位置
		index := i
		for j := i - 1; j >= 0; j-- {
			if temp < arr[j] {
				arr[j+1] = arr[j]
				index = j
			} else {
				break
			}
		}
		arr[index] = temp
	}
}

type SelectionSort struct {
}

// 选择排序，每一次选择剩下的最小的元素来交换
// 非稳定算法
func (s *SelectionSort) Sort(arr []int) {
	length := len(arr)
	for i := 0; i < length-1; i++ {
		// 选择arr[i+1:]最小值
		min := i + 1
		for j := i + 2; j < length; j++ {
			if arr[j] < arr[min] {
				min = j
			}
		}
		// 最小值和arr[i]比较 交换
		if arr[i] > arr[min] {
			arr[i], arr[min] = arr[min], arr[i]
		}
	}
}

type ShellSort struct {
}

// 希尔排序 非稳定排序【当步长大于1时】
// 希尔排序通过将比较的全部元素分为几个区域来提升插入排序的性能。这样可以让一个元素可以一次性地朝最终位置前进一大步。然后算法再取越来越小的步长进行排序，算法的最后一步就是普通的插入排序，但是到了这步，需排序的数据几乎是已排好的了（此时插入排序较快）
func (s *ShellSort) Sort(arr []int) {
	n := len(arr)
	if n < 2 {
		return
	}

	key := n / 2
	for key > 0 {
		for i := key; i < n; i++ {
			j := i
			for j >= key && arr[j] < arr[j-key] {
				arr[j], arr[j-key] = arr[j-key], arr[j]
				j = j - key
			}
		}
		key = key / 2
	}


	/********************************************************
	当key为1时，就是插入排序
		for i := 1; i < n; i++ {
			j := i
			for j >= 1 && arr[j] < arr[j-1] {
				arr[j], arr[j-1] = arr[j-1], arr[j]
				j = j - 1
			}
		}
		fmt.Println(1,arr)
	*******************************************************/


}
