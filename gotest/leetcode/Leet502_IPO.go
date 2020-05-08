package main

import (
	"container/heap"
	"fmt"
)

// https://leetcode.com/problems/ipo/

// Suppose LeetCode will start its IPO soon. In order to sell a good price of its
// shares to Venture Capital, LeetCode would like to work on some projects to increase
// its capital before the IPO. Since it has limited resources, it can only finish
// at most k distinct projects before the IPO. Help LeetCode design the best way to
// maximize its total capital after finishing at most k distinct projects.
// You are given several projects. For each project i, it has a pure profit Pi and
// a minimum capital of Ci is needed to start the corresponding project. Initially,
// you have W capital. When you finish a project, you will obtain its pure profit
// and the profit will be added to your total capital.
// To sum up, pick a list of at most k distinct projects from given projects to maximize
// your final capital, and output your final maximized capital.
// Example 1:
//   Input: k=2, W=0, Profits=[1,2,3], Capital=[0,1,1].
//   Output: 4
//   Explanation: Since your initial capital is 0, you can only start the project indexed 0.
//     After finishing it you will obtain profit 1 and your capital becomes 1.
//     With capital 1, you can either start the project indexed 1 or the project indexed 2.
//     Since you can choose at most 2 projects, you need to finish the project indexed 2
//     to get the maximum capital. Therefore, output the final maximized capital, which is 0+1+3=4.
// Note:
//   You may assume all numbers in the input are non-negative integers.
//   The length of Profits array and Capital array will not exceed 50,000.
//   The answer is guaranteed to fit in a 32-bit signed integer.

func findMaximizedCapital(k int, W int, Profits []int, Capital []int) int {
    // for each operation, we always pick the i that W>=Capital[i] and Profit[i] is the maximal.
	p := pairs(make([]pair, 0, len(Profits)))
	for i:=0; i<len(Profits); i++ {  // push capital profit pair to min heap based on capital
		heap.Push(&p, pair{Profits[i], Capital[i]})
	}
	op := 0
	is := ints(make([]int, 0, len(Profits)))   // max heap for profit
	for op<k {
		for len(p)>0 && p[0].capital <= W {
			pi := heap.Pop(&p).(pair)     // pop all i that capital[i]<=W
			heap.Push(&is, pi.profit)
		}
		if is.Len()==0 {   // no profit can be used
			break
		}
		W += heap.Pop(&is).(int)
		op++
	}
	return W
}

type pair struct {
	profit, capital int
}

type pairs []pair

func (p pairs) Len() int {
	return len(p)
}

func (p pairs) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func (p pairs) Less(i, j int) bool {
	return p[i].capital < p[j].capital
}

func (p *pairs) Push(x interface{}) {
	*p = append(*p, x.(pair))
}

func (p *pairs) Pop() interface{}  {
	x := (*p)[len(*p)-1]
	*p = (*p)[:len(*p)-1]
	return x
}

type ints []int

func (p ints) Len() int {
	return len(p)
}

func (p ints) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func (p ints) Less(i, j int) bool {
	return p[i] > p[j]
}

func (p *ints) Push(x interface{}) {
	*p = append(*p, x.(int))
}

func (p *ints) Pop() interface{}  {
	x := (*p)[len(*p)-1]
	*p = (*p)[:len(*p)-1]
	return x
}

func main() {
	fmt.Println(findMaximizedCapital(2, 0, []int{1,2,3}, []int{0,1,1}))
	fmt.Println(findMaximizedCapital(4, 0, []int{1,2,3}, []int{0,1,1}))
	fmt.Println(findMaximizedCapital(2, 0, []int{1,2,3}, []int{0,3,2}))
}