package find_kth

import (
	"fmt"
	"math/rand"
	"sort"
	"testing"
	"time"
)

type intSlice struct{ sort.IntSlice }

func (p intSlice) Get(i int) interface{} { return []int(p.IntSlice)[i] }

func TestFindKthElemQuickSort(t *testing.T) {
	rand.Seed(time.Now().Unix())
	var slicer intSlice
	slicer = intSlice{sort.IntSlice(rand.Perm(100))}
	fmt.Println(FindKthElemQuickSort(slicer, 1))
	slicer = intSlice{sort.IntSlice(rand.Perm(100))}
	fmt.Println(FindKthElemQuickSort(slicer, 2))
	slicer = intSlice{sort.IntSlice(rand.Perm(100))}
	fmt.Println(FindKthElemQuickSort(slicer, 33))
	slicer = intSlice{sort.IntSlice(rand.Perm(100))}
	fmt.Println(FindKthElemQuickSort(slicer, 100))
}

func TestFindKthElemQuickSelect(t *testing.T) {
	rand.Seed(time.Now().Unix())
	var slicer intSlice
	slicer = intSlice{sort.IntSlice(rand.Perm(100))}
	fmt.Println(FindKthElemQuickSelect(slicer, 1))
	slicer = intSlice{sort.IntSlice(rand.Perm(100))}
	fmt.Println(FindKthElemQuickSelect(slicer, 2))
	slicer = intSlice{sort.IntSlice(rand.Perm(100))}
	fmt.Println(FindKthElemQuickSelect(slicer, 33))
	slicer = intSlice{sort.IntSlice(rand.Perm(100))}
	fmt.Println(FindKthElemQuickSelect(slicer, 100))
}
