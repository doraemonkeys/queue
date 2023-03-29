package queue

type Iterator[T any] interface {
	Begin()
	End()
	Next() bool
	Prev() bool
	Index() int
	Value() T
}

// LessFn is a function that returns whether 'a' is less than 'b'.
type LessFn[T any] func(a, b T) bool
