package arrayqueue

/*
除了队列为空的情况，last和first都指向已经有值的位置，
当队列 只有一个值 或 没有值 的时候，first与last都是是重合的，
当队列 只有一个值 的时候，first与last是重合的，但位置不固定
当队列 没有值 的时候，last和first都指向0
每个函数都要考略的边界情况：
1.空队列调用
2.队列中只有一个值的时候调用
*/

//循环队列，底层是切片，初始容量为100,当使用Pop()发生扩容时,采用与append相同的策略
func New[T any]() *queue[T] {
	var aq queue[T]
	initCap := 100                 //初始容量
	aq.data = make([]T, initCap+1) //第0个位置不放数据
	//第0个位置不放数据
	//当队列 只有一个值 或 没有值 的时候，first与last是重合的，需要特别注意
	aq.last = 0
	aq.first = 0
	aq.cap = initCap
	aq.len = 0
	return &aq
}

func (Q *queue[T]) Push(value T) {
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
	//容量已满,进行扩容
	if next == Q.first {
		newCap := 0
		if Q.cap < 1024 {
			newCap = Q.cap * 2
		} else {
			newCap = int(float32(Q.cap) * 1.25)
		}
		Q.Resize(newCap)
		next = (Q.last + 1) % (Q.cap + 1)
	}
	Q.data[next] = value
	Q.last = next
	Q.len++
}

//直接遍历底层切片发送到channel,可能比迭代器遍历更快
func (Q *queue[T]) GetValueFromChannel() chan T {
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

func (Q *queue[T]) sendValue(ch chan<- T) {
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

func (Q *queue[T]) Pop() (value T) {
	if Q.len == 0 {
		panic("queue is empty!")
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

//返回队列第一个元素(最先插入),空队列调用会panic
func (Q *queue[T]) Front() T {
	if Q.len == 0 {
		panic("queue is empty!")
	}
	return Q.data[Q.first]
}

//返回队列最后一个元素(最后插入),空队列调用会panic
func (Q *queue[T]) Back() T {
	if Q.len == 0 {
		panic("queue is empty!")
	}
	return Q.data[Q.last]
}

func (Q *queue[T]) Empty() bool {
	return Q.len == 0
}

func (Q *queue[T]) Len() int {
	return Q.len
}

//重新分配队列底层内存空间,设置容量为n(n最小为2，减少push的判断条件,最大为makeslice的长度)，
//队列元素在n的范围内保持不变
func (Q *queue[T]) Resize(newCap int) {
	//newCap must be non-negative
	if newCap < 2 {
		newCap = 2
	}
	var newAq queue[T] = *Q
	newAq.data = make([]T, newCap+1) //注意第0位不保存元素
	newAq.cap = newCap

	//队列元素在旧底层切片中最后一个元素的索引
	endIndex := int(min(int64(Q.len)+int64(Q.first), int64(Q.cap)))
	copy(newAq.data[1:], Q.data[Q.first:endIndex+1])
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

//第一个元素之前index = -1，
//最后一个元素之后index = -2，
//索引为0则程序可能出现了错误。

//返回一个队列的迭代器，默认处于begin的位置。
//遍历时不要对队列调用Push()和Pop(),否者可能会出现不可预料的错误。
func (Q *queue[T]) Iterator() *aqIterator[T] {
	var it = &aqIterator[T]{
		que:   Q,
		index: -1,
	}
	return it
}

//将迭代器指向第一个元素之前，第一个元素之前index = -1
func (a *aqIterator[T]) Begin() {
	a.index = -1
}

//迭代器当前所指元素的索引(第几个元素),计数从1开始,空队列返回0
func (a *aqIterator[T]) Index() int {
	if a.index == -1 || a.index == -2 {
		return 0
	}
	//按我的设想index不应该为0
	if a.index >= a.que.first {
		return a.index - a.que.first + 1
	} else {
		return (a.que.cap - a.que.first + 1) + a.index
	}
}

//将迭代器指向最后一个元素之后,最后一个元素之后index = -2
func (a *aqIterator[T]) End() {
	a.index = -2
}

//迭代器当前所指向元素的值,调用之前应该确保迭代器没有指向队列首部之前或末尾之后
func (a *aqIterator[T]) Value() T {
	return a.que.data[a.index]
}

//将迭代器指向下一个元素，如果迭代器所指的位置没有下一个元素，则调用Next()后会返回false
func (a *aqIterator[T]) Next() bool {
	//最后一个元素之后index = -2，
	if a.index == -2 {
		return false
	}
	if a.index == -1 {
		if a.que.len == 0 {
			return false
		} else {
			a.index = a.que.first
			//按我的设想index不应该为0
			return true
		}
	}
	if a.index == a.que.last {
		a.index = -2
		return false
	}
	a.index = (a.index + 1) % (a.que.cap + 1)
	if a.index == 0 {
		a.index = 1
	}
	return true
}

//将迭代器指向上一个元素，如果迭代器所指的位置没有上一个元素，则调用Prev()后会返回false
func (a *aqIterator[T]) Prev() bool {
	if a.index == -1 {
		return false
	}
	if a.index == -2 {
		if a.que.len == 0 {
			return false
		} else {
			a.index = a.que.last
			//按我的设想index不应该为0
			return true
		}
	}
	if a.index == a.que.first {
		a.index = -1
		return false
	}
	a.index = (a.index - 1) % (a.que.cap + 1)
	if a.index == 0 {
		a.index = a.que.cap
	}
	return true
}