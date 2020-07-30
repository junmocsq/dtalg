package linked

import (
	"errors"
	"fmt"
)

// 指针或引用的含义、警惕指针丢失和内存泄漏、利用哨兵简化实现难度、重点留意边界条件处理，以及举例画图、辅助思考，还有多写多练
type single struct {
	val  interface{}
	next *single
}

var ERROR_SINGLE_UNSORTED error = errors.New("singly linked list is not sorted")

type SingleLinkedLister interface {
	Compare(a, b interface{}) (int, error)
	Tail() *single
	AddToTail(val interface{}) *single
	CheckSort() int
	AddSort(val interface{}) (*single, error) // 添加数据到有序链表
	DeleteByValue(val interface{})            // 根据值删除
	Find()                                    // 根据值查找
	Update()                                  //
	Reverse()                                 // 链表反转
	CheckLoop()                               // 寻找环
	MergeSort(single *single)                 // 有序链表合并
	DeleteByIndex(index int)                  // 按序号删除
	Middle()                                  // 找出链表中间单元
}

func SingleInit() *single {
	return &single{}
}

// 大于为1 小于为-1 相等为0
func (s *single) Compare(a, b interface{}) (int, error) {
	return Compare(a, b)
}

func (s *single) Tail() *single {
	temp := s
	for temp.next != nil {
		temp = temp.next
	}
	return temp
}

func (s *single) AddToTail(val interface{}) *single {
	node := &single{val: val}
	tail := s.Tail()
	tail.next = node
	return node
}

// 升序为1 降序为-1 无序为0
func (s *single) CheckSort() int {
	temp := s
	res := 1
	flag := 0

	for {
		if temp.next == nil || temp.next.next == nil {
			return res
		}
		compareRes, err := s.Compare(temp.next.val, temp.next.next.val)
		if err != nil {
			return 0
		}
		if flag == 0 {
			if compareRes == 1 {
				res = -1
				flag = 1
			} else if compareRes == 0 {

			} else {
				res = 1
				flag = 1
			}
		}
		// 升序
		if res == 1 {
			if compareRes != 1 {
				temp = temp.next
			} else {
				return 0
			}

		}
		if res == -1 {
			if compareRes != -1 {
				temp = temp.next
			} else {
				return 0
			}
		}
	}
}

func (s *single) AddSort(val interface{}) (*single, error) {
	sortType := s.CheckSort()
	if sortType == 0 {
		return nil, ERROR_SINGLE_UNSORTED
	}
	node := &single{val: val}
	return node, nil
}

func (s *single) Print() {
	temp := s
	for temp.next != nil {
		temp = temp.next
		fmt.Printf("%v ", temp.val)
	}
	fmt.Println()
}
