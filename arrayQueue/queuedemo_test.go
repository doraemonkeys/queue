package arrayQueue

import (
	"fmt"
)

func Example() {
	que := New[int]()
	que.Push(1)
	que.Push(2)
	que.Push(3)
	que.Pop()
	que.Push(99)
	fmt.Println(que.Front())
	fmt.Println(que.Back())
	it := que.Iterator()
	for it.Next() {
		fmt.Println(it.Value())
	}
	// Output:
	// 2
	// 99
	// 2
	// 3
	// 99
}
