package arrayQueue

import (
	"fmt"
	"testing"

	aq "github.com/emirpasic/gods/queues/arrayqueue"
)

type operate struct {
	op string
	n  int
}

// var tests = []operate{
// 	{"push", 1},
// 	{"pop", 1},
// }

func TestQueue(t *testing.T) {
	var que = aq.New()
	var myque *Queue[int] = New[int]()
	var tests = []operate{
		{"push", 1},
		{"pop", 1},
		{"push", 80},
		{"pop", 80},
		{"push", 80},
		{"pop", 50},
		{"push", 100},
		{"pop", 130},
		{"push", 200},
		{"pop", 199},
		{"push", 10},
		{"pop", 11},
		{"push", 1044},
		{"pop", 800},
		{"push", 2002},
		{"pop", 2003},
	}
	for _, v := range tests {
		fmt.Println("start", v.op, v.n)
		if v.op == "push" {
			for i := 0; i < v.n; i++ {
				que.Enqueue(i + 1)
				myque.Push(i + 1)
				tem, _ := que.Peek()
				if tem.(int) != myque.Front() {
					t.Errorf("expected:%v, got:%v", tem, myque.Front())
				}
				if myque.Len() != que.Size() {
					t.Errorf("expected:%v, got:%v", que.Size(), myque.Len())
				}
			}
		} else if v.op == "pop" {
			for i := 0; i < v.n; i++ {
				value, _ := que.Dequeue()
				myValue := myque.Pop()
				if value.(int) != myValue {
					t.Errorf("expected:%v, got:%v", value.(int), myValue)
				}
				tem, ok := que.Peek()
				if !ok {
					if ok != (!myque.Empty()) {
						t.Errorf("expected:%v, got:%v", ok, !myque.Empty())
					}
				} else {
					if tem.(int) != myque.Front() {
						t.Errorf("expected:%v, got:%v", tem, myque.Front())
					}
				}
			}
		}
		if que.Empty() != myque.Empty() {
			t.Errorf("expected:%v, got:%v", que.Empty(), myque.Empty())
		}
		queIt := que.Iterator()
		myqueIt := myque.Iterator()
		i := 1
		for queIt.Next() {
			if myqueIt.Next() == false {
				t.Errorf("expected:%v, got:%v", true, false)
			}
			if myqueIt.Index() != i {
				t.Errorf("expected:%v, got:%v", i, myqueIt.Index())
			}
			if queIt.Value() != myqueIt.Value() {
				t.Errorf("expected:%v, got:%v", queIt.Value(), myqueIt.Value())
			}
			i++
		}
		if myqueIt.Next() == true {
			t.Errorf("expected:%v, got:%v", false, true)
		}
		if myqueIt.Next() == true {
			t.Errorf("expected:%v, got:%v", false, true)
		}
		queIt.Begin()
		if que.Empty() != myque.Empty() {
			t.Errorf("expected:%v, got:%v", que.Empty(), myque.Empty())
		}
		queIt2 := que.Iterator()
		myqueIt2 := myque.Iterator()
		i = que.Size()
		queIt2.End()
		myqueIt2.End()
		for queIt2.Prev() {
			if myqueIt2.Prev() == false {
				t.Errorf("expected:%v, got:%v", true, false)
			}
			if myqueIt2.Index() != i {
				t.Errorf("expected:%v, got:%v", i, myqueIt2.Index())
			}
			if queIt2.Value() != myqueIt2.Value() {
				t.Errorf("expected:%v, got:%v", queIt2.Value(), myqueIt2.Value())
			}
			i--
		}
		if myqueIt2.Prev() == true {
			t.Errorf("expected:%v, got:%v", false, true)
		}
	}
}

func TestQueue2(t *testing.T) {
	var myque *Queue[int] = New[int]()
	myqueIt := myque.Iterator()
	if myqueIt.Prev() == true {
		t.Errorf("expected:%v, got:%v", false, true)
	}
	if myqueIt.Next() == true {
		t.Errorf("expected:%v, got:%v", false, true)
	}
	if myqueIt.index != -1 {
		t.Errorf("expected:%v, got:%v", -1, myqueIt.index)
	}
	if myqueIt.Index() != 0 {
		t.Errorf("expected:%v, got:%v", 0, myqueIt.Index())
	}
	myqueIt.Begin()
	myqueIt = myque.Iterator()
	if myqueIt.Prev() == true {
		t.Errorf("expected:%v, got:%v", false, true)
	}
	if myqueIt.Next() == true {
		t.Errorf("expected:%v, got:%v", false, true)
	}
	if myqueIt.index != -1 {
		t.Errorf("expected:%v, got:%v", -1, myqueIt.index)
	}
	if myqueIt.Index() != 0 {
		t.Errorf("expected:%v, got:%v", 0, myqueIt.Index())
	}

	newcap := 2
	myque.Resize(newcap)
	if myque.cap != newcap || myque.len != 0 || myque.first != 0 || myque.last != 0 {
		t.Error("que.Resize() error")
	}
	myque.Push(99)
	myque.Resize(newcap)
	if myque.cap != newcap || myque.len != 1 || myque.first != 1 || myque.last != 1 {
		t.Error("que.Resize() error")
	}
	if myque.Front() != myque.Back() {
		t.Error("myque.Front() != myque.Back()")
	}
	myque.Push(99)
	myque.Resize(newcap)
	if myque.cap != newcap || myque.len != 2 || myque.first != 1 || myque.last != 2 {
		t.Error("que.Resize() error")
	}
	myque.Push(99)
	myque.Resize(newcap)
	if myque.cap != newcap || myque.len != 2 || myque.first != 1 || myque.last != 2 {
		t.Error("que.Resize() error")
	}
	myque.Pop()
	myque.Pop()
	newcap = 3
	myque.Resize(newcap)
	if myque.cap != newcap || myque.len != 0 || myque.first != 0 || myque.last != 0 {
		t.Error("que.Resize() error")
	}
	myque.Push(99)
	myque.Resize(newcap)
	if myque.cap != newcap || myque.len != 1 || myque.first != 1 || myque.last != 1 {
		t.Error("que.Resize() error")
	}
	myque.Push(99)
	myque.Resize(newcap)
	if myque.cap != newcap || myque.len != 2 || myque.first != 1 || myque.last != 2 {
		t.Error("que.Resize() error")
	}
	myque.Push(99)
	myque.Resize(newcap)
	if myque.cap != newcap || myque.len != 3 || myque.first != 1 || myque.last != 3 {
		t.Error("que.Resize() error")
	}

	var myque2 *Queue[int] = New[int]()
	var tests = []operate{
		{"push", 80},
		{"pop", 50},
		{"push", 50},
	}
	for _, v := range tests {
		fmt.Println("start", v.op, v.n)
		if v.op == "push" {
			for i := 0; i < v.n; i++ {
				myque2.Push(i + 1)
			}
		} else if v.op == "pop" {
			for i := 0; i < v.n; i++ {
				myque2.Pop()
			}
		}
	}
	newcap = 60
	myque2.Resize(newcap)
	if myque2.cap != newcap || myque2.len != 60 || myque2.first != 1 || myque2.last != 60 {
		t.Error("que.Resize() error")
	}
	myque.Resize(0)
	//myque.Resize(-1)
}

func TestQueue3(t *testing.T) {
	que := New[int]()
	for i := 0; i < 90; i++ {
		que.Push(i + 1)
	}
	for i := 0; i < 80; i++ {
		que.Pop()
	}
	for i := 100; i <= 120; i++ {
		que.Push(i)
	}
	it := que.Iterator()
	ok := it.MoveTo(1)
	if !ok {
		t.Errorf("expected:%v, got:%v", true, false)
	}
	for i := 0; i < 10; i++ {
		if it.Value() != 81+i {
			t.Errorf("expected:%v, got:%v", 81+i, it.Value())
		}
		it.Next()
	}
	for i := 0; i < 21; i++ {
		if it.Value() != 100+i {
			t.Errorf("expected:%v, got:%v", 100+i, it.Value())
		}
		it.Next()
	}
	que2 := New[int]()
	for i := 0; i < 30; i++ {
		que2.Push(i + 1)
	}
	it2 := que2.Iterator()
	for i := 0; i < 30; i++ {
		it2.MoveTo(i + 1)
		if it2.Value() != i+1 {
			t.Errorf("expected:%v, got:%v", i+1, it2.Value())
		}
	}
}

func TestQueue4(t *testing.T) {
	que := New[int]()
	que.Clear()
	for i := 0; i < 90; i++ {
		que.Push(i + 1)
	}
	for i := 0; i < 80; i++ {
		que.Pop()
	}
	for i := 100; i <= 120; i++ {
		que.Push(i)
	}
	it := que.Iterator()
	ok := it.MoveTo(1)
	if !ok {
		t.Errorf("expected:%v, got:%v", true, false)
	}
	for i := 0; i < 10; i++ {
		if it.Value() != 81+i {
			t.Errorf("expected:%v, got:%v", 81+i, it.Value())
		}
		if it.Index() != i+1 {
			t.Errorf("expected:%v, got:%v", 1+i, it.Index())
		}
		it.Next()
	}
	for i := 0; i < 21; i++ {
		if it.Value() != 100+i {
			t.Errorf("expected:%v, got:%v", 100+i, it.Value())
		}
		if it.Index() != 11+i {
			t.Errorf("expected:%v, got:%v", 11+i, it.Index())
		}
		it.Next()
	}
	ok = it.MoveTo(-1)
	if ok {
		t.Errorf("expected:%v, got:%v", false, true)
	}
	ok = it.MoveTo(21)
	if !ok {
		t.Errorf("expected:%v, got:%v", true, false)
	}
	for i := 0; i < 11; i++ {
		if it.Value() != 110+i {
			t.Errorf("expected:%v, got:%v", 110+i, it.Value())
		}
		if it.Index() != 21+i {
			t.Errorf("expected:%v, got:%v", 21+i, it.Index())
		}
		it.Next()
	}

	que.Clear()
	for i := 0; i < 30; i++ {
		que.Push(i + 1)
	}
	it2 := que.Iterator()
	for i := 0; i < 30; i++ {
		it2.MoveTo(i + 1)
		if it2.Value() != i+1 {
			t.Errorf("expected:%v, got:%v", i+1, it2.Value())
		}
	}
	que.Clear()
	que.Push(22)
	que.Push(22)
	que.Push(22)
	it = que.Iterator()
	for i := 0; i < 3; i++ {
		que.Pop()
	}
	if it.Next() != false {
		t.Errorf("expected:%v, got:%v", false, true)
	}
	if it.Prev() != false {
		t.Errorf("expected:%v, got:%v", false, true)
	}
	que.Clear()
	que.Push(22)
	que.Push(22)
	que.Push(22)
	it = que.Iterator()
	it.MoveTo(2)
	for i := 0; i < 3; i++ {
		que.Pop()
	}
	if it.Next() != false {
		t.Errorf("expected:%v, got:%v", false, true)
	}
	if it.Prev() != false {
		t.Errorf("expected:%v, got:%v", false, true)
	}
	que.Clear()
	que.Push(22)
	que.Push(22)
	que.Push(22)
	it = que.Iterator()
	it.MoveTo(2)
	for i := 0; i < 3; i++ {
		que.Pop()
	}
	if it.Prev() != false {
		t.Errorf("expected:%v, got:%v", false, true)
	}
	if it.Next() != false {
		t.Errorf("expected:%v, got:%v", false, true)
	}
	que.Clear()
	que.Push(22)
	que.Push(22)
	que.Push(22)
	it = que.Iterator()
	it.End()
	for i := 0; i < 3; i++ {
		que.Pop()
	}
	if it.Next() != false {
		t.Errorf("expected:%v, got:%v", false, true)
	}
	if it.Prev() != false {
		t.Errorf("expected:%v, got:%v", false, true)
	}
	que3 := New[int]()
	que3.Push(22)
	que3.Push(22)
	que3.Push(22)
	it = que3.Iterator()
	it.End()
	for i := 0; i < 3; i++ {
		que3.Pop()
	}
	if it.Next() != false {
		t.Errorf("expected:%v, got:%v", false, true)
	}
	if it.Prev() != false {
		t.Errorf("expected:%v, got:%v", false, true)
	}
}

func TestQueue5(t *testing.T) {
	que := New[int]()
	que.Push(1)
	que.Push(2)
	que.Resize(2)
	it := que.Iterator()
	if it.Next() {
		if it.Value() != 1 {
			t.Errorf("expected:%v, got:%v", 1, it.Value())
		}
	}
	if it.Next() {
		if it.Value() != 2 {
			t.Errorf("expected:%v, got:%v", 2, it.Value())
		}
	}
	if it.Next() {
		t.Errorf("expected:%v, got:%v", false, true)
	}
	que.Resize(2)
	it = que.Iterator()
	if it.Next() {
		if it.Value() != 1 {
			t.Errorf("expected:%v, got:%v", 1, it.Value())
		}
	}
	if it.Next() {
		if it.Value() != 2 {
			t.Errorf("expected:%v, got:%v", 2, it.Value())
		}
	}
	if it.Next() {
		t.Errorf("expected:%v, got:%v", false, true)
	}
	que2 := new(Queue[int])
	que2.Resize(2)
}

func TestQueue6(t *testing.T) {
	que := New[int]()
	que.Resize(0)
	if que.cap != 0 {
		t.Errorf("expected:%v, got:%v", 0, que.cap)
	}
	if que.len != 0 {
		t.Errorf("expected:%v, got:%v", 0, que.len)
	}
	if que.Empty() != true {
		t.Errorf("expected:%v, got:%v", true, que.Empty())
	}
	que.Resize(1)
	if que.cap != 1 {
		t.Errorf("expected:%v, got:%v", 1, que.cap)
	}
	if que.len != 0 {
		t.Errorf("expected:%v, got:%v", 0, que.len)
	}
	if que.Empty() != true {
		t.Errorf("expected:%v, got:%v", true, que.Empty())
	}
	que = New[int]()
	que.Resize(0)
	que.Push(1)
	if que.cap != 1 {
		t.Errorf("expected:%v, got:%v", 1, que.cap)
	}
	if que.len != 1 {
		t.Errorf("expected:%v, got:%v", 1, que.len)
	}
	if que.Empty() != false {
		t.Errorf("expected:%v, got:%v", false, que.Empty())
	}
	que.Resize(1)
	if que.cap != 1 {
		t.Errorf("expected:%v, got:%v", 1, que.cap)
	}
	if que.len != 1 {
		t.Errorf("expected:%v, got:%v", 1, que.len)
	}
	if que.Empty() != false {
		t.Errorf("expected:%v, got:%v", false, que.Empty())
	}
}

func TestQueue7(t *testing.T) {
	que := New[int]()
	que.Resize(0)
	if que.Cap() != 0 {
		t.Errorf("expected:%v, got:%v", 0, que.Cap())
	}
	que.Resize(10)
	for i := 0; i < 10; i++ {
		que.Push(i)
	}
	if que.Cap() != 10 {
		t.Errorf("expected:%v, got:%v", 10, que.Cap())
	}
	for i := 0; i < 5; i++ {
		que.Pop()
	}
	for i := 0; i < 5; i++ {
		que.Push(i)
	}
	if que.Cap() != 10 {
		t.Errorf("expected:%v, got:%v", 10, que.Cap())
	}
	it := que.Iterator()
	it.MoveTo(4)
	for i := 4; it.Next(); i++ {
		if it.Index() != i+1 {
			t.Errorf("expected:%v, got:%v", i+1, it.Index())
		}
	}
	it.MoveTo(6)
	if it.Index() != 6 {
		t.Errorf("expected:%v, got:%v", 6, it.Index())
	}
	for i := 6; it.Prev(); i-- {
		if it.Index() != i-1 {
			t.Errorf("expected:%v, got:%v", i+1, it.Index())
		}
	}
	for i := 0; i < 10; i++ {
		que.Pop()
	}
	que2 := New[int]()
	que2.Resize(0)
	if que2.Cap() != 0 {
		t.Errorf("expected:%v, got:%v", 0, que2.Cap())
	}
	que2.Resize(10)
	for i := 0; i < 10; i++ {
		que2.Push(i)
	}
	if que2.Cap() != 10 {
		t.Errorf("expected:%v, got:%v", 10, que2.Cap())
	}
	for i := 0; i < 5; i++ {
		que2.Pop()
	}
	for i := 0; i < 5; i++ {
		que2.Push(i)
	}
	if que2.Cap() != 10 {
		t.Errorf("expected:%v, got:%v", 10, que2.Cap())
	}
	que2.Resize(7)
	if que2.Cap() != 7 {
		t.Errorf("expected:%v, got:%v", 7, que2.Cap())
	}
	if que2.Len() != 7 {
		t.Errorf("expected:%v, got:%v", 7, que2.Len())
	}
}

func Test_Resize(t *testing.T) {
	//没有元素的情况
	que := New[int]()
	que.Resize(0)
	if que.cap != 0 {
		t.Errorf("expected:%v, got:%v", 0, que.cap)
	}
	if que.len != 0 {
		t.Errorf("expected:%v, got:%v", 0, que.len)
	}
	if que.Empty() != true {
		t.Errorf("expected:%v, got:%v", true, que.Empty())
	}
	que.Push(1)
	if que.cap != 1 {
		t.Errorf("expected:%v, got:%v", 1, que.cap)
	}
	if que.len != 1 {
		t.Errorf("expected:%v, got:%v", 1, que.len)
	}
	if que.Empty() != false {
		t.Errorf("expected:%v, got:%v", false, que.Empty())
	}
	que.Push(2)
	if que.cap != 2 {
		t.Errorf("expected:%v, got:%v", 2, que.cap)
	}
	if que.len != 2 {
		t.Errorf("expected:%v, got:%v", 2, que.len)
	}
	if que.Empty() != false {
		t.Errorf("expected:%v, got:%v", false, que.Empty())
	}
	que.Push(3)
	if que.cap != 4 {
		t.Errorf("expected:%v, got:%v", 4, que.cap)
	}
	if que.len != 3 {
		t.Errorf("expected:%v, got:%v", 3, que.len)
	}
	if que.Empty() != false {
		t.Errorf("expected:%v, got:%v", false, que.Empty())
	}
	for i := 0; i < 3; i++ {
		que.Pop()
		if que.cap != 4 {
			t.Errorf("expected:%v, got:%v", 4, que.cap)
		}
		if que.len != 2-i {
			t.Errorf("expected:%v, got:%v", 2-i, que.len)
		}
	}
	//0 resize 0
	que = New[int]()
	que.Resize(0)
	if que.cap != 0 {
		t.Errorf("expected:%v, got:%v", 0, que.cap)
	}
	if que.len != 0 {
		t.Errorf("expected:%v, got:%v", 0, que.len)
	}
	if que.Empty() != true {
		t.Errorf("expected:%v, got:%v", true, que.Empty())
	}
	//0 resize 1
	que = New[int]()
	que.Resize(1)
	if que.cap != 1 {
		t.Errorf("expected:%v, got:%v", 1, que.cap)
	}
	if que.len != 0 {
		t.Errorf("expected:%v, got:%v", 0, que.len)
	}
	if que.Empty() != true {
		t.Errorf("expected:%v, got:%v", true, que.Empty())
	}
	if que.first != 0 {
		t.Errorf("expected:%v, got:%v", 0, que.first)
	}
	if que.last != 0 {
		t.Errorf("expected:%v, got:%v", 0, que.last)
	}
	//0 resize 2
	que = New[int]()
	que.Resize(2)
	if que.cap != 2 {
		t.Errorf("expected:%v, got:%v", 2, que.cap)
	}
	if que.len != 0 {
		t.Errorf("expected:%v, got:%v", 0, que.len)
	}
	if que.Empty() != true {
		t.Errorf("expected:%v, got:%v", true, que.Empty())
	}
	if que.first != 0 {
		t.Errorf("expected:%v, got:%v", 0, que.first)
	}
	if que.last != 0 {
		t.Errorf("expected:%v, got:%v", 0, que.last)
	}
	//1 resize 0
	que = New[int]()
	que.Push(1)
	que.Resize(0)
	if que.cap != 0 {
		t.Errorf("expected:%v, got:%v", 0, que.cap)
	}
	if que.len != 0 {
		t.Errorf("expected:%v, got:%v", 0, que.len)
	}
	if que.Empty() != true {
		t.Errorf("expected:%v, got:%v", true, que.Empty())
	}
	if que.first != 0 {
		t.Errorf("expected:%v, got:%v", 0, que.first)
	}
	if que.last != 0 {
		t.Errorf("expected:%v, got:%v", 0, que.last)
	}
	//1 resize 1
	que = New[int]()
	que.Push(1)
	que.Resize(1)
	if que.cap != 1 {
		t.Errorf("expected:%v, got:%v", 1, que.cap)
	}
	if que.len != 1 {
		t.Errorf("expected:%v, got:%v", 1, que.len)
	}
	if que.Empty() != false {
		t.Errorf("expected:%v, got:%v", false, que.Empty())
	}
	if que.first != 1 {
		t.Errorf("expected:%v, got:%v", 1, que.first)
	}
	if que.last != 1 {
		t.Errorf("expected:%v, got:%v", 1, que.last)
	}
	if que.Front() != 1 {
		t.Errorf("expected:%v, got:%v", 1, que.Front())
	}
	if que.Back() != 1 {
		t.Errorf("expected:%v, got:%v", 1, que.Back())
	}
	//1 resize 2
	que = New[int]()
	que.Push(1)
	que.Resize(2)
	if que.cap != 2 {
		t.Errorf("expected:%v, got:%v", 2, que.cap)
	}
	if que.len != 1 {
		t.Errorf("expected:%v, got:%v", 1, que.len)
	}
	if que.Empty() != false {
		t.Errorf("expected:%v, got:%v", false, que.Empty())
	}
	if que.first != 1 {
		t.Errorf("expected:%v, got:%v", 1, que.first)
	}
	if que.last != 1 {
		t.Errorf("expected:%v, got:%v", 1, que.last)
	}
	if que.Front() != 1 {
		t.Errorf("expected:%v, got:%v", 1, que.Front())
	}
	if que.Back() != 1 {
		t.Errorf("expected:%v, got:%v", 1, que.Back())
	}
	//2 resize 0
	que = New[int]()
	que.Push(1)
	que.Push(2)
	que.Resize(0)
	if que.cap != 0 {
		t.Errorf("expected:%v, got:%v", 0, que.cap)
	}
	if que.len != 0 {
		t.Errorf("expected:%v, got:%v", 0, que.len)
	}
	if que.Empty() != true {
		t.Errorf("expected:%v, got:%v", true, que.Empty())
	}
	if que.first != 0 {
		t.Errorf("expected:%v, got:%v", 0, que.first)
	}
	if que.last != 0 {
		t.Errorf("expected:%v, got:%v", 0, que.last)
	}
	//2 resize 1
	que = New[int]()
	que.Push(1)
	que.Push(2)
	que.Resize(1)
	if que.cap != 1 {
		t.Errorf("expected:%v, got:%v", 1, que.cap)
	}
	if que.len != 1 {
		t.Errorf("expected:%v, got:%v", 1, que.len)
	}
	if que.Empty() != false {
		t.Errorf("expected:%v, got:%v", false, que.Empty())
	}
	if que.first != 1 {
		t.Errorf("expected:%v, got:%v", 1, que.first)
	}
	if que.last != 1 {
		t.Errorf("expected:%v, got:%v", 1, que.last)
	}
	if que.Front() != 1 {
		t.Errorf("expected:%v, got:%v", 1, que.Front())
	}
	if que.Back() != 1 {
		t.Errorf("expected:%v, got:%v", 1, que.Back())
	}
	//2 resize 2
	que = New[int]()
	que.Push(1)
	que.Push(2)
	que.Resize(2)
	if que.cap != 2 {
		t.Errorf("expected:%v, got:%v", 2, que.cap)
	}
	if que.len != 2 {
		t.Errorf("expected:%v, got:%v", 2, que.len)
	}
	if que.Empty() != false {
		t.Errorf("expected:%v, got:%v", false, que.Empty())
	}
	if que.first != 1 {
		t.Errorf("expected:%v, got:%v", 1, que.first)
	}
	if que.last != 2 {
		t.Errorf("expected:%v, got:%v", 2, que.last)
	}
	if que.Front() != 1 {
		t.Errorf("expected:%v, got:%v", 1, que.Front())
	}
	if que.Back() != 2 {
		t.Errorf("expected:%v, got:%v", 2, que.Back())
	}
	//2 resize 3
	que = New[int]()
	que.Push(1)
	que.Push(2)
	que.Resize(3)
	if que.cap != 3 {
		t.Errorf("expected:%v, got:%v", 3, que.cap)
	}
	if que.len != 2 {
		t.Errorf("expected:%v, got:%v", 2, que.len)
	}
	if que.Empty() != false {
		t.Errorf("expected:%v, got:%v", false, que.Empty())
	}
	if que.first != 1 {
		t.Errorf("expected:%v, got:%v", 1, que.first)
	}
	if que.last != 2 {
		t.Errorf("expected:%v, got:%v", 2, que.last)
	}
	if que.Front() != 1 {
		t.Errorf("expected:%v, got:%v", 1, que.Front())
	}
	if que.Back() != 2 {
		t.Errorf("expected:%v, got:%v", 2, que.Back())
	}
}
