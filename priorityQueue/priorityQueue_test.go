package priorityQueue

import (
	"fmt"
	"reflect"
	"sort"
	"testing"
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
