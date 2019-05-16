package catalan

import "fmt"

// https://en.wikipedia.org/wiki/Catalan_number
// catalan number, C(0)=1, C(n)=C(0)C(n-1)+C(1)C(n-2)+...+C(n-1)C(0)
// catalan number series: 1, 1, 2, 5, 14, 42, 132, 429, 1430, 4862
// the approach to calculate catalan number:
// (1). dynamic programming, O(n^2)
func getCatalanDp(n int) int {
	catalans := make([]int, n+1)
	catalans[0] = 1
	for i:=1; i<=n; i++ {
		for j:=0; j<i; j++ {
			catalans[i] += catalans[j]*catalans[i-1-j]
		}
	}
	return catalans[n]
}

// (2). formula C(n)=C(n, 2n)/(n+1) = 2n!/(n!*n!)/(n+1) = (n+1)*(n+2)*...*(2n)/(1*2*...*n)/(n+1)
func getCatalanFormula(n int) int {
	c := 1
	for i:=1; i<=n; i++ {
		c = c*(n+i)
		c = c/i
	}
	return c/(n+1)
}

// catalan numbers application:
// (1). the balanced parentheses, C(n) is the number of balanced parenthesis of size 2n
//      n=1 C(1)=1    ()
//      n=2 C(2)=2    ()(), (())
//      n=3 C(3)=5    (()()), ((())), ()()(), ()(()), (())()
func printBalancedParenthesis(pair int) {
	chars := make([]byte, pair<<1)
	printBalancedParenthesisHelper(chars, 0, 0, 0)
}

func printBalancedParenthesisHelper(chars []byte, a, b, i int) {
	if a==len(chars)>>1 {   // all 'a' is set, fill remain with 'b'
		for i<len(chars) {
			chars[i] = 'B'
			i++
		}
		fmt.Println(string(chars))
		return
	}
	if a==b {               // a, b is the same number, can only set a
		chars[i] = 'A'
		printBalancedParenthesisHelper(chars, a+1, b, i+1)
		return
	}
	chars[i] = 'A'          // a>b, can set a or set b
	printBalancedParenthesisHelper(chars, a+1, b, i+1)
	chars[i] = 'B'
	printBalancedParenthesisHelper(chars, a, b+1, i+1)
}