package p4

import "fmt"

// 伸展树

type splayNode struct {
	ele   int
	left  *splayNode
	right *splayNode
}

type splayTree splayNode

func NewSplayTree() *splayTree {
	return &splayTree{}
}

// 向左旋
func (s *splayTree) leftSingleRotation(k1 *splayNode) *splayNode {
	k2 := k1.right
	k1.right = k2.left
	k2.left = k1
	return k2
}

// 向右旋
func (s *splayTree) rightSingleRotation(k1 *splayNode) *splayNode {
	k2 := k1.left
	k1.left = k2.right
	k2.right = k1
	return k2
}

// 左-右旋转
func (s *splayTree) leftRightDoubleRotation(k1 *splayNode) *splayNode {
	k1.left = s.leftSingleRotation(k1.left)
	return s.rightSingleRotation(k1)
}

// 右-左旋转
func (s *splayTree) rightLeftDoubleRotation(k1 *splayNode) *splayNode {
	k1.right = s.rightSingleRotation(k1.right)
	return s.leftSingleRotation(k1)
}

// 右一字型左变换
func (s *splayTree) oneLeftRotation(k1 *splayNode) *splayNode {
	k2 := k1.right
	k3 := k2.right
	k1.right = k2.left
	k2.right = k3.left
	k2.left = k1
	k3.left = k2
	return k3
}

// 左一字型有变换
func (s *splayTree) oneRightRotation(k1 *splayNode) *splayNode {
	k2 := k1.left
	k3 := k2.left
	k1.left = k2.right
	k2.left = k3.right
	k2.right = k1
	k3.right = k2
	return k3
}

func (s *splayTree) Insert(ele int) {
	node := &splayNode{ele: ele}
	if s.left == nil {
		s.left = node
	} else {
		root := s.left
		for {
			if root.ele == ele {
				return
			} else if root.ele > ele {
				// 左子树
				if root.left == nil {
					root.left = node
					break
				} else {
					root = root.left
				}
			} else {
				if root.right == nil {
					root.right = node
					break
				} else {
					root = root.right
				}
			}
		}
	}
}

func (s *splayTree) Find(ele int) *splayNode {
	nodeArr := make([]*splayNode, 1)
	nodeArr[0] = (*splayNode)(s)
	root := s.left
	for root != nil {
		nodeArr = append(nodeArr, root)
		if root.ele > ele {
			// 左子树
			root = root.left
		} else if root.ele < ele {
			root = root.right
		} else {
			break
		}
	}
	// 未搜索到数据
	if root == nil {
		return nil
	}
	length := len(nodeArr)
	for index := length - 1; index > 1; index -= 2 {
		if index >= 3 {

			if nodeArr[index-2].left == nodeArr[index-1] {
				if nodeArr[index-1].left == nodeArr[index] {
					// 左一字型
					s.oneRightRotation(nodeArr[index-2])
				} else {
					// 左右型
					s.leftRightDoubleRotation(nodeArr[index-2])
				}
			} else {
				if nodeArr[index-1].right == nodeArr[index] {
					// 右一字型
					s.oneLeftRotation(nodeArr[index-2])
				} else {
					s.rightLeftDoubleRotation(nodeArr[index-2])
				}
			}
			if nodeArr[index-3].left == nodeArr[index-2] {
				nodeArr[index-3].left = nodeArr[index]
			} else {
				nodeArr[index-3].right = nodeArr[index]
			}
			nodeArr[index-2] = nodeArr[index]
		} else {
			if nodeArr[1].left == nodeArr[2] {
				// 左节点
				s.rightSingleRotation(nodeArr[1])
			} else {
				s.leftSingleRotation(nodeArr[1])
			}
			s.left = nodeArr[2]
		}
	}
	return s.left
}

func (s *splayTree) inOrderPrint(root *splayNode) {
	if root == nil {
		return
	}
	fmt.Printf("%d ", root.ele)
	s.inOrderPrint(root.left)
	s.inOrderPrint(root.right)

}

func (s *splayTree) midOrderPrint(root *splayNode) {
	if root == nil {
		return
	}
	s.midOrderPrint(root.left)
	fmt.Printf("%d ", root.ele)
	s.midOrderPrint(root.right)
}

func (s *splayTree) postOrderPrint(root *splayNode) {
	if root == nil {
		return
	}
	s.postOrderPrint(root.left)
	s.postOrderPrint(root.right)
	fmt.Printf("%d ", root.ele)
}
