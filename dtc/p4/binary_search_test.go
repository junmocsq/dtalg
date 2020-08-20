package p4

import (
	"fmt"
	"testing"
)

func TestBinarySearch(t *testing.T) {
	tree := NewBinarySearchTree()
	tree.Insert(100)
	tree.Insert(80)
	tree.Insert(150)
	tree.Insert(90)
	tree.Insert(50)
	tree.Insert(120)
	tree.Insert(180)
	tree.Insert(20)
	tree.Insert(70)
	tree.Insert(85)
	tree.inOrderPrint(tree.left)
	fmt.Println()
	tree.midOrderPrint(tree.left)
	fmt.Println()
	tree.postOrderPrint(tree.left)
	fmt.Println()

	//tree.Delete(80)
	tree.DeleteRecursion(150, tree.left, (*binarySearchNode)(tree))
	tree.inOrderPrint(tree.left)
	fmt.Println()
	tree.layerPrint(tree.left)
}
