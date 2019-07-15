package tiling_dominoes

// f(n) = f(n-1) + f(n-2), f(1)=1, f(2)=2
func board2xnDomino2x1(n int) int {
	if n == 1 || n == 2 {
		return n
	}
	a, b := 1, 2
	for i := 3; i <= n; i++ {
		a, b = b, a+b
	}
	return b
}

// 2×1 dominoes to fill 3×n board
/* can only be three situation:
   A(n) means number of ways to fill 3×n board completely
   B(n) means number of ways to fill 3×n board with upper left corner missing

                XX         XX          Y
   A(n) = A(n-2)YY + B(n-1) Y + C(n-1) Y
                ZZ          Y         XX

                           ZZ
   B(n) = A(n-1)X  + B(n-2) XX
                X           YY

i.e., A(n) = A(n-2) + 2*B(n-1)
      B(n) = A(n-1) + B(n-2)
base case: A(0)=1, A(1)=0, B(0)=0, B(1)=1
*/
func board3xnDomino2x1(n int) int {
	a := make([]int, n+1)
	b := make([]int, n+1)
	a[0], a[1] = 1, 0
	b[0], b[1] = 0, 1
	for i := 2; i <= n; i++ {
		a[i] = a[i-2] + 2*b[i-1]
		b[i] = a[i-1] + b[i-2]
	}
	if n%2 == 0 {
		return a[n] // n is even, can fill full board
	} else {
		return 2 * b[n] // n is odd, can not fill full board
	}
}

// https://www.geeksforgeeks.org/painting-fence-algorithm/
// given a fence with n posts and k colors, find out the number of ways of painting
// the fence such that at most 2 adjacent posts have the same color.
// Since answer can be large return it modulo 10^9 + 7.
func paintFence(n, k int) int {
	// for position 0 and 1, we both have k choice,
	// for position 0 and 1, we have k same two nodes, and k*(k-1) different two nodes,
	// we store same[1]=k and diff[1]=k*(k-1)
	// total[i] = same[i] + diff[i]
	// same[i]  = diff[i-1]
	// diff[i]  = (diff[i-1] + diff[i-2]) * (k-1)
	//          = total[i-1] * (k-1)
	modolus := 1000000007
	if n == 0 {
		return k
	}
	diff := make([]int, n)
	same := make([]int, n)
	diff[1] = k * (k - 1)
	same[1] = k
	for i := 2; i < n; i++ {
		same[i] = diff[i-1]
		diff[i] = ((same[i-1] + diff[i-1]) * (k - 1)) % modolus
	}
	return (same[n-1] + diff[n-1]) % modolus
}
