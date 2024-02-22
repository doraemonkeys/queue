package arrayQueue

/*
除了队列为空的情况，last和first都指向已经有值的位置，
当队列 只有一个值 或 没有值 的时候，first与last都是是重合的，
当队列 只有一个值 的时候，first与last是重合的，但位置不固定
当队列 没有值 的时候，last和first都指向0
每个函数都要考虑的边界情况：
1.空队列调用
2.队列中只有一个值的时候调用
*/

//推荐通过queue的New()方法获取新队列，若不是，请使用Resize()初始化底层切片

// Recommand to get a new queue by queue's New() method,
// if not, please use Resize() to initialize the underlying slice.
type Queue[T any] struct {
	data  []T
	last  int //最新插入的元素,first和last在队列不为空时指向已经有值的位置
	first int //最先插入的元素,first和last在队列不为空时指向已经有值的位置
	len   int //队列元素长度
	//表示已申请的内存能容纳元素的最大个数
	//容量 cap = 切片真实长度 - 1 (第0个位置不放数据)
	cap int
}

//arrayQueue的迭代器

// arrayQueue's iterator
type AqIterator[T any] struct {
	que   *Queue[T]
	index int //此处的index是队列元素在底层切片的实际索引
}

//循环队列，底层是切片，初始容量为1,当使用Pop()发生扩容时,采用与append相同的策略

// Queue is a circular buffer queue, implemented using a slice.
// The initial capacity of the queue is one,
// when using Pop() to expand, the same strategy as append is used.
func New[T any]() *Queue[T] {
	var cq Queue[T]
	initCap := 1                   //初始容量
	cq.data = make([]T, initCap+1) //第0个位置不放数据
	//第0个位置不放数据
	//当队列 只有一个值 或 没有值 的时候，first与last是重合的，需要特别注意
	cq.last = 0
	cq.first = 0
	cq.cap = initCap
	cq.len = 0
	return &cq
}

func (Q *Queue[T]) Push(value T) {
	if Q.len == 0 {
		if Q.cap == 0 {
			Q.Resize(1)
		}
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
	//容量已满,进行扩容
	if next == Q.first {
		// 扩容策略与append相同，不使用append是因为first和last的位置不固定
		newCap := 0
		if Q.cap < 1024 {
			newCap = Q.cap * 2
		} else {
			//扩容到原来的1.25倍,超过int的最大值时,扩容到int的最大值
			newCapF := float32(Q.cap) * 1.25
			if newCapF > float32(int(^uint(0)>>1)) {
				newCap = int(^uint(0) >> 1)
			} else {
				newCap = int(newCapF)
			}
		}
		Q.Resize(newCap)
		next = (Q.last + 1) % (Q.cap + 1)
	}
	Q.data[next] = value
	Q.last = next
	Q.len++
}

//对空队列调用会导致panic,
//pop不会释放内存，没有太大性能消耗，释放内存可以调用Resize()。

// Pop removes and returns the first element of the queue.
// It panics if the queue is empty.
// Pop does not release memory, there is no significant performance loss,
// and memory can be released by calling Resize().
func (Q *Queue[T]) Pop() (value T) {
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

//清空队列
//Clear不会释放内存，没有太大性能消耗。

// Clear removes all elements from the queue.
// Clear does not release memory, there is no significant performance loss,
// and memory can be released by calling Resize().
func (Q *Queue[T]) Clear() {
	if Q.len == 0 {
		return
	}
	Q.len = 0
	Q.first = 0
	Q.last = 0
}

//返回队列第一个元素(最先插入),空队列调用会panic

// Front returns the first element of the queue.
// It panics if the queue is empty.
func (Q *Queue[T]) Front() T {
	if Q.len == 0 {
		panic("queue is empty")
	}
	return Q.data[Q.first]
}

//返回队列最后一个元素(最后插入),空队列调用会panic

// Back returns the last element of the queue.
// It panics if the queue is empty.
func (Q *Queue[T]) Back() T {
	if Q.len == 0 {
		panic("queue is empty")
	}
	return Q.data[Q.last]
}

// Empty returns true if the queue is empty.
func (Q *Queue[T]) Empty() bool {
	return Q.len == 0
}

//返回队列的长度

// Len returns the number of elements of the queue.
func (Q *Queue[T]) Len() int {
	return Q.len
}

//返回队列的容量

// Cap returns the capacity of the queue.
func (Q *Queue[T]) Cap() int {
	return Q.cap
}

//重新分配队列底层内存空间,设置容量为newCap(newCap最小0,最大为makeslice的长度),
//队列元素在newCap的范围内保持不变。

// Resize resizes the queue to the specified capacity.
// The capacity must be non-negative.
// The elements in the queue are kept within the new capacity.
func (Q *Queue[T]) Resize(newCap int) {
	//newCap must be non-negative
	if newCap < 0 {
		panic("newCap must be non-negative")
	}
	var newAq Queue[T] = *Q
	newAq.data = make([]T, newCap+1) //注意第0位不保存元素
	newAq.cap = newCap
	if newCap == 0 {
		newAq.len = 0
		newAq.first = 0
		newAq.last = 0
		*Q = newAq
		return
	}
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

// 返回一个队列的迭代器，默认处于begin的位置。
//  遍历时不要对队列调用 Push()、Pop()、Resize(),否者可能会出现不可预料的错误,
//  若必须调用这些方法，则需在调用后重新获取迭代器,或者通过 MoveTo()移动到原来的位置。
//  e.g.:
//  it := que.Iterator()
//  index := it.Index()
//  que.Pop()
//  success := it.MoveTo(index)
//  ......

// Iterator returns an iterator of the queue.
// The iterator is at the begin position by default.
// Do not call Push(), Pop(), Resize() on the queue while iterating,
// otherwise it may cause unpredictable errors.
// If you must call these methods, you need to get the iterator again,
// or move the iterator to the original position by calling MoveTo().
//  e.g.:
//  it := que.Iterator()
//  index := it.Index()
//  que.Pop()
//  success := it.MoveTo(index)
//  ......
func (Q *Queue[T]) Iterator() *AqIterator[T] {
	var it = &AqIterator[T]{
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
func (c *AqIterator[T]) Begin() {
	c.index = -1
}

//迭代器当前所指元素的索引(队列中的第几个元素),计数从1开始,空队列返回0

// Index returns the index of the element the iterator points to.
// The index is counted from 1.
// It returns 0 if the queue is empty.
func (c *AqIterator[T]) Index() int {
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
func (c *AqIterator[T]) End() {
	c.index = -2
}

//迭代器当前所指向元素的值,调用之前应该确保迭代器没有指向队列首部之前或末尾之后

// Value returns the value of the element the iterator points to.
// It panics if the iterator points to the position before the first element or after the last element.
func (c *AqIterator[T]) Value() T {
	return c.que.data[c.index]
}

//将迭代器指向下一个元素，如果迭代器所指的位置没有下一个元素，则调用Next()后会返回false

// Next moves the iterator to the next element and returns true.
// It returns false if there is no next element.
func (c *AqIterator[T]) Next() bool {
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
func (c *AqIterator[T]) Prev() bool {
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
func (c *AqIterator[T]) MoveTo(index int) bool {
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
