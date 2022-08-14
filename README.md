

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

type AqIterator

func (a *AqIterator[T]) Begin()

func (a *AqIterator[T]) End()

func (a *AqIterator[T]) Index() int

func (a *AqIterator[T]) MoveTo(index int) bool

func (a *AqIterator[T]) Next() bool

func (a *AqIterator[T]) Prev() bool

func (a *AqIterator[T]) Value() T


type Queue

func New() *Queue[T]

func (Q *Queue[T]) Back() T

func (Q *Queue[T]) Clear()

func (Q *Queue[T]) Empty() bool

func (Q *Queue[T]) Front() T

func (Q *Queue[T]) GetValueFromChannel() chan T

func (Q *Queue[T]) Iterator() *AqIterator[T]

func (Q *Queue[T]) Len() int

func (Q *Queue[T]) Pop() (value T)

func (Q *Queue[T]) Push(value T)

func (Q *Queue[T]) Resize(newCap int)


