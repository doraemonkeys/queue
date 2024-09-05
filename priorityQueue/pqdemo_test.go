package priorityQueue

import "fmt"

func Example() {
	pq := New(func(a, b int) bool {
		return a < b
	})
	pq.Push(1)
	pq.Push(5)
	pq.Push(3)

	pq2 := pq.ToTopK(pq.Len())
	pq2.Push(7)
	for !pq2.IsEmpty() {
		fmt.Println(pq2.Pop())
	}
	// Output:
	// 3
	// 5
	// 7
}
