package sort

import "sort"

type SelectionSorter sort.Interface

// on each loop, select the littlest value and swap to the front
func SelectionSort(arr QuickSorter) {
	selectionSort(arr, 0, arr.Len()-1)
}

func selectionSort(arr QuickSorter, low, high int) {
	for i:=low; i<=high; i++ {
		lowest := i
		for j:=i+1; j<=high; j++ {
			if arr.Less(j, lowest) {
				lowest = j
			}
		}
		arr.Swap(lowest, i)
	}
}

// iterate the array, find a place to insert to previous values, which are already ordered
func InsertionSort(arr QuickSorter) {
	insertionSort(arr, 0, arr.Len()-1)
}

func insertionSort(arr QuickSorter, low, high int) {
	for i:=low+1; i<=high; i++ {
		for j:=i-1; j>=low && arr.Less(j+1, j); j-- {
			arr.Swap(j+1, j)
		}
	}
}

// one each loop, swap adjacent values if they are wrong-ordered,
// by which it can 'bubble' the greatest value to end
func BubbleSort(arr QuickSorter) {
	bubbleSort(arr, 0, arr.Len()-1)
}

func bubbleSort(arr QuickSorter, low, high int) {
	// TODO implements
}
