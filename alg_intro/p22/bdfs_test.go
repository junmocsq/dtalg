package p22

import (
	"testing"
)

func TestNewBdfsNode(t *testing.T) {
	b := NewBdfs()
	b.CreateSccGraph(false)
	b.Scc()
	t.Log(b.nameNodeMap)
	//b.printPath("s", "w")
}
