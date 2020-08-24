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

// 入度数组
func (a *adjacencyList) InDegree() map[string]int {
	inDegree := make(map[string]int)
	for name, id := range a.nameMap {
		if inDegree[name] == 0 {
			inDegree[name] = 0
		}
		index := a.arr[id]

		for index != nil {
			inDegree[index.name]++
			index = index.next
		}
	}
	return inDegree
}

// 拓扑排序
func (a *adjacencyList) TopSort() []string {
	inDegree := a.InDegree()
	nameArr := make([]string, 0)
	queue := make([]string, 0)
	for name, v := range inDegree {
		if v == 0 {
			queue = a.push(queue, name)
			delete(inDegree, name)
		}
	}
	count := 0
	for len(queue) != 0 {
		var name string
		name, queue = a.pop(queue)
		nameArr = append(nameArr, name)
		count++
		id := a.getNameId(name)
		head := a.arr[id]
		for head != nil {
			if inDegree[head.name] == 1 {
				queue = a.push(queue, head.name)
				delete(inDegree, head.name)
			} else {
				inDegree[head.name]--
			}
			head = head.next
		}
	}
	if count != len(a.nameMap) {
		panic("Graph has a cycle")
	}
	return nameArr
}
func (a *adjacencyList) push(queue []string, name string) []string {
	return append(queue, name)
}
func (a *adjacencyList) pop(queue []string) (string, []string) {
	if len(queue) == 0 {
		return "", nil
	}
	name := queue[0]
	queue = queue[1:]
	return name, queue
}
