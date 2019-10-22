package leetcode

import "fmt"

// https://leetcode.com/problems/can-i-win/

// In the "100 game," two players take turns adding, to a running total, any integer from 1..10.
// The player who first causes the running total to reach or exceed 100 wins.
// What if we change the game so that players cannot re-use integers?
// For example, two players might take turns drawing from a common pool of numbers
// of 1..15 without replacement until they reach a total >= 100.
// Given an integer maxChoosableInteger and another integer desiredTotal,
// determine if the first player to move can force a win, assuming both players play optimally.
// You can always assume that maxChoosableInteger will not be larger than 20 and
// desiredTotal will not be larger than 300.
// Example
//   Input: maxChoosableInteger = 10  desiredTotal = 11
//   Output: false
// Explanation:
//   No matter which integer the first player choose, the first player will lose.
//   The first player can choose an integer from 1 up to 10.
//   If the first player choose 1, the second player can only choose integers from 2 up to 10.
//   The second player will win by choosing 10 and get a total = 11, which is >= desiredTotal.
//   Same with other integers chosen by the first player, the second player will always win.

func canIWin(maxChoosableInteger int, desiredTotal int) bool {
	if (1+maxChoosableInteger)*maxChoosableInteger/2 < desiredTotal {
		return false // sum less than target, both can't win
	}
	// use recursive function with memory, use hash map to store the result,
	// use integer as key (i.e. use bitmap to represent an integer set).
	// let c(iset, j) be true if i can win using integers in 'iset' to get target 'j'.
	// then if max(iset)>=j, then c(iset, j)=true; else,
	// c(iset, j) = not( c(iset-{i1}, j-i1) and c(iset-{i2}, j-i2) ....)
	canMap := make(map[int]bool)
	key := 0xfffff
	return canIWinMap(canMap, key>>uint(20-maxChoosableInteger), desiredTotal)
}

func canIWinMap(canMap map[int]bool, key int, target int) bool {
	if can, ok := canMap[key]; ok {
		return can
	}
	mask := 0x80000
	sig := 20
	for key&mask == 0 { // find the most significant bit,
		mask >>= 1 // e.g., for key=0b10101, sig=5, mask=0b10000
		sig--
	}
	if sig >= target {
		canMap[key] = true // cache result in map
		return true
	}
	for i := sig; i > 0; i-- {
		if key&mask > 0 {
			if !canIWinMap(canMap, key&(^mask), target-i) {
				canMap[key] = true
				return true
			}
		}
		mask >>= 1
	}
	canMap[key] = false
	return false
}

func main() {
	fmt.Println(canIWin(4, 8))
	fmt.Println(canIWin(4, 10))
	fmt.Println(canIWin(4, 4))
	fmt.Println(canIWin(4, 5))
	fmt.Println(canIWin(5, 10))
	fmt.Println(canIWin(20, 300))
	fmt.Println(canIWin(20, 200))
}
