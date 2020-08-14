package linked

import (
	"testing"
)

func TestBinarySearch(t *testing.T)  {
	t.Log((&binarySearchSquare{}).square(10))
	arr := []int{3,5,7,7,7,7,9,11,13,16,19,32}
	/**
	1 查找第一个等于给定值的元素
	2 查找最后一个等于给定值的元素
	3 查找第一个大于等于给定值的元素
	4 查找最后一个小于等于给定值的元素
	 */
	if (&binarySearchSquare{}).search(arr,7,1)!=2{
		t.Error("查找第一个等于给定值的元素7 falied")
	}
	if (&binarySearchSquare{}).search(arr,7,2)!=5{
		t.Error("查找最后一个等于给定值的元素7 falied")
	}
	if (&binarySearchSquare{}).search(arr,7,3)!=2{
		t.Error("查找第一个大于等于给定值的元素7 falied")
	}
	if (&binarySearchSquare{}).search(arr,7,4)!=5{
		t.Error("查找最后一个小于等于给定值的元素7 falied")
	}

	if (&binarySearchSquare{}).search(arr,8,3)!=6{
		t.Error("查找第一个大于等于给定值的元素7 falied")
	}
	if (&binarySearchSquare{}).search(arr,8,4)!=5{
		t.Error("查找最后一个小于等于给定值的元素7 falied")
	}
}