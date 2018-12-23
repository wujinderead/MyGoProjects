package sort

import (
	"testing"
	"fmt"
	"math/rand"
	"sort"
)

func TestQuickSort(t *testing.T) {
	num := 20
	permed := rand.Perm(num)
	ref := make([]int, num)
	copy(ref, permed)
	fmt.Println(permed)
	fmt.Println(ref)
	QuickSort(permed, 0, len(permed)-1)
	fmt.Println(permed)
	sort.Ints(ref)
	fmt.Println(ref)
}

func TestPartition(t *testing.T) {
	arr := []int{10, 80, 30, 90, 40, 50, 70}
	pi := partition(arr, 0, len(arr)-1)
	fmt.Println(pi)
	fmt.Println(arr)
}