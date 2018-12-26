package sort

import (
	"sort"
)

type MergeSorter sort.Interface

func MergeSort(arr MergeSorter) {
	mergeSort(arr, 0, arr.Len()-1)
}

// merge two sorted list
func mergeSort(arr MergeSorter, low, high int) {
	// TODO merge sort
}