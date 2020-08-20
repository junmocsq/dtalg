package iterator

import (
	"time"
)

type Vector interface {
	Add(val interface{}) (int, bool)
	Remove(i int) (interface{}, bool)
	ValueOf(i int) interface{}
	Length() int
	Next(int) (int, bool)
	SetFindTime(int64)
}

// 采用删除时间和添加时间标记元素，实现支持按时间的遍历的容器
type vector struct {
	arr        []interface{}
	addTimeArr []int64
	delTimeArr []int64
	findTime   int64
}

func NewVector() Vector {
	return &vector{}
}

func (v *vector) Add(val interface{}) (int, bool) {
	if val == nil {
		return 0, false
	}
	currTime := time.Now().UnixNano()
	v.arr = append(v.arr, val)
	v.addTimeArr = append(v.addTimeArr, currTime)
	v.delTimeArr = append(v.delTimeArr, 0)
	return len(v.arr) - 1, true
}

func (v *vector) Remove(i int) (interface{}, bool) {
	if i >= len(v.arr) {
		return 0, false
	}
	currTime := time.Now().UnixNano()
	v.delTimeArr[i] = currTime
	return v.arr[i], true
}

func (v *vector) ValueOf(i int) interface{} {
	if i >= len(v.arr) || i < 0 {
		return nil
	}
	val := v.arr[i]
	delTime := v.delTimeArr[i]
	addTime := v.addTimeArr[i]
	if addTime > v.findTime || delTime != 0 && delTime < v.findTime {
		return nil
	}
	return val
}

func (v *vector) Length() int {
	length := 0
	for i := 0; i < len(v.arr); i++ {
		if v.ValueOf(i) != nil {
			length++
		}
	}
	return length
}

func (v *vector) Next(i int) (int, bool) {
	i++
	for ; i < len(v.arr); i++ {
		if v.ValueOf(i) != nil {
			return i, true
		}
	}
	return 0, false
}

func (v *vector) SetFindTime(t int64) {
	v.findTime = t
}
