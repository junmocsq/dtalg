package p6

import (
	"fmt"
	"testing"
)

func TestBinomialQueue(t *testing.T) {
	binomial := newBinQueue(15)

	for i := 0; i < 27; i++ {
		binomial.Insert(i + 1)
	}
	t.Log(binomial.trees)
	for i := 0; i < 27; i++ {
		fmt.Println(binomial.DeleteMin())
	}
}
