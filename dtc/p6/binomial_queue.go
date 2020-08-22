package p6

import (
	"errors"
	"math"
)

// 二项式队列：不是一棵堆序的树，而是堆序树的集合，称为森林。
// 二项树的每一个节点将包含数据、第一个儿子以及第一个儿子的右兄弟，二项树种诸儿子以递减次序排列

type binNode struct {
	ele         int
	left        *binNode
	nextSibling *binNode
}

type binTree binNode

type binQueue struct {
	size  int
	trees []*binTree
}

func newBinQueue(cap int) *binQueue {
	return &binQueue{trees: make([]*binTree, cap)}
}

// 合并两个相同的二项树
func (q *binQueue) combineTrees(t1, t2 *binTree) *binTree {
	if t1.ele > t2.ele {
		return q.combineTrees(t2, t1)
	}
	t2.nextSibling = t1.left
	t1.left = (*binNode)(t2)
	return t1
}

func (q *binQueue) merge(h2 *binQueue) error {
	var t1, t2, carry *binTree

	if math.Pow(2, float64(len(q.trees)))-1 <= float64(q.size+h2.size) {
		return errors.New("二项式队列已满")
	}
	q.size += h2.size
	h2.size = 0
	for i, j := 0, 1; j <= q.size; i, j = i+1, j*2 {
		t1 = q.trees[i]
		t2 = h2.trees[i]
		a1, a2, a3 := 0, 0, 0
		if t1 != nil {
			a1 = 1
		}
		if t2 != nil {
			a2 = 1
		}
		if carry != nil {
			a3 = 1
		}

		switch a1 + 2*a2 + 4*a3 {
		case 0, 1:
		case 2:
			q.trees[i] = t2
			h2.trees[i] = nil
		case 3:
			carry = q.combineTrees(t1, t2)
			q.trees[i] = nil
			h2.trees[i] = nil
		case 4:
			q.trees[i] = carry
			h2.trees[i] = nil
			carry = nil
		case 5:
			carry = q.combineTrees(t1, carry)
			q.trees[i] = nil
		case 6:
			carry = q.combineTrees(t2, carry)
			h2.trees[i] = nil
		case 7:
			carry = q.combineTrees(t2, carry)
			h2.trees[i] = nil
		}
	}
	return nil
}
func (q *binQueue) Insert(ele int) {
	h2 := newBinQueue(len(q.trees))
	h2.size = 1
	h2.trees[0] = &binTree{ele: ele}
	q.merge(h2)
}
func (q *binQueue) DeleteMin() (*binNode, error) {
	if q.isEmpty() {
		return nil, errors.New("二项式队列为空")
	}
	var delNode *binNode
	var index int
	for i, j := 0, 1; j <= q.size; i, j = i+1, j*2 {
		if q.trees[i] != nil {
			top := (*binNode)(q.trees[i])
			if delNode == nil || delNode.ele > top.ele {
				delNode = top
				index = i
			}
		}
	}
	h2Size := int(math.Pow(2, float64(index)))
	q.size -= h2Size
	q.trees[index] = nil
	h2 := newBinQueue(len(q.trees))
	temp := delNode.left
	for i := index - 1; i >= 0; i-- {
		h2.trees[i] = (*binTree)(temp)
		temp = temp.nextSibling
		h2.trees[i].nextSibling = nil
	}
	h2.size = h2Size - 1
	q.merge(h2)
	delNode.left = nil
	delNode.nextSibling = nil
	return delNode, nil

}

func (q *binQueue) isEmpty() bool {
	for i, j := 0, 1; j <= q.size; i, j = i+1, j*2 {
		if q.trees[i] != nil {
			return false
		}
	}
	return true
}
