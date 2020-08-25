package p9

import "testing"

func TestBdfs(t *testing.T) {
	b := NewBdfs()
	b.CreateGraph()
	b.print()
	b.bfs("v0", "v5")
	b.dfs("v0", "v5")
}
