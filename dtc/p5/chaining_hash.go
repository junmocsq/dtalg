package p5

import "fmt"

// 拉链法
type listNode struct {
	key  string
	next *listNode
	val  interface{}
}

type HashTable struct {
	size int
	list []*listNode
	hash HashFunc
}

func NewChainingHashTable(size int, hashFunc HashFunc) *HashTable {
	return &HashTable{
		size: size,
		list: make([]*listNode, size),
		hash: hashFunc,
	}
}

func (h *HashTable) Find(key string) *listNode {
	head := h.getHashSlotHead(key)
	for head != nil {
		if head.key == key {
			return head
		}
		head = head.next
	}
	return nil
}

func (h *HashTable) Insert(key string, val interface{}) {
	p := h.Find(key)
	// 已经存在则替换
	if p != nil {
		p.val = val
		return
	}
	head := h.getHashSlotHead(key)
	if head == nil {
		h.list[h.getHashSlot(key)] = &listNode{key: key, val: val}
	} else {
		head.next = &listNode{key: key, val: val, next: head.next}
	}

}

func (h *HashTable) Dkeyte(key string) {
	head := h.getHashSlotHead(key)
	for head.next != nil {
		if head.next.key == key {
			head.next = head.next.next
			break
		}
	}
}

func (h *HashTable) print() {
	for k, v := range h.list {
		if v != nil {
			fmt.Printf("%d:\t", k)
			for v != nil {
				fmt.Printf("%s:%v ", v.key, v.val)
				v = v.next
			}
			fmt.Println()
		}
	}
}

func (h *HashTable) getHashSlotHead(key string) *listNode {
	return h.list[h.getHashSlot(key)]
}

func (h *HashTable) getHashSlot(key string) int {
	return h.hash.Hash(key, h.size)
}
