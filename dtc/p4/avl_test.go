package p4

import (
	"fmt"
	"testing"
)

func TestAvlTree(t *testing.T) {
	tree := NewAvl()
	tree.Insert(3)
	tree.Insert(2)
	tree.Insert(1)
	printAvlTree(tree)
	tree.Insert(4)
	tree.Insert(5)
	printAvlTree(tree)
	tree.Insert(6)
	tree.Insert(7)
	printAvlTree(tree)
	tree.Insert(16)
	tree.Insert(15)
	printAvlTree(tree)
	tree.Insert(14)
	tree.Insert(13)
	tree.Insert(12)
	tree.Insert(11)
	tree.Insert(10)
	tree.Insert(9)
	tree.Insert(8)
	printAvlTree(tree)

}

func printAvlTree(tree *avlTree) {
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
