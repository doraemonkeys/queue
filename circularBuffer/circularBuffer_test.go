package circularBuffer

import (
	"testing"
)

func TestCircularBuffer(t *testing.T) {
	// Test NewCircularBuffer
	cb := New[int](3)
	if cb.Len() != 0 {
		t.Errorf("expected length 0, got %d", cb.Len())
	}
	if cb.Cap() != 3 {
		t.Errorf("expected capacity 3, got %d", cb.Cap())
	}

	// Test PushBack
	cb.PushBack(1)
	cb.PushBack(2)
	cb.PushBack(3)
	if cb.Len() != 3 {
		t.Errorf("expected length 3, got %d", cb.Len())
	}
	if cb.Back() != 3 {
		t.Errorf("expected back element 3, got %d", cb.Back())
	}

	// Test PushFront
	cb.PushFront(4)
	if cb.Len() != 3 {
		t.Errorf("expected length 3, got %d", cb.Len())
	}
	if cb.Front() != 4 {
		t.Errorf("expected front element 4, got %d", cb.Front())
	}

	// Test PopBack
	val := cb.PopBack()
	if val != 2 {
		t.Errorf("expected popped element 2, got %d", val)
	}
	if cb.Len() != 2 {
		t.Errorf("expected length 2, got %d", cb.Len())
	}
	if cb.Back() != 1 {
		t.Errorf("expected back element 1, got %d", cb.Back())
	}

	// Test PopFront
	val = cb.PopFront()
	if val != 4 {
		t.Errorf("expected popped element 4, got %d", val)
	}
	if cb.Len() != 1 {
		t.Errorf("expected length 1, got %d", cb.Len())
	}
	if cb.Front() != 1 {
		t.Errorf("expected front element 1, got %d", cb.Front())
	}
	if cb.Back() != 1 {
		t.Errorf("expected back element 1, got %d", cb.Back())
	}

	// Test iterator
	it := cb.Iterator()
	vals := []int{}
	for {
		if !it.Next() {
			break
		}
		val := it.Value()
		vals = append(vals, val)
	}
	if len(vals) != 1 {
		t.Errorf("expected 1 value from iterator, got %d", len(vals))
	}
	if vals[0] != 1 {
		t.Errorf("expected iterator value 1, got %d", vals[0])
	}
	it.Begin()
	vals = []int{}
	for {
		if !it.Next() {
			break
		}
		val := it.Value()
		vals = append(vals, val)
	}
	if len(vals) != 1 {
		t.Errorf("expected 1 value from iterator after reset, got %d", len(vals))
	}
	if vals[0] != 1 {
		t.Errorf("expected iterator value 1 after reset, got %d", vals[0])
	}

}

func TestCircularBuffer01(t *testing.T) {
	// Test NewCircularBuffer
	cb := New[int](3)
	if cb.Len() != 0 {
		t.Errorf("expected length 0, got %d", cb.Len())
	}
	if cb.Cap() != 3 {
		t.Errorf("expected capacity 3, got %d", cb.Cap())
	}

	// Test PushBack
	cb.PushBack(1)
	cb.PushBack(2)
	cb.PushBack(3)
	if cb.Len() != 3 {
		t.Errorf("expected length 3, got %d", cb.Len())
	}
	if cb.Back() != 3 {
		t.Errorf("expected back element 3, got %d", cb.Back())
	}

	// Test PushFront
	cb.PushFront(4)
	if cb.Len() != 3 {
		t.Errorf("expected length 3, got %d", cb.Len())
	}
	if cb.Front() != 4 {
		t.Errorf("expected front element 4, got %d", cb.Front())
	}

	// Test PopBack
	val := cb.PopBack()
	if val != 2 {
		t.Errorf("expected popped element 2, got %d", val)
	}
	if cb.Len() != 2 {
		t.Errorf("expected length 2, got %d", cb.Len())
	}
	if cb.Back() != 1 {
		t.Errorf("expected back element 1, got %d", cb.Back())
	}

	// Test PopFront
	val = cb.PopFront()
	if val != 4 {
		t.Errorf("expected popped element 4, got %d", val)
	}
	if cb.Len() != 1 {
		t.Errorf("expected length 1, got %d", cb.Len())
	}
	if cb.Front() != 1 {
		t.Errorf("expected front element 1, got %d", cb.Front())
	}
	if cb.Back() != 1 {
		t.Errorf("expected back element 1, got %d", cb.Back())
	}

	// Test iterator
	it := cb.Iterator()
	vals := []int{}
	for {
		if !it.Next() {
			break
		}
		val := it.Value()
		vals = append(vals, val)
	}
	if len(vals) != 1 {
		t.Errorf("expected 1 value from iterator, got %d", len(vals))
	}
	if vals[0] != 1 {
		t.Errorf("expected iterator value 1, got %d", vals[0])
	}
	it.Begin()
	vals = []int{}
	for {
		if !it.Next() {
			break
		}
		val := it.Value()
		vals = append(vals, val)
	}
	if len(vals) != 1 {
		t.Errorf("expected 1 value from iterator after reset, got %d", len(vals))
	}
	if vals[0] != 1 {
		t.Errorf("expected iterator value 1 after reset, got %d", vals[0])
	}

}

// TestCircularBuffer2
// TestCircularBuffer2
func TestCircularBuffer2(t *testing.T) {
	// Test PopBack on empty buffer
	cb := New[int](3)
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("expected panic on PopBack from empty buffer")
		}
	}()
	cb.PopBack()
}

// TestCircularBuffer3
func TestCircularBuffer3(t *testing.T) {
	// Test PopFront on empty buffer
	cb := New[int](3)
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("expected panic on PopFront from empty buffer")
		}
	}()
	cb.PopFront()
}

// TestCircularBuffer4
func TestCircularBuffer4(t *testing.T) {
	// Test PushBack on full buffer
	cb := New[int](3)
	cb.PushBack(1)
	cb.PushBack(2)
	cb.PushBack(3)
	cb.PushBack(4)
	if cb.Len() != 3 {
		t.Errorf("expected length 3, got %d", cb.Len())
	}
	if cb.Back() != 4 {
		t.Errorf("expected back element 4, got %d", cb.Back())
	}
}

// TestCircularBuffer5
func TestCircularBuffer5(t *testing.T) {
	// Test PushFront on full buffer
	cb := New[int](3)
	cb.PushFront(1)
	cb.PushFront(2)
	cb.PushFront(3)
	cb.PushFront(4)
	if cb.Len() != 3 {
		t.Errorf("expected length 3, got %d", cb.Len())
	}
	if cb.Front() != 4 {
		t.Errorf("expected front element 4, got %d", cb.Front())
	}

	// TestCircularBuffer3
	// Test Back on empty buffer
	cb = New[int](3)
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("expected panic on Back from empty buffer")
		}
	}()
	cb.Back()
	// Test Front on empty buffer
	cb = New[int](3)
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("expected panic on Front from empty buffer")
		}
	}()
	cb.Front()
	// Test Len on empty buffer
	cb = New[int](3)
	if cb.Len() != 0 {
		t.Errorf("expected length 0, got %d", cb.Len())
	}
	// Test Cap on empty buffer
	cb = New[int](3)
	if cb.Cap() != 3 {
		t.Errorf("expected capacity 3, got %d", cb.Cap())
	}
	// Test Iterator on empty buffer
	cb = New[int](3)
	it := cb.Iterator()
	if it.Next() {
		t.Errorf("expected no values from iterator on empty buffer")
	}
	it.Begin()
	if it.Next() {
		t.Errorf("expected no values from iterator on empty buffer after reset")
	}

	// TestCircularBuffer4
	// Test Back on empty buffer
	cb = New[int](3)
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("expected panic on Back from empty buffer")
		}
	}()
	cb.Back()
	// Test Front on empty buffer
	cb = New[int](3)
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("expected panic on Front from empty buffer")
		}
	}()
	cb.Front()
	// Test Len on empty buffer
	cb = New[int](3)
	if cb.Len() != 0 {
		t.Errorf("expected length 0, got %d", cb.Len())
	}
	// Test Cap on empty buffer
	cb = New[int](3)
	if cb.Cap() != 3 {
		t.Errorf("expected capacity 3, got %d", cb.Cap())
	}
	// Test Iterator on empty buffer
	cb = New[int](3)
	it = cb.Iterator()
	if it.Next() {
		t.Errorf("expected no values from iterator on empty buffer")
	}
	it.Begin()
	if it.Next() {
		t.Errorf("expected no values from iterator on empty buffer after reset")
	}
	// Test PushBack on buffer with one element
	cb = New[int](3)
	cb.PushBack(1)
	if cb.Len() != 1 {
		t.Errorf("expected length 1, got %d", cb.Len())
	}
	if cb.Back() != 1 {
		t.Errorf("expected back element 1, got %d", cb.Back())
	}
	// Test PushFront on buffer with one element
	cb = New[int](3)
	cb.PushFront(1)
	if cb.Len() != 1 {
		t.Errorf("expected length 1, got %d", cb.Len())
	}
	if cb.Front() != 1 {
		t.Errorf("expected front element 1, got %d", cb.Front())
	}
	// Test PopBack on buffer with one element
	cb = New[int](3)
	cb.PushBack(1)
	val := cb.PopBack()
	if val != 1 {
		t.Errorf("expected popped element 1, got %d", val)
	}
	if cb.Len() != 0 {
		t.Errorf("expected length 0, got %d", cb.Len())
	}

	// TestCircularBuffer5
	// Test Back on buffer with multiple elements
	cb = New[int](3)
	cb.PushBack(1)
	cb.PushBack(2)
	cb.PushBack(3)
	if cb.Back() != 3 {
		t.Errorf("expected back element 3, got %d", cb.Back())
	}

	// Test Front on buffer with multiple elements
	cb = New[int](3)
	cb.PushBack(1)
	cb.PushBack(2)
	cb.PushBack(3)
	if cb.Front() != 1 {
		t.Errorf("expected front element 1, got %d", cb.Front())
	}

	// Test Len on buffer with multiple elements
	cb = New[int](3)
	cb.PushBack(1)
	cb.PushBack(2)
	cb.PushBack(3)
	if cb.Len() != 3 {
		t.Errorf("expected length 3, got %d", cb.Len())
	}

	// Test Cap on buffer with multiple elements
	cb = New[int](3)
	cb.PushBack(1)
	cb.PushBack(2)
	cb.PushBack(3)
	if cb.Cap() != 3 {
		t.Errorf("expected capacity 3, got %d", cb.Cap())
	}

	// Test Iterator on buffer with multiple elements
	cb = New[int](3)
	cb.PushBack(1)
	cb.PushBack(2)
	cb.PushBack(3)
	it = cb.Iterator()
	vals := []int{}
	for {
		if !it.Next() {
			break
		}
		val := it.Value()
		vals = append(vals, val)
	}
	if len(vals) != 3 {
		t.Errorf("expected 3 values from iterator, got %d", len(vals))
	}
	if vals[0] != 1 || vals[1] != 2 || vals[2] != 3 {
		t.Errorf("expected iterator values [1 2 3], got %v", vals)
	}
	it.Begin()
	vals = []int{}
	for {
		if !it.Next() {
			break
		}
		val := it.Value()
		vals = append(vals, val)
	}
	if len(vals) != 3 {
		t.Errorf("expected 3 values from iterator after reset, got %d", len(vals))
	}
	if vals[0] != 1 || vals[1] != 2 || vals[2] != 3 {
		t.Errorf("expected iterator values [1 2 3] after reset, got %v", vals)
	}
}

// TestCircularBuffer6
func TestCircularBuffer6(t *testing.T) {
	// Test PushBack and PopFront alternately
	cb := New[int](3)
	cb.PushBack(1)
	cb.PushBack(2)
	if cb.PopFront() != 1 {
		t.Errorf("expected popped element 1, got %d", cb.PopFront())
	}
	cb.PushBack(3)
	if cb.PopFront() != 2 {
		t.Errorf("expected popped element 2, got %d", cb.PopFront())
	}
	if cb.PopFront() != 3 {
		t.Errorf("expected popped element 3, got %d", cb.PopFront())
	}
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("expected panic on PopFront from empty buffer")
		}
	}()
	cb.PopFront()
}

// TestCircularBuffer7
func TestCircularBuffer7(t *testing.T) {
	// Test PushFront and PopBack alternately
	cb := New[int](3)
	cb.PushFront(1)
	cb.PushFront(2)
	if cb.PopBack() != 1 {
		t.Errorf("expected popped element 1, got %d", cb.PopBack())
	}
	cb.PushFront(3)
	if cb.PopBack() != 2 {
		t.Errorf("expected popped element 2, got %d", cb.PopBack())
	}
	if cb.PopBack() != 3 {
		t.Errorf("expected popped element 3, got %d", cb.PopBack())
	}
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("expected panic on PopBack from empty buffer")
		}
	}()
	cb.PopBack()
}

func TestCircularBuffer8(t *testing.T) {
	cb := New[int](1)
	var f = func() {
		defer func() {
			if r := recover(); r == nil {
				t.Errorf("expected panic on PopBack from empty buffer")
			}
		}()
		cb.PopBack()
	}
	f()
	cb.PushBack(1)
	if cb.Back() != 1 || cb.Front() != 1 {
		t.Errorf("expected popped element 1, got %d", cb.PopBack())
	}
	cb.PushBack(2)
	if cb.Back() != 2 || cb.Front() != 2 {
		t.Errorf("expected popped element 2, got %d", cb.PopBack())
	}
	cb.PushFront(3)
	if cb.Back() != 3 || cb.Front() != 3 {
		t.Errorf("expected popped element 3, got %d", cb.PopBack())
	}
	cb.PushFront(4)
	if cb.Back() != 4 || cb.Front() != 4 {
		t.Errorf("expected popped element 4, got %d", cb.PopBack())
	}
}

func TestCircularBuffer9(t *testing.T) {
	cb := New[int](3)
	cb.PushBack(1)
	cb.PushBack(2)
	cb.PushBack(3)
	// 1 2 3
	it := cb.Iterator()
	for i := 0; i < 3; i++ {
		it.Next()
		if it.Value() != i+1 {
			t.Errorf("expected iterator value %d, got %d", i+1, it.Value())
		}
	}
	valCh := cb.GetValueFromChannel()
	want := []int{1, 2, 3}
	for i := 0; i < 3; i++ {
		if val := <-valCh; val != want[i] {
			t.Errorf("expected value %d, got %d", want[i], val)
		}
	}
	cb.PushFront(4)
	// 4 1 2
	it.Begin()
	if !it.Next() || it.Value() != 4 {
		t.Errorf("expected iterator value 4, got %d", it.Value())
	}
	if !it.Next() || it.Value() != 1 {
		t.Errorf("expected iterator value 1, got %d", it.Value())
	}
	if !it.Next() || it.Value() != 2 {
		t.Errorf("expected iterator value 2, got %d", it.Value())
	}
	want = []int{4, 1, 2}
	valCh = cb.GetValueFromChannel()
	for i := 0; i < 3; i++ {
		if val := <-valCh; val != want[i] {
			t.Errorf("expected value %d, got %d", want[i], val)
		}
	}

	cb.PushFront(5)
	// 5 4 1
	it.Begin()
	if !it.Next() || it.Value() != 5 {
		t.Errorf("expected iterator value 5, got %d", it.Value())
	}
	if !it.Next() || it.Value() != 4 {
		t.Errorf("expected iterator value 4, got %d", it.Value())
	}
	if !it.Next() || it.Value() != 1 {
		t.Errorf("expected iterator value 1, got %d", it.Value())
	}
	want = []int{5, 4, 1}
	valCh = cb.GetValueFromChannel()
	for i := 0; i < 3; i++ {
		if val := <-valCh; val != want[i] {
			t.Errorf("expected value %d, got %d", want[i], val)
		}
	}

	cb.PushBack(6)
	// 4 1 6
	it.Begin()
	if !it.Next() || it.Value() != 4 {
		t.Errorf("expected iterator value 4, got %d", it.Value())
	}
	if !it.Next() || it.Value() != 1 {
		t.Errorf("expected iterator value 1, got %d", it.Value())
	}
	if !it.Next() || it.Value() != 6 {
		t.Errorf("expected iterator value 6, got %d", it.Value())
	}
	want = []int{4, 1, 6}
	valCh = cb.GetValueFromChannel()
	for i := 0; i < 3; i++ {
		if val := <-valCh; val != want[i] {
			t.Errorf("expected value %d, got %d", want[i], val)
		}
	}
	cb.Resize(6)
	// 4 1 6
	it.Begin()
	if !it.Next() || it.Value() != 4 {
		t.Errorf("expected iterator value 4, got %d", it.Value())
	}
	if !it.Next() || it.Value() != 1 {
		t.Errorf("expected iterator value 1, got %d", it.Value())
	}
	if !it.Next() || it.Value() != 6 {
		t.Errorf("expected iterator value 6, got %d", it.Value())
	}
	want = []int{4, 1, 6}
	valCh = cb.GetValueFromChannel()
	for i := 0; i < 3; i++ {
		if val := <-valCh; val != want[i] {
			t.Errorf("expected value %d, got %d", want[i], val)
		}
	}

	cb.Resize(2)
	// 4 1
	it.Begin()
	if !it.Next() || it.Value() != 4 {
		t.Errorf("expected iterator value 4, got %d", it.Value())
	}
	if !it.Next() || it.Value() != 1 {
		t.Errorf("expected iterator value 1, got %d", it.Value())
	}
	want = []int{4, 1}
	valCh = cb.GetValueFromChannel()
	for i := 0; i < 2; i++ {
		if val := <-valCh; val != want[i] {
			t.Errorf("expected value %d, got %d", want[i], val)
		}
	}
	cb.Clear()
	// empty
	it.Begin()
	if it.Next() {
		t.Errorf("expected iterator to be empty")
	}
	cb.PushBack(1)
	cb.PushBack(2)
	cb.PushBack(3)
	// 2 3
	want = []int{2, 3}
	it.Begin()
	for i := 0; i < 2; i++ {
		if !it.Next() || it.Value() != want[i] {
			t.Errorf("expected iterator value %d, got %d", want[i], it.Value())
		}
	}
}
