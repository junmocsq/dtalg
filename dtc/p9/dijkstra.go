package p9

import (
	"fmt"
)

type TableEntry struct {
	id     int
	name   string
	weight int
}

type Table struct {
	arr     [][]*TableEntry
	index   int
	idMap   map[int]string
	nameMap map[string]int
}

func NewTable() *Table {
	return &Table{
		idMap:   make(map[int]string),
		nameMap: make(map[string]int),
	}
}

func (t *Table) getId() int {
	t.index++
	return t.index - 1
}

func (t *Table) getIdByName(name string) int {
	if id, ok := t.nameMap[name]; ok {
		return id
	} else {
		id = t.getId()
		t.idMap[id] = name
		t.nameMap[name] = id
		return id
	}
}

func (t *Table) getNameById(id int) string {
	return t.idMap[id]
}

func (t *Table) addEdge(v1, v2 string, w2 int) {
	headIndex := t.getIdByName(v1)
	length := len(t.arr)
	for i := 0; i <= headIndex-length; i++ {
		t.arr = append(t.arr, nil)
	}
	for _, v := range t.arr[headIndex] {
		if v.name == v2 {
			return
		}
	}

	id := t.getIdByName(v2)
	v := &TableEntry{id: id, name: v2, weight: w2}
	t.arr[headIndex] = append(t.arr[headIndex], v)
}

func (t *Table) CreateGraph() {
	t.addEdge("v1", "v2", 2)
	t.addEdge("v1", "v4", 1)
	t.addEdge("v2", "v5", 10)
	t.addEdge("v2", "v4", 3)
	t.addEdge("v3", "v6", 5)
	t.addEdge("v3", "v1", 4)
	t.addEdge("v4", "v3", 2)
	t.addEdge("v4", "v6", 8)
	t.addEdge("v4", "v7", 4)
	t.addEdge("v4", "v5", 2)
	t.addEdge("v5", "v7", 6)
	t.addEdge("v7", "v6", 1)
}

func (t *Table) print() {
	for k, v := range t.arr {
		name := t.getNameById(k)
		fmt.Printf("index:%s ", name)
		for _, _v := range v {
			fmt.Printf("%s %d\t", _v.name, _v.weight)
		}
		fmt.Println()
	}
}

func (t *Table) Dijkstra(v1, v2 string) {
	known := make(map[string]bool)
	weights := make(map[string]int)
	prev := make(map[string]string)
	id := t.getIdByName(v1)
	known[v1] = true
	weights[v1] = 0
	prev[v1] = ""
	for _, v := range t.arr[id] {
		known[v.name] = false
		weights[v.name] = v.weight
		prev[v.name] = v1
	}
	for {
		min := -1
		node := ""
		for name, ok := range known {
			if !ok {
				if min == -1 || weights[name] < min {
					min = weights[name]
					node = name
				}
			}
		}
		if min == -1 {
			break
		}

		known[node] = true
		baseWeight := weights[node]
		for _, v := range t.arr[t.getIdByName(node)] {
			if !known[v.name] {
				known[v.name] = false
				plusWeight := baseWeight + v.weight
				if w, ok := weights[v.name]; ok {
					if w > plusWeight {
						weights[v.name] = plusWeight
						prev[v.name] = node
					}
				} else {
					weights[v.name] = plusWeight
					prev[v.name] = node
				}
			}
		}
	}
	fmt.Println(weights, known, prev)

}
