package p5

import "fmt"

type hashEntry struct {
	key   string
	val   interface{}
	State int
}

const (
	OPEN_DELETE     = 1
	OPEN_LEGITIMATE = 0
)

type OpenHashTable struct {
	size int
	list []*hashEntry
	hash HashFunc
}

func NewOpenHashTable(size int, hashFunc HashFunc) *OpenHashTable {
	return &OpenHashTable{
		size: size,
		list: make([]*hashEntry, size),
		hash: hashFunc,
	}
}

func (o *OpenHashTable) getHashSlot(key string) int {
	return o.hash.Hash(key, o.size)
}

func (o *OpenHashTable) Find(key string) *hashEntry {
	entry := o.find(key)
	if entry != nil && entry.State != OPEN_DELETE {
		return entry
	} else {
		return nil
	}
}

func (o *OpenHashTable) find(key string) *hashEntry {
	base := o.getHashSlot(key)
	entry := o.list[base]
	index := 1
	for entry != nil {
		if entry.key != key {
			base += 2*index - 1
			base = base % o.size
			index++
			entry = o.list[base]
		} else {
			return entry
		}
	}
	return nil
}

func (o *OpenHashTable) Insert(key string, val interface{}) {
	base := o.getHashSlot(key)
	entry := o.list[base]
	index := 1
	for entry != nil {
		if entry.key != key {
			base += 2*index - 1
			base = base % o.size
			index++
			entry = o.list[base]
		} else {
			entry.val = val
			entry.State = OPEN_LEGITIMATE
			return
		}
	}
	o.list[base] = &hashEntry{key: key, val: val, State: OPEN_LEGITIMATE}
	return
}

func (o *OpenHashTable) Delete(key string) {
	entry := o.find(key)
	if entry != nil {
		entry.State = OPEN_DELETE
	}
}

func (o *OpenHashTable) print() {
	for k, v := range o.list {
		if v != nil {
			fmt.Printf("%d %s %v %d\n", k, v.key, v.val, v.State)
		}
	}
}
