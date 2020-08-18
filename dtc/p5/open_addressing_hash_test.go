package p5

import "testing"

func TestNewOpenHashTable(t *testing.T) {
	o := NewOpenHashTable(10, NewChapter())
	o.Insert("junmo", "a1")
	o.Insert("junmo1", "a2")
	o.Insert("junmo2", "a3")
	o.Delete("junmo2")
	o.print()
}
