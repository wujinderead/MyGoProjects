package sort

import (
	"fmt"
	"math/rand"
	"sort"
	"testing"
	"time"
)

func TestMergeSort(t *testing.T) {
	rand.Seed(time.Now().Unix())
	for i:=0; i<10; i++ {
		num := rand.Int31n(100)
		permed := rand.Perm(int(num))
		ref := make([]int, num)
		copy(ref, permed)
		//fmt.Println(permed)
		//fmt.Println(ref)
		MergeSort(sort.IntSlice(permed))
		//fmt.Println(permed)
		sort.Ints(ref)
		//fmt.Println(ref)
		fmt.Println(intSlieceEqual(permed, ref))
	}
}

func TestStringMergeSorter(t *testing.T) {
	sorter := newStringSorter("摄影语言爱好者的学习家园。")
	MergeSort(sorter)
	fmt.Println(sorter)
}

// test boundary values, i.e., arrays with small length
func TestBoundaryMergeSort(t *testing.T) {
	zero := []int{}
	one := []int{3}
	two := []int{3, 2}
	two_s := []int{2, 3}
	three := []int{2, 1, 3}
	arrs := [][]int{zero, one, two, two_s, three}
	for i := range arrs {
		MergeSort(sort.IntSlice(arrs[i]))
		fmt.Println(arrs[i])
	}
}