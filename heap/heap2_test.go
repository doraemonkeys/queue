package heap

import (
	"github.com/doraemonkeys/queue"
	"reflect"
	"testing"
)

func TestIsHeap(t *testing.T) {
	tests := []struct {
		name     string
		array    []int
		less     queue.LessFn[int]
		expected bool
	}{
		{
			name:     "Empty heap",
			array:    []int{},
			less:     func(a, b int) bool { return a < b },
			expected: true,
		},
		{
			name:     "Single element heap",
			array:    []int{1},
			less:     func(a, b int) bool { return a < b },
			expected: true,
		},
		{
			name:     "Valid min heap",
			array:    []int{1, 3, 2, 4, 5, 6, 7},
			less:     func(a, b int) bool { return a < b },
			expected: true,
		},
		{
			name:     "Valid max heap",
			array:    []int{7, 5, 6, 4, 2, 1, 3},
			less:     func(a, b int) bool { return a > b },
			expected: true,
		},
		{
			name:     "min heap",
			array:    []int{1, 3, 2, 5, 4, 6, 7},
			less:     func(a, b int) bool { return a < b },
			expected: true,
		},
		{
			name:     "max heap",
			array:    []int{7, 5, 6, 4, 3, 1, 2},
			less:     func(a, b int) bool { return a > b },
			expected: true,
		},
		{
			name:     "Heap with duplicate elements",
			array:    []int{1, 2, 2, 3, 3, 3, 3},
			less:     func(a, b int) bool { return a < b },
			expected: true,
		},
		{
			name:     "Heap with negative numbers",
			array:    []int{-10, -5, -8, -1, -2, -6, -7},
			less:     func(a, b int) bool { return a < b },
			expected: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := IsHeap(tt.array, tt.less)
			if result != tt.expected {
				t.Errorf("IsHeap() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestPushHeapTopK(t *testing.T) {
	tests := []struct {
		name     string
		heap     []int
		v        int
		k        int
		expected []int
	}{
		{
			name:     "Empty heap, k=1",
			heap:     []int{},
			v:        5,
			k:        1,
			expected: []int{5},
		},
		{
			name:     "Heap with one element, k=1, new element smaller",
			heap:     []int{5},
			v:        3,
			k:        1,
			expected: []int{5},
		},
		{
			name:     "Heap with one element, k=1, new element larger",
			heap:     []int{5},
			v:        7,
			k:        1,
			expected: []int{7},
		},
		{
			name:     "Heap with multiple elements, k=3, new element in middle",
			heap:     []int{3, 5, 7},
			v:        4,
			k:        3,
			expected: []int{4, 5, 7},
		},
		{
			name:     "Heap with multiple elements, k=3, new element largest",
			heap:     []int{3, 5, 7},
			v:        8,
			k:        3,
			expected: []int{5, 7, 8},
		},
		{
			name:     "Heap with multiple elements, k=3, new element smallest",
			heap:     []int{3, 5, 7},
			v:        1,
			k:        3,
			expected: []int{3, 5, 7},
		},
		{
			name:     "Heap with k-1 elements, k=5",
			heap:     []int{2, 4, 6, 8},
			v:        5,
			k:        5,
			expected: []int{2, 4, 5, 6, 8},
		},
		{
			name:     "Heap with k elements, k=5, new element not in top k",
			heap:     []int{1, 3, 5, 7, 9},
			v:        4,
			k:        5,
			expected: []int{3, 4, 5, 7, 9},
		},
		{
			name:     "Heap with k elements, k=5, new element in top k",
			heap:     []int{1, 3, 5, 7, 9},
			v:        6,
			k:        5,
			expected: []int{3, 5, 6, 7, 9},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			heap := tt.heap
			PushHeapTopK(&heap, tt.v, func(a, b int) bool { return a < b }, tt.k)
			var actual []int
			for len(heap) != 0 {
				v := PopHeap(&heap, func(a, b int) bool { return a < b })
				actual = append(actual, v)
			}
			if !reflect.DeepEqual(actual, tt.expected) {
				t.Errorf("PushHeapTopK() = %v, want %v", actual, tt.expected)
			}
			if !IsHeap(heap, func(a, b int) bool { return a < b }) {
				t.Errorf("Result is not a valid heap: %v", heap)
			}
		})
	}
}

func TestPushHeapTopKWithCustomComparator(t *testing.T) {
	type Person struct {
		name string
		age  int
	}

	comparator := func(a, b Person) bool {
		return a.age < b.age // Min heap based on age
	}

	tests := []struct {
		name     string
		heap     []Person
		v        Person
		k        int
		expected []Person
	}{
		{
			name: "Custom type heap, k=3, new element in middle",
			heap: []Person{
				{"Alice", 25},
				{"Bob", 30},
				{"Charlie", 35},
			},
			v: Person{"David", 28},
			k: 3,
			expected: []Person{
				{"David", 28},
				{"Bob", 30},
				{"Charlie", 35},
			},
		},
		{
			name: "Custom type heap, k=3, new element oldest",
			heap: []Person{
				{"Alice", 25},
				{"Bob", 30},
				{"Charlie", 35},
			},
			v: Person{"David", 40},
			k: 3,
			expected: []Person{
				{"Bob", 30},
				{"Charlie", 35},
				{"David", 40},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			heap := tt.heap
			PushHeapTopK(&heap, tt.v, comparator, tt.k)
			var actual []Person
			for len(heap) != 0 {
				v := PopHeap(&heap, comparator)
				actual = append(actual, v)
			}
			if !reflect.DeepEqual(actual, tt.expected) {
				t.Errorf("PushHeapTopK() = %v, want %v", actual, tt.expected)
			}
			if !IsHeap(heap, comparator) {
				t.Errorf("Result is not a valid heap: %v", heap)
			}
		})
	}
}

func TestPushHeapTopKWithLargeK(t *testing.T) {
	heap := make([]int, 1000)
	for i := 0; i < 1000; i++ {
		heap[i] = 1000 - i
	}
	MakeHeap(heap, func(a, b int) bool { return a < b })

	if !IsHeap(heap, func(a, b int) bool { return a < b }) {
		t.Errorf("Result is not a valid heap")
	}

	PushHeapTopK(&heap, 500, func(a, b int) bool { return a < b }, 2000)

	if len(heap) != 1001 {
		t.Errorf("Expected heap length to be 1001, got %d", len(heap))
	}

	if !IsHeap(heap, func(a, b int) bool { return a < b }) {
		t.Errorf("Result is not a valid heap")
	}

	// Check if 500 is in the heap
	found := false
	for _, v := range heap {
		if v == 500 {
			found = true
			break
		}
	}
	if !found {
		t.Errorf("Expected 500 to be in the heap")
	}
}

func TestPushHeapTopK2(t *testing.T) {
	// Helper function to create a min-heap
	lessInt := func(a, b int) bool { return a < b }

	tests := []struct {
		name     string
		heap     []int
		val      int
		k        int
		expected []int
	}{
		{
			name:     "Push to empty heap",
			heap:     []int{},
			val:      5,
			k:        3,
			expected: []int{5},
		},
		{
			name:     "Push to non-full heap",
			heap:     []int{3, 7},
			val:      5,
			k:        3,
			expected: []int{3, 5, 7},
		},
		{
			name:     "Push smaller value to full heap",
			heap:     []int{3, 7, 9},
			val:      5,
			k:        3,
			expected: []int{5, 7, 9},
		},
		{
			name:     "Push larger value to full heap",
			heap:     []int{3, 7, 9},
			val:      10,
			k:        3,
			expected: []int{7, 9, 10},
		},
		{
			name:     "Push to heap with k=1",
			heap:     []int{5},
			val:      3,
			k:        1,
			expected: []int{5},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			heap := tt.heap
			PushHeapTopK(&heap, tt.val, queue.LessFn[int](lessInt), tt.k)

			// Check if the result is correct
			var result []int
			for len(heap) != 0 {
				result = append(result, PopHeap(&heap, lessInt))
			}
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("PushHeapTopK() = %v, want %v", result, tt.expected)
			}
			// Check if the heap is valid
			if !IsHeap(heap, queue.LessFn[int](lessInt)) {
				t.Errorf("Result is not a valid heap")
			}
		})
	}
}

func TestPushHeapTopKPanic(t *testing.T) {
	lessInt := func(a, b int) bool { return a < b }

	heap := []int{1, 2, 3, 4}
	k := 3

	defer func() {
		if r := recover(); r != nil {
			t.Errorf("PushHeapTopK() panicked, error: %v", r)
		}
	}()

	PushHeapTopK(&heap, 5, queue.LessFn[int](lessInt), k)

	expected := []int{3, 4, 5}
	result := []int{}
	for len(heap) != 0 {
		result = append(result, PopHeap(&heap, lessInt))
	}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("PushHeapTopK() = %v, want %v", result, expected)
	}
}

func TestPushHeapTopKWithCustomType(t *testing.T) {
	type Person struct {
		Name string
		Age  int
	}

	lessPerson := func(a, b Person) bool { return a.Age < b.Age }

	heap := []Person{
		{"Alice", 30},
		{"Bob", 40},
	}
	k := 3

	PushHeapTopK(&heap, Person{"Charlie", 35}, queue.LessFn[Person](lessPerson), k)

	expected := []Person{
		{"Alice", 30},
		{"Bob", 40},
		{"Charlie", 35},
	}

	if !reflect.DeepEqual(heap, expected) {
		t.Errorf("PushHeapTopK() = %v, want %v", heap, expected)
	}

	if !IsHeap(heap, queue.LessFn[Person](lessPerson)) {
		t.Errorf("Result is not a valid heap")
	}
}

func TestPushHeapTopKMaintainsOrder(t *testing.T) {
	lessInt := func(a, b int) bool { return a < b }

	heap := []int{1, 3, 5}
	k := 3

	for i := 0; i < 100; i++ {
		PushHeapTopK(&heap, i, queue.LessFn[int](lessInt), k)
		if !IsHeap(heap, queue.LessFn[int](lessInt)) {
			t.Errorf("Heap property violated after pushing %d", i)
		}
		if len(heap) > k {
			t.Errorf("Heap size exceeded k after pushing %d", i)
		}
	}

	expected := []int{97, 98, 99}
	if !reflect.DeepEqual(heap, expected) {
		t.Errorf("Final heap = %v, want %v", heap, expected)
	}
}

func TestPushHeapTopK3(t *testing.T) {
	// Define a less function for integers (min heap)
	lessInt := func(a, b int) bool { return a < b }

	tests := []struct {
		name     string
		heap     []int
		v        int
		k        int
		expected []int
	}{
		{
			name:     "Empty heap, k=1",
			heap:     []int{},
			v:        5,
			k:        1,
			expected: []int{5},
		},
		{
			name:     "Heap with one element, k=1, new element smaller",
			heap:     []int{5},
			v:        3,
			k:        1,
			expected: []int{5},
		},
		{
			name:     "Heap with one element, k=1, new element larger",
			heap:     []int{5},
			v:        7,
			k:        1,
			expected: []int{7},
		},
		{
			name:     "Heap with multiple elements, k=3, new element in middle",
			heap:     []int{3, 7, 5},
			v:        4,
			k:        3,
			expected: []int{4, 7, 5},
		},
		{
			name:     "Heap with multiple elements, k=3, new element largest",
			heap:     []int{3, 7, 5},
			v:        8,
			k:        3,
			expected: []int{5, 7, 8},
		},
		{
			name:     "Heap with more elements than k, k=3",
			heap:     []int{2, 4, 6, 8, 10},
			v:        5,
			k:        3,
			expected: []int{6, 8, 10},
		},
		{
			name:     "Heap with fewer elements than k, k=5",
			heap:     []int{3, 5},
			v:        4,
			k:        5,
			expected: []int{3, 5, 4},
		},
		{
			name:     "Heap with k elements, k=5, new element not in top k",
			heap:     []int{1, 3, 5, 7, 9},
			v:        4,
			k:        5,
			expected: []int{3, 4, 5, 7, 9},
		},
		{
			name:     "Heap with k elements, k=3, new element in top k",
			heap:     []int{1, 3, 5, 7, 9},
			v:        6,
			k:        3,
			expected: []int{6, 7, 9},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			heap := tt.heap
			PushHeapTopK(&heap, tt.v, queue.LessFn[int](lessInt), tt.k)

			if !reflect.DeepEqual(heap, tt.expected) {
				t.Errorf("PushHeapTopK() = %v, want %v", heap, tt.expected)
			}

			// Additional check to ensure the result is a valid heap
			if !IsHeap(heap, queue.LessFn[int](lessInt)) {
				t.Errorf("Result is not a valid heap: %v", heap)
			}
		})
	}
}
