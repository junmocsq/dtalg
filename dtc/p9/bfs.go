package p9

import (
	"errors"
	"fmt"
)

// 广度优先搜索（Breadth-First-Search）

type bdfs struct {
	nameMap map[string]int
	idMap   map[int]string
	arr     [][]string
	index   int
	stack   []string
	queue   []string
}

func NewBdfs() *bdfs {
	return &bdfs{
		nameMap: make(map[string]int),
		idMap:   make(map[int]string),
		arr:     make([][]string, 0),
		index:   0,
		stack:   make([]string, 0),
	}
}

func (b *bdfs) getId() int {
	b.index++
	return b.index - 1
}
func (b *bdfs) getIdByName(name string) int {
	if id, ok := b.nameMap[name]; ok {
		return id
	} else {
		id = b.getId()
		b.nameMap[name] = id
		b.idMap[id] = name
		return b.nameMap[name]
	}
}

func (b *bdfs) getNameById(id int) string {
	return b.idMap[id]
}

func (b *bdfs) addEdge(v1, v2 string) {
	headIndex := b.getIdByName(v1)
	for i := 0; i <= headIndex-len(b.arr); i++ {
		b.arr = append(b.arr, nil)
	}
	flag := false
	for _, v := range b.arr[headIndex] {
		if v == v2 {
			flag = true
			break
		}
	}
	if !flag {
		b.arr[headIndex] = append(b.arr[headIndex], v2)
	}
}

func (b *bdfs) AddEdge(v1, v2 string) {
	b.addEdge(v1, v2)
	b.addEdge(v2, v1)
}

func (b *bdfs) CreateGraph() {
	b.AddEdge("v0", "v1")
	b.AddEdge("v1", "v2")
	b.AddEdge("v0", "v3")
	b.AddEdge("v1", "v4")
	b.AddEdge("v2", "v5")
	b.AddEdge("v3", "v4")
	b.AddEdge("v4", "v5")
	b.AddEdge("v4", "v6")
	b.AddEdge("v5", "v7")
	b.AddEdge("v6", "v7")
}

func (b *bdfs) print() {
	for k, v := range b.arr {
		name := b.getNameById(k)
		fmt.Println(name, v)
	}
}

// 广度优先算法
func (b *bdfs) bfs(v1, v2 string) {
	// 队列
	b.queueInit()
	b.enqueue(v1)
	// 搜索路径
	path := make([]string, 0)
	findEleMap := make(map[string]bool)
	for !b.queueIsEmpty() {
		name, err := b.dequeue()
		if err != nil {
			fmt.Println(err)
			break
		}
		if !findEleMap[name] {
			path = append(path, name)
			if name == v2 {
				for _, v := range path {
					fmt.Println("++", v)
				}
				return
			}
		}
		findEleMap[name] = true
		for _, v := range b.arr[b.getIdByName(name)] {
			if !findEleMap[v] {
				b.enqueue(v)
			}
		}
	}
}

// 深度优先算法
func (b *bdfs) dfs(v1, v2 string) {
	b.stackInit()
	b.push(v1)
	nameSearchMap := make(map[string]bool)
	path := make([]string, 0)
	for !b.stackIsEmpty() {
		name, err := b.pop()
		if err != nil {
			fmt.Println(err)
			return
		}
		if !nameSearchMap[name] {
			nameSearchMap[name] = true
			path = append(path, name)
			//fmt.Println("==", name)
			if name == v2 {
				for _, v := range path {
					fmt.Println("==", v)
				}
				return
			}
			for _, v := range b.arr[b.getIdByName(name)] {
				if !nameSearchMap[v] {
					b.push(v)
				}
			}
		}

	}
}

func (b *bdfs) stackInit() {
	b.stack = make([]string, 0)
}
func (b *bdfs) stackIsEmpty() bool {
	return len(b.stack) == 0
}
func (b *bdfs) push(name string) {
	b.stack = append(b.stack, name)
}
func (b *bdfs) pop() (string, error) {
	if len(b.stack) == 0 {
		return "", errors.New("栈为空")
	}
	length := len(b.stack)
	name := b.stack[length-1]
	b.stack = b.stack[:length-1]
	return name, nil
}

func (b *bdfs) queueInit() {
	b.queue = make([]string, 0)
}
func (b *bdfs) queueIsEmpty() bool {
	return len(b.queue) == 0
}
func (b *bdfs) enqueue(name string) {
	b.queue = append(b.queue, name)
}
func (b *bdfs) dequeue() (string, error) {
	if len(b.queue) == 0 {
		return "", errors.New("队列为空")
	}
	name := b.queue[0]
	b.queue = b.queue[1:]
	return name, nil
}
