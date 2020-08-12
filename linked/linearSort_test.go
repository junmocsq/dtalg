package linked

import (
	"math/rand"
	"testing"
	"time"
)

func TestLinear(t *testing.T) {

	counting := &CountingSort{}
	a := make([]int, 0)
	for i := 0; i < 15; i++ {
		rand.Seed(time.Now().UnixNano())
		a = append(a, rand.Intn(100)-50)
	}
	t.Log(a)
	counting.Sort(a)
	t.Log(a)

	b:= []byte("123hghgFFGash2hGGsd")
	(&NumberCharacterCategory{}).Sort(b)
}
