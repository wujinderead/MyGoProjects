package sort

import (
	"sort"
)

type ShellSorter sort.Interface

var gaps = []int{1, 4, 10, 23, 57, 132, 301, 701}

func ShellSort(arr ShellSorter) {
	shellSort(arr, 0, arr.Len()-1)
}

// insertion sort with gap
func shellSort(arr ShellSorter, low, high int) {
	for gi := len(gaps) - 1; gi >= 0; gi-- {
		gap := gaps[gi]
		for i := low + gap; i <= high; i++ {
			for i-gap >= low && arr.Less(i, i-gap) {
				arr.Swap(i, i-gap)
				i -= gap
			}
		}
	}
}
