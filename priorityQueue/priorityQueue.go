package priorityQueue

import (
	"github.com/doraemonkeys/queue"
	"github.com/doraemonkeys/queue/heap"
)

// PriorityQueue is an queue with priority.
type PriorityQueue[T any] struct {
	heap []T
	less queue.LessFn[T]
}

// PriorityQueueTopK is a priority queue that maintains the top k elements.
type PriorityQueueTopK[T any] struct {
	*PriorityQueue[T]
	k int
}

// NewTopK creates a new PriorityQueueTopK from an existing PriorityQueue.
// It requires that k is greater than or equal to the current number of elements.
func (pq *PriorityQueue[T]) NewTopK(k int) *PriorityQueueTopK[T] {
	if k < len(pq.heap) {
		panic("PriorityQueueTopK: k must be greater than or equal to the current number of elements")
	}
	return &PriorityQueueTopK[T]{PriorityQueue: pq, k: k}
}

// NewTopKOn creates a new PriorityQueueTopK from a slice.
// If the length of the slice is greater than k, only the first k elements are used.
func NewTopKOn[T any](slice []T, k int, less queue.LessFn[T]) *PriorityQueueTopK[T] {
	if k < len(slice) {
		slice = slice[0:k]
	}
	return &PriorityQueueTopK[T]{PriorityQueue: NewOn(slice, less), k: k}
}

// Push adds an element to the priority queue, maintaining only the top k elements.
// If the queue already contains k elements, the new element is added only if it is
// greater (in a min-heap) or lesser (in a max-heap) than the current top element.
func (pq *PriorityQueueTopK[T]) Push(v T) {
	heap.PushHeapTopK(&pq.heap, v, pq.less, pq.k)
}

// New creates an empty priority object.
func New[T any](less queue.LessFn[T]) *PriorityQueue[T] {
	pq := &PriorityQueue[T]{}
	pq.less = less
	return pq
}

// NewOn creates a new priority object on the specified slices.
// The slice become a heap after the call.
func NewOn[T any](slice []T, less queue.LessFn[T]) *PriorityQueue[T] {
	heap.MakeHeap(slice, less)
	pq := &PriorityQueue[T]{}
	pq.heap = slice
	pq.less = less
	return pq
}

// NewOf creates a new priority object with specified initial elements.
func NewOf[T any](less queue.LessFn[T], elements ...T) *PriorityQueue[T] {
	return NewOn(elements, less)
}

// Len returns the number of elements in the priority queue.
func (pq *PriorityQueue[T]) Len() int {
	return len(pq.heap)
}

// Cap returns the capacity of the priority queue.
func (pq *PriorityQueue[T]) Cap() int {
	return cap(pq.heap)
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
	heap.PushHeap(&pq.heap, v, pq.less)
}

// Pop removes the top element in the priority queue.
func (pq *PriorityQueue[T]) Pop() T {
	return heap.PopHeap(&pq.heap, pq.less)
}
