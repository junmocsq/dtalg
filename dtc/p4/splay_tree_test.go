package p4

import (
	"fmt"
	"testing"
)

func TestSplayTree(t *testing.T) {
	tree := NewSplayTree()
	for i := 32; i > 0; i-- {
		tree.Insert(i)
	}
	printSplayTree(tree)
	tree.Find(1)
	tree.Find(2)
	tree.Find(3)
	tree.Find(4)
	//tree.Delete(13)
	//tree.Delete(2)
	printSplayTree(tree)
}

func printSplayTree(tree *splayTree) {
	root := tree.left
	fmt.Printf("前序遍历：")
	tree.inOrderPrint(root)
	fmt.Println()
	fmt.Printf("中序遍历：")
	tree.midOrderPrint(root)
	fmt.Println()
	fmt.Printf("后序遍历：")
	tree.postOrderPrint(root)
	fmt.Println()
}
