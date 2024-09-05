

# arrayQueue、circularBuffer、priorityQueue

[![Go Reference](https://pkg.go.dev/badge/github.com/doraemonkeys/queue.svg)](https://pkg.go.dev/github.com/doraemonkeys/queue) [![Go Report Card](https://goreportcard.com/badge/github.com/doraemonkeys/queue)](https://goreportcard.com/report/github.com/doraemonkeys/queue)

## Features

- The data structure implemented using generics
- Higher performance
- arrayQueue is scalable, automatic expanding like slice



## Quick Start

### priorityQueue

```bash
go get -u github.com/doraemonkeys/queue/priorityQueue
```


```go
package main

import (
	"fmt"

	pq "github.com/doraemonkeys/queue/priorityQueue"
)

func main() {
	q := pq.New(func(a, b int) bool {
		return a < b
	})
	q.Push(1)
	q.Push(5)
	q.Push(3)

	q2 := q.ToTopK(q.Len())
	q2.Push(7)
	for !q2.IsEmpty() {
		fmt.Println(q2.Pop())
	}
	// Output:
	// 3
	// 5
	// 7
}
```


### arrayQueue

```bash
go get -u github.com/doraemonkeys/queue/arrayQueue
```

```go
package main

import (
	"fmt"

	aq "github.com/doraemonkeys/queue/arrayQueue"
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



### circularBuffer

```bash
go get -u github.com/doraemonkeys/queue/circularBuffer
```

```go
package main

import (
	"fmt"

	cb "github.com/doraemonkeys/queue/circularBuffer"
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









