package main

import (
	"fmt"
	"sort"
)

// https://leetcode.com/problems/maximum-height-by-stacking-cuboids/

// Given n cuboids where the dimensions of the ith cuboid is cuboids[i] =
// [widthi, lengthi, heighti] (0-indexed). Choose a subset of cuboids and place
// them on each other.
// You can place cuboid i on cuboid j if widthi <= widthj and lengthi <= lengthj
// and heighti <= heightj. You can rearrange any cuboid's dimensions by rotating
// it to put it on another cuboid.
// Return the maximum height of the stacked cuboids.
// Example 1:
//   Input: cuboids = [[50,45,20],[95,37,53],[45,23,12]]
//   Output: 190
//   Explanation:
//     Cuboid 1 is placed on the bottom with the 53x37 side facing down with height 95.
//     Cuboid 0 is placed next with the 45x20 side facing down with height 50.
//     Cuboid 2 is placed next with the 23x12 side facing down with height 45.
//     The total height is 95 + 50 + 45 = 190.
// Example 2:
//   Input: cuboids = [[38,25,45],[76,35,3]]
//   Output: 76
//   Explanation:
//     You can't place any of the cuboids on the other.
//     We choose cuboid 1 and rotate it so that the 35x3 side is facing down and its height is 76.
// Example 3:
//   Input: cuboids = [[7,11,17],[7,17,11],[11,7,17],[11,17,7],[17,7,11],[17,11,7]]
//   Output: 102
//   Explanation:
//     After rearranging the cuboids, you can see that all cuboids have the same dimension.
//     You can place the 11x7 side down on all cuboids so their heights are 17.
//     The maximum height of stacked cuboids is 6 * 17 = 102.
// Constraints:
//   n == cuboids.length
//   1 <= n <= 100
//   1 <= widthi, lengthi, heighti <= 100

func maxHeight(cuboids [][]int) int {
	// sort 3 dimensions of a cuboid
	for i := range cuboids {
		if cuboids[i][0] > cuboids[i][1] {
			cuboids[i][0], cuboids[i][1] = cuboids[i][1], cuboids[i][0]
		}
		if cuboids[i][0] > cuboids[i][2] {
			cuboids[i][0], cuboids[i][2] = cuboids[i][2], cuboids[i][0]
		}
		if cuboids[i][1] > cuboids[i][2] {
			cuboids[i][1], cuboids[i][2] = cuboids[i][2], cuboids[i][1]
		}
	}
	// sort cuboids by 3 dimensions
	sort.Slice(cuboids, func(i, j int) bool {
		if cuboids[i][0] != cuboids[j][0] {
			return cuboids[i][0] < cuboids[j][0]
		}
		if cuboids[i][1] != cuboids[j][1] {
			return cuboids[i][1] < cuboids[j][1]
		}
		return cuboids[i][2] < cuboids[j][2]
	})

	// dp: let h[i] be the max height by cuboids[0..i] with cuboids[i] included
	h := make([]int, len(cuboids))
	max := 0
	for i := 0; i < len(cuboids); i++ {
		h[i] = cuboids[i][2]
		for j := 0; j < i; j++ {
			// if cuboids[i] can contain cuboids[j], check if we can get a higher height
			if cuboids[j][0] <= cuboids[i][0] && cuboids[j][1] <= cuboids[i][1] && cuboids[j][2] <= cuboids[i][2] {
				if h[j]+cuboids[i][2] > h[i] {
					h[i] = h[j] + cuboids[i][2]
				}
			}
		}
		if h[i] > max {
			max = h[i]
		}
	}
	return max
}

func main() {
	for _, v := range []struct {
		cuboids [][]int
		ans     int
	}{
		{[][]int{{1, 2, 5}, {1, 4, 6}, {2, 1, 9}, {2, 3, 6}, {3, 1, 7}, {3, 2, 5}}, 16},
		{[][]int{{50, 45, 20}, {95, 37, 53}, {45, 23, 12}}, 190},
		{[][]int{{38, 25, 45}, {76, 35, 3}}, 76},
		{[][]int{{7, 11, 17}, {7, 17, 11}, {11, 7, 17}, {11, 17, 7}, {17, 7, 11}, {17, 11, 7}}, 102},
	} {
		fmt.Println(maxHeight(v.cuboids), v.ans)
	}
}
