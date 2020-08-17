package p4

import "fmt"

type binaryExpressionNode struct {
	ele   string
	left  *binaryExpressionNode
	right *binaryExpressionNode
}

// 表达式树
type expressionTree struct {
	head      *binaryExpressionNode
	stack     []string
	stackNode []*binaryExpressionNode
}

func newExpressionTree() *expressionTree {
	return &expressionTree{}
}

func (e *expressionTree) push(ele string) {

	e.stack = append(e.stack, ele)

}

func (e *expressionTree) pop() string {

	if len(e.stack) > 0 {
		length := len(e.stack)
		ele := e.stack[length-1]
		e.stack = e.stack[:length-1]
		return ele
	}

	return ""
}

func (e *expressionTree) pushNode(node *binaryExpressionNode) {

	e.stackNode = append(e.stackNode, node)

}

func (e *expressionTree) popNode() *binaryExpressionNode {

	if len(e.stackNode) > 0 {
		length := len(e.stackNode)
		ele := e.stackNode[length-1]
		e.stackNode = e.stackNode[:length-1]
		return ele
	}
	return nil
}

func (e *expressionTree) parseStrToArr(str string) []string {

	// 解析表达式为中缀数组
	midArr := make([]string, 0)
	s := ""
	preIsNumber := false
	for _, v := range []byte(str) {

		switch v {
		case '+', '-', '*', '/', '(', ')':
			if preIsNumber {
				midArr = append(midArr, s)
			}
			s = ""
			preIsNumber = false
			midArr = append(midArr, string(v))
		default:
			if v >= '0' && v <= '9' {
				preIsNumber = true
				s += string(v)
			}
		}
	}
	if preIsNumber {
		midArr = append(midArr, s)
	}
	return midArr
}

func (e *expressionTree) suffixArr(midArr []string) []string {

	// 转换为后缀表达式数组
	suffixArr := make([]string, 0)
	// 和栈顶元素相比，对于 "+-*/" 优先级高入栈，'('括号也入栈，其他的出栈
	for _, v := range midArr {
		//fmt.Println("====", e.stack, suffixArr)
		switch v {
		case "+", "-":
			for len(e.stack) != 0 {
				top := e.pop()
				if top == "(" {
					e.push(top)
					break
				}
				suffixArr = append(suffixArr, top)
			}
			e.push(v)

		case "*", "/":
			if len(e.stack) == 0 {
				e.push(v)
			} else {
				for len(e.stack) != 0 {
					top := e.pop()
					if top == "+" || top == "-" || top == "(" {
						e.push(top)
						break
					} else {
						suffixArr = append(suffixArr, top)
					}
				}
				e.push(v)
			}
		case "(":
			e.push(v)
		case ")":
			for len(e.stack) != 0 {
				top := e.pop()
				if top != "(" {
					suffixArr = append(suffixArr, top)
				} else {
					break
				}
			}
		default:
			suffixArr = append(suffixArr, v)
		}
	}
	// 清空剩余栈中的操作符
	for len(e.stack) != 0 {
		suffixArr = append(suffixArr, e.pop())
	}
	return suffixArr
}

// 后缀表达式转化为树
func (e *expressionTree) tree(suffixArr []string) *binaryExpressionNode {
	for _, ele := range suffixArr {
		node := &binaryExpressionNode{ele: ele}
		switch ele {
		case "+", "-", "*", "/":
			b := e.popNode()
			a := e.popNode()
			node.left = a
			node.right = b
		}
		e.pushNode(node)
	}
	return e.popNode()
}

// 前序遍历
func (e *expressionTree) inOrderPrint(root *binaryExpressionNode) {
	if root != nil {
		fmt.Printf("%s ", root.ele)
		e.inOrderPrint(root.left)
		e.inOrderPrint(root.right)
	}
}
func (e *expressionTree) midOrderPrint(root *binaryExpressionNode) {
	if root != nil {
		e.midOrderPrint(root.left)
		fmt.Printf("%s ", root.ele)
		e.midOrderPrint(root.right)
	}
}
func (e *expressionTree) postOrderPrint(root *binaryExpressionNode) {
	if root != nil {
		e.postOrderPrint(root.left)
		e.postOrderPrint(root.right)
		fmt.Printf("%s ", root.ele)
	}
}
