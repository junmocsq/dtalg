package p4

import "fmt"

// 一棵AVL树是其每个节点的左子树和右子树的高度最多差1的二叉查找树
type avlNode struct {
	ele    int
	left   *avlNode
	right  *avlNode
	height int
}

type AvlTreeInterface interface {
	Find(ele int) *avlNode
	FindMin() *avlNode
	FindMax() *avlNode
	Insert(ele int)
	Delete(ele int)
	Retrieve(node *avlNode) int
}

type avlTree avlNode

func NewAvl() *avlTree {
	return &avlTree{}
}

// 向左单旋
// 右儿子的右子树插入的旋转修正
func (a *avlTree) leftSingleRotation(node *avlNode) *avlNode {
	right := node.right
	node.right = right.left
	right.left = node

	node.height = a.max(a.Height(node.left), a.Height(node.right)) + 1
	right.height = a.max(a.Height(right.right), node.height) + 1
	return right
}

func (a *avlTree) max(h1, h2 int) int {
	if h1 > h2 {
		return h1
	} else {
		return h2
	}
}

// 向右单旋
// 左儿子的左子树插入的旋转修正
func (a *avlTree) rightSingleRotation(node *avlNode) *avlNode {
	left := node.left
	node.left = left.right
	left.right = node

	node.height = a.max(a.Height(node.left), a.Height(node.right)) + 1
	left.height = a.max(a.Height(left.left), node.height) + 1
	return left

}

// 左儿子的右子树插入数据的双旋转修正
func (a *avlTree) leftRightDoubleRotation(node *avlNode) *avlNode {
	// 向左单旋
	node.left = a.leftSingleRotation(node.left)

	return a.rightSingleRotation(node)

}

// 右儿子的左子树插入数据的双旋转修正
func (a *avlTree) rightLeftDoubleRotation(node *avlNode) *avlNode {
	// 向右单旋
	node.right = a.rightSingleRotation(node.right)
	return a.leftSingleRotation(node)
}

func (a *avlTree) Height(position *avlNode) int {
	if position == nil {
		return -1
	} else {
		return position.height
	}
}

func (a *avlTree) NewAvlTree() *avlTree {
	return &avlTree{}
}

func (a *avlTree) Insert(ele int) {
	node := &avlNode{ele: ele}
	if a.left == nil {
		a.left = node
	} else {
		a.left = a.insert(ele, a.left)
	}
}

func (a *avlTree) insert(ele int, T *avlNode) *avlNode {

	n := &avlNode{ele: ele}
	if T == nil {
		T = n
	} else if T.ele > ele {
		// 左子树
		T.left = a.insert(ele, T.left)
		if a.Height(T.left)-a.Height(T.right) == 2 {
			if ele < T.left.ele {
				// 向右单旋
				T = a.rightSingleRotation(T)
			} else {
				T = a.leftRightDoubleRotation(T)
			}
		}
	} else if T.ele < ele {
		// 右子树
		T.right = a.insert(ele, T.right)
		if a.Height(T.right)-a.Height(T.left) == 2 {
			if ele > T.right.ele {
				T = a.leftSingleRotation(T)
			} else {
				T = a.rightLeftDoubleRotation(T)
			}
		}
	}
	T.height = a.max(a.Height(T.left), a.Height(T.right)) + 1
	return T

}

// TODO AVL树删除
func (a *avlTree) Delete(ele int) {

}

func (s *avlTree) inOrderPrint(root *avlNode) {
	if root == nil {
		return
	}
	fmt.Printf("%d(%d) ", root.ele, root.height)
	s.inOrderPrint(root.left)
	s.inOrderPrint(root.right)

}

func (s *avlTree) midOrderPrint(root *avlNode) {
	if root == nil {
		return
	}
	s.midOrderPrint(root.left)
	fmt.Printf("%d(%d) ", root.ele, root.height)
	s.midOrderPrint(root.right)
}

func (s *avlTree) postOrderPrint(root *avlNode) {
	if root == nil {
		return
	}
	s.postOrderPrint(root.left)
	s.postOrderPrint(root.right)
	fmt.Printf("%d(%d) ", root.ele, root.height)
}
