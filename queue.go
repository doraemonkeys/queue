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
//
// Returns true if 'a' is less than 'b'.
type LessFn[T any] func(a, b T) bool

// Comparator is a function that compares two elements.
//
// Returns -1 if x < y, 0 if x == y, 1 if x > y.
type Comparator[T any] func(x, y T) int

// Less returns true if a < b.
func (c Comparator[T]) Less(a, b T) bool {
	return c(a, b) < 0
}

// Greater returns true if a > b.
func (c Comparator[T]) Greater(a, b T) bool {
	return c(a, b) > 0
}

// Equal returns true if a == b.
func (c Comparator[T]) Equal(a, b T) bool {
	return c(a, b) == 0
}

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
