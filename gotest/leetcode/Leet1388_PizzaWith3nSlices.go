package main

import (
	"fmt"
	"math/rand"
)

// https://leetcode.com/problems/pizza-with-3n-slices/

// There is a pizza with 3n slices of varying size, you and your friends will take
// slices of pizza as follows:
//   You will pick any pizza slice.
//   Your friend Alice will pick next slice in anti clockwise direction of your pick.
//   Your friend Bob will pick next slice in clockwise direction of your pick.
//   Repeat until there are no more slices of pizzas.
// Sizes of Pizza slices is represented by circular array slices in clockwise direction.
// Return the maximum possible sum of slice sizes which you can have.
// Example 1:
//   Input: slices = [1,2,3,4,5,6]
//   Output: 10
//   Explanation: Pick pizza slice of size 4, Alice and Bob will pick slices with size
//     3 and 5 respectively. Then Pick slices with size 6, finally Alice and Bob will
//     pick slice of size 2 and 1 respectively. Total = 4 + 6.
// Example 2:
//   Input: slices = [8,9,8,6,1,1]
//   Output: 16
//   Explanation: Pick pizza slice of size 8 in each turn. If you pick slice with size 9
//     your partners will pick slices of size 8.
// Example 3:
//   Input: slices = [4,1,2,5,8,3,1,9,7]
//   Output: 21
// Example 4:
//   Input: slices = [3,1,2]
//   Output: 3
// Constraints:
//   1 <= slices.length <= 500
//   slices.length % 3 == 0
//   1 <= slices[i] <= 1000

// IMPROVEMENT: to avoid circle, we need to calculate F(2...n-2, k-1) and F(1...n-1, k)
// the general problem is to select k non-adjacent numbers from F(arr[x...y], k)
// we can make a function of it, like F(arr []int, k int).

func maxSizeSlices(s []int) int {
	if len(s)==3 {
		return max(s[0], max(s[1], s[2]))
	}
    // the problem is equal to select N/3 non-adjacent numbers from N numbers with max sum.
	// denote s = [0, 1, 2, ..., n-1], target=n/3.
	// let F(0...n-1, k) be the max sum of selected k non-adjacent numbers from s[0...n-1]
	// then, if select s[0], we got sum: s[0] + F(2...n-2, k-1)
	// if not select s[0], we continue to search: F(1...n-1, k)
	// for F(2...n-2, k-1), we has candidates: s[2]+F(4...n-2, k-2) and F(3...n-2, k-1)
	// for F(1...n-1, k),   we has candidates: s[1]+F(3...n-1, k-1) and F(2...n-1, k)
	// we can observe that in F(x...n-1, k) and F(x...n-2, k), only x and k are varied.
	// base case F(xxx, 0)=0, F(n-1...n-1, 1) = s[n-1]. return F(0...n-1, k)
	n := len(s)
	k := n/3
	oldn1, newn1 := make([]int, n), make([]int, n)   // F(xxx, 0) = 0
	oldn2, newn2 := make([]int, n), make([]int, n)
	for i:=1; i<=k-1; i++ {  // from F(2...n-2, k-1) to F(n-2...n-2, 1)
		for j:=n-2; j>=2; j-- {   // to calculate F(j...n-2, i)
			if n-2-j+1 < 2*i-1 {  // we need 1 nums to select 1 num, 3 nums to select 2 nums ...
				newn2[j] = 0
				continue
			}
			// F(j...n-2, i) = max(s[j]+F(j+2...n-2, i-1), F(j+1...n-2, i))
			tmp1 := 0
			if j+2<=n-2 {            // F(x...n-2, i)=0 if x>n-2
				tmp1 = oldn2[j+2]
			}
			tmp2 := 0
			if j+1<=n-2 {
				tmp2 = newn2[j+1]
			}
			newn2[j] = max(s[j]+tmp1, tmp2)
			// fmt.Printf("F(%v, %d)=%d\n", s[j: n-1], i, newn2[j])
		}
		oldn2, newn2 = newn2, oldn2
	}
	for i:=1; i<=k; i++ {  // from F(1...n-1, k) to F(n-1...n-1, 1)
		for j:=n-1; j>=1; j-- {   // to calculate F(j...n-1, i)
			if n-1-j+1 < 2*i-1 {  // we need 1 nums to select 1 num, 3 nums to select 2 nums ...
				newn1[j] = 0
				continue
			}
			// F(j...n-1, i) = max(s[j]+F(j+2...n-1, i-1), F(j+1...n-1, i))
			tmp1 := 0
			if j+2<=n-1 {            // F(x...n-1, i)=0 if x>n-1
				tmp1 = oldn1[j+2]
			}
			tmp2 := 0
			if j+1<=n-1 {
				tmp2 = newn1[j+1]
			}
			newn1[j] = max(s[j]+tmp1, tmp2)
			// fmt.Printf("F(%v, %d)=%d\n", s[j: n], i, newn1[j])
		}
		oldn1, newn1 = newn1, oldn1
	}
	// after we got F(2...n-2, k-1) and F(1...n-1, k)
	// the answer is the max of: s[0] + F(2...n-2, k-1),  F(1...n-1, k)
	// NOTE: why calculating separately? because it avoids circle
	return max(s[0]+oldn2[2], oldn1[1])
}

func max(a, b int) int {
	if a>b {
		return a
	}
	return b
}

func main() {
	fmt.Println(maxSizeSlices([]int{1,2,3,4,5,6}))
	fmt.Println(maxSizeSlices([]int{8,9,8,6,1,1}))
	fmt.Println(maxSizeSlices([]int{4,1,2,5,8,3,1,9,7}))
	fmt.Println(maxSizeSlices([]int{3,1,2}))
	arr := rand.Perm(20)[:12]
	fmt.Println(arr)
	fmt.Println(maxSizeSlices(arr))
	fmt.Println(maxSizeSlices([]int{9,5,1,7,8,4,4,5,5,8,7,7}))
}