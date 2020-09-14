package p2

import "testing"

func TestNewInversion(t *testing.T) {
	i := NewInversion()
	arr := []int{}
	for n := 10; n > 0; n-- {
		arr = append(arr, n)
	}
	t.Log(i.iv(arr))
	t.Log(arr)
}
