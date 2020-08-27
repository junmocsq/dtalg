package p22

import (
	"errors"
	"fmt"
	"log"
)

const (
	WHITE = 1
	GRAY  = 2
	BLACK = 2
)

type bdfsNode struct {
	color  int
	id     int
	name   string
	weight int
	prev   *bdfsNode
}
type Bdfs struct {
	arr         [][]*bdfsNode
	index       int
	nameNodeMap map[string]*bdfsNode
	idNameMap   map[int]string
	stack       []*bdfsNode
	queue       []*bdfsNode
}

func (b *Bdfs) init() {
	b.index = 0
	b.arr = make([][]*bdfsNode, 0)
	b.nameNodeMap = make(map[string]*bdfsNode)
	b.idNameMap = make(map[int]string)
}

func (b *Bdfs) getId() int {
	b.index++
	return b.index - 1
}

func (b *Bdfs) getNodeByName(name string) *bdfsNode {
	if node, ok := b.nameNodeMap[name]; ok {
		return node
	} else {
		return b.newNode(name)
	}
}
func (b *Bdfs) getNameById(id int) string {
	return b.idNameMap[id]
}

func NewBdfs() *Bdfs {
	return &Bdfs{
		arr:         make([][]*bdfsNode, 0),
		nameNodeMap: make(map[string]*bdfsNode),
		idNameMap:   make(map[int]string),
	}
}
func (b *Bdfs) newNode(v1 string) *bdfsNode {
	id := b.getId()
	node := &bdfsNode{
		id:     id,
		name:   v1,
		color:  WHITE,
		weight: 99999,
	}
	b.nameNodeMap[v1] = node
	return node
}
func (b *Bdfs) addEdge(v1, v2 string) {
	length := len(b.arr)
	id := b.getNodeByName(v1).id
	for i := length; i <= id; i++ {
		b.arr = append(b.arr, nil)
	}
	for _, v := range b.arr[id] {
		if v.name == v2 {
			return
		}
	}
	b.arr[id] = append(b.arr[id], b.getNodeByName(v2))
}

func (b *Bdfs) AddEdge(v1, v2 string) {
	b.addEdge(v1, v2)
	b.addEdge(v2, v1)
}

func (b *Bdfs) CreateGraph() {
	b.arr = make([][]*bdfsNode, 0)
	b.AddEdge("v", "r")
	b.AddEdge("r", "s")
	b.AddEdge("s", "w")
	b.AddEdge("w", "t")
	b.AddEdge("w", "x")
	b.AddEdge("x", "t")
	b.AddEdge("x", "u")
	b.AddEdge("x", "y")
	b.AddEdge("t", "u")
	b.AddEdge("u", "y")
}

func (b *Bdfs) Bfs(s string) {
	b.init()
	b.queueInt()
	b.CreateGraph()
	src := b.nameNodeMap[s]
	if src == nil {
		log.Fatal(s + " node is not exists!")
	}
	src.color = GRAY
	src.weight = 0
	b.queuePush(src)
	for !b.queueIsEmpty() {
		node, err := b.queuePop()
		if err != nil {
			return
		}
		node.color = BLACK
		id := node.id
		for _, v := range b.arr[id] {
			if v.color != WHITE {
				continue
			}
			v.color = GRAY
			v.weight = node.weight + 1
			v.prev = node
			b.queuePush(v)
		}
	}

	// b.nameNodeMap 广度优先树
	for k, v := range b.nameNodeMap {
		if v.prev != nil {
			fmt.Println(k, v, v.prev.name)
		} else {
			fmt.Println(k, v)
		}

	}
}

func (b *Bdfs) printPath(s, v string) {
	if v == s {
		fmt.Println(s)
	} else if b.getNodeByName(v).prev == nil {
		fmt.Println("no path from " + s + " to " + v)
	} else {
		b.printPath(s, b.getNodeByName(v).prev.name)
		fmt.Println(v)
	}
}

func (b *Bdfs) stackInt() {
	b.stack = make([]*bdfsNode, 0)
}

func (b *Bdfs) stackIsEmpty() bool {
	return len(b.stack) == 0
}

func (b *Bdfs) stackPush(v *bdfsNode) {
	b.stack = append(b.stack, v)
}

func (b *Bdfs) stackPop() (*bdfsNode, error) {
	if b.stackIsEmpty() {
		return nil, errors.New("stack is empty")
	}
	length := len(b.stack)
	name := b.stack[length-1]
	b.stack = b.stack[:length-1]
	return name, nil
}

func (b *Bdfs) queueInt() {
	b.queue = make([]*bdfsNode, 0)
}

func (b *Bdfs) queueIsEmpty() bool {
	return len(b.queue) == 0
}

func (b *Bdfs) queuePush(v *bdfsNode) {
	b.queue = append(b.queue, v)
}

func (b *Bdfs) queuePop() (*bdfsNode, error) {
	if b.queueIsEmpty() {
		return nil, errors.New("queue is empty")
	}
	name := b.queue[0]
	b.queue = b.queue[1:]
	return name, nil
}
