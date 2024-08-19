package queue

import "time"

type Iterator[T any] interface {
	Begin()
	End()
	Next() bool
	Prev() bool
	Index() int
	Value() T
}

// LessFn is a function that returns whether 'a' is less than 'b'.
// Returns true if 'a' is less than 'b'.
type LessFn[T any] func(a, b T) bool

// Comparator is a function that compares two elements.
// Returns -1 if x < y, 0 if x == y, 1 if x > y.
type Comparator[T any] func(x, y T) int

// TimeComparator provides a basic comparison on time.Time
func TimeComparator(a, b time.Time) int {
	switch {
	case a.After(b):
		return 1
	case a.Before(b):
		return -1
	default:
		return 0
	}
}
