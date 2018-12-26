package sort

import (
	"fmt"
	"math/rand"
	"sort"
	"testing"
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
	sorter := newStringSorter("摄影语言爱好者的学习家园。")
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
	sorter := newStringSorter("摄影语言爱好者的学习家园。")
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
	sorter := newStringSorter("摄影语言爱好者的学习家园。")
	BubbleSort(sorter)
	fmt.Println(sorter)
}

// test boundary values, i.e., arrays with small length
func TestBoundaryNSquare(t *testing.T) {
	{
		zero := []int{}
		one := []int{3}
		two := []int{3, 2}
		two_s := []int{2, 3}
		three := []int{2, 1, 3}
		arrs := [][]int{zero, one, two, two_s, three}
		for i := range arrs {
			SelectionSort(sort.IntSlice(arrs[i]))
			fmt.Println(arrs[i])
		}
	}
	{
		zero := []int{}
		one := []int{3}
		two := []int{3, 2}
		two_s := []int{2, 3}
		three := []int{2, 1, 3}
		arrs := [][]int{zero, one, two, two_s, three}
		for i := range arrs {
			InsertionSort(sort.IntSlice(arrs[i]))
			fmt.Println(arrs[i])
		}
	}
	{
		zero := []int{}
		one := []int{3}
		two := []int{3, 2}
		two_s := []int{2, 3}
		three := []int{2, 1, 3}
		arrs := [][]int{zero, one, two, two_s, three}
		for i := range arrs {
			BubbleSort(sort.IntSlice(arrs[i]))
			fmt.Println(arrs[i])
		}
	}
}
