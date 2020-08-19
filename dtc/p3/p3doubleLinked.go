package p3

type doubleLinkedNode struct {
	ele      int
	next     *doubleLinkedNode
	previous *doubleLinkedNode
}

type DoubleLinkedLister interface {
	MakeEmpty()
	IsEmpty() bool
	IsLast(position *doubleLinkedNode) bool
	Find(ele int) *doubleLinkedNode
	Delete(ele int) bool
	Insert(ele int, position *doubleLinkedNode) bool
	Add(ele int) bool
	Header() *doubleLinkedNode
	First() *doubleLinkedNode
	Retrieve(position *doubleLinkedNode) int
}

func NewDoubleLinkedList() *doubleLinkedNode {
	return &doubleLinkedNode{}
}

func (d *doubleLinkedNode) MakeEmpty() {
	d.next = nil
}

func (d *doubleLinkedNode) IsEmpty() bool {
	return d.next == nil
}
func (d *doubleLinkedNode) IsLast(position *doubleLinkedNode) bool {
	return position.next == nil
}
func (d *doubleLinkedNode) Find(ele int) *doubleLinkedNode {
	tmp := d.next
	for tmp != nil && tmp.ele != ele {
		tmp = tmp.next
	}
	return tmp
}
func (d *doubleLinkedNode) Delete(ele int) bool {
	node := d.Find(ele)
	if node != nil {
		if d.IsLast(node) {
			// 最后一个元素
			node.previous.next = nil
		} else {
			pre := node.previous
			next := node.next
			pre.next = next
			next.previous = pre
		}
		return true
	}
	return false
}
func (d *doubleLinkedNode) Insert(ele int, position *doubleLinkedNode) bool {
	node := &doubleLinkedNode{ele: ele}
	if d.IsLast(position) {
		position.next = node
		node.previous = position
	} else {
		next := position.next
		position.next = node
		node.previous = position
		node.next = next
		next.previous = node
	}
	return true

}
func (d *doubleLinkedNode) Add(ele int) bool {
	tmp := d
	for tmp.next != nil {
		tmp = tmp.next
	}
	return d.Insert(ele, tmp)
}
func (d *doubleLinkedNode) Header() *doubleLinkedNode {
	return d
}
func (d *doubleLinkedNode) First() *doubleLinkedNode {
	return d.next
}
func (d *doubleLinkedNode) Retrieve(position *doubleLinkedNode) int {
	return position.ele
}
