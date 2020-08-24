package p9

import (
	"fmt"
	"testing"
)

func TestGraph(t *testing.T) {
	g := NewAdjacencyList()
	g.addEdge("1", "2")
	g.addEdge("1", "4")
	g.addEdge("1", "3")
	g.addEdge("2", "4")
	g.addEdge("2", "5")
	g.addEdge("3", "6")
	g.addEdge("4", "6")
	g.addEdge("4", "7")
	g.addEdge("4", "3")
	g.addEdge("5", "4")
	g.addEdge("5", "7")
	g.addEdge("7", "6")
	fmt.Println(g.arr[0])

	fmt.Println(g.InDegree())
	t.Log(g.TopSort())
}
