package sort

import "sort"

type SelectionSorter sort.Interface
type InsertionSorter sort.Interface
type BubbleSorter sort.Interface

// on each loop, select the littlest value and swap to the front
func SelectionSort(arr SelectionSorter) {
	selectionSort(arr, 0, arr.Len()-1)
}

func selectionSort(arr SelectionSorter, low, high int) {
	for i := low; i <= high; i++ {
		lowest := i
		for j := i + 1; j <= high; j++ {
			if arr.Less(j, lowest) {
				lowest = j
			}
		}
		arr.Swap(lowest, i)
	}
}

// iterate the array, find a place to insert to previous values, which are already ordered
func InsertionSort(arr InsertionSorter) {
	insertionSort(arr, 0, arr.Len()-1)
}

func insertionSort(arr InsertionSorter, low, high int) {
	for i := low + 1; i <= high; i++ {
		for j := i - 1; j >= low && arr.Less(j+1, j); j-- {
			arr.Swap(j+1, j)
		}
	}
}

// one each loop, swap adjacent values if they are wrong-ordered,
// by which it can 'bubble' the greatest value to end
func BubbleSort(arr BubbleSorter) {
	bubbleSort(arr, 0, arr.Len()-1)
}

func bubbleSort(arr BubbleSorter, low, high int) {
	for i := low; i <= high; i++ {
		for j := low; j < high-i+low; j++ {
			if arr.Less(j+1, j) {
				arr.Swap(j+1, j)
			}
		}
	}
}
