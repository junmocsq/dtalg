package p22

import (
	"errors"
	"fmt"
	"log"
)

const (
	WHITE = 0
	GRAY  = 1
	BLACK = 2
)

type bdfsNode struct {
	color  int
	id     int
	name   string
	weight int
	prev   *bdfsNode
	time   int
	ftime  int
}

type stackDfsNode struct {
	node *bdfsNode
	edge []*bdfsNode
}
type Bdfs struct {
	arr         [][]*bdfsNode
	index       int
	nameNodeMap map[string]*bdfsNode
	idNameMap   map[int]string
	stack       []*bdfsNode
	stack2      []*stackDfsNode
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
	b.idNameMap[id] = v1
	b.nameNodeMap[v1] = node
	return node
}
func (b *Bdfs) addEdge(v1, v2 string) {
	length := len(b.arr)
	id := b.getNodeByName(v1).id
	for i := length; i <= id; i++ {
		b.arr = append(b.arr, nil)
	}
	if v2 == "" {
		return
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
	b.init()
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
	b.queueInt()
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
		//node.color = BLACK
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

func (b *Bdfs) CreateDigraph() {
	b.init()
	b.arr = make([][]*bdfsNode, 0)
	for _, v := range []string{"q", "r", "s", "t", "u", "v", "w", "x", "y", "z"} {
		b.getNodeByName(v)
	}

	b.addEdge("q", "s")
	b.addEdge("q", "w")
	b.addEdge("s", "v")
	b.addEdge("v", "w")
	b.addEdge("w", "s")

	b.addEdge("q", "t")
	b.addEdge("t", "x")
	b.addEdge("x", "z")
	b.addEdge("z", "x")
	b.addEdge("t", "y")
	b.addEdge("y", "q")
	b.addEdge("r", "y")
	b.addEdge("r", "u")
	b.addEdge("u", "y")
}

func (b *Bdfs) Dfs() {

	b.stackInt()
	b.stack2Int()
	time := 0
	for i := 0; i < len(b.idNameMap); i++ {
		v := b.getNodeByName(b.getNameById(i))
		if v.color == WHITE {
			v.weight = 0
			//b.dfs_visit(v,&time)
			b.dfs_visit_stack2(v, &time)
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

// 采用栈来模拟
func (b *Bdfs) dfs_visit_stack(node *bdfsNode, time *int) {
	node.color = GRAY
	*time++
	node.time = *time
	sn := &stackDfsNode{node: node, edge: b.arr[node.id]}
	b.stack2Push(sn)
	for !b.stack2IsEmpty() {
		sn, _ = b.stack2Top()
		flag := false
		for _, v := range sn.edge {
			if v.color == WHITE {
				v.weight = node.weight + 1
				v.prev = node
				v.color = GRAY
				*time++
				v.time = *time
				b.stack2Push(&stackDfsNode{node: v, edge: b.arr[v.id]})
				flag = true
				break
			}
		}
		if !flag {
			sn, _ = b.stack2Pop()
			node.color = BLACK
			*time++
			sn.node.ftime = *time
		}

	}

}

func (b *Bdfs) dfs_visit_stack2(node *bdfsNode, time *int) {
	node.color = GRAY
	*time++
	node.time = *time
	b.stackPush(node)
	for !b.stackIsEmpty() {
		node, _ = b.stackTop()
		flag := false
		for _, v := range b.arr[node.id] {
			if v.color == WHITE {
				v.weight = node.weight + 1
				v.prev = node
				v.color = GRAY
				*time++
				v.time = *time
				b.stackPush(v)
				flag = true
				break
			}
		}
		if !flag {
			node, _ = b.stackPop()
			node.color = BLACK
			*time++
			node.ftime = *time
		}

	}

}

func (b *Bdfs) dfs_visit(node *bdfsNode, time *int) {
	*time++
	node.time = *time
	node.color = GRAY
	for _, v := range b.arr[node.id] {
		if v.color == WHITE {
			fmt.Println("树边" + node.name + "-" + v.name)
			v.weight = node.weight + 1
			v.prev = node
			b.dfs_visit(v, time)
		} else if v.color == GRAY && node.time >= v.time {
			fmt.Println("后向边" + node.name + "-" + v.name)
		} else {
			if node.time < v.time {
				fmt.Println("前向边" + node.name + "-" + v.name)
			} else {
				fmt.Println("横向边" + node.name + "-" + v.name)
			}

		}

	}
	node.color = BLACK
	*time++
	node.ftime = *time

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

// p365
func (b *Bdfs) CreateTopOLogicalSortGraph() {
	b.init()
	b.addEdge("内裤", "裤子")
	b.addEdge("内裤", "鞋")
	b.addEdge("裤子", "鞋")
	b.addEdge("裤子", "腰带")
	b.addEdge("腰带", "夹克")
	b.addEdge("衬衣", "腰带")
	b.addEdge("衬衣", "领带")
	b.addEdge("领带", "夹克")
	b.addEdge("袜子", "鞋")
	b.addEdge("手表", "")
}

// p366 22.4-1
func (b *Bdfs) CreateTopOLogicalSortPtGraph() {
	b.init()
	for i := 0; i <= 13; i++ {
		//b.getNodeByName(string('m'+i))
		b.addEdge(string('m'+i), "")
	}
	fmt.Println(b.idNameMap)
	b.addEdge("m", "x")
	b.addEdge("m", "q")
	b.addEdge("m", "r")
	b.addEdge("n", "q")
	b.addEdge("n", "o")
	b.addEdge("n", "u")

	b.addEdge("o", "r")
	b.addEdge("o", "v")
	b.addEdge("o", "s")

	b.addEdge("p", "o")
	b.addEdge("p", "s")
	b.addEdge("p", "z")

	b.addEdge("q", "t")

	b.addEdge("r", "u")
	b.addEdge("r", "y")

	b.addEdge("s", "r")

	b.addEdge("u", "t")

	b.addEdge("v", "x")
	b.addEdge("v", "w")

	b.addEdge("w", "z")
	b.addEdge("y", "u")
}

// 拓扑排序
func (b *Bdfs) TopOLogicalSort() {
	b.stackInt()
	time := 0
	arr := make([]string, 0)
	for i := 0; i < len(b.idNameMap); i++ {
		node := b.getNodeByName(b.getNameById(i))
		if node.color == WHITE {
			node.weight = 0
			node.color = GRAY
			time++
			node.time = time
			b.stackPush(node)
			for !b.stackIsEmpty() {
				node, _ = b.stackTop()
				flag := false
				for _, v := range b.arr[node.id] {
					if v.color == WHITE {
						v.weight = node.weight + 1
						v.prev = node
						v.color = GRAY
						time++
						v.time = time
						b.stackPush(v)
						flag = true
						break
					}
				}
				if !flag {
					node, _ = b.stackPop()
					node.color = BLACK
					time++
					node.ftime = time
					fmt.Println("===", node)
					arr = append(arr, node.name)
				}
			}
		}
	}

	a := make([]string, 2*len(b.nameNodeMap)+2)
	for _, v := range b.nameNodeMap {
		a[v.ftime] = v.name
	}
	for _, v := range a {
		if v != "" {
			fmt.Printf("%s ", v)
		}
	}
	fmt.Println()
	for i := len(arr) - 1; i >= 0; i-- {
		fmt.Printf("%s ", arr[i])
	}
	fmt.Println()
}

func (b *Bdfs) addEdgeT(v1, v2 string) {
	b.addEdge(v2, v1)
}

func (b *Bdfs) CreateSccGraph(isT bool) {
	b.init()
	//for i := 0; i < 8; i++ {
	//	b.addEdge(string(i+'a'), "")
	//}
	if !isT {
		b.addEdge("a", "b")
		b.addEdge("b", "e")
		b.addEdge("e", "a")
		b.addEdge("b", "f")
		b.addEdge("e", "f")
		b.addEdge("f", "g")
		b.addEdge("g", "f")
		b.addEdge("g", "h")
		b.addEdge("h", "h")
		b.addEdge("c", "g")
		b.addEdge("c", "d")
		b.addEdge("d", "c")
		b.addEdge("d", "h")
	} else {
		b.addEdgeT("a", "b")
		b.addEdgeT("b", "e")
		b.addEdgeT("e", "a")
		b.addEdgeT("b", "f")
		b.addEdgeT("e", "f")
		b.addEdgeT("f", "g")
		b.addEdgeT("g", "f")
		b.addEdgeT("g", "h")
		b.addEdgeT("h", "h")
		b.addEdgeT("c", "g")
		b.addEdgeT("c", "d")
		b.addEdgeT("d", "c")
		b.addEdgeT("d", "h")
	}
	//if !isT {
	//	b.addEdge("内裤", "裤子")
	//	b.addEdge("内裤", "鞋")
	//	b.addEdge("裤子", "鞋")
	//	b.addEdge("裤子", "腰带")
	//	b.addEdge("腰带", "夹克")
	//	b.addEdge("衬衣", "腰带")
	//	b.addEdge("衬衣", "领带")
	//	b.addEdge("领带", "夹克")
	//	b.addEdge("袜子", "鞋")
	//	b.addEdge("手表", "")
	//}else {
	//	b.addEdgeT("内裤", "裤子")
	//	b.addEdgeT("内裤", "鞋")
	//	b.addEdgeT("裤子", "鞋")
	//	b.addEdgeT("裤子", "腰带")
	//	b.addEdgeT("腰带", "夹克")
	//	b.addEdgeT("衬衣", "腰带")
	//	b.addEdgeT("衬衣", "领带")
	//	b.addEdgeT("领带", "夹克")
	//	b.addEdgeT("袜子", "鞋")
	//	b.addEdge("手表", "")
	//}
}

// 强连通分量
func (b *Bdfs) Scc() {
	// dfs搜索
	b.CreateSccGraph(false)
	b.stackInt()
	time := 0
	for i := 0; i < len(b.idNameMap); i++ {
		node := b.getNodeByName(b.getNameById(i))
		if node.color == WHITE {
			node.weight = 0
			b.dfs_visit_stack(node, &time)
		}
	}

	// 按发现时间ftime倒序排列
	a := make([]string, 2*len(b.nameNodeMap)+2)
	for _, v := range b.nameNodeMap {
		a[v.ftime] = v.name
	}

	sortArr := make([]string, 0)
	for i := len(a) - 1; i >= 0; i-- {
		if a[i] != "" {
			sortArr = append(sortArr, a[i])
		}
	}

	// 转置图按上述的发现时间ftime倒序排列sortArr进行深度优先，找出分割节点
	b.CreateSccGraph(true)
	b.stackInt()
	time = 0
	separateArr := make([]string, 0) // 分割节点数据
	for _, v := range sortArr {
		node := b.getNodeByName(v)
		if node.color == WHITE {
			separateArr = append(separateArr, node.name)
			node.weight = 0
			b.dfs_visit_stack(node, &time)
		}
	}

	// 节点分割
	resArr := make([][]string, 0)
	index := 0
	separateIndexArr := make([]int, 0)
	for k, v := range sortArr {
		if v == separateArr[index] {
			separateIndexArr = append(separateIndexArr, k)
			index++
		}
	}
	separateIndexArr = append(separateIndexArr, len(sortArr))

	for i := 0; i < len(separateIndexArr)-1; i++ {
		resArr = append(resArr, sortArr[separateIndexArr[i]:separateIndexArr[i+1]])
	}
	fmt.Println(separateArr, sortArr, resArr)
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

func (b *Bdfs) stackTop() (*bdfsNode, error) {
	if b.stackIsEmpty() {
		return nil, errors.New("stack is empty")
	}
	length := len(b.stack)
	return b.stack[length-1], nil
}

func (b *Bdfs) stack2Int() {
	b.stack2 = make([]*stackDfsNode, 0)
}

func (b *Bdfs) stack2IsEmpty() bool {
	return len(b.stack2) == 0
}

func (b *Bdfs) stack2Push(v *stackDfsNode) {
	b.stack2 = append(b.stack2, v)
}

func (b *Bdfs) stack2Pop() (*stackDfsNode, error) {
	if b.stack2IsEmpty() {
		return nil, errors.New("stack is empty")
	}
	length := len(b.stack2)
	name := b.stack2[length-1]
	b.stack2 = b.stack2[:length-1]
	return name, nil
}

func (b *Bdfs) stack2Top() (*stackDfsNode, error) {
	if b.stack2IsEmpty() {
		return nil, errors.New("stack is empty")
	}
	length := len(b.stack2)
	return b.stack2[length-1], nil
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
