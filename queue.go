package arrayQueue

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

type Iterator[T any] interface {
	Begin()
	End()
	Next() bool
	Prev() bool
	Index() int
	Value() T
}

//arrayQueue的迭代器

// arrayQueue's iterator
type AqIterator[T any] struct {
	que   *Queue[T]
	index int //此处的index是队列元素在底层切片的实际索引
}
