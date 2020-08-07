package linked

import "testing"

func TestArrayQueue(t *testing.T) {
	aq := &arrayQueue{}
	queueTest(aq, t)
	queueTest(newLinkedQueue(), t)
	queueTest(newLoopArrayQueue(), t)

}

func queueTest(q Queue, t *testing.T) {
	end := 200
	for i := 10; i < end; i++ {
		q.Enqueue(i * i)
	}
	for i := 10; i < end; i++ {
		e, _ := q.Dequeue()
		if e != i*i {
			t.Fatalf("want:%d actual:%d", i*i, e)
		}
	}
}
