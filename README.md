

# queue - arrayqueue

## QuickStart

```go
import (
	"fmt"

	aq "github.com/Doraemonkeys/arrayqueue"
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

func New[T any]\() *queue[T]



func (Q *queue[T]) Back() T

func (Q *queue[T]) Empty() bool

func (Q *queue[T]) Front() T

func (Q *queue[T]) GetValueFromChannel() chan T

func (Q *queue[T]) Iterator() *aqIterator[T]

func (Q *queue[T]) Len() int

func (Q *queue[T]) Pop() (value T)

func (Q *queue[T]) Push(value T)

func (Q *queue[T]) Resize(newCap int)

func (Q *queue[T]) sendValue(ch chan<- T)



func (a *aqIterator[T]) Begin()

func (a *aqIterator[T]) End()

func (a *aqIterator[T]) Index() int

func (a *aqIterator[T]) Next() bool

func (a *aqIterator[T]) Prev() bool

func (a *aqIterator[T]) Value() T


