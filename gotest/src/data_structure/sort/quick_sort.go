package sort

import "sort"

type QuickSorter sort.Interface

func QuickSort(arr QuickSorter) {
	quickSort(arr, 0, arr.Len()-1)
}

func quickSort(arr QuickSorter, low, high int) {
	if low < high {
		pi := partition(arr, low, high)
		quickSort(arr, low, pi-1)
		quickSort(arr, pi+1, high)
	}
}

func partition(arr QuickSorter, low, high int) int {
	n := low-1
	for i:=low; i<high; i++ {
		if arr.Less(i, high) {
			n++
			arr.Swap(n, i)
		}
	}
	arr.Swap(n+1, high)
	return n+1
}