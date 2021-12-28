package main

import "fmt"

// https://leetcode.com/problems/intervals-between-identical-elements/

// You are given a 0-indexed array of n integers arr.
//
// The interval between two elements in arr is defined as the absolute difference between their indices.
// More formally, the interval between arr[i] and arr[j] is |i - j|.
// Return an array intervals of length n where intervals[i] is the sum of intervals between arr[i] and
// each element in arr with the same value as arr[i].
// Note: |x| is the absolute value of x.
// Example 1:
//   Input: arr = [2,1,3,1,2,3,3]
//   Output: [4,2,7,2,4,4,5]
//   Explanation:
//     - Index 0: Another 2 is found at index 4. |0 - 4| = 4
//     - Index 1: Another 1 is found at index 3. |1 - 3| = 2
//     - Index 2: Two more 3s are found at indices 5 and 6. |2 - 5| + |2 - 6| = 7
//     - Index 3: Another 1 is found at index 1. |3 - 1| = 2
//     - Index 4: Another 2 is found at index 0. |4 - 0| = 4
//     - Index 5: Two more 3s are found at indices 2 and 6. |5 - 2| + |5 - 6| = 4
//     - Index 6: Two more 3s are found at indices 2 and 5. |6 - 2| + |6 - 5| = 5
// Example 2:
//   Input: arr = [10,5,10,10]
//   Output: [5,0,3,4]
//   Explanation:
//     - Index 0: Two more 10s are found at indices 2 and 3. |0 - 2| + |0 - 3| = 5
//     - Index 1: There is only one 5 in the array, so its sum of intervals to identical elements is 0.
//     - Index 2: Two more 10s are found at indices 0 and 3. |2 - 0| + |2 - 3| = 3
//     - Index 3: Two more 10s are found at indices 0 and 2. |3 - 0| + |3 - 2| = 4
// Constraints:
//   n == arr.length
//   1 <= n <= 10⁵
//   1 <= arr[i] <= 10⁵

// for an array with length 4, say the distance between items are a, b, c,
// x0---a---x1---b---x2---c---x3
// then dist(x0) = 0 +   a + a+b + a+b+c
//      dist(x1) = a +   0 +  b  + b+c      =  dist(x0)+a-3a
//      dist(x2) = a+b + b +  0  + c        =  dist(x0)+2b-2b
// ...
// so dist(x[i])=dist(x[i-1]) + i*(x[i]-x[i-1]) - (len(x)-i)*(x[i]-x[i-1])
func getDistances(arr []int) []int64 {
	mapp := make(map[int][]int)
	for i, v := range arr {
		mapp[v] = append(mapp[v], i)
	}
	ans := make([]int64, len(arr))
	for _, v := range mapp {
		sum := int64(0)
		for i := 1; i < len(v); i++ {
			sum += int64(v[i] - v[0])
		}
		ans[v[0]] = sum
		for i := 1; i < len(v); i++ {
			sum = sum + int64(i+i-len(v))*int64(v[i]-v[i-1])
			ans[v[i]] = sum
		}
	}
	return ans
}

func main() {
	for _, v := range []struct {
		arr []int
		ans []int64
	}{
		{[]int{2, 1, 3, 1, 2, 3, 3}, []int64{4, 2, 7, 2, 4, 4, 5}},
		{[]int{10, 5, 10, 10}, []int64{5, 0, 3, 4}},
	} {
		fmt.Println(getDistances(v.arr), v.ans)
	}
}
