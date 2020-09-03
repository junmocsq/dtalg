package linked

import (
	"math/rand"
	"testing"
	"time"
)

func TestNewEightQueen(t *testing.T) {
	t.Log(NewEightQueen().EightQueens())

	b := NewBagZeroOne(100)
	arr := []int{}
	for i := 0; i < 10; i++ {
		rand.Seed(time.Now().UnixNano())
		arr = append(arr, rand.Intn(30))
	}
	t.Log(arr)
	t.Log(b.BagZeroOne(arr))
	t.Log(b.res)

	p := NewPattern("abc*?ddd")
	t.Log(p.match("abc*?ddd"))
	t.Log(p.match("ababcddd"))
}
