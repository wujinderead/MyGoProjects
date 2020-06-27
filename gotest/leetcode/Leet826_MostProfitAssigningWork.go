package main

import (
	"fmt"
	"sort"
)

// We have jobs: difficulty[i] is the difficulty of the ith job, and profit[i] is
// the profit of the ith job. Now we have some workers. worker[i] is the ability of 
// the ith worker, which means that this worker can only complete a job with difficulty 
// at most worker[i]. Every worker can be assigned at most one job, but one job can be 
// completed multiple times. For example, if 3 people attempt the same job that pays $1, 
// then the total profit will be $3. If a worker cannot complete any job, his profit is $0. 
// What is the most profit we can make? 
// Example 1: 
//   Input: difficulty = [2,4,6,8,10], profit = [10,20,30,40,50], worker = [4,5,6,7]
//   Output: 100 
//   Explanation: Workers are assigned jobs of difficulty [4,4,6,6] and 
//     they get profit of [20,20,30,30] seperately. 
// Notes: 
//   1 <= difficulty.length = profit.length <= 10000 
//   1 <= worker.length <= 10000 
//   difficulty[i], profit[i], worker[i] are in range [1, 10^5] 

func maxProfitAssignment(difficulty []int, profit []int, worker []int) int {
	t := tuple{difficulty, profit}
	sort.Sort(t)
	sort.Sort(sort.IntSlice(worker))
	s := -1
	allprofit := 0
	curprofit := 0
	for _, w := range worker {
		// find first s that w>=difficulty[s]
		for s+1<len(difficulty) && w>=difficulty[s+1] {
			s++
			if curprofit<profit[s] {   // the profit we have seen
				curprofit = profit[s]
			}
		} 
		allprofit += curprofit
	}
	return allprofit
}

type tuple struct {
	difficulty, profit []int
}

func (t tuple) Len() int {
	return len(t.difficulty)
}

func (t tuple) Less(i, j int) bool {
	return t.difficulty[i]<t.difficulty[j]
} 

func (t tuple) Swap(i, j int) {
	t.difficulty[i], t.difficulty[j] = t.difficulty[j], t.difficulty[i]
	t.profit[i], t.profit[j] = t.profit[j], t.profit[i]	
}

func main() {
	fmt.Println(maxProfitAssignment([]int{2,4,6,8,10}, []int{10,20,30,20,50}, []int{1,4,5,6,7,8,11}))
}