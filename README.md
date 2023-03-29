

# arrayQueue、circularBuffer、priorityQueue

## Features

- The data structure implemented using generics
- Higher performance
- arrayQueue is scalable, automatic expanding like slice

## QuickStart - arrayQueue

```bash
go get -u github.com/Doraemonkeys/queue/arrayQueue
```



```go
package main

import (
	"fmt"

	aq "github.com/Doraemonkeys/queue/arrayQueue"
)

func main() {
	que := aq.New[int]()
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
```



## QuickStart - circularBuffer

```go
package main

import (
	"fmt"

	cb "github.com/Doraemonkeys/queue/circularBuffer"
)

func main() {
	c := cb.New[int](3)
	c.PushBack(1)
	c.PushBack(2)
	c.PushBack(3)
	// 1 2 3
	c.PushBack(4)
	// 2 3 4
	c.PushFront(5)
	// 5 2 3
	it := c.Iterator()
	for it.Next() {
		fmt.Println(it.Value())
	}
	// Output:
	// 5
	// 2
	// 3
}
```



## QuickStart - priorityQueue

```go
package main

import (
	"fmt"

	pq "github.com/Doraemonkeys/queue/priorityQueue"
)

func main() {
	q := pq.New(func(a, b int) bool {
		return a < b
	})
	q.Push(1)
	q.Push(5)
	q.Push(3)
	for !q.IsEmpty() {
		fmt.Println(q.Pop())
	}
	// Output:
	// 1
	// 3
	// 5
}
```











