package heap

import (
	"github.com/doraemonkeys/queue"
)

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
		if less(array[child], array[parent]) {
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

// PushHeapTopK pushes a element v into the heap, keeping the top k elements.
// If the heap is a min heap, PushHeapTopK will keep the top k elements maximum,
// otherwise it will keep the top k elements minimum.
//
// Note that k should be a positive integer and k is the maximum length of the input heap.
// If k is less than the current length of the heap, it will panic.
//
// Complexity: O(log k).
func PushHeapTopK[T any](heap *[]T, v T, less queue.LessFn[T], k int) {
	if len(*heap) < k {
		PushHeap(heap, v, less)
		return
	}
	h := *heap
	if len(h) > k {
		panic("PushHeapTopK: heap length > k")
	}
	if !less(v, HeapTop(h)) {
		h[0] = v
		heapDown(h, 0, k, less)
	}
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

func HeapTop[T any](heap []T) T {
	return heap[0]
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

// heapDown maintains the heap property by moving the element at index i0
// downwards in the heap until it is in the correct position.
//
// Parameters:
// heap - a slice of type T representing the heap.
// i0 - the index of the element to be moved down.
// n - the number of elements in the heap.
//
// Returns:
// A boolean indicating whether the element at index i0 was moved.
//
// The function works as follows:
//  1. Start with the element at index i0.
//  2. Loop until the element is in the correct position:
//     a. Calculate the index of the left child (j1).
//     b. If j1 is out of bounds, break the loop.
//     c. Assume the left child (j1) is the smaller child.
//     d. Check if there is a right child (j2) and if it is smaller than the left child.
//     If so, update j to be j2.
//     e. If the current element is less than or equal to the smallest child, break the loop.
//     f. Swap the current element with the smallest child.
//     g. Update the current index to be the index of the smallest child.
//  3. Return true if the element was moved from its original position.
func heapDown[T any](heap []T, i0, n int, less queue.LessFn[T]) bool {
	i := i0
	for {
		j1 := 2*i + 1          // left child
		if j1 >= n || j1 < 0 { // j1 < 0 after int overflow
			break
		}
		j := j1
		if j2 := j1 + 1; j2 < n && less(heap[j2], heap[j1]) {
			j = j2 // = 2*i + 2  right child
		}
		if !less(heap[j], heap[i]) {
			break
		}
		heapSwap(heap, i, j)
		i = j
	}
	return i > i0
}

// LevelOrder performs a level-order traversal of a heap represented as a slice.
// It returns a 2D slice where each inner slice represents a level of the tree.
func LevelOrder[T any](heap []T) [][]T {
	if len(heap) == 0 {
		return nil
	}

	var result [][]T
	level := 0
	for start := 0; start < len(heap); {
		// Calculate the end index for the current level
		// The total nodes up to level n is 2^(n+1)-1
		end := min(len(heap), 1<<(level+1)-1)
		result = append(result, heap[start:end])
		start = end
		// Move to the next level
		level++
	}
	return result
}
