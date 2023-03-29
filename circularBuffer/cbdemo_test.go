package circularBuffer

import "fmt"

func Example() {
	cb := New[int](3)
	cb.PushBack(1)
	cb.PushBack(2)
	cb.PushBack(3)
	// 1 2 3
	cb.PushBack(4)
	// 2 3 4
	cb.PushFront(5)
	// 5 2 3
	it := cb.Iterator()
	for it.Next() {
		fmt.Println(it.Value())
	}
	// Output:
	// 5
	// 2
	// 3
}
