package sort

func QuickSort(arr []int, low, high int) {
	if low < high {
		pi := partition(arr, low, high)
		QuickSort(arr, low, pi-1)
		QuickSort(arr, pi+1, high)
	}
}

func partition(arr []int, low, high int) int {
	pivot := arr[high]
	n := low-1
	for i:=low; i<high; i++ {
		if arr[i] <= pivot {
			n++
			arr[n], arr[i] = arr[i], arr[n]
		}
	}
	arr[n+1], arr[high] = arr[high], arr[n+1]
	return n+1
}