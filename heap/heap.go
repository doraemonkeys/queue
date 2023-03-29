package heap

import "github.com/Doraemonkeys/queue"

func heapSwap[T any](heap []T, i, j int) {
	heap[i], heap[j] = heap[j], heap[i]
}

// MakeHeap build a min-heap on slice array with compare function less.
//
// Complexity: O(len(array))
func MakeHeap[T any](array []T, less queue.LessFn[T]) {
	// heapify
	n := len(array)
	for i := n/2 - 1; i >= 0; i-- {
		heapDown(array, i, n, less)
	}
}

// IsHeap checks whether the elements in slice array are a min heap (accord to less).
//
// Complexity: O(len(array)).
func IsHeap[T any](array []T, less queue.LessFn[T]) bool {
	parent := 0
	for child := 1; child < len(array); child++ {
		if !less(array[parent], array[child]) {
			return false
		}

		if (child & 1) == 0 {
			parent++
		}

	}
	return true
}

// PushHeap pushes a element v into the heap.
//
// Complexity: O(log(len(*heap))).
func PushHeap[T any](heap *[]T, v T, less queue.LessFn[T]) {
	*heap = append(*heap, v)
	heapUp(*heap, len(*heap)-1, less)
}

// PopHeap removes and returns the minimum (according to less) element from the heap.
//
// Complexity: O(log n) where n = len(*heap).
func PopHeap[T any](heap *[]T, less queue.LessFn[T]) T {
	h := *heap
	n := len(h) - 1
	heapSwap(h, 0, n)
	heapDown(h, 0, n, less)
	*heap = h[0:n]
	return h[n]
}

// RemoveHeap removes and returns the element at index i from the heap.
//
// Complexity: is O(log(n)) where n = len(*heap).
func RemoveHeap[T any](heap *[]T, i int, less queue.LessFn[T]) T {
	h := *heap
	n := len(h) - 1
	if n != i {
		heapSwap(h, i, n)
		if !heapDown(h, i, n, less) {
			heapUp(h, i, less)
		}
	}
	*heap = h[0:n]
	return h[n]
}

func heapUp[T any](heap []T, j int, less queue.LessFn[T]) {
	for {
		i := (j - 1) / 2 // parent
		if i == j || !less(heap[j], heap[i]) {
			break
		}
		heapSwap(heap, i, j)
		j = i
	}
}

func heapDown[T any](heap []T, i0, n int, less queue.LessFn[T]) bool {
	i := i0
	for {
		j1 := 2*i + 1
		if j1 >= n || j1 < 0 { // j1 < 0 after int overflow
			break
		}
		j := j1 // left child
		if j2 := j1 + 1; j2 < n && less(heap[j2], heap[j1]) {
			j = j2 // = 2*i + 2  // right child
		}
		if !less(heap[j], heap[i]) {
			break
		}
		heapSwap(heap, i, j)
		i = j
	}
	return i > i0
}
