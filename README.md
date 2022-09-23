

# queue - arrayQueue

 - The data structure implemented using generics.

## QuickStart

```go
import (
	"fmt"

	aq "github.com/Doraemonkeys/arrayQueue"
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





## overview

type AqIterator

    func (c *AqIterator[T]) Begin()

    func (c *AqIterator[T]) End()

    func (c *AqIterator[T]) Index() int

    func (c *AqIterator[T]) MoveTo(index int) bool

    func (c *AqIterator[T]) Next() bool

    func (c *AqIterator[T]) Prev() bool

    func (c *AqIterator[T]) Value() T



type Queue

    func New[T any]() *Queue[T]

    func (Q *Queue[T]) Back() T

    func (Q *Queue[T]) Clear()

    func (Q *Queue[T]) Empty() bool

    func (Q *Queue[T]) Front() T

    func (Q *Queue[T]) GetValueFromChannel() chan T

    func (Q *Queue[T]) Iterator() *CqIterator[T]

    func (Q *Queue[T]) Len() int

    func (Q *Queue[T]) Pop() (value T)

    func (Q *Queue[T]) Push(value T)
	
    func (Q *Queue[T]) Resize(newCap int)


