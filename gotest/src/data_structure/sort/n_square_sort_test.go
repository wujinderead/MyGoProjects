package sort

import (
	"sort"
	"testing"
	"math/rand"
	"fmt"
)

func TestSelectionSort(t *testing.T) {
	num := 20
	permed := rand.Perm(num)
	ref := make([]int, num)
	copy(ref, permed)
	fmt.Println(permed)
	fmt.Println(ref)
	SelectionSort(sort.IntSlice(permed))
	fmt.Println(permed)
	sort.Ints(ref)
	fmt.Println(ref)
}

func TestStringSelectionSorter(t *testing.T) {
	sorter := newStringQuickSorter("摄影语言爱好者的学习家园。")
	SelectionSort(sorter)
	fmt.Println(sorter)
}

func TestInsertionSort(t *testing.T) {
	num := 20
	permed := rand.Perm(num)
	ref := make([]int, num)
	copy(ref, permed)
	fmt.Println(permed)
	fmt.Println(ref)
	InsertionSort(sort.IntSlice(permed))
	fmt.Println(permed)
	sort.Ints(ref)
	fmt.Println(ref)
}

func TestStringInsertionSorter(t *testing.T) {
	sorter := newStringQuickSorter("摄影语言爱好者的学习家园。")
	InsertionSort(sorter)
	fmt.Println(sorter)
}

func TestBubbleSort(t *testing.T) {
	num := 20
	permed := rand.Perm(num)
	ref := make([]int, num)
	copy(ref, permed)
	fmt.Println(permed)
	fmt.Println(ref)
	BubbleSort(sort.IntSlice(permed))
	fmt.Println(permed)
	sort.Ints(ref)
	fmt.Println(ref)
}

func TestStringBubbleSorter(t *testing.T) {
	sorter := newStringQuickSorter("摄影语言爱好者的学习家园。")
	BubbleSort(sorter)
	fmt.Println(sorter)
}