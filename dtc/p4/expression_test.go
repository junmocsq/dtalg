package p4

import (
	"fmt"
	"testing"
)

func TestBinaryTree(t *testing.T) {
	b := newExpressionTree()
	expr := "100 + 32*12*(33 * 43 + 122)*134"
	midArr := b.parseStrToArr(expr)
	suffixArr := b.suffixArr(midArr)
	t.Log("suffixArr", suffixArr)
	root := b.tree(suffixArr)
	b.inOrderPrint(root)
	fmt.Println()
	b.midOrderPrint(root)
	fmt.Println()
	b.postOrderPrint(root)
	fmt.Println()
}
