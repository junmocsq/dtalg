package linked

import (
	"fmt"
	"math/rand"
)

// Reference https://github.com/wangzheng0822/algo/blob/master/go/17_skiplist/skiplist.go
// 跳表

type skipListNode struct {
	val      interface{}
	score    int
	level    int
	forwards []*skipListNode
}

type skipList struct {
	//跳表头结点
	head *skipListNode
	//跳表当前层数
	level int
	//跳表长度
	length int
}

const MAX_LEVEL = 16

func NewSkipList() *skipList {
	node := &skipListNode{forwards: make([]*skipListNode, MAX_LEVEL)}
	return &skipList{head: node}
}

func (s *skipList) newNode(v interface{}, score, level int) *skipListNode {
	return &skipListNode{level: level, score: score, val: v, forwards: make([]*skipListNode, MAX_LEVEL)}
}

func (s *skipList) randLevel() int {
	level := 1
	for i := 1; i < MAX_LEVEL; i++ {
		if rand.Int31()%7 == 1 {
			level++
		}
	}
	return level
}

func (s *skipList) Insert(val interface{}, score int) int {
	if nil == val {
		return 0
	}
	cur := s.head
	i := s.level - 1
	// 记录前驱节点
	update := make([]*skipListNode, MAX_LEVEL)
	for ; i >= 0; i-- {
		for cur.forwards[i] != nil {
			// 不能插入分数和数据完全一样的节点
			if cur.forwards[i].val == val && cur.forwards[i].score == score {
				return 2
			}
			if cur.forwards[i].score > score {
				break
			}
			cur = cur.forwards[i]
		}
		update[i] = cur
	}
	level := s.randLevel()
	if s.level < level {
		s.level = level
	}
	s.length++
	node := s.newNode(val, score, level)
	for i = level - 1; i >= 0; i-- {
		if update[i] == nil {
			s.head.forwards[i] = node
		} else {
			forward := update[i].forwards[i]
			update[i].forwards[i] = node
			node.forwards[i] = forward
		}
	}
	return 1
}

func (s *skipList) Find(val interface{}) *skipListNode {
	if val == nil {
		return nil
	}
	head := s.head
	for head.forwards[0] != nil {
		if head.forwards[0].val == val {
			return head.forwards[0]
		}
		head = head.forwards[0]
	}
	return nil
}

func (s *skipList) FindByScore(val interface{}, score int) *skipListNode {
	temp := s.head
	if temp == nil && val == nil {
		return nil
	}
	isFind := false
	// 查找路径
	path := make([]*skipListNode, s.level)
	for i := s.level - 1; i >= 0; i-- {
		if !isFind {
			for temp.forwards[i] != nil {
				if temp.forwards[i].score == score && temp.forwards[i].val == val {
					isFind = true
					temp = temp.forwards[i]
					break
				}
				if temp.forwards[i].score > score {
					break
				}
				temp = temp.forwards[i]
			}
		}
		path[i] = temp
	}
	for i := s.level - 1; i >= 0; i-- {
		fmt.Printf("index:%d val:%v level:%d score:%d\n", i, path[i].val, path[i].level, path[i].score)
	}
	if isFind {
		return temp
	}
	return nil
}

func (s *skipList) DeleteByScore(val interface{}, score int) int {
	temp := s.head
	if temp == nil && val == nil {
		return 0
	}
	isFind := false
	// 记录前驱路径
	path := make([]*skipListNode, s.level)
	for i := s.level - 1; i >= 0; i-- {
		for temp.forwards[i] != nil {
			if temp.forwards[i].score == score && temp.forwards[i].val == val {
				isFind = true
				break
			}
			temp = temp.forwards[i]
		}
		path[i] = temp
	}
	if isFind {
		deleteNode := path[0].forwards[0]
		for i := deleteNode.level - 1; i >= 0; i-- {
			path[i].forwards[i] = deleteNode.forwards[i]
		}
		if s.head.forwards[s.level] == nil {
			s.level--
		}
		s.length--
		return 1
	}
	return 0
}

func (s *skipList) print() {
	head := s.head
	for head != nil {
		fmt.Printf("%p %v\n", head, head)
		head = head.forwards[0]
	}
}
