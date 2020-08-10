package linked

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestSort(t *testing.T) {
	length := 20
	arr := make([]int, length)
	for i := 0; i < length; i++ {
		rand.Seed(time.Now().UnixNano())
		arr[i] = rand.Intn(10000)
	}
	var arr1 = make([]int, length)
	var arr2 = make([]int, length)
	copy(arr1, arr)
	copy(arr2, arr)
	bubble := &BubbleSort{}
	bubble.Sort(arr)
	fmt.Println("bubble", arr)
	insertion := &InsertionSort{}
	insertion.Sort(arr1)
	for k, v := range arr {
		if arr1[k] != v {
			t.Error("bubbleSort or insertionSort is failed")
		}
	}
	selection := &SelectionSort{}
	selection.Sort(arr2)
	for k, v := range arr {
		if arr2[k] != v {
			t.Error("bubbleSort or selectionSort is failed")
		}
	}

	shell := &ShellSort{}
	a := []int{55, 94, 87, 1, 4, 32, 11, 77, 39, 42, 64, 53, 70, 12, 9}
	shell.Sort(a)
	if !checkSort(a) {
		t.Error("shellSort is failed")
	}

}

func checkSort(arr []int) bool {
	for i := 0; i < len(arr)-1; i++ {
		if arr[i] > arr[i+1] {
			return false
		}
	}
	return true
}
