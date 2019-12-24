package heap

import (
	"container/heap"
	"fmt"
	"math/rand"
	"sort"
	"testing"
	"time"
)

// e.g. initial array keys=[b, d, a, e, c], original index [0, 1, 2, 3, 4]
// heapify, heap=[a, c, b, d, e], original index oriind=[2, 4, 0, 1, 3], means heap[i] == keys[h.oriind[i]];
// inheap=[2, 3, 0, 4, 1], means keys[i] == heap[inheap[i]], i.e. the i-th key is at inheap[i] position.

func TestDoubleHeap(t *testing.T) {
	n := 100
	f := 20
	strs := make([]string, n)
	for i := range strs {
		strs[i] = getRandomStr(i + 1)
	}
	rander.Shuffle(n, func(i, j int) {
		strs[i], strs[j] = strs[j], strs[i]
	})
	// get top 30 max keys (in lexicographical order) while keep the string length in another order
	lexi := newHeap(strs[:f], n)
	leng := newOHeap(strs[:f], n)
	heap.Init(lexi)
	heap.Init(leng)
	for i := f; i < n; i++ {
		if strs[i] < lexi.keys[0] { // current > min-in-heap, should substitute
			poporiind := lexi.oriind[0]
			popstr := lexi.keys[0]
			assertTrue(t, leng.keys[leng.inheap[poporiind]] == lexi.keys[lexi.inheap[poporiind]])
			assertTrue(t, leng.keys[leng.inheap[poporiind]] == popstr)
			assertTrue(t, strs[poporiind] == popstr)
			heap.Pop(lexi)
			heap.Push(lexi, struct {
				string
				int
			}{strs[i], i})
			pop2 := heap.Remove(leng, leng.inheap[poporiind]).(struct { // find position in pop2 and delete it
				string
				int
			})
			assertTrue(t, strs[pop2.int] == popstr)
			assertTrue(t, pop2.string == popstr)
			heap.Push(leng, struct {
				string
				int
			}{strs[i], i})
		}
	}
	fmt.Println()

	drainlexi := make([]string, f)
	drainleng := make([]string, f)
	for i := 0; i < f; i++ {
		drainlexi[i] = heap.Pop(lexi).(struct {
			string
			int
		}).string
		drainleng[i] = heap.Pop(leng).(struct {
			string
			int
		}).string
	}
	for i := 0; i < f; i++ {
		fmt.Println(drainlexi[i])
	}
	fmt.Println()
	for i := 0; i < f; i++ {
		fmt.Println(drainleng[i])
	}
	fmt.Println()
	fmt.Println(checkStrsSameSet(drainlexi, drainleng))
	sort.Strings(strs)
	fmt.Println(checkStrsEqual(drainlexi, strs[n-f:]))
}

func newHeap(s []string, orilen int) *myheap {
	h := new(myheap)
	h.keys = make([]string, len(s)) // keys and oriind is associated with heap size
	copy(h.keys, s)
	h.oriind = make([]int, len(s))
	h.inheap = make([]int, orilen) // inheap is associated with the original array
	for i := range s {
		h.oriind[i] = i
		h.inheap[i] = i
	}
	return h
}

func newOHeap(s []string, orilen int) *oheap {
	return &oheap{newHeap(s, orilen)}
}

type myheap struct {
	keys   []string
	oriind []int
	inheap []int
}

type oheap struct {
	*myheap
}

// two heap according to different criteria
func (m *myheap) Less(i, j int) bool {
	return m.keys[i] < m.keys[j] // in lexicographical order
}

func (o *oheap) Less(i, j int) bool {
	return len(o.keys[i]) < len(o.keys[j]) // in string length order
}

func (m *myheap) Swap(i, j int) {
	m.keys[i], m.keys[j] = m.keys[j], m.keys[i]
	m.oriind[i], m.oriind[j] = m.oriind[j], m.oriind[i]
	m.inheap[m.oriind[i]] = i
	m.inheap[m.oriind[j]] = j
}

func (m *myheap) Len() int {
	return len(m.keys)
}

func (m *myheap) Push(x interface{}) {
	xx := x.(struct {
		string
		int
	})
	m.keys = append(m.keys, xx.string)
	m.oriind = append(m.oriind, xx.int)
	m.inheap[xx.int] = len(m.keys) - 1
}

func (m *myheap) Pop() interface{} {
	s := m.keys[len(m.keys)-1]
	i := m.oriind[len(m.oriind)-1]
	m.keys = m.keys[:len(m.keys)-1]
	m.oriind = m.oriind[:len(m.oriind)-1]
	return struct {
		string
		int
	}{s, i}
}

func getRandomStr(length int) string {
	runes := make([]rune, length)
	for i := 0; i < length; i++ {
		runes[i] = 'a' + rander.Int31n(26)
	}
	return string(runes)
}

func checkStrsSameSet(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	m := make(map[string]struct{})
	for i := range a {
		m[a[i]] = struct{}{}
	}
	for i := range b {
		if _, ok := m[b[i]]; !ok {
			return false
		}
	}
	return true
}

func checkStrsEqual(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

var rander = rand.New(rand.NewSource(time.Now().Unix()))

func assertTrue(t *testing.T, b bool) {
	if !b {
		t.FailNow()
	}
}
