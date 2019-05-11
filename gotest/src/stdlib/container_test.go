package stdlib

import (
	"container/heap"
	"container/list"
	"container/ring"
	"fmt"
	"math/rand"
	"testing"
	"time"
	"unsafe"
)

var printList = func(l *list.List) {
	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Print(e.Value, " ")
	}
	fmt.Println()
}

// list.List is a doubly linked list
func TestList(t *testing.T) {
	l := list.New()
	ea := l.PushBack("a")
	eb := l.PushBack("b")
	ec := l.PushBack("c")
	e1 := l.PushFront("1")
	e2 := l.PushFront("2")
	e3 := l.PushFront("3")
	printList(l)
	fmt.Println(l.Len())
	fmt.Println(l.Front().Value)
	fmt.Println(l.Back().Value)
	ed := l.InsertAfter("d", eb)
	ee := l.InsertBefore("e", ea)
	printList(l)
	l.MoveAfter(e3, ec)
	l.MoveBefore(ea, e2)
	printList(l)
	l.MoveToFront(e1)
	l.MoveToBack(ee)
	l.Remove(ed)
	printList(l)
}

type intHeap []int

func (h intHeap) Len() int           { return len(h) }
func (h intHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h intHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *intHeap) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(int))
}
func (h *intHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

// list.Heap implements heap
func TestHeap(t *testing.T) {
	rand.Seed(time.Now().Unix())
	rands := rand.Perm(15)
	cands := rands[10:]
	ints := intHeap(rands[:10])
	fmt.Println("ints: ", ints, ", cands: ", cands)
	heap.Init(&ints)
	fmt.Println("heapified: ", ints)

	heap.Push(&ints, cands[0])
	fmt.Println("ist", cands[0], ":", ints)

	pop := heap.Pop(&ints)
	fmt.Println("pop", pop, ":", ints)

	heap.Push(&ints, cands[1])
	fmt.Println("ist", cands[1], ":", ints)

	pop = heap.Pop(&ints)
	fmt.Println("pop", pop, ":", ints)

	heap.Push(&ints, cands[2])
	fmt.Println("ist", cands[2], ":", ints)

	pop = heap.Pop(&ints)
	fmt.Println("pop", pop, ":", ints)

	pop = heap.Remove(&ints, 7)
	fmt.Println("rmv", pop, ":", ints)

	([]int)(ints)[6] = cands[3]
	heap.Fix(&ints, 6)
	fmt.Println("fix", cands[3], ":", ints)

	pop = heap.Remove(&ints, 5)
	fmt.Println("rmv", pop, ":", ints)

	([]int(ints))[4] = cands[4]
	heap.Fix(&ints, 4)
	fmt.Println("fix", cands[4], ":", ints)
}

// heap.Ring implements circular list
func TestRing(t *testing.T) {
	r := ring.New(5)
	n := r.Len()
	// Initialize the ring with some integer values
	for i := 0; i < n; i++ {
		r.Value = i
		r = r.Next()
	}
	// Iterate through the ring and print its contents
	r.Do(func(p interface{}) {
		fmt.Print(p.(int), " ")
	})
	fmt.Println()

	r = r.Next()
	r.Do(func(p interface{}) {
		fmt.Print(p.(int), " ")
	})
	fmt.Println()

	r = r.Move(2)
	r.Do(func(p interface{}) {
		fmt.Print(p.(int), " ")
	})
	fmt.Println()

	for i := 0; i < n; i++ {
		fmt.Println(unsafe.Pointer(r))
		r = r.Next()
	}
}