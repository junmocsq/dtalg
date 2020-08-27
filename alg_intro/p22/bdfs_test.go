package p22

import (
	"testing"
)

func TestNewBdfsNode(t *testing.T) {
	b := NewBdfs()
	b.Bfs("s")
	t.Log(b.nameNodeMap)
	b.printPath("s", "w")
}
