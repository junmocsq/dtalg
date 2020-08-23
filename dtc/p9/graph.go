package p9

type vertex struct {
	id   int
	name string
	val  interface{}
	next *vertex
}

// 邻接表表示
type adjacencyList struct {
	arr     []*vertex
	index   int
	nameMap map[string]int
}

func NewAdjacencyList() *adjacencyList {
	return &adjacencyList{nameMap: map[string]int{}}
}

func (a *adjacencyList) generateId() int {
	a.index++
	return a.index - 1
}

func (a *adjacencyList) getNameId(name string) int {
	if id, ok := a.nameMap[name]; ok {
		return id
	} else {
		id := a.generateId()
		a.nameMap[name] = id
		return id
	}
}

// 添加name name2的边
func (a *adjacencyList) addEdge(name string, name2 string) {
	headIndex := a.getNameId(name)
	length := len(a.arr)
	for i := 0; i <= headIndex-length; i++ {
		a.arr = append(a.arr, nil)
	}

	id := a.getNameId(name2)
	v := &vertex{id: id, name: name2}
	head := a.arr[headIndex]
	if head == nil {
		a.arr[headIndex] = v
	} else {
		for head.next != nil {
			head = head.next
		}
		head.next = v
	}
}
