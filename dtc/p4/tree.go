package p4

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

// 深度 从root节点开始向leaf数
// 高度 从leaf节点向root节点数

type treeNode struct {
	ele         string
	firstChild  *treeNode
	nextSibling *treeNode
	childNum    int
}

type generalTree treeNode

func listDir(root string) []string {
	cmd := exec.Command("ls", root)
	out, _ := cmd.Output()
	return strings.Split(string(out), "\n")
}

func newGeneralTree() *generalTree {
	return &generalTree{}
}

func (t *generalTree) loopDir(dir string, parent *treeNode) {
	arr := listDir(dir)
	if len(arr) == 0 {
		return
	}
	for _, v := range arr {
		if v == "" {
			continue
		}
		tempParent := t.add(v, parent)
		temp := dir + "/" + v
		if f, _ := os.Stat(temp); f.IsDir() {
			t.loopDir(temp, tempParent)
		}
	}
}

func (t *generalTree) add(ele string, parent *treeNode) *treeNode {
	node := &treeNode{ele: ele}
	if parent.firstChild == nil {
		parent.firstChild = node
	} else {
		temp := parent.firstChild
		for temp.nextSibling != nil {
			temp = temp.nextSibling
		}
		temp.nextSibling = node
	}
	parent.childNum++
	return node
}

func (t *generalTree) print(depth int, parent *treeNode) {
	sibling := parent.firstChild
	for {
		if sibling != nil {
			for i := 0; i < depth-1; i++ {
				fmt.Printf("\t")
			}
			if depth >= 1 {
				fmt.Printf("|_")
			}

			t.print(depth+1, sibling)
			fmt.Println(sibling.ele, "(", sibling.childNum, ")")
			sibling = sibling.nextSibling
		} else {
			break
		}
	}
}
