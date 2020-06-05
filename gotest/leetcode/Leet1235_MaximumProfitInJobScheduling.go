package main

import (
    "fmt"
    "sort"
)

// https://leetcode.com/problems/maximum-profit-in-job-scheduling/

// We have n jobs, where every job is scheduled to be done from startTime[i] to endTime[i], 
// obtaining a profit of profit[i]. You're given the startTime , endTime and profit arrays, 
// you need to output the maximum profit you can take such that there are no 2 jobs in the 
// subset with overlapping time range. If you choose a job that ends at time X you will be 
// able to start another job that starts at time X. 
// Example 1: 
//           [-40--]       
//        [--10-]     
//     [--50-][---70--]   
//     1  2  3  4  5  6
//   Input: startTime = [1,2,3,3], endTime = [3,4,5,6], profit = [50,10,40,70]
//   Output: 120
//   Explanation: The subset chosen is the first and fourth job. 
//     Time range [1-3]+[3-6] , we get profit of 120 = 50 + 70.
// Example 2: 
//   Input: startTime = [1,2,3,4,6], endTime = [3,5,10,6,9], profit = [20,20,100,70,60]
//   Output: 150
//   Explanation: The subset chosen is the first, fourth and fifth job. 
//     Profit obtained 150 = 20 + 70 + 60.
// Example 3: 
//     [------4-------]       
//     [----6----]     
//     [--5-]   
//     1    2    3    4
//   Input: startTime = [1,1,1], endTime = [2,3,4], profit = [5,6,4]
//   Output: 6
// Constraints: 
//   1 <= startTime.length == endTime.length == profit.length <= 5 * 10^4 
//   1 <= startTime[i] < endTime[i] <= 10^9 
//   1 <= profit[i] <= 10^4 

// it's just 01-knapsack. but the problem is if we choose a job, which job is overlapping?
// we need to sort the job by start time. we use binary search to find endtime in the sorted start time. 
// time O(nlogn) for sorting, O(nlogn) for binary search for each job, O(n) for dp. space O(n).
func jobScheduling(startTime []int, endTime []int, profit []int) int {
    // sort jobs by startTime
    sort.Sort(&jobs{startTime, endTime, profit})

    n := len(startTime)
    endInStart := make([]int, n)   // endInStart[i] is the index of endTime[i] in startTime
    for i, v := range endTime {
    	if v>startTime[n-1] {
    		endInStart[i] = n
    		continue
    	}
    	l, r := 0, n-1
    	for l<r {
    		mid := (l+r)/2
    		if v>startTime[mid] {
    			l = mid+1
    		} else {
    			r = mid    			
    		}
    	}
    	endInStart[i] = l
    }
    // fmt.Println(startTime, endTime, endInStart)

    // dp[i] is the overall profit for job[i:]. if not use job[i], dp[i] = dp[i+1];
    // if use job[i], dp[i] = profit[i] + dp[endInStart[i]]; we compare and find the max.
    dp := make([]int, n+1)
    for i:=n-1; i>=0; i-- {
    	dp[i] = dp[i+1]
    	if profit[i]+dp[endInStart[i]] > dp[i] {
    		dp[i] = profit[i]+dp[endInStart[i]]
    	}
    }
    return dp[0]
}

type jobs struct {
	start, end, profit []int
}

func (job *jobs) Len() int {
	return len(job.start)
}

func (job *jobs) Swap(i, j int) {
	job.start[i], job.start[j] = job.start[j], job.start[i]
	job.end[i], job.end[j] = job.end[j], job.end[i]
	job.profit[i], job.profit[j] = job.profit[j], job.profit[i]
}

func (job *jobs) Less(i, j int) bool {
	return job.start[i]<job.start[j]
}

func main() {
	fmt.Println(jobScheduling([]int{1,2,3,3}, []int{3,4,5,6}, []int{50,10,40,70}))
	fmt.Println(jobScheduling([]int{2,3,1,3}, []int{4,5,3,6}, []int{10,40,50,70}))
	fmt.Println(jobScheduling([]int{1,2,3,4,6}, []int{3,5,10,6,9}, []int{20,20,100,70,60}))
	fmt.Println(jobScheduling([]int{1,1,1}, []int{2,3,4}, []int{5,6,4}))
}