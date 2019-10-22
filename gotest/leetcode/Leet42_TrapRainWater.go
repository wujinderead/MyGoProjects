package leetcode

import "fmt"

func trapDp(height []int) int {
	if len(height) < 1 {
		return 0
	}
	trapped := 0
	leftmax := make([]int, len(height))
	rightmax := make([]int, len(height))
	leftmax[0] = height[0]
	for i := 1; i < len(height); i++ {
		leftmax[i] = max(height[i], leftmax[i-1])
	}
	rightmax[len(height)-1] = height[len(height)-1]
	for i := len(height) - 2; i >= 0; i-- {
		rightmax[i] = max(height[i], rightmax[i+1])
	}
	for i := range height {
		trapped += min(leftmax[i], rightmax[i]) - height[i]
	}
	return trapped
}

// actually do not need to store left max and right max in array,
// because they are only needed for current position
func trapTwoPointers(height []int) int {
	if len(height) < 1 {
		return 0
	}
	trapped := 0
	i := 0
	j := len(height) - 1
	leftmax := height[i]
	rightmax := height[j]
	for i < j {
		if height[i] < height[j] { // when left height smaller, trapped depends on left
			// when height[i]>leftmax, cannot trap water, just update leftmax
			leftmax = max(leftmax, height[i])
			trapped += leftmax - height[i]
			i++
		} else {
			// when height[j]>rightmax, cannot trap water, just update rightmax
			rightmax = max(rightmax, height[j])
			trapped += rightmax - height[j]
			j--
		}
	}
	return trapped
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	arr := []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 0, 2, 1}
	fmt.Println(trapDp(arr))
	fmt.Println(trapTwoPointers(arr))
	arr = []int{}
	fmt.Println(trapDp(arr))
	fmt.Println(trapTwoPointers(arr))
}
