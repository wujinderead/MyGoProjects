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

// 2×1 dominos to fill 3×n board
/* can only be three situation:
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
