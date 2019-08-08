package ostest

import (
	"fmt"
	"container/list"
	"testing"
)

type cacheReplacer interface {
	Get(interface{}) uintptr
	Put(interface{}, uintptr)
	Remove(interface{})
}

type fifoPage struct {
	key interface{}
	addr uintptr
	elem *list.Element
}

type fifo struct {
	mapper map[interface{}]*fifoPage
	queue *list.List
	size int
}

func newFifo(size int) *fifo {
	if size<1 {
		panic("wrong cache size")
	}
	return &fifo{make(map[interface{}]*fifoPage), list.New(), size}
}

func (fifo *fifo) Get(key interface{}) uintptr {
	if page, ok:= fifo.mapper[key]; ok {
		return page.addr
	}
	return 0
}

func (fifo *fifo) Put(key interface{}, addr uintptr) {
	if page, ok:= fifo.mapper[key]; ok {  // if exist
		page.addr = addr
		return
	}
	// queue full, remove first in queue
	if fifo.queue.Len() == fifo.size {
		first := fifo.queue.Front()
		delete(fifo.mapper, fifo.queue.Remove(first).(*fifoPage).key)  // remove in map
	}
	page := &fifoPage{key, addr, nil}
	elem := fifo.queue.PushBack(page)
	page.elem = elem
	fifo.mapper[key] = page
}

func (fifo *fifo) Remove(key interface{}) {
	if page, ok:= fifo.mapper[key]; ok {
		elem := page.elem
		fifo.queue.Remove(elem)      // remove in queue
		delete(fifo.mapper, key)     // remove in map
	}
}

func TestFifoReplacer(t *testing.T) {
	cache := newFifo(4)
	display := func(fifo *fifo, str string) {
		fmt.Println(str)
		for k, v := range fifo.mapper {
			fmt.Println(k, ", ", v)
		}
		for elem := fifo.queue.Front(); elem != nil; elem = elem.Next() {
			fmt.Print(elem.Value.(*fifoPage).key, elem.Value.(*fifoPage).addr, ", ")
		}
		fmt.Println()
		fmt.Println()
	}
	display(cache, "empty")

	cache.Put("a", 1)
	cache.Put("b", 2)
	fmt.Println(cache.Get("a"))
	fmt.Println(cache.Get("c"))
	display(cache, "put ab")

	cache.Put("c", 3)
	cache.Put("d", 4)
	display(cache, "put cd")

	cache.Put("e", 5)
	fmt.Println(cache.Get("c"))
	display(cache, "put e")

	cache.Put("f", 6)
	cache.Put("c", 8)

	display(cache, "put fc")

	cache.Remove("d")
	display(cache, "remove d")

	cache.Remove("e")
	display(cache, "remove e")
}

type lruPage struct {
	key interface{}
	addr uintptr
	elem *list.Element
}

type lru struct {
	mapper map[interface{}]*lruPage
	queue *list.List
	size int
}

func newLru(size int) *lru {
	if size<1 {
		panic("wrong cache size")
	}
	return &lru{make(map[interface{}]*lruPage), list.New(), size}
}

func (lru *lru) Get(key interface{}) uintptr {
	if page, ok:= lru.mapper[key]; ok {
		// move page to front to make it most recent
		if page.elem != lru.queue.Front() {  // no need to remove when its already front
			lru.queue.Remove(page.elem)
			elem := lru.queue.PushFront(page)
			page.elem = elem
		}
		return page.addr
	}
	return 0
}

func (lru *lru) Put(key interface{}, addr uintptr) {
	if page, ok:= lru.mapper[key]; ok {      // if exist
		if page.elem != lru.queue.Front() {
			lru.queue.Remove(page.elem)
			elem := lru.queue.PushFront(page)
			page.elem = elem
		}
		page.addr = addr
		return
	}
	// queue full, remove last in queue, i.e. the least recently used page
	if lru.queue.Len() == lru.size {
		last := lru.queue.Back()
		delete(lru.mapper, lru.queue.Remove(last).(*lruPage).key)  // remove in map
	}
	// push to queue front
	page := &lruPage{key, addr, nil}
	elem := lru.queue.PushFront(page)
	page.elem = elem
	lru.mapper[key] = page
}

func (lru *lru) Remove(key interface{}) {
	if page, ok:= lru.mapper[key]; ok {
		elem := page.elem
		lru.queue.Remove(elem)      // remove in queue
		delete(lru.mapper, key)     // remove in map
	}
}

func TestLruReplacer(t *testing.T) {
	cache := newLru(4)
	display := func(lru *lru, str string) {
		fmt.Println(str)
		for k, v := range lru.mapper {
			fmt.Println(k, ", ", v)
		}
		for elem := lru.queue.Front(); elem != nil; elem = elem.Next() {
			fmt.Print(elem.Value.(*lruPage).key, elem.Value.(*lruPage).addr, ", ")
		}
		fmt.Println()
		fmt.Println()
	}
	display(cache, "empty")

	cache.Put("a", 1)
	cache.Put("b", 2)
	display(cache, "put ab")

	cache.Put("c", 3)
	cache.Put("d", 4)
	display(cache, "put cd")

	fmt.Println(cache.Get("b"))
	display(cache, "get b")

	cache.Put("e", 5)
	display(cache, "put e")

	fmt.Println(cache.Get("d"))
	display(cache, "get d")

	cache.Put("b", 7)
	display(cache, "put b")

	cache.Remove("e")
	display(cache, "remove e")
}

