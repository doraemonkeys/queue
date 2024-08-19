package arrayQueue

import (
	"testing"

	aq "github.com/emirpasic/gods/queues/arrayqueue"
	cq "github.com/emirpasic/gods/queues/circularbuffer"
)

func BenchmarkMyQueue(b *testing.B) {
	var myque *Queue[int] = New[int]()
	myque.Resize(1000000)
	for i := 0; i < b.N; i++ {
		myque.Push(i)
	}
	for !myque.Empty() {
		myque.Pop()
	}
}

func BenchmarkArrayQueue(b *testing.B) {
	var que = aq.New()
	N := 100000
	for i := 0; i < N; i++ {
		que.Enqueue(i)
	}
	for !que.Empty() {
		que.Dequeue()
	}
}

func BenchmarkQueue(b *testing.B) {
	var que = cq.New(1000000)
	for i := 0; i < b.N; i++ {
		que.Enqueue(i)
	}
	for !que.Empty() {
		que.Dequeue()
	}
}
