package main

import (
    "fmt"
)

// https://leetcode.com/problems/student-attendance-record-ii/

// Given a positive integer n, return the number of all possible attendance records with length n, 
// which will be regarded as rewardable. The answer may be very large, return it after mod 10^9 + 7.
// A student attendance record is a string that only contains the following three characters:
//   'A' : Absent.
//   'L' : Late.
//   'P' : Present.
// A record is regarded as rewardable if it doesn't contain more than one 'A' (absent) or 
// more than two continuous 'L' (late).
// Example 1:
//   Input: n = 2
//   Output: 8 
//   Explanation:
//     There are 8 records with length 2 will be regarded as rewardable:
//     "PP" , "AP", "PA", "LP", "PL", "AL", "LA", "LL"
//     Only "AA" won't be regarded as rewardable owing to more than one absent times. 

func checkRecord(n int) int {
	if n<2 {
		return 3*n
	}
	// arr[i] is the number of ways to end with P or L or LL with length i 
    p, l, ll := make([]int, n+2), make([]int, n+2), make([]int, n+2)
   	p[1] = 1
   	l[1] = 1
   	ll[1] = 0
    for i:=2; i<=n+1; i++ {
    	p[i] = (p[i-1]+l[i-1]+ll[i-1]) % (1e9+7)      // end at P
    	l[i] = p[i-1]                     // end at L
    	ll[i] = l[i-1]                    // end at LL
    }
    sum := 0
    for i:=0; i<n; i++ {        //  (i=0 to n-1 with only PL, i.e. p[i+1]) A (len=n-1-i with only PL)
    	sum += p[i+1]*p[n-1-i+1]
    	sum = sum % (1e9+7)  
    }
    sum += p[n+1]
    return sum % (1e9+7)
}

func main() {
    fmt.Println(checkRecord(0))
    fmt.Println(checkRecord(1))
    fmt.Println(checkRecord(2))
    fmt.Println(checkRecord(3))
    fmt.Println(checkRecord(4))
    fmt.Println(checkRecord(5)) 
    fmt.Println(checkRecord(1000))       
}