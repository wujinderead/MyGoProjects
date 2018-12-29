package kthElem

import (
	"math/rand"
	"sort"
	"time"
)

var rander = rand.New(rand.NewSource(time.Now().Unix()))

type Array interface {
	sort.Interface
	Get(i int) interface{}
}

// method1:
// quick sort and return the kth elem
func FindKthElemQuickSort(arr Array, k int) interface{} {
	if k < 1 && k > arr.Len() {
		panic("k out of range")
	}
	sort.Sort(arr)
	return arr.Get(k - 1)
}

// two heap methods:
// method2: heapify the arr as a min heap, and pop k times
// method3. heapify the first k elements as a max heap; then for each the remaining (n-k) elements,
//    if it's smaller than heap's head, pop the head and push the value; finally, pop the heap

// method4:
// quick select:
// only sort the partition that contains the kth elem, use random pivot to overcome worst case (already sorted)
func FindKthElemQuickSelect(arr Array, k int) interface{} {
	if k < 1 && k > arr.Len() {
		panic("k out of range")
	}
	var partition = func(arr Array, low, high int) int {
		pivot := low + rander.Intn(high-low+1)
		arr.Swap(high, pivot)
		i := low
		for j := low; j <= high-1; j++ {
			if arr.Less(j, high) {
				arr.Swap(i, j)
				i++
			}
		}
		arr.Swap(i, high)
		return i
	}
	// first to declare the func variable to use it recursively
	var quickSelect func(Array, int, int, int) int
	quickSelect = func(arr Array, low, high, k int) int {
		pi := partition(arr, low, high)
		if pi-low == k-1 { // left part length is k-1, means that pivot is the kth
			return pi
		} else if pi-low < k-1 { // left part length less than k-1, search the right part
			return quickSelect(arr, pi+1, high, k-(pi-low+1))
		} else { // search the left part
			return quickSelect(arr, low, pi-1, k)
		}
	}
	return arr.Get(quickSelect(arr, 0, arr.Len()-1, k))
}

/*
method5: get the median of the whole array, so that to evenly part the array
         the complexity of this method is worst case O(n). implementation see:
         https://www.geeksforgeeks.org/kth-smallestlargest-element-unsorted-array-set-3-worst-case-linear-time/

kthSmallest(arr[0..n-1], k)
1) Divide arr[] into ⌈n/5⌉ groups where size of each group is 5
except possibly the last group which may have less than 5 elements.

2) Sort the above created ⌈n/5⌉ groups and find median
of all groups. Create an auxiliary array 'median[]' and store medians
of all ⌈n/5⌉ groups in this median array.

// Recursively call this method to find median of median[0..⌈n/5⌉-1]
3) medOfMed = kthSmallest(median[0..⌈n/5⌉-1], ⌈n/10⌉)

4) Partition arr[] around medOfMed and obtain its position.
pos = partition(arr, n, medOfMed)

5) If pos == k return medOfMed
6) If pos > k return kthSmallest(arr[l..pos-1], k)
7) If pos < k return kthSmallest(arr[pos+1..r], k-pos+l-1)
*/
