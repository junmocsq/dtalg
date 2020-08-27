package p9

import "testing"

func TestNewTable(t *testing.T) {
	d := NewTable()
	d.CreateGraph()
	d.print()
	d.Dijkstra("v1", "v7")

}
