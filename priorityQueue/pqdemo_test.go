package priorityQueue

import "fmt"

func Example() {
	pq := New(func(a, b int) bool {
		return a < b
	})
	pq.Push(1)
	pq.Push(5)
	pq.Push(3)
	for !pq.IsEmpty() {
		fmt.Println(pq.Pop())
	}
	// Output:
	// 1
	// 3
	// 5
}
