package leetcode

import "fmt"

// https://leetcode.com/problems/dice-roll-simulation/

// A die simulator generates a random number from 1 to 6 for each roll.
// You introduced a constraint to the generator such that it cannot
// roll the number i more than rollMax[i] (1-indexed) consecutive times.
// Given an array of integers rollMax and an integer n,
// return the number of distinct sequences that can be obtained with exact n rolls.
// Two sequences are considered different if at least one element differs from each other.
// Since the answer may be too large, return it modulo 10^9 + 7.
// Example 1:
//   Input: n = 2, rollMax = [1,1,2,2,2,3]
//   Output: 34
//   Explanation: There will be 2 rolls of die, if there are no constraints on the die,
//   there are 6 * 6 = 36 possible combinations. In this case, looking at rollMax array,
//   the numbers 1 and 2 appear at most once consecutively, therefore sequences
//   (1,1) and (2,2) cannot occur, so the final answer is 36-2 = 34.
// Example 2:
//   Input: n = 2, rollMax = [1,1,1,1,1,1]
//   Output: 30
// Example 3:
//   Input: n = 3, rollMax = [1,1,1,2,2,3]
//   Output: 181
// Constraints:
//    1 <= n <= 5000
//    rollMax.length == 6
//    1 <= rollMax[i] <= 15

func dieSimulator(n int, rollMax []int) int {
	old := make([][]int, 6)
	new := make([][]int, 6)
	for i := 0; i < 6; i++ {
		old[i] = make([]int, rollMax[i])
		new[i] = make([]int, rollMax[i])
	}
	for i := 0; i < 6; i++ {
		old[i][0] = 1
	}

	// calculate the count
	for phase := 2; phase <= n; phase++ { // append number to reach n numbers
		// set single to 0
		for i := 0; i < 6; i++ {
			new[i][0] = 0
		}
		for i := 0; i < 6; i++ { // i-th die
			if rollMax[i] > 1 { // first count those can be consecutive
				for j := rollMax[i] - 1; j > 0; j-- {
					new[i][j] = old[i][j-1]
					for k := 0; k < 6; k++ {
						if k != i {
							new[k][0] += old[i][j]
						}
					}
				}
			}
			for k := 0; k < 6; k++ {
				if k != i {
					new[k][0] += old[i][0]
				}
			}
		}
		for i := 0; i < 6; i++ {
			new[i][0] %= 1000000007
		}
		old, new = new, old
	}

	// sum the result
	sum := 0
	for i := range rollMax {
		for j := range old[i] {
			sum += old[i][j]
		}
	}
	return sum % 1000000007
}

func main() {
	fmt.Println(dieSimulator(1, []int{1, 2, 3, 4, 5, 6}))
	fmt.Println(dieSimulator(2, []int{1, 1, 2, 2, 2, 3}))
	fmt.Println(dieSimulator(2, []int{1, 1, 1, 1, 1, 1}))
	fmt.Println(dieSimulator(200, []int{3, 4, 5, 7, 8, 9}))
}
