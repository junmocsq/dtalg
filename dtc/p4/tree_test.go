package p4

import (
	"testing"
)

func TestTreeNode(t *testing.T) {
	listDir("/")
	tree := newGeneralTree()
	root := "/Users/junmo/www/dtalg"
	parent := tree.add(root, (*treeNode)(tree))
	tree.loopDir(root, parent)
	tree.print(0, (*treeNode)(tree))

}
