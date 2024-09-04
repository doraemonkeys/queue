package priorityQueue

import (
	"fmt"
	"reflect"
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

func less1(a, b int) bool {
	return a > b
}

func TestNewPriorityQueue(t *testing.T) {
	que := New(less1)
	var Nums []int = []int{99, 67, 45, 22, 7, 84, 4, 4, 21, 2, 1}

	for _, v := range Nums {
		que.Push(v)
	}
	var elements []int
	for !que.IsEmpty() {
		elements = append(elements, que.Pop())
	}

	sort.Slice(Nums, func(i, j int) bool {
		return Nums[i] > Nums[j]
	})
	fmt.Println(Nums, elements)
	if !reflect.DeepEqual(elements, Nums) {
		t.Errorf("Expected %v, got %v", Nums, elements)
	}

}

func less2(a, b int) bool {
	return a < b
}

func TestNewPriorityQueue2(t *testing.T) {
	que := New(less2)
	var Nums []int = []int{99, 67, 45, 22, 7, 84, 4, 4, 21, 2, 1}

	for _, v := range Nums {
		que.Push(v)
	}
	var elements []int
	for !que.IsEmpty() {
		elements = append(elements, que.Pop())
	}

	sort.Slice(Nums, func(i, j int) bool {
		return Nums[i] < Nums[j]
	})
	fmt.Println(Nums, elements)
	if !reflect.DeepEqual(elements, Nums) {
		t.Errorf("Expected %v, got %v", Nums, elements)
	}

}

func lessInt(a, b int) bool {
	return a < b
}

func TestPriorityQueue(t *testing.T) {
	t.Run("New", func(t *testing.T) {
		pq := New(lessInt)
		assert.NotNil(t, pq)
		assert.Equal(t, 0, pq.Len())
		assert.True(t, pq.IsEmpty())
	})

	t.Run("NewOn", func(t *testing.T) {
		slice := []int{3, 1, 4, 1, 5, 9}
		pq := NewOn(slice, lessInt)
		assert.NotNil(t, pq)
		assert.Equal(t, 6, pq.Len())
		assert.Equal(t, 1, pq.Top())
	})

	t.Run("NewOf", func(t *testing.T) {
		pq := NewOf(lessInt, 3, 1, 4, 1, 5, 9)
		assert.NotNil(t, pq)
		assert.Equal(t, 6, pq.Len())
		assert.Equal(t, 1, pq.Top())
	})

	t.Run("Len and Cap", func(t *testing.T) {
		pq := NewOf(lessInt, 1, 2, 3)
		assert.Equal(t, 3, pq.Len())
		assert.GreaterOrEqual(t, pq.Cap(), 3)
	})

	t.Run("IsEmpty", func(t *testing.T) {
		pq := New(lessInt)
		assert.True(t, pq.IsEmpty())
		pq.Push(1)
		assert.False(t, pq.IsEmpty())
	})

	t.Run("Clear", func(t *testing.T) {
		pq := NewOf(lessInt, 1, 2, 3)
		pq.Clear()
		assert.True(t, pq.IsEmpty())
		assert.Equal(t, 0, pq.Len())
	})

	t.Run("Top", func(t *testing.T) {
		pq := NewOf(lessInt, 3, 1, 4)
		assert.Equal(t, 1, pq.Top())
		pq.Push(0)
		assert.Equal(t, 0, pq.Top())
	})

	t.Run("Push", func(t *testing.T) {
		pq := New(lessInt)
		pq.Push(3)
		pq.Push(1)
		pq.Push(4)
		assert.Equal(t, 3, pq.Len())
		assert.Equal(t, 1, pq.Top())
	})

	t.Run("Pop", func(t *testing.T) {
		pq := NewOf(lessInt, 3, 1, 4, 1, 5, 9)
		assert.Equal(t, 1, pq.Pop())
		assert.Equal(t, 1, pq.Pop())
		assert.Equal(t, 3, pq.Pop())
		assert.Equal(t, 3, pq.Len())
	})

	t.Run("Integration test", func(t *testing.T) {
		pq := New(lessInt)
		numbers := []int{3, 1, 4, 1, 5, 9, 2, 6, 5, 3, 5}
		for _, num := range numbers {
			pq.Push(num)
		}
		assert.Equal(t, 11, pq.Len())

		sortedNumbers := make([]int, 0, len(numbers))
		for !pq.IsEmpty() {
			sortedNumbers = append(sortedNumbers, pq.Pop())
		}
		assert.Equal(t, []int{1, 1, 2, 3, 3, 4, 5, 5, 5, 6, 9}, sortedNumbers)
		assert.True(t, pq.IsEmpty())
	})
}

func TestTopKPanicOnInvalidK(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		}
	}()

	less := func(a, b int) bool { return a < b }
	pq := NewOf(less, 5, 2, 7, 1, 3)
	pq.NewTopK(2) // This should panic
}

func TestNewTopK(t *testing.T) {
	pq := NewOf(func(a, b int) bool { return a < b }, 1, 2, 3, 4, 5)
	topK := pq.NewTopK(5)

	if topK.Len() != 5 {
		t.Errorf("Expected length 5, got %d", topK.Len())
	}

	if topK.k != 5 {
		t.Errorf("Expected k to be 5, got %d", topK.k)
	}
}

func TestNewTopKPanic(t *testing.T) {
	pq := NewOf(func(a, b int) bool { return a < b }, 1, 2, 3, 4, 5)

	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		}
	}()

	pq.NewTopK(4) // This should panic
}

func TestNewTopKOn(t *testing.T) {
	slice := []int{5, 2, 8, 1, 9, 3, 7}
	less := func(a, b int) bool { return a < b } // Min heap
	topK := NewTopKOn(slice, 5, less)

	if topK.Len() != 5 {
		t.Errorf("Expected length 5, got %d", topK.Len())
	}

	if topK.Top() != 1 {
		t.Errorf("Expected top element to be 1, got %d", topK.Top())
	}
}

func TestPushTopK(t *testing.T) {
	less := func(a, b int) bool { return a < b } // Min heap
	topK := NewTopKOn([]int{5, 2, 8, 1, 9}, 5, less)

	// Push an element smaller than the current top (should not be added)
	topK.Push(0)
	if topK.Len() != 5 || topK.Top() != 1 {
		t.Errorf("Unexpected state after pushing 0")
	}

	// Push an element larger than the current top (should be added)
	topK.Push(10)
	if topK.Len() != 5 || topK.Top() != 2 {
		t.Errorf("Unexpected state after pushing 10")
	}
}

func TestPushTopKMaxHeap(t *testing.T) {
	less := func(a, b int) bool { return a > b } // Max heap
	topK := NewTopKOn([]int{5, 2, 8, 1, 9}, 5, less)

	// Push an element larger than the current top (should not be added)
	topK.Push(10)
	if topK.Len() != 5 || topK.Top() != 9 {
		t.Errorf("Unexpected state after pushing 10")
	}

	// Push an element smaller than the current top (should be added)
	topK.Push(3)
	if topK.Len() != 5 || topK.Top() != 8 {
		t.Errorf("Unexpected state after pushing 3")
	}
}

func TestTopKMaintainsOrder(t *testing.T) {
	less := func(a, b int) bool { return a < b } // Min heap
	topK := NewTopKOn([]int{}, 3, less)

	elements := []int{4, 1, 7, 3, 8, 2, 6, 5}
	for _, e := range elements {
		topK.Push(e)
	}

	expected := []int{6, 7, 8}
	for i := 0; i < 3; i++ {
		if topK.Pop() != expected[i] {
			t.Errorf("Unexpected order of elements")
		}
	}
}

func TestTopKWithCustomType(t *testing.T) {
	type Person struct {
		Name string
		Age  int
	}

	less := func(a, b Person) bool { return a.Age < b.Age }
	topK := NewTopKOn([]Person{}, 3, less)

	people := []Person{
		{"Alice", 30},
		{"Bob", 25},
		{"Charlie", 35},
		{"David", 20},
		{"Eve", 40},
	}

	for _, p := range people {
		topK.Push(p)
	}

	if topK.Len() != 3 {
		t.Errorf("Expected length 3, got %d", topK.Len())
	}

	oldest := topK.Pop()
	if oldest.Age != 30 || oldest.Name != "Alice" {
		t.Errorf("Expected oldest person to be Alice, got %s", oldest.Name)
	}
}
