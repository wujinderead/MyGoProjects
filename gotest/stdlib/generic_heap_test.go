package stdlib

import (
	"container/heap"
	"fmt"
	"math"
	"testing"

	"golang.org/x/exp/constraints"
)

type genericHeap[K any] []K

func (h genericHeap[K]) Len() int {
	return len(h)
}

func (h genericHeap[K]) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h genericHeap[K]) Less(i, j int) bool {
	return false
}

func (h *genericHeap[K]) Push(x any) {
	*h = append(*h, x.(K))
}

func (h *genericHeap[K]) Pop() any {
	x := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return x
}

type HeapOrdered[K constraints.Ordered] struct {
	genericHeap[K]
}

func NewHeapOrdered[K constraints.Ordered](arr []K) heap.Interface {
	return &HeapOrdered[K]{
		genericHeap: genericHeap[K](arr),
	}
}

func (h HeapOrdered[K]) Less(i, j int) bool {
	return h.genericHeap[i] < h.genericHeap[j]
}

type HeapLesser[K Lesser[K]] struct {
	genericHeap[K]
}

func (h HeapLesser[K]) Less(i, j int) bool {
	return h.genericHeap[i].Less(h.genericHeap[j])
}

func NewHeapLesser[K Lesser[K]](arr []K) heap.Interface {
	return &HeapLesser[K]{
		genericHeap: genericHeap[K](arr),
	}
}

type HeapByLessFunc[K any] struct {
	genericHeap[K]
	less func(a, b K) bool
}

func (h HeapByLessFunc[K]) Less(i, j int) bool {
	return h.less(h.genericHeap[i], h.genericHeap[j])
}

func NewHeapByLessFunc[K any](arr []K, less func(a, b K) bool) heap.Interface {
	return &HeapByLessFunc[K]{
		genericHeap: genericHeap[K](arr),
		less:        less,
	}
}

type HeapByKeyFunc[K any, V constraints.Ordered] struct {
	genericHeap[K]
	key func(a K) V
}

func (h HeapByKeyFunc[K, V]) Less(i, j int) bool {
	return h.key(h.genericHeap[i]) < h.key(h.genericHeap[j])
}

func NewHeapByKeyFunc[K any, V constraints.Ordered](arr []K, key func(a K) V) heap.Interface {
	return &HeapByKeyFunc[K, V]{
		genericHeap: genericHeap[K](arr),
		key:         key,
	}
}

func TestHeapOrdered(t *testing.T) {
	{
		var hh heap.Interface = NewHeapOrdered[int]([]int{4, 5, 1, 2, 8, 6})
		heap.Init(hh)
		for i := 0; i < 2; i++ {
			var x int = heap.Pop(hh).(int)
			fmt.Println(x)
		}
		heap.Push(hh, 2)
		heap.Push(hh, 3)
		for hh.Len() > 0 {
			var x int = heap.Pop(hh).(int)
			fmt.Println(x)
		}
	}
	{
		var hh heap.Interface = NewHeapByLessFunc[int](
			[]int{-4, 5, -1, 2, -8, 6},
			func(a, b int) bool { return a*a < b*b },
		)
		heap.Init(hh)
		for i := 0; i < 2; i++ {
			var x int = heap.Pop(hh).(int)
			fmt.Println(x)
		}
		heap.Push(hh, -2)
		heap.Push(hh, 3)
		for hh.Len() > 0 {
			var x int = heap.Pop(hh).(int)
			fmt.Println(x)
		}
	}
	{
		type Obj struct {
			a int
			b string
		}
		var hh heap.Interface = NewHeapByKeyFunc[Obj, float64](
			[]Obj{{-4, "a"}, {5, "b"}, {-1, "c"}, {-8, "d"}},
			func(x Obj) float64 { return math.Abs(float64(x.a)) },
		)
		heap.Init(hh)
		for i := 0; i < 2; i++ {
			var x Obj = heap.Pop(hh).(Obj)
			fmt.Println(x)
		}
		heap.Push(hh, Obj{-2, "x"})
		heap.Push(hh, Obj{3, "t"})
		for hh.Len() > 0 {
			var x Obj = heap.Pop(hh).(Obj)
			fmt.Println(x)
		}
	}
}

type MyHeap[K any] interface {
	Push(k K)
	Pop() K
	Len() int
}

type MyHeapImpl[K any] struct {
	h heap.Interface
}

func NewMyHeapOrdered[K constraints.Ordered](arr []K) MyHeap[K] {
	x := MyHeapImpl[K]{
		h: NewHeapOrdered[K](arr),
	}
	heap.Init(x.h)
	return x
}

func NewMyHeapByLessFunc[K any](arr []K, less func(a, b K) bool) MyHeap[K] {
	x := MyHeapImpl[K]{
		h: NewHeapByLessFunc[K](arr, less),
	}
	heap.Init(x.h)
	return x
}

func NewMyHeapByKeyFunc[K any, V constraints.Ordered](arr []K, key func(a K) V) MyHeap[K] {
	x := MyHeapImpl[K]{
		h: NewHeapByKeyFunc[K](arr, key),
	}
	heap.Init(x.h)
	return x
}

func NewMyHeapLesser[K Lesser[K]](arr []K) MyHeap[K] {
	x := MyHeapImpl[K]{
		h: NewHeapLesser[K](arr),
	}
	heap.Init(x.h)
	return x
}

func (h MyHeapImpl[K]) Push(k K) {
	heap.Push(h.h, k)
}

func (h MyHeapImpl[K]) Pop() K {
	return heap.Pop(h.h).(K)
}

func (h MyHeapImpl[K]) Len() int {
	return h.h.Len()
}

func TestMyHeapOrdered(t *testing.T) {
	{
		var h MyHeap[int] = NewMyHeapOrdered[int]([]int{4, 3, 7, 6, 8, 1, 9})
		for i := 0; i < 2; i++ {
			var x int = h.Pop()
			fmt.Println(x)
		}
		for i := 0; i < 2; i++ {
			var x int = i + 2
			h.Push(x)
		}
		for h.Len() > 0 {
			var x int = h.Pop()
			fmt.Println(x)
		}
		// compile failed: dummy does not implement constraints.Ordered
		//type dummy struct {}
		//var _ MyHeap[dummy] = NewMyHeapOrdered[dummy]([]dummy{})
	}
	{
		type Obj struct {
			a int
			b string
		}
		var h MyHeap[Obj] = NewMyHeapByLessFunc[Obj](
			[]Obj{{-4, "a"}, {5, "b"}, {1, "c"}, {-8, "d"}},
			func(a, b Obj) bool { return a.a*a.a < b.a*b.a },
		)
		for h.Len() > 0 {
			var x Obj = h.Pop()
			fmt.Println(x)
		}
	}
	{
		type Obj struct {
			a int
			b string
		}
		var h MyHeap[Obj] = NewMyHeapByKeyFunc[Obj, float64](
			[]Obj{{-4, "a"}, {5, "b"}, {1, "c"}, {-8, "d"}},
			func(a Obj) float64 { return math.Abs(float64(a.a)) },
		)
		for h.Len() > 0 {
			var x Obj = h.Pop()
			fmt.Println(x)
		}

		// compile failed: vObj does not implement constraints.Ordered
		/*
			type vObj struct {
				aa int
				bb string
			}
			var _ MyHeap[Obj] = NewMyHeapByKeyFunc[Obj, vObj](
				[]Obj{{-4, "a"}, {5, "b"}, {1, "c"}, {-8, "d"}},
				func(a Obj) vObj { return vObj{} },
			)
		*/
	}
}

/* can't compile: `cannot use io.Reader in union`, we can't use an interface with func in a type union
type MyOrdered interface {
	constraints.Ordered | io.Reader
}
*/

type Lesser[K any] interface {
	Less(other K) bool
}

type Obj struct {
	A int
	B string
}

func (o Obj) Less(other Obj) bool {
	return o.A*o.A < other.A*other.A
}

type MyInt int

func (a MyInt) Less(other MyInt) bool {
	return a*a < other*other
}

func TestLesser(t *testing.T) {
	var a Lesser[Obj] = Obj{A: -2, B: ""}
	fmt.Println(a.Less(Obj{A: -2, B: ""}))
	fmt.Println(a.Less(Obj{A: 1, B: ""}))
	fmt.Println(a.Less(Obj{A: -3, B: ""}))

	var b Lesser[MyInt] = MyInt(-2)
	var c = MyInt(3)
	fmt.Println(b.Less(1))
	fmt.Println(b.Less(c))
	fmt.Println(b.Less(-3))

	objs := []Obj{{A: 1, B: ""}, {A: -4, B: ""}, {A: 3, B: ""}, {A: -2, B: ""}}
	var h heap.Interface = NewHeapLesser[Obj](objs)
	for h.Len() > 0 {
		fmt.Println(heap.Pop(h).(Obj))
	}

	{
		objs := []Obj{{A: 1, B: ""}, {A: -4, B: ""}, {A: 3, B: ""}, {A: -2, B: ""}}
		var h MyHeap[Obj] = NewMyHeapLesser[Obj](objs)
		for h.Len() > 0 {
			var obj Obj = h.Pop()
			fmt.Println(obj)
		}
	}
	{
		myInts := []MyInt{-3, 1, -5, 7, -4, 9}
		var h MyHeap[MyInt] = NewMyHeapLesser[MyInt](myInts)
		var i1 MyInt = h.Pop()
		fmt.Println(i1)
		var i2 MyInt = h.Pop()
		fmt.Println(i2)
		h.Push(MyInt(2))
		h.Push(MyInt(-6))
		for h.Len() > 0 {
			var i MyInt = h.Pop()
			fmt.Println(i)
		}
	}
}
