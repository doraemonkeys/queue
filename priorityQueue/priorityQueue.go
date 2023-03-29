package priorityQueue

import (
	"github.com/Doraemonkeys/queue"
	"github.com/Doraemonkeys/queue/heap"
)

// PriorityQueue is an queue with priority.
type PriorityQueue[T any] struct {
	heap []T
	impl pqImpl[T]
}

// New creates an empty priority object.
func New[T any](less queue.LessFn[T]) *PriorityQueue[T] {
	pq := pqFunc[T]{}
	pq.impl = (pqImpl[T])(&pq)
	pq.less = less
	return &pq.PriorityQueue
}

// NewOn creates a new priority object on the specified slices.
// The slice become a heap after the call.
func NewOn[T any](slice []T, less queue.LessFn[T]) *PriorityQueue[T] {
	heap.MakeHeap(slice, less)
	pq := pqFunc[T]{}
	pq.heap = slice
	pq.impl = pqImpl[T](&pq)
	pq.less = less
	return &pq.PriorityQueue
}

// NewOf creates a new priority object with specified initial elements.
func NewOf[T any](less queue.LessFn[T], elements ...T) *PriorityQueue[T] {
	return NewOn(elements, less)
}

// Len returns the number of elements in the priority queue.
func (pq *PriorityQueue[T]) Len() int {
	return len(pq.heap)
}

// IsEmpty checks whether priority queue has no elements.
func (pq *PriorityQueue[T]) IsEmpty() bool {
	return len(pq.heap) == 0
}

// Clear clear the priority queue.
func (pq *PriorityQueue[T]) Clear() {
	pq.heap = pq.heap[0:0]
}

// Top returns the top element in the priority queue.
func (pq *PriorityQueue[T]) Top() T {
	return pq.heap[0]
}

// Push pushes the given element v to the priority queue.
func (pq *PriorityQueue[T]) Push(v T) {
	pq.impl.Push(v)
}

// Pop removes the top element in the priority queue.
func (pq *PriorityQueue[T]) Pop() T {
	return pq.impl.Pop()
}

type pqImpl[T any] interface {
	Push(v T)
	Pop() T
}

// funcHeap is a min-heap of T compared with less.
type pqFunc[T any] struct {
	PriorityQueue[T]
	less queue.LessFn[T]
}

func (pq *pqFunc[T]) Push(v T) {
	heap.PushHeap(&pq.heap, v, pq.less)
}

func (pq *pqFunc[T]) Pop() T {
	return heap.PopHeap(&pq.heap, pq.less)
}
