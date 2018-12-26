package sort

import (
	"sort"
)

type MergeSorter sort.Interface

func MergeSort(arr MergeSorter) {
	mergeSort(arr, 0, arr.Len()-1)
}

// split array to two parts, sort separated parts individually, then merge two sorted parts
func mergeSort(arr MergeSorter, low, high int) {
	if high-low<6 {  // selection sort
		for i:=low; i<=high; i++ {
			for j:=i+1; j<=high; j++ {
				if arr.Less(j, i) {
					arr.Swap(i, j)
				}
			}
		}
		return
	}
	mid := low+(high-low)/2;
	mergeSort(arr, low, mid)
	mergeSort(arr, mid+1, high)
	mergeTwoSorted(arr, low, mid, mid+1, high)
}

// merge two adjacent sorted array in place
func mergeTwoSorted(arr MergeSorter, low1, high1, low2, high2 int) {
	i, j := low1, low2
	for i<=high1 && j<=high2 {
		if arr.Less(i, j) {
			i++
		} else {
			for k:=j; k>i; k-- {
				arr.Swap(k, k-1)
			}
			i++
			j++
			high1++
		}
	}
}