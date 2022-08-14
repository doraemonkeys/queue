

# queue - circularQueue

 - The data structure implemented using generics.

## QuickStart

```go
import (
	"fmt"

	cq "github.com/Doraemonkeys/circularQueue"
)

func main() {
	que := cq.New[int]()
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

type CqIterator

    func (c *CqIterator[T]) Begin()

    func (c *CqIterator[T]) End()

    func (c *CqIterator[T]) Index() int

    func (c *CqIterator[T]) MoveTo(index int) bool

    func (c *CqIterator[T]) Next() bool

    func (c *CqIterator[T]) Prev() bool

    func (c *CqIterator[T]) Value() T



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


