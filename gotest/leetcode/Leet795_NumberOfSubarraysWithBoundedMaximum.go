package leetcode

import "fmt"

func numSubarrayBoundedMax(A []int, L int, R int) int {
	var ans = 0
	var front = -1
	for end := front + 1; end < len(A); end++ {
		if A[end] > R {
			var prev = front
			for i := front + 1; i < end; i++ {
				if A[i] >= L {
					ans += (i - prev) * (end - i)
					prev = i
				}
			}
			front = end
		}
	}
	var prev = front
	for i := front + 1; i < len(A); i++ {
		if A[i] >= L {
			ans += (i - prev) * (len(A) - i)
			prev = i
		}
	}
	return ans
}

func main() {
	fmt.Println(numSubarrayBoundedMax([]int{4, 0, 1, 2, 1, 3, 0, 1, 2, 0, 4, 0, 1, 2, 1, 2}, 2, 3))
	fmt.Println(numSubarrayBoundedMax([]int{}, 2, 3))
	fmt.Println(numSubarrayBoundedMax([]int{0, 1, 2}, 2, 3))
	fmt.Println(numSubarrayBoundedMax([]int{2, 4}, 2, 3))
	fmt.Println(numSubarrayBoundedMax([]int{2}, 2, 3))
	fmt.Println(numSubarrayBoundedMax([]int{4, 2}, 2, 3))
}
