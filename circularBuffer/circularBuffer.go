package circularBuffer

//推荐通过New()方法获取，若不是，请使用Resize()初始化底层切片

// Recommand to get a new circularBuffer by New() method,
// if not, please use Resize() to initialize the underlying slice.
type Buffer[T any] struct {
	data  []T
	last  int //最新插入的元素,first和last在队列不为空时指向已经有值的位置
	first int //最先插入的元素,first和last在队列不为空时指向已经有值的位置
	len   int //队列元素长度
	//表示已申请的内存能容纳元素的最大个数
	//容量 cap = 切片真实长度 - 1 (第0个位置不放数据)
	cap int
}

// circularBuffer's iterator
type CbIterator[T any] struct {
	que   *Buffer[T]
	index int //此处的index是队列元素在底层切片的实际索引
}

func New[T any](cap int) *Buffer[T] {
	if cap <= 0 {
		panic("capacity must be positive")
	}
	var cq Buffer[T]
	initCap := cap                 //初始容量
	cq.data = make([]T, initCap+1) //第0个位置不放数据
	//第0个位置不放数据
	//当队列 只有一个值 或 没有值 的时候，first与last是重合的，需要特别注意
	cq.last = 0
	cq.first = 0
	cq.cap = initCap
	cq.len = 0
	return &cq
}

func (Q *Buffer[T]) PushBack(value T) {
	if Q.len == 0 {
		Q.data[1] = value
		//last和first都指向已经有值的位置(没有元素除外)
		Q.last = 1
		Q.first = 1
		Q.len = 1
		return
	}
	next := (Q.last + 1) % (Q.cap + 1)
	if next == 0 {
		next = 1
	}
	//容量已满
	if next == Q.first {
		Q.first = (next + 1) % (Q.cap + 1)
		if Q.first == 0 {
			Q.first = 1
		}
	} else {
		Q.len++
	}
	Q.data[next] = value
	Q.last = next

}

func (Q *Buffer[T]) PushFront(value T) {
	if Q.len == 0 {
		Q.data[1] = value
		//last和first都指向已经有值的位置(没有元素除外)
		Q.last = 1
		Q.first = 1
		Q.len = 1
		return
	}
	next := Q.first - 1
	if next == 0 {
		next = Q.cap
	}
	//容量已满
	if next == Q.last {
		Q.last = Q.last - 1
		if Q.last == 0 {
			Q.last = Q.cap
		}
	} else {
		Q.len++
	}
	Q.data[next] = value
	Q.first = next
}

//直接遍历底层切片发送到channel,可能比迭代器遍历更快

// GetValueFromChannel returns a channel that can be used to iterate over the values of the queue.
// The channel is closed when all values have been sent.
func (Q *Buffer[T]) GetValueFromChannel() chan T {
	//fmt.Println("len:", Q.len, "cap:", Q.cap, "last:", Q.last, "first:", Q.first, "front:", Q.data[Q.first])
	if Q.len == 0 {
		//fmt.Println("空队列")
		ch := make(chan T, Q.len)
		close(ch)
		return ch
	}
	ch := make(chan T, Q.len)
	go Q.sendValue(ch)
	return ch
}

func (Q *Buffer[T]) sendValue(ch chan<- T) {
	temp := Q.first
	for i := 0; i < Q.len; i++ {
		ch <- Q.data[temp]
		temp = (temp + 1) % (Q.cap + 1)
		if temp == 0 {
			temp = 1
		}
	}
	close(ch)
}

// PopFront removes and returns the first element of the circularBuffer.
// It panics if the circularBuffer is empty.
// PopFront does not release memory, there is no significant performance loss,
// and memory can be released by calling Resize().
func (Q *Buffer[T]) PopFront() (value T) {
	if Q.len == 0 {
		panic("queue is empty")
	}
	value = Q.data[Q.first]
	//当队列 只有一个值 或 没有值 的时候，first与last是重合的，需要特别注意
	if Q.len == 1 {
		Q.len = 0
		Q.first = 0
		Q.last = 0
		return value
	}
	Q.len--
	Q.first = (Q.first + 1) % (Q.cap + 1)
	if Q.first == 0 {
		Q.first = 1
	}
	return value
}

// PopBack removes and returns the last element of the circularBuffer.
// It panics if the circularBuffer is empty.
// PopBack does not release memory, there is no significant performance loss,
// and memory can be released by calling Resize().
func (Q *Buffer[T]) PopBack() (value T) {
	if Q.len == 0 {
		panic("queue is empty")
	}
	value = Q.data[Q.last]
	//当队列 只有一个值 或 没有值 的时候，first与last是重合的，需要特别注意
	if Q.len == 1 {
		Q.len = 0
		Q.first = 0
		Q.last = 0
		return value
	}
	Q.len--
	Q.last = (Q.last - 1) % (Q.cap + 1)
	if Q.last == 0 {
		Q.last = Q.cap
	}
	return value
}

//Clear不会释放内存，没有太大性能消耗。

// Clear removes all elements from the queue.
// Clear does not release memory, there is no significant performance loss,
// and memory can be released by calling Resize().
func (Q *Buffer[T]) Clear() {
	if Q.len == 0 {
		return
	}
	Q.len = 0
	Q.first = 0
	Q.last = 0
}

//返回队列第一个元素(最先插入),空队列调用会panic

// Front returns the first element of the circularBuffer.
// It panics if the circularBuffer is empty.
func (Q *Buffer[T]) Front() T {
	if Q.len == 0 {
		panic("queue is empty")
	}
	return Q.data[Q.first]
}

//返回队列最后一个元素(最后插入),空队列调用会panic

// Back returns the last element of the circularBuffer.
// It panics if the circularBuffer is empty.
func (Q *Buffer[T]) Back() T {
	if Q.len == 0 {
		panic("queue is empty")
	}
	return Q.data[Q.last]
}

// Empty returns true if the circularBuffer is empty.
func (Q *Buffer[T]) Empty() bool {
	return Q.len == 0
}

//返回队列的长度

// Len returns the number of elements of the circularBuffer.
func (Q *Buffer[T]) Len() int {
	return Q.len
}

//返回队列的容量

// Cap returns the capacity of the circularBuffer.
func (Q *Buffer[T]) Cap() int {
	return Q.cap
}

// Resize resizes the circularBuffer to the specified capacity.
// The capacity must be positive.
// The elements in the circularBuffer are kept within the new capacity.
func (Q *Buffer[T]) Resize(newCap int) {
	//newCap must be non-negative
	if newCap <= 0 {
		panic("newCap must be positive")
	}
	var newAq Buffer[T] = *Q
	newAq.data = make([]T, newCap+1) //注意第0位不保存元素
	newAq.cap = newCap
	//队列元素在旧底层切片中最后一个元素的索引,注意第0位不保存元素
	endIndex := int(min(int64(Q.len)+int64(Q.first), int64(Q.cap)))
	if endIndex != 0 {
		copy(newAq.data[1:], Q.data[Q.first:endIndex+1])
	}
	//count是已经复制的元素个数，但不一定是全都复制到了新的底层切片中，因为新的底层切片容量可能比count小
	count := endIndex - Q.first + 1
	if newCap > Q.len {
		//Q.len 为总共的元素，count为已经复制的元素个数
		if count < Q.len {
			copy(newAq.data[count+1:], Q.data[1:Q.last+1])
		}
		newAq.last = Q.len
		newAq.len = Q.len
	} else {
		//Q.len 为总共的元素，count为已经复制的元素个数
		if count < Q.len && count < newCap {
			copy(newAq.data[count+1:], Q.data[1:Q.last+1])
		}
		newAq.last = newCap
		newAq.len = newCap
	}
	if newCap != 0 && Q.len != 0 {
		newAq.first = 1
	} else {
		newAq.first = 0
	}
	*Q = newAq
}

type minType interface {
	~int | ~int32 | ~int64
}

func min[T minType](a, b T) T {
	if a < b {
		return a
	}
	return b
}

// Iterator returns an iterator of the circularBuffer.
// The iterator is at the begin position by default.
// Do not call Push(), Pop(), Resize() on the circularBuffer while iterating,
// otherwise it may cause unpredictable errors.
// If you must call these methods, you need to get the iterator again,
// or move the iterator to the original position by calling MoveTo().
//  e.g.:
//  it := cb.Iterator()
//  index := it.Index()
//  que.Pop()
//  success := it.MoveTo(index)
//  ......
func (Q *Buffer[T]) Iterator() *CbIterator[T] {
	var it = &CbIterator[T]{
		que: Q,
		//第一个元素之前index = -1，
		//最后一个元素之后index = -2，
		//索引为0则程序可能出现了错误。
		index: -1,
	}
	return it
}

//将迭代器指向第一个元素之前，第一个元素之前index = -1

// Begin moves the iterator to the position before the first element.
func (c *CbIterator[T]) Begin() {
	c.index = -1
}

//迭代器当前所指元素的索引(队列中的第几个元素),计数从1开始,空队列返回0

// Index returns the index of the element the iterator points to.
// The index is counted from 1.
// It returns 0 if the queue is empty.
func (c *CbIterator[T]) Index() int {
	if c.index == -1 || c.index == -2 {
		return 0
	}
	//按我的设想index不应该为0
	if c.index >= c.que.first {
		return c.index - c.que.first + 1
	} else {
		return (c.que.cap - c.que.first + 1) + c.index
	}
}

//将迭代器指向最后一个元素之后,最后一个元素之后index = -2

// End moves the iterator to the position after the last element.
func (c *CbIterator[T]) End() {
	c.index = -2
}

//迭代器当前所指向元素的值,调用之前应该确保迭代器没有指向队列首部之前或末尾之后

// Value returns the value of the element the iterator points to.
// It panics if the iterator points to the position before the first element or after the last element.
func (c *CbIterator[T]) Value() T {
	return c.que.data[c.index]
}

//将迭代器指向下一个元素，如果迭代器所指的位置没有下一个元素，则调用Next()后会返回false

// Next moves the iterator to the next element and returns true.
// It returns false if there is no next element.
func (c *CbIterator[T]) Next() bool {
	//最后一个元素之后index = -2，
	if c.index == -2 {
		return false
	}
	if c.index == -1 {
		if c.que.len == 0 {
			return false
		} else {
			c.index = c.que.first
			//按我的设想index不应该为0
			return true
		}
	}
	//可能的情况，MoveTo失败且没有进行迭代器归位
	if c.Index() > c.que.len {
		c.index = -2
		return false
	}
	//迭代结束
	if c.index == c.que.last {
		c.index = -2
		return false
	}
	c.index = (c.index + 1) % (c.que.cap + 1)
	if c.index == 0 {
		c.index = 1
	}
	return true
}

//将迭代器指向上一个元素，如果迭代器所指的位置没有上一个元素，则调用Prev()后会返回false

// Prev moves the iterator to the previous element and returns true.
// It returns false if there is no previous element.
func (c *CbIterator[T]) Prev() bool {
	if c.index == -1 {
		return false
	}
	if c.index == -2 {
		if c.que.len == 0 {
			return false
		} else {
			c.index = c.que.last
			//按我的设想index不应该为0
			return true
		}
	}
	//可能的情况，MoveTo失败且没有进行迭代器归位
	if c.Index() > c.que.len {
		c.index = -1
		return false
	}
	if c.index == c.que.first {
		c.index = -1
		return false
	}
	c.index = (c.index - 1) % (c.que.cap + 1)
	if c.index == 0 {
		c.index = c.que.cap
	}
	return true
}

//将迭代器移动到目标索引(len >= index > 0)的位置，若索引不合法，则迭代器状态不变并返回false

// MoveTo moves the iterator to the position of the specified index.
// The index is counted from 1.
// It returns false if the index is invalid.
func (c *CbIterator[T]) MoveTo(index int) bool {
	if index > c.que.len || index <= 0 {
		return false
	}
	c.Begin()
	countFirstToSliceEnd := c.que.cap - c.que.first + 1
	if countFirstToSliceEnd >= index {
		c.index = c.que.first + index - 1
	} else {
		c.index = index - countFirstToSliceEnd
	}
	return true
}
