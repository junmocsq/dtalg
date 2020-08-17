package p3

import (
	"math/rand"
	"testing"
	"time"
)

func TestNewRadix(t *testing.T) {
	arr := make([]int, 100)
	for i := 0; i < 100; i++ {
		rand.Seed(time.Now().UnixNano())
		arr[i] = rand.Intn(10000)
	}
	t.Log(arr)
	a := NewRadix(arr)
	a.radixSort()
	t.Log(a)
}
