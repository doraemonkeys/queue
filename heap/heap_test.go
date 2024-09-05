package heap_test

import (
	"reflect"
	"testing"

	"github.com/doraemonkeys/queue/heap"
)

type IntHeap []int

func (h IntHeap) Less(i, j int) bool {
	return i < j
}

func TestMakeHeap(t *testing.T) {
	data := []int{5, 3, 8, 4, 1, 2}
	heap.MakeHeap(data, IntHeap(data).Less)

	if !heap.IsHeap(data, IntHeap(data).Less) {
		t.Errorf("MakeHeap failed: %v is not a valid heap", data)
	}
}

func TestIsHeap(t *testing.T) {
	data := []int{1, 3, 2, 4, 5, 8}
	if !heap.IsHeap(data, IntHeap(data).Less) {
		t.Errorf("IsHeap failed: %v should be a valid heap", data)
	}

	data = []int{3, 1, 2, 4, 5, 8}
	if heap.IsHeap(data, IntHeap(data).Less) {
		t.Errorf("IsHeap failed: %v should not be a valid heap", data)
	}
}

func TestPushHeap(t *testing.T) {
	data := []int{1, 3, 2, 4, 5, 8}
	heap.PushHeap(&data, 0, IntHeap(data).Less)

	if !heap.IsHeap(data, IntHeap(data).Less) {
		t.Errorf("PushHeap failed: %v is not a valid heap", data)
	}
}

func TestPopHeap(t *testing.T) {
	data := []int{1, 3, 2, 4, 5, 8}
	heap.MakeHeap(data, IntHeap(data).Less)
	min := heap.PopHeap(&data, IntHeap(data).Less)

	if min != 1 {
		t.Errorf("PopHeap failed: expected 1, got %d", min)
	}
	if !heap.IsHeap(data, IntHeap(data).Less) {
		t.Errorf("PopHeap failed: %v is not a valid heap", data)
	}
}

func TestRemoveHeap(t *testing.T) {
	data := []int{1, 3, 2, 4, 5, 8}
	heap.MakeHeap(data, IntHeap(data).Less)
	removed := heap.RemoveHeap(&data, 2, IntHeap(data).Less)

	if removed != 2 {
		t.Errorf("RemoveHeap failed: expected 2, got %d", removed)
	}
	if !heap.IsHeap(data, IntHeap(data).Less) {
		t.Errorf("RemoveHeap failed: %v is not a valid heap", data)
	}
}

func TestHeapPeek(t *testing.T) {
	data := []int{1, 3, 2, 4, 5, 8}
	heap.MakeHeap(data, IntHeap(data).Less)
	peek := heap.HeapTop(data)

	if peek != 1 {
		t.Errorf("HeapPeek failed: expected 1, got %d", peek)
	}
}

func TestLevelOrder(t *testing.T) {
	tests := []struct {
		name     string
		heap     []int
		expected [][]int
	}{
		{
			name:     "Empty heap",
			heap:     []int{},
			expected: nil,
		},
		{
			name:     "Single element heap",
			heap:     []int{1},
			expected: [][]int{{1}},
		},
		{
			name:     "Complete binary heap",
			heap:     []int{1, 2, 3, 4, 5, 6, 7},
			expected: [][]int{{1}, {2, 3}, {4, 5, 6, 7}},
		},
		{
			name:     "Incomplete binary heap",
			heap:     []int{1, 2, 3, 4, 5},
			expected: [][]int{{1}, {2, 3}, {4, 5}},
		},
		{
			name:     "Large heap",
			heap:     []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15},
			expected: [][]int{{1}, {2, 3}, {4, 5, 6, 7}, {8, 9, 10, 11, 12, 13, 14, 15}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := heap.LevelOrder(tt.heap)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("LevelOrder() = %v, want %v", result, tt.expected)
			}
		})
	}
}

// TestLevelOrderString tests the LevelOrder function with a string slice
func TestLevelOrderString(t *testing.T) {
	h := []string{"A", "B", "C", "D", "E"}
	expected := [][]string{{"A"}, {"B", "C"}, {"D", "E"}}

	result := heap.LevelOrder(h)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("LevelOrder() = %v, want %v", result, expected)
	}
}

func TestPushHeapTopKWithZeroOrNegativeK(t *testing.T) {
	less := func(a, b int) bool {
		return a < b
	}

	tests := []struct {
		name    string
		initial []int
		value   int
		k       int
	}{
		{
			name:    "k is zero",
			initial: []int{3, 2, 1},
			value:   4,
			k:       0,
		},
		{
			name:    "k is negative",
			initial: []int{3, 2, 1},
			value:   4,
			k:       -1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer func() {
				if r := recover(); r == nil {
					t.Errorf("expected panic for k = %d, but did not panic", tt.k)
				}
			}()

			heapArray := make([]int, len(tt.initial))
			copy(heapArray, tt.initial)

			heap.PushHeapTopK(&heapArray, tt.value, less, tt.k)
		})
	}
}
