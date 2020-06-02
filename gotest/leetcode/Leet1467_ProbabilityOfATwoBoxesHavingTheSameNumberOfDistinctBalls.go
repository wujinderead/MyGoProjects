package main

import "fmt"

// https://leetcode.com/problems/probability-of-a-two-boxes-having-the-same-number-of-distinct-balls/

// Given 2n balls of k distinct colors. You will be given an integer array balls 
// of size k where balls[i] is the number of balls of color i. 
// All the balls will be shuffled uniformly at random, then we will distribute the 
// first n balls to the first box and the remaining n balls to the other box 
// (Please read the explanation of the second example carefully). 
// Please note that the two boxes are considered different. For example, if we have 
// two balls of colors a and b, and two boxes [] and (), then the distribution 
// [a] (b) is considered different than the distribution [b] (a) (Please read the 
// explanation of the first example carefully). 
// We want to calculate the probability that the two boxes have the same number 
// of distinct balls. 
// Example 1: 
//   Input: balls = [1,1]
//   Output: 1.00000
//   Explanation: Only 2 ways to divide the balls equally:
//     - A ball of color 1 to box 1 and a ball of color 2 to box 2
//     - A ball of color 2 to box 1 and a ball of color 1 to box 2
//     In both ways, the number of distinct colors in each box is equal. The probability is 2/2 = 1
// Example 2: 
//   Input: balls = [2,1,1]
//   Output: 0.66667
//   Explanation: We have the set of balls [1, 1, 2, 3]
//     This set of balls will be shuffled randomly and we may have one of the 12 distinct shuffles 
//     with equale probability (i.e. 1/12):
//       [1,1 / 2,3], [1,1 / 3,2], [1,2 / 1,3], [1,2 / 3,1], [1,3 / 1,2], [1,3 / 2,1], 
//       [2,1 / 1,3], [2,1 / 3,1], [2,3 / 1,1], [3,1 / 1,2], [3,1 / 2,1], [3,2 / 1,1]
//     After that we add the first two balls to the first box and the second two balls to the second box.
//     We can see that 8 of these 12 possible random distributions have the same number of distinct colors 
//     of balls in each box. Probability is 8/12 = 0.66667
// Example 3: 
//   Input: balls = [1,2,1,2]
//   Output: 0.60000
//   Explanation: The set of balls is [1, 2, 2, 3, 4, 4]. It is hard to display all
//     the 180 possible random shuffles of this set but it is easy to check that 108 
//     of them will have the same number of distinct colors in each box.
//     Probability = 108 / 180 = 0.6
// Example 4:  
//   Input: balls = [3,2,1]
//   Output: 0.30000
//   Explanation: The set of balls is [1, 1, 1, 2, 2, 3]. It is hard to display all
//     the 60 possible random shuffles of this set but it is easy to check that 18 of 
//     them will have the same number of distinct colors in each box.
//     Probability = 18 / 60 = 0.3
// Example 5: 
//   Input: balls = [6,6,6,6,6,6]
//   Output: 0.90327
// Constraints:  
//   1 <= balls.length <= 8 
//   1 <= balls[i] <= 6 
//   sum(balls) is even. 
//   Answers within 10^-5 of the actual value will be accepted as correct. 

// we can use backtracking to find all valid partitions. the backtracking need π(balls) time, at most 6^8.
// e.g., for balls [5,5,6], we have a valid patition as [2,4,2 | 3,1,4]. 
// if count non-duplicated permutions, we have: 
//    all 16!/5!5!6! permutations;
//    for this partition, 8!/2!4!2! * 8!/3!4! permutions.
// if deem the ball as all distinct, we have:
//    to select 8 elements from all 16 to the left part, the number of way is C(16, 8);
//    to select 2 from all 5 of first-type to the left, we have C(5,2); so the total is C(5,2)*C(5,4)*C(6,3)
// actually the both is equal.
//                           5!    5!    6!             8!         8!
// C(5,2)*C(5,4)*C(6,2)     --- * --- * ----          -----   *  -----   
//                          3!2!   4!   4!2!          2!4!2!      3!4! 
// -------------------- = --------------------- = -------------------------
//                               16!                        16!
//     C(16,8)                  -----                     -------
//                               8!8!                      5!5!6!
// the repeated is easy to count, as the most is C(6,3)^8=12^8, while 24!/6!6!6!6! will overflow.
func getProbability(balls []int) float64 {
	n := 0
	for i := range balls {
		n += balls[i]
	}
	n = n/2
	curball := make([]int, len(balls))
	all := 0
	patition(n, 0, 0, balls, curball, &all)

	// the final answer is all/(2n!/n!n!) = all*(n/2n)*(n-1/2n-1)*...(1/n+1)
	ret := float64(all)
	for i:=1; i<=n; i++ {
		ret *= float64(i)/float64(n+i)
	}
    return ret
}

func patition(n, ind, curnum int, balls, curball []int, all *int) {
	if curnum==n {     
		if check(balls, curball) {      // if it's a valid partition, add the number
			*all += getfraction(balls, curball) 
		}
		return
	}
	if ind==len(balls) {
		return
	}
	for i:=0; i<=balls[ind]; i++ {
		if curnum+i>n {
			break
		}
		curball[ind] = i
		patition(n, ind+1, curnum+i, balls, curball, all)
	}
	curball[ind] = 0
}

func getfraction(balls, curball []int) int {
	frac := []int{1,1,2,6,24,120,720}
	s := 1
	for i := range balls {
		s *= frac[balls[i]]/frac[curball[i]]/frac[balls[i]-curball[i]]  // C(balls[i], curballs[i])
	}
	return s
}

func check(balls, curball []int) bool {
	a, b := 0, 0 
	for i:=range balls {
		if curball[i]>0 {
			a++
		}
		if balls[i]-curball[i]>0 {
			b++
		}
	}
	return a==b
}

func main() {
	fmt.Println(getProbability([]int{1,1}))
	fmt.Println(getProbability([]int{2,1,1}))
	fmt.Println(getProbability([]int{1,2,1,2}))
	fmt.Println(getProbability([]int{3,2,1}))
	fmt.Println(getProbability([]int{6,6,6,6,6,6}))
}
