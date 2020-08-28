package linked

import (
	"fmt"
	"github.com/sirupsen/logrus"
)

type trieNode struct {
	char     byte
	children [26]*trieNode
	isEnd    bool
}

type trie trieNode

func NewTrie() *trie {
	return &trie{char: '/'}
}

func (t *trie) Insert(str string) {

	var root = (*trieNode)(t)
	isNew := false
	for _, c := range str {
		if c < 'a' || c > 'z' {
			logrus.Error("insert char is failed ")
			continue
		}
		if root.children[c-'a'] == nil {
			root.children[c-'a'] = &trieNode{char: byte(c)}
			isNew = true
		}
		root.isEnd = false
		root = root.children[c-'a']
	}
	if isNew{
		root.isEnd = true
	}

}

func (t *trie) find(str string) {
	var root = (*trieNode)(t)
	for _, c := range str {
		if c < 'a' || c > 'z' {
			logrus.Error("insert char is failed ")
			continue
		}
		if root.children[c-'a'] == nil {
			logrus.Info("未找到")
			return
		}
		root = root.children[c-'a']
	}
	if root.isEnd{
		logrus.Info("找到")
	}else {
		logrus.Info("前缀匹配")
		t.findPrefix(str,root.children)
	}

}

func (t *trie) findPrefix(str string,children [26]*trieNode) {
	for _,v := range children{
		if v!=nil{
			if v.isEnd{
				fmt.Println(str+string(v.char))
			}else {
				t.findPrefix(str+string(v.char),v.children)
			}

		}
	}

}