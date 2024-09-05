package priorityQueue

import (
	"github.com/doraemonkeys/queue"
	"github.com/doraemonkeys/queue/heap"
)

// PQueue is an queue with priority.
type PQueue[T any] struct {
	heap []T
	less queue.LessFn[T]
}

// PQueueTopK is a priority queue that maintains the top k elements.
type PQueueTopK[T any] struct {
	*PQueue[T]
	k int
}

// ToTopK creates a new priority queue that maintains the top k elements.
// It returns a object that wraps the existing priority queue.
//
// If the heap exceeds size k, for a min heap, the smallest element is removed,
// otherwise the largest element is removed.
func (pq *PQueue[T]) ToTopK(k int) *PQueueTopK[T] {
	if k <= 0 {
		panic("k should be a positive integer")
	}
	for k < pq.Len() {
		pq.Pop()
	}
	return &PQueueTopK[T]{PQueue: pq, k: k}
}

// Push adds an element to the priority queue while maintaining only the top k elements.
// It returns true if the element is added, otherwise false.
//
// If the heap is a min heap, Push will keep the k largest elements,
// otherwise it will keep the k smallest elements.
func (pq *PQueueTopK[T]) Push(v T) bool {
	return heap.PushHeapTopK(&pq.heap, v, pq.less, pq.k)
}

// New creates an empty priority object.
func New[T any](less queue.LessFn[T]) *PQueue[T] {
	pq := &PQueue[T]{}
	pq.less = less
	return pq
}

// NewOn creates a new priority object on the specified slices.
// The slice become a heap after the call.
func NewOn[T any](slice []T, less queue.LessFn[T]) *PQueue[T] {
	heap.MakeHeap(slice, less)
	pq := &PQueue[T]{}
	pq.heap = slice
	pq.less = less
	return pq
}

// NewOf creates a new priority object with specified initial elements.
func NewOf[T any](less queue.LessFn[T], elements ...T) *PQueue[T] {
	return NewOn(elements, less)
}

// Len returns the number of elements in the priority queue.
func (pq *PQueue[T]) Len() int {
	return len(pq.heap)
}

// Cap returns the capacity of the priority queue.
func (pq *PQueue[T]) Cap() int {
	return cap(pq.heap)
}

// IsEmpty checks whether priority queue has no elements.
func (pq *PQueue[T]) IsEmpty() bool {
	return len(pq.heap) == 0
}

// Clear clear the priority queue.
func (pq *PQueue[T]) Clear() {
	pq.heap = pq.heap[0:0]
}

// Top returns the top element in the priority queue.
func (pq *PQueue[T]) Top() T {
	return pq.heap[0]
}

// Push pushes the given element v to the priority queue.
func (pq *PQueue[T]) Push(v T) {
	heap.PushHeap(&pq.heap, v, pq.less)
}

// Pop removes the top element in the priority queue.
func (pq *PQueue[T]) Pop() T {
	return heap.PopHeap(&pq.heap, pq.less)
}
