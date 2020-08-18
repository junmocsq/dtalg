package p5

import (
	"fmt"
	"testing"
)

func TestHashFunc(t *testing.T) {
	h := NewHorner()
	fmt.Println(h.Hash("junmo", 100))
}
