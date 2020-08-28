package linked

import (
	"fmt"
	"testing"
)

func TestNewTrie(t *testing.T) {
	trie := NewTrie()
	fmt.Println(trie)
	trie.Insert("how")
	trie.Insert("hi")
	trie.Insert("her")
	trie.Insert("hello")
	trie.Insert("hello world")
	trie.Insert("hella")
	trie.Insert("hellb")
	trie.Insert("hellc")
	trie.Insert("so")
	trie.Insert("see")
	trie.find("h")
}
