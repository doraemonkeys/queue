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
