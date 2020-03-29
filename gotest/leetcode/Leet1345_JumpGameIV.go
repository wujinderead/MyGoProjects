package main

import (
	"container/list"
	"fmt"
)

// Given an array of integers arr, you are initially positioned at the first
// index of the array.
// In one step you can jump from index i to index:
//   i + 1 where: i + 1 < arr.length.
//   i - 1 where: i - 1 >= 0.
//   j where: arr[i] == arr[j] and i != j.
// Return the minimum number of steps to reach the last index of the array.
// Notice that you can not jump outside of the array at any time.
// Example 1:
//   Input: arr = [100,-23,-23,404,100,23,23,23,3,404]
//   Output: 3
//   Explanation: You need three jumps from index 0 --> 4 --> 3 --> 9. Note that
//     index 9 is the last index of the array.
// Example 2:
//   Input: arr = [7]
//   Output: 0
//   Explanation: Start index is the last index. You don't need to jump.
// Example 3:
//   Input: arr = [7,6,9,6,9,6,9,7]
//   Output: 1
//   Explanation: You can jump directly from index 0 to index 7 which is last index
//     of the array.
// Example 4:
//   Input: arr = [6,1,9]
//   Output: 2
// Example 5:
//   Input: arr = [11,22,7,7,7,7,7,7,7,22,13]
//   Output: 3
// Constraints:
//   1 <= arr.length <= 5 * 10^4
//   -10^8 <= arr[i] <= 10^8

func minJumps(arr []int) int {
	if len(arr) == 1 {
		return 0
	}
	mapp := map[int][]int{arr[0]: {0}}
	for i := 1; i < len(arr); i++ {
		if i == len(arr)-1 || arr[i] != arr[i-1] || arr[i] != arr[i+1] { // ignore intermediate
			poss, ok := mapp[arr[i]]
			if ok {
				mapp[arr[i]] = append(poss, i)
			} else {
				mapp[arr[i]] = []int{i}
			}
		}
	}
	visited := make([]bool, len(arr))
	queue := list.New()
	queue.PushBack([2]int{0, 0}) // (index, step) pair
	visited[0] = true
	for queue.Len() > 0 {
		pair := queue.Remove(queue.Front()).([2]int)
		curind := pair[0]
		curstep := pair[1]
		curval := arr[curind]
		if curind-1 > 0 && !visited[curind-1] {
			visited[curind-1] = true
			queue.PushBack([2]int{curind - 1, curstep + 1})
		}
		if curind+1 < len(arr) && !visited[curind+1] {
			if curind+1 == len(arr)-1 {
				return curstep + 1
			}
			visited[curind+1] = true
			queue.PushBack([2]int{curind + 1, curstep + 1})
		}
		pos := mapp[curval]
		for _, v := range pos {
			if !visited[v] {
				if v == len(arr)-1 {
					return curstep + 1
				}
				visited[v] = true
				queue.PushBack([2]int{v, curstep + 1})
			}
		}

	}
	return len(arr) - 1
}

func main() {
	fmt.Println(minJumps([]int{100, -23, -23, 404, 100, 23, 23, 23, 3, 404}))
	fmt.Println(minJumps([]int{7}))
	fmt.Println(minJumps([]int{7, 6, 9, 6, 9, 6, 9, 7}))
	fmt.Println(minJumps([]int{6, 1, 9}))
	fmt.Println(minJumps([]int{11, 22, 7, 7, 7, 7, 7, 7, 7, 22, 13}))
}
