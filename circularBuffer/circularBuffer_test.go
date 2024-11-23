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
	cb.Clear()
	// empty
	it.Begin()
	if it.Next() {
		t.Errorf("expected iterator to be empty")
	}
	cb.PushBack(1)
	cb.PushBack(2)
	cb.PushBack(3)
	if cb.PopFront() != 1 {
		t.Errorf("expected popped element 1, got %d", cb.PopFront())
	}
	// 2 3
	want := []int{2, 3}
	it.Begin()
	for i := 0; i < 2; i++ {
		if !it.Next() || it.Value() != want[i] {
			t.Errorf("expected iterator value %d, got %d", want[i], it.Value())
		}
	}
}

func TestResize(t *testing.T) {
	t.Run("Resize to larger capacity", func(t *testing.T) {
		buf := New[int](5)
		for i := 1; i <= 5; i++ {
			buf.PushBack(i)
		}
		buf.Resize(10)
		if buf.Cap() != 10 {
			t.Errorf("Expected capacity 10, got %d", buf.Cap())
		}
		if buf.Len() != 5 {
			t.Errorf("Expected length 5, got %d", buf.Len())
		}
		for i := 1; i <= 5; i++ {
			if buf.PopFront() != i {
				t.Errorf("Expected %d, got %d", i, buf.PopFront())
			}
		}
	})

	t.Run("Resize to smaller capacity", func(t *testing.T) {
		buf := New[int](10)
		for i := 1; i <= 10; i++ {
			buf.PushBack(i)
		}
		buf.Resize(5)
		if buf.Cap() != 5 {
			t.Errorf("Expected capacity 5, got %d", buf.Cap())
		}
		if buf.Len() != 5 {
			t.Errorf("Expected length 5, got %d", buf.Len())
		}
		for i := 1; i <= 5; i++ {
			if buf.PopFront() != i {
				t.Errorf("Expected %d, got %d", i, buf.PopFront())
			}
		}
	})

	t.Run("Resize to same capacity", func(t *testing.T) {
		buf := New[int](5)
		for i := 1; i <= 5; i++ {
			buf.PushBack(i)
		}
		buf.Resize(5)
		if buf.Cap() != 5 {
			t.Errorf("Expected capacity 5, got %d", buf.Cap())
		}
		if buf.Len() != 5 {
			t.Errorf("Expected length 5, got %d", buf.Len())
		}
		for i := 1; i <= 5; i++ {
			if buf.PopFront() != i {
				t.Errorf("Expected %d, got %d", i, buf.PopFront())
			}
		}
	})

	t.Run("Resize empty buffer", func(t *testing.T) {
		buf := New[int](5)
		buf.Resize(10)
		if buf.Cap() != 10 {
			t.Errorf("Expected capacity 10, got %d", buf.Cap())
		}
		if buf.Len() != 0 {
			t.Errorf("Expected length 0, got %d", buf.Len())
		}
	})

	t.Run("Resize with wrapped elements", func(t *testing.T) {
		buf := New[int](5)
		for i := 1; i <= 5; i++ {
			buf.PushBack(i)
		}
		buf.PopFront() // Remove 1
		buf.PopFront() // Remove 2
		buf.PushBack(6)
		buf.PushBack(7)
		buf.Resize(7)
		if buf.Cap() != 7 {
			t.Errorf("Expected capacity 7, got %d", buf.Cap())
		}
		if buf.Len() != 5 {
			t.Errorf("Expected length 5, got %d", buf.Len())
		}
		expected := []int{3, 4, 5, 6, 7}
		for _, v := range expected {
			if buf.PopFront() != v {
				t.Errorf("Expected %d, got %d", v, buf.PopFront())
			}
		}
	})

	t.Run("Resize to 1", func(t *testing.T) {
		buf := New[int](5)
		for i := 1; i <= 5; i++ {
			buf.PushBack(i)
		}
		buf.Resize(1)
		if buf.Cap() != 1 {
			t.Errorf("Expected capacity 1, got %d", buf.Cap())
		}
		if buf.Len() != 1 {
			t.Errorf("Expected length 1, got %d", buf.Len())
		}
		if buf.PopFront() != 1 {
			t.Errorf("Expected 5, got %d", buf.PopFront())
		}
	})

	t.Run("Panic on resize to 0", func(t *testing.T) {
		buf := New[int](5)
		defer func() {
			if r := recover(); r == nil {
				t.Errorf("The code did not panic")
			}
		}()
		buf.Resize(0)
	})

	t.Run("Panic on resize to negative", func(t *testing.T) {
		buf := New[int](5)
		defer func() {
			if r := recover(); r == nil {
				t.Errorf("The code did not panic")
			}
		}()
		buf.Resize(-1)
	})

	t.Run("Resize with wrapped elements and partial copy", func(t *testing.T) {
		buf := New[int](7)
		// 填充缓冲区并使其环绕
		for i := 1; i <= 7; i++ {
			buf.PushBack(i)
		}
		// 移除前两个元素并添加两个新元素,使缓冲区环绕
		buf.PopFront() // 移除 1
		buf.PopFront() // 移除 2
		buf.PushBack(8)
		buf.PushBack(9)

		// 此时缓冲区状态: [8 9 3 4 5 6 7]
		//                     ^first

		// 调整大小到触发部分复制逻辑
		buf.Resize(6)

		if buf.Cap() != 6 {
			t.Errorf("Expected capacity 5, got %d", buf.Cap())
		}
		if buf.Len() != 6 {
			t.Errorf("Expected length 5, got %d", buf.Len())
		}

		// 验证元素
		expected := []int{3, 4, 5, 6, 7, 8}
		for i, v := range expected {
			if got := buf.PopFront(); got != v {
				t.Errorf("At index %d: expected %d, got %d", i, v, got)
			}
		}

		if !buf.Empty() {
			t.Errorf("Buffer should be empty after popping all elements")
		}
	})
}

func TestIterator(t *testing.T) {
	// 创建一个容量为5的Buffer
	cb := New[int](5)

	// 测试空Buffer的迭代器
	t.Run("Empty Buffer", func(t *testing.T) {
		it := cb.Iterator()
		if it.Index() != 0 {
			t.Errorf("Expected index 0 for empty buffer, got %d", it.Index())
		}
		if it.Next() {
			t.Error("Next() should return false for empty buffer")
		}
		if it.Prev() {
			t.Error("Prev() should return false for empty buffer")
		}
	})

	// 添加元素并测试
	cb.PushBack(1)
	cb.PushBack(2)
	cb.PushBack(3)

	t.Run("Filled Buffer", func(t *testing.T) {
		it := cb.Iterator()

		// 测试Begin()和Index()
		it.Begin()
		if it.Index() != 0 {
			t.Errorf("Expected index 0 after Begin(), got %d", it.Index())
		}

		// 测试Next()和Value()
		expected := []int{1, 2, 3}
		for i, exp := range expected {
			if !it.Next() {
				t.Errorf("Next() returned false at index %d", i)
			}
			if it.Value() != exp {
				t.Errorf("Expected value %d at index %d, got %d", exp, i, it.Value())
			}
			if it.Index() != i+1 {
				t.Errorf("Expected index %d, got %d", i+1, it.Index())
			}
		}

		// 测试End()
		it.End()
		if it.Next() {
			t.Error("Next() should return false after End()")
		}

		// 测试Prev()
		for i := len(expected) - 1; i >= 0; i-- {
			if !it.Prev() {
				t.Errorf("Prev() returned false at index %d", i)
			}
			if it.Value() != expected[i] {
				t.Errorf("Expected value %d at index %d, got %d", expected[i], i, it.Value())
			}
		}

		if it.Prev() {
			t.Error("Prev() should return false at the beginning")
		}
	})

	t.Run("MoveTo", func(t *testing.T) {
		it := cb.Iterator()

		// 测试有效移动
		if !it.MoveTo(2) {
			t.Error("MoveTo(2) should return true")
		}
		if it.Value() != 2 {
			t.Errorf("Expected value 2 after MoveTo(2), got %d", it.Value())
		}

		// 测试无效移动
		if it.MoveTo(0) {
			t.Error("MoveTo(0) should return false")
		}
		if it.MoveTo(4) {
			t.Error("MoveTo(4) should return false")
		}
	})

	t.Run("MoveTo2", func(t *testing.T) {
		it := cb.Iterator()

		// 测试有效移动
		if !it.MoveTo2(1) {
			t.Error("MoveTo2(2) should return true")
		}
		if it.Value() != 2 {
			t.Errorf("Expected value 2 after MoveTo(2), got %d", it.Value())
		}
	})

	t.Run("Wraparound", func(t *testing.T) {
		// 填满Buffer并造成环绕
		cb.PushBack(4)
		cb.PushBack(5)
		cb.PushBack(6) // 这会导致1被覆盖

		it := cb.Iterator()
		expected := []int{2, 3, 4, 5, 6}

		for i, exp := range expected {
			if !it.Next() {
				t.Errorf("Next() returned false at index %d", i)
			}
			if it.Value() != exp {
				t.Errorf("Expected value %d at index %d, got %d", exp, i, it.Value())
			}
		}

		if it.Next() {
			t.Error("Next() should return false after last element")
		}
	})
}

// 辅助函数:创建一个包含给定元素的Buffer
func createBuffer(elements ...int) *Buffer[int] {
	buffer := New[int](len(elements))
	for _, e := range elements {
		buffer.PushBack(e)
	}
	return buffer
}

// 测试Iterator()方法
func TestIterator2(t *testing.T) {
	buffer := createBuffer(1, 2, 3)
	it := buffer.Iterator()

	if it.index != -1 {
		t.Errorf("Iterator should start at index -1, got %d", it.index)
	}
}

// 测试Begin()方法
func TestBegin(t *testing.T) {
	buffer := createBuffer(1, 2, 3)
	it := buffer.Iterator()
	it.Next()
	it.Begin()

	if it.index != -1 {
		t.Errorf("Begin() should set index to -1, got %d", it.index)
	}
}

// 测试Index()方法
func TestIndex(t *testing.T) {
	buffer := createBuffer(1, 2, 3)
	it := buffer.Iterator()

	if idx := it.Index(); idx != 0 {
		t.Errorf("Index() should return 0 for iterator at beginning, got %d", idx)
	}

	it.Next()
	if idx := it.Index(); idx != 1 {
		t.Errorf("Index() should return 1 for first element, got %d", idx)
	}

	it.End()
	if idx := it.Index(); idx != 0 {
		t.Errorf("Index() should return 0 for iterator at end, got %d", idx)
	}

	expectedPrevEle := []int{3, 2, 1}
	for i, exp := range expectedPrevEle {
		if !it.Prev() {
			t.Errorf("Prev() should return true at index %d", i)
		}
		if it.Value() != exp {
			t.Errorf("Expected value %d at index %d, got %d", exp, i, it.Value())
		}
	}
}

// 测试End()方法
func TestEnd(t *testing.T) {
	buffer := createBuffer(1, 2, 3)
	it := buffer.Iterator()
	it.End()

	if it.index != -2 {
		t.Errorf("End() should set index to -2, got %d", it.index)
	}
}

// 测试Value()方法
func TestValue(t *testing.T) {
	buffer := createBuffer(1, 2, 3)
	it := buffer.Iterator()
	it.Next()

	if value := it.Value(); value != 1 {
		t.Errorf("Value() should return 1, got %d", value)
	}
}

// 测试Next()方法
func TestNext(t *testing.T) {
	buffer := createBuffer(1, 2, 3)
	it := buffer.Iterator()

	if !it.Next() {
		t.Error("Next() should return true for first element")
	}
	if it.Value() != 1 {
		t.Errorf("Next() should move to first element (1), got %d", it.Value())
	}

	it.Next()
	it.Next()
	if it.Next() {
		t.Error("Next() should return false after last element")
	}
}

// 测试Prev()方法
func TestPrev(t *testing.T) {
	buffer := createBuffer(1, 2, 3)
	it := buffer.Iterator()
	it.End()

	if !it.Prev() {
		t.Error("Prev() should return true for last element")
	}
	if it.Value() != 3 {
		t.Errorf("Prev() should move to last element (3), got %d", it.Value())
	}

	it.Prev()
	it.Prev()
	if it.Prev() {
		t.Error("Prev() should return false before first element")
	}
}

// 测试MoveTo()方法
func TestMoveTo(t *testing.T) {
	buffer := createBuffer(1, 2, 3, 4, 5)
	it := buffer.Iterator()

	if !it.MoveTo(3) {
		t.Error("MoveTo(3) should return true")
	}
	if it.Value() != 3 {
		t.Errorf("MoveTo(3) should move to element 3, got %d", it.Value())
	}

	if it.MoveTo(0) {
		t.Error("MoveTo(0) should return false")
	}

	if it.MoveTo(6) {
		t.Error("MoveTo(6) should return false")
	}
}

// 测试空Buffer的情况
func TestEmptyBuffer(t *testing.T) {
	buffer := New[int](5)
	it := buffer.Iterator()

	if it.Next() {
		t.Error("Next() should return false for empty buffer")
	}

	if it.Prev() {
		t.Error("Prev() should return false for empty buffer")
	}

	if it.Index() != 0 {
		t.Errorf("Index() should return 0 for empty buffer, got %d", it.Index())
	}
}

// 测试在迭代过程中修改Buffer
func TestModifyingBufferDuringIteration(t *testing.T) {
	buffer := createBuffer(1, 2, 3)
	it := buffer.Iterator()
	index := it.Index()
	it.Next()

	buffer.PopBack()
	if it.MoveTo(index) {
		t.Error("MoveTo() should return false after modifying buffer")
	}
}

// 测试Prev()方法的边界情况
func TestPrevEdgeCases(t *testing.T) {
	buffer := createBuffer(1, 2, 3)
	it := buffer.Iterator()

	// 模拟MoveTo失败的情况
	it.index = buffer.cap + 1 // 设置一个超出范围的index

	if it.Prev() {
		t.Error("Prev() should return false when index is out of range")
	}

	if it.index != -1 {
		t.Errorf("Prev() should set index to -1 when index was out of range, got %d", it.index)
	}
	buffer2 := New[int](1)
	it2 := buffer2.Iterator()
	it2.End()
	if it2.Prev() {
		t.Error("Prev() should return false for empty buffer")
	}
	if it2.index != -2 {
		t.Errorf("Prev() should set index to -2 for empty buffer, got %d", it2.index)
	}
}

// 测试MoveTo()方法的边界情况
func TestMoveToEdgeCases(t *testing.T) {
	// 创建一个Buffer,使其内部数组环绕
	buffer := New[int](5)
	for i := 1; i <= 7; i++ {
		buffer.PushBack(i)
	}
	// 此时buffer内容应为: [6 7 3 4 5], first = 2, last = 1

	it := buffer.Iterator()

	// 测试移动到第4个元素 (应该是值5)
	if !it.MoveTo(4) {
		t.Error("MoveTo(4) should return true")
	}
	if it.Value() != 6 {
		t.Errorf("MoveTo(4) should move to element 6, got %d", it.Value())
	}

	// 测试移动到第5个元素 (应该是值6)
	if !it.MoveTo(5) {
		t.Error("MoveTo(5) should return true")
	}
	if it.Value() != 7 {
		t.Errorf("MoveTo(5) should move to element 7, got %d", it.Value())
	}

	// 验证内部索引计算
	expectedIndex := 2
	if it.index != expectedIndex {
		t.Errorf("Internal index should be %d, got %d", expectedIndex, it.index)
	}
}

func TestIteratorNextEdgeCases(t *testing.T) {
	// 创建一个容量为 5 的 Buffer
	cb := New[int](5)

	// 添加一些元素
	cb.PushBack(1)
	cb.PushBack(2)
	cb.PushBack(3)

	it := cb.Iterator()

	// 测试 Next() 当 Index() > que.len 的情况
	it.index = cb.cap // 设置一个无效的索引
	if it.Next() {
		t.Errorf("Next() should return false when index > que.len")
	}
	if it.index != -2 {
		t.Errorf("index should be set to -2 when Next() fails due to invalid index")
	}

}
func TestCbIteratorPrev(t *testing.T) {
	// 创建一个容量为5的缓冲区
	cb := New[int](5)

	// 测试空缓冲区
	it := cb.Iterator()
	if it.Prev() {
		t.Error("Expected Prev() to return false on empty buffer")
	}

	// 填充缓冲区
	for i := 1; i <= 5; i++ {
		cb.PushBack(i)
	}

	// 测试从末尾开始的Prev
	it = cb.Iterator()
	it.End()
	if !it.Prev() {
		t.Error("Expected Prev() to return true at the end of non-empty buffer")
	}
	if it.Value() != 5 {
		t.Errorf("Expected last value to be 5, got %d", it.Value())
	}

	// 测试连续调用Prev
	for i := 4; i >= 1; i-- {
		if !it.Prev() {
			t.Errorf("Expected Prev() to return true for element %d", i)
		}
		if it.Value() != i {
			t.Errorf("Expected value %d, got %d", i, it.Value())
		}
	}

	// 测试到达开头
	if it.Prev() {
		t.Error("Expected Prev() to return false at the beginning of buffer")
	}

	// 测试在开头调用Prev
	it.Begin()
	if it.Prev() {
		t.Error("Expected Prev() to return false when called at the beginning")
	}

	// 测试循环缓冲区的情况
	cb.PopFront()  // 移除一个元素
	cb.PushBack(6) // 添加一个新元素，使缓冲区形成循环

	it = cb.Iterator()
	it.End()
	for i := 6; i >= 2; i-- {
		if !it.Prev() {
			t.Errorf("Expected Prev() to return true for element %d in circular buffer", i)
		}
		if it.Value() != i {
			t.Errorf("Expected value %d, got %d in circular buffer", i, it.Value())
		}
	}

	// 测试MoveTo后的Prev
	it.MoveTo(3)
	if !it.Prev() {
		t.Error("Expected Prev() to return true after MoveTo(3)")
	}
	if it.Value() != 3 {
		t.Errorf("Expected value 3 after MoveTo(3) and Prev(), got %d", it.Value())
	}

}

func TestGet_EmptyBuffer(t *testing.T) {
	buf := New[int](5)

	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Get on empty buffer should panic")
		}
	}()

	buf.Get(0)
}

func TestGet_SingleElement(t *testing.T) {
	buf := New[int](5)
	buf.PushBack(42)

	if got := buf.Get(0); got != 42 {
		t.Errorf("Get(0) = %v, want 42", got)
	}
}

func TestGet_MultipleElements(t *testing.T) {
	buf := New[int](5)
	buf.PushBack(1)
	buf.PushBack(2)
	buf.PushBack(3)

	testCases := []struct {
		index int
		want  int
	}{
		{0, 1},
		{1, 2},
		{2, 3},
	}

	for _, tc := range testCases {
		if got := buf.Get(tc.index); got != tc.want {
			t.Errorf("Get(%d) = %v, want %v", tc.index, got, tc.want)
		}
	}
}

func TestGet_WrappedBuffer(t *testing.T) {
	buf := New[int](3)
	buf.PushBack(1)
	buf.PushBack(2)
	buf.PushBack(3)
	buf.PushBack(4) // This should wrap around

	testCases := []struct {
		index int
		want  int
	}{
		{0, 2},
		{1, 3},
		{2, 4},
	}

	for _, tc := range testCases {
		if got := buf.Get(tc.index); got != tc.want {
			t.Errorf("Get(%d) = %v, want %v", tc.index, got, tc.want)
		}
	}
}

func TestGet_IndexOutOfRange(t *testing.T) {
	buf := New[int](5)
	buf.PushBack(1)
	buf.PushBack(2)

	testCases := []int{-1, 2, 5}

	for _, index := range testCases {
		func() {
			defer func() {
				if r := recover(); r == nil {
					t.Errorf("Get(%d) should panic", index)
				}
			}()
			buf.Get(index)
		}()
	}
}

func TestGet_AfterClear(t *testing.T) {
	buf := New[int](5)
	buf.PushBack(1)
	buf.PushBack(2)
	buf.Clear()

	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Get on cleared buffer should panic")
		}
	}()

	buf.Get(0)
}

func TestGet_AfterPopFront(t *testing.T) {
	buf := New[int](5)
	buf.PushBack(1)
	buf.PushBack(2)
	buf.PushBack(3)
	buf.PopFront()

	if got := buf.Get(0); got != 2 {
		t.Errorf("Get(0) after PopFront = %v, want 2", got)
	}
	if got := buf.Get(1); got != 3 {
		t.Errorf("Get(1) after PopFront = %v, want 3", got)
	}
}

func TestGet_AfterPopBack(t *testing.T) {
	buf := New[int](5)
	buf.PushBack(1)
	buf.PushBack(2)
	buf.PushBack(3)
	buf.PopBack()

	if got := buf.Get(0); got != 1 {
		t.Errorf("Get(0) after PopBack = %v, want 1", got)
	}
	if got := buf.Get(1); got != 2 {
		t.Errorf("Get(1) after PopBack = %v, want 2", got)
	}
}

func TestGet_LargeBuffer(t *testing.T) {
	buf := New[int](1000)
	for i := 0; i < 1000; i++ {
		buf.PushBack(i)
	}

	for i := 0; i < 1000; i++ {
		if got := buf.Get(i); got != i {
			t.Errorf("Get(%d) = %v, want %d", i, got, i)
		}
	}
}

func TestBinarySearch(t *testing.T) {
	// Helper function to create and populate a buffer
	createBuffer := func(elements ...int) *Buffer[int] {
		buf := New[int](len(elements))
		for _, e := range elements {
			buf.PushBack(e)
		}
		return buf
	}

	t.Run("Found in middle", func(t *testing.T) {
		buf := createBuffer(1, 3, 5, 7, 9)
		index, found := buf.BinarySearch(buf.Len(), func(v int) int {
			return 5 - v
		})
		if !found || index != 2 {
			t.Errorf("Expected (2, true), got (%d, %v)", index, found)
		}
	})

	t.Run("Found at beginning", func(t *testing.T) {
		buf := createBuffer(1, 3, 5, 7, 9)
		index, found := buf.BinarySearch(buf.Len(), func(v int) int {
			return 1 - v
		})
		if !found || index != 0 {
			t.Errorf("Expected (0, true), got (%d, %v)", index, found)
		}
	})

	t.Run("Found at end", func(t *testing.T) {
		buf := createBuffer(1, 3, 5, 7, 9)
		index, found := buf.BinarySearch(buf.Len(), func(v int) int {
			return 9 - v
		})
		if !found || index != 4 {
			t.Errorf("Expected (4, true), got (%d, %v)", index, found)
		}
	})

	t.Run("Not found - between elements", func(t *testing.T) {
		buf := createBuffer(1, 3, 5, 7, 9)
		index, found := buf.BinarySearch(buf.Len(), func(v int) int {
			return 4 - v
		})
		if found || index != 2 {
			t.Errorf("Expected (2, false), got (%d, %v)", index, found)
		}
	})

	t.Run("Not found - less than all", func(t *testing.T) {
		buf := createBuffer(1, 3, 5, 7, 9)
		index, found := buf.BinarySearch(buf.Len(), func(v int) int {
			return 0 - v
		})
		if found || index != 0 {
			t.Errorf("Expected (0, false), got (%d, %v)", index, found)
		}
	})

	t.Run("Not found - greater than all", func(t *testing.T) {
		buf := createBuffer(1, 3, 5, 7, 9)
		index, found := buf.BinarySearch(buf.Len(), func(v int) int {
			return 10 - v
		})
		if found || index != 5 {
			t.Errorf("Expected (5, false), got (%d, %v)", index, found)
		}
	})

	t.Run("Single element - found", func(t *testing.T) {
		buf := createBuffer(5)
		index, found := buf.BinarySearch(buf.Len(), func(v int) int {
			return 5 - v
		})
		if !found || index != 0 {
			t.Errorf("Expected (0, true), got (%d, %v)", index, found)
		}
	})

	t.Run("Single element - not found", func(t *testing.T) {
		buf := createBuffer(5)
		index, found := buf.BinarySearch(buf.Len(), func(v int) int {
			return 3 - v
		})
		if found || index != 0 {
			t.Errorf("Expected (0, false), got (%d, %v)", index, found)
		}
	})

	t.Run("Duplicate elements", func(t *testing.T) {
		buf := createBuffer(1, 3, 3, 3, 5)
		index, found := buf.BinarySearch(buf.Len(), func(v int) int {
			return 3 - v
		})
		if !found || index < 1 || index > 3 {
			t.Errorf("Expected (1-3, true), got (%d, %v)", index, found)
		}
	})
}
