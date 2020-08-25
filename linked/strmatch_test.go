package linked

import "testing"

func TestBFAndRK_Match(t *testing.T) {
	s := "junmocsqandfdjhfjhasjakdshgshddsbjsdhsjds"
	p := "bjsdhs"
	index := NewBF().Match(s, p)
	if index == -1 {
		t.Log("BF match failed")
	}
	t.Log(index)

	r := NewRk()
	i := r.Match(s, p)
	t.Log(i)
}
