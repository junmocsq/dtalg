package p4

import "fmt"

type binarySearchNode struct {
	ele   int
	left  *binarySearchNode
	right *binarySearchNode
}

type binarySearchTree binarySearchNode

func NewBinarySearchTree() *binarySearchTree {
	return &binarySearchTree{}
}

type binarySearcher interface {
	Find(ele int) *binarySearchNode
	FindMin(position *binarySearchNode) *binarySearchNode
	FindMax(position *binarySearchNode) *binarySearchNode
	Insert(ele int)
	Delete(ele int)
	Retrieve(node *binarySearchNode) int
}

func (s *binarySearchTree) Find(ele int) *binarySearchNode {
	tree := s.left
	if tree == nil {
		return nil
	}
	for tree != nil {
		if tree.ele > ele {
			// 左子树
			tree = tree.left
		} else if tree.ele == ele {
			return tree
		} else {
			tree = tree.right
		}
	}
	return nil
}

func (s *binarySearchTree) FindParent(ele int) *binarySearchNode {
	tree := s.left
	if tree == nil {
		return nil
	}
	parent := (*binarySearchNode)(s)
	for tree != nil {
		if tree.ele > ele {
			parent = tree
			// 左子树
			tree = tree.left
		} else if tree.ele == ele {
			return parent
		} else {
			parent = tree
			tree = tree.right
		}
	}
	return nil
}

func (s *binarySearchTree) FindMin(position *binarySearchNode) *binarySearchNode {

	for position != nil && position.left != nil {
		position = position.left
	}
	return position
}

func (s *binarySearchTree) FindMinParent(position *binarySearchNode) *binarySearchNode {
	var parent *binarySearchNode
	for position != nil && position.left != nil {
		parent = position
		position = position.left
	}
	return parent
}

func (s *binarySearchTree) FindMax(position *binarySearchNode) *binarySearchNode {
	for position != nil && position.right != nil {
		position = position.right
	}
	return position
}

func (s *binarySearchTree) Insert(ele int) {
	tree := s.left
	node := &binarySearchNode{ele: ele}
	if tree == nil {
		s.left = node
		return
	}
	for tree != nil {
		if tree.ele == ele {
			return
		} else if tree.ele > ele {
			// 左子树
			if tree.left != nil {
				tree = tree.left
			} else {
				tree.left = node
				return
			}
		} else {
			// 右子树
			if tree.right != nil {
				tree = tree.right
			} else {
				tree.right = node
				return
			}
		}
	}
}

// 叶子节点直接删除
// 只有一个子节点的直接用删除节点的父节点指向它
// 两个子节点的用删除节点的右节点的最小值替换它，再删除被替换的节点
func (s *binarySearchTree) Delete(ele int) {
	root := s.left
	if root == nil {
		return
	}
	// 查找删除节点的父节点
	parent := s.FindParent(ele)
	if parent == nil {
		return
	}

	if parent.left.ele == ele {
		// 左子树
		child := parent.left
		if child.right == nil && child.left == nil {
			parent.left = nil
		} else if child.right != nil && child.left != nil {
			minNodeParent := s.FindMinParent(child.right)
			var minNode *binarySearchNode
			if minNodeParent == nil {
				child.ele = child.right.ele
				child.right = child.right.right
			} else {
				minNode = minNodeParent.left
				child.ele = minNode.ele
				minNodeParent.left = minNodeParent.left.right
			}
		} else {
			if child.left != nil {
				parent.left = child.left
			} else {
				parent.left = child.right
			}
		}
		return
	} else {
		// 右子树
		child := parent.right
		if child.right == nil && child.left == nil {
			parent.right = nil
			return
		} else if child.right != nil && child.left != nil {
			minNodeParent := s.FindMinParent(child.right)
			var minNode *binarySearchNode
			if minNodeParent == nil {
				child.ele = child.right.ele
				child.right = child.right.right
			} else {
				minNode = minNodeParent.left
				child.ele = minNode.ele
				minNodeParent.left = minNodeParent.left.right
			}
		} else {
			if child.left != nil {
				parent.right = child.left
			} else {
				parent.right = child.right
			}
		}
	}

}

func (s *binarySearchTree) DeleteRecursion(ele int, root *binarySearchNode, parent *binarySearchNode) {
	if root == nil {
		return
	} else if root.ele > ele {
		s.DeleteRecursion(ele, root.left, root)
	} else if root.ele < ele {
		s.DeleteRecursion(ele, root.right, root)
	} else if root.left != nil && root.right != nil {
		minNodeParent := s.FindMinParent(root.right)
		var minNode *binarySearchNode
		if minNodeParent == nil {
			root.ele = root.right.ele
			root.right = root.right.right
		} else {
			minNode = minNodeParent.left
			root.ele = minNode.ele
			minNodeParent.left = minNodeParent.left.right
		}
	} else {
		if parent.left == root {
			if root.left != nil {
				parent.left = root.left
			} else {
				parent.left = root.right
			}
		} else {
			if root.left != nil {
				parent.right = root.left
			} else {
				parent.right = root.right
			}
		}

	}

}

func (s *binarySearchTree) Retrieve(node *binarySearchNode) int {
	return node.ele
}

func (s *binarySearchTree) inOrderPrint(root *binarySearchNode) {
	if root == nil {
		return
	}
	fmt.Printf("%d ", root.ele)
	s.inOrderPrint(root.left)
	s.inOrderPrint(root.right)

}

func (s *binarySearchTree) midOrderPrint(root *binarySearchNode) {
	if root == nil {
		return
	}
	s.midOrderPrint(root.left)
	fmt.Printf("%d ", root.ele)
	s.midOrderPrint(root.right)
}

func (s *binarySearchTree) postOrderPrint(root *binarySearchNode) {
	if root == nil {
		return
	}
	s.postOrderPrint(root.left)
	s.postOrderPrint(root.right)
	fmt.Printf("%d ", root.ele)
}

// 层次遍历
func (s *binarySearchTree) layerPrint(root *binarySearchNode) {
	if root == nil {
		return
	}
	arr := make([]*binarySearchNode, 0)
	arr = s.push(arr, root)
	front := len(arr)
	rear := 0
	floor := 1
	fmt.Println("level", front)
	var node *binarySearchNode
	for len(arr) != 0 {
		arr, node = s.pop(arr)
		front--
		fmt.Println(floor, node.ele)
		if node.left != nil {
			arr = s.push(arr, node.left)
		}
		if node.right != nil {
			arr = s.push(arr, node.right)
		}
		if front == rear {
			front = len(arr)
			rear = 0
			floor++
		}
	}
}

func (s *binarySearchTree) push(arr []*binarySearchNode, node *binarySearchNode) []*binarySearchNode {
	arr = append(arr, node)
	return arr
}

func (s *binarySearchTree) pop(arr []*binarySearchNode) ([]*binarySearchNode, *binarySearchNode) {
	if len(arr) == 0 {
		return nil, nil
	}
	node := arr[0]
	arr = arr[1:]
	return arr, node
}
