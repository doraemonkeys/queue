package circularQueue

//first和last在队列不为空时指向已经有值的位置
type Queue[T any] struct {
	data  []T
	last  int //最新插入的元素
	first int //最先插入的元素
	len   int //队列元素长度
	//表示已申请的内存能容纳元素的最大个数
	//容量 cap = 切片真实长度 - 1 (第0个位置不放数据)
	cap int
}

type Iterator[T any] interface {
	Begin()
	End()
	Next() bool
	Prev() bool
	Index() int
	Value() T
}

//circularQueue的迭代器
type CqIterator[T any] struct {
	que   *Queue[T]
	index int//此处的index是队列元素在底层切片的实际索引
}
