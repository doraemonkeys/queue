

# arrayQueue、circularBuffer、priorityQueue

[Doc](https://pkg.go.dev/github.com/Doraemonkeys/queue)



## Features

- The data structure implemented using generics
- Higher performance
- arrayQueue is scalable, automatic expanding like slice

## QuickStart

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

















