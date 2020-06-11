package main

import (
    "fmt"
)

// https://leetcode.com/problems/nth-magical-number/

// A positive integer is magical if it is divisible by either A or B. 
// Return the N-th magical number. Since the answer may be very large, return it modulo 10^9 + 7. 
// Example 1:  
//   Input: N = 1, A = 2, B = 3
//   Output: 2
// Example 2: 
//   Input: N = 4, A = 2, B = 3
//   Output: 6
// Example 3: 
//   Input: N = 5, A = 2, B = 4
//   Output: 10
// Example 4: 
//   Input: N = 3, A = 6, B = 4
//   Output: 8
// Note: 
//   1 <= N <= 10^9 
//   2 <= A <= 40000 
//   2 <= B <= 40000 

func nthMagicalNumber(N int, A int, B int) int {
	if A>B {
		A, B = B, A
	}
	l, r := A, A*N
	lcd := A/gcd(A,B)*B
	for l<=r {
		mid := (l+r)/2
		count := mid/A + mid/B - mid/lcd  // get the count of numbers that can either divide A or B
		// when we want the first mid that can satisfy some condition, (e.g., for count=32, we have 96~99, but we want 96) 
		// we use l<=r as for condition, and l=mid+1 and r=mid-1 in loop.
		if count>=N {    // for A=4 B=6: mid  count                                   
			r = mid-1                 //  95   31                    
		} else {                      //  96   32              
			l = mid+1                 //  97   32                   
		}                             //  98   32       
                                      //  99   32           
	}                                 // 100   33
    return l%(1e9+7)                  // 101   33                   
}

func gcd(a, b int) int {
	if b==0 {
		return a
	}
	return gcd(b, a%b)
}

func main() {
	for i:=1; i<=6; i++ {
		fmt.Println(i, nthMagicalNumber(i, 4, 6))
	}
	for i:=31; i<=34; i++ {
		fmt.Println(i, nthMagicalNumber(i, 4, 6))  // 100 is 33th
	}
	fmt.Println(nthMagicalNumber(1000000, 12345, 33333), 9209844)
	fmt.Println(nthMagicalNumber(1, 2, 3), 2)
	fmt.Println(nthMagicalNumber(4, 2, 3), 6)
	fmt.Println(nthMagicalNumber(5, 2, 4), 10)
}