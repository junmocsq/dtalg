package iterator

// 迭代器模式
import (
	"time"
)

type Iterator interface {
	Init()
	hasNext() bool
	currentItem() interface{}
}

type iterator struct {
	vector Vector
	index  int
}

func NewIterator(c Vector) Iterator {
	return &iterator{vector: c}
}

func (i *iterator) Init() {
	i.vector.SetFindTime(time.Now().UnixNano())
	i.index = -1
}

func (i *iterator) hasNext() bool {
	if next, ok := i.vector.Next(i.index); ok {
		i.index = next
		return true
	}
	return false
}

func (i *iterator) currentItem() interface{} {
	return i.vector.ValueOf(i.index)
}
