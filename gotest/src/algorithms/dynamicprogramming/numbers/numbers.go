package numbers

// fibonacci numbers
// F(0)=0, F(1)=1, F(n)=F(n-1)+F(n-2)
// sequence: 0, 1, 1, 2, 3, 5, 8, 13, 21, 34, 55, 89, 144, ...
// return n-th fibonacci number (0-indexed)
func fibonacci(n int) int {
	a := 0
	b := 1
	for i := 0; i < n; i++ {
		a, b = b, a+b
	}
	return a
}

// catalan numbers
// C(0)=1, C(1)=1, C(n+1)=C(0)C(n)+C(1)C(n-1)+...+C(n)C(0)
// sequence: 1, 1, 2, 5, 14, 42, 132, 429, 1430, 4862, ...
// return n-th catalan number (0-indexed)
func catalanDp(n int) int {
	if n == 0 || n == 1 {
		return 1
	}
	catalans := make([]int, n+1)
	catalans[0] = 1
	catalans[1] = 1
	for i := 2; i <= n; i++ {
		for j := 0; j < i; j++ {
			catalans[i] += catalans[j] * catalans[i-1-j]
		}
	}
	return catalans[n]
}

// use formula: C(n)=nCr(n, 2n)/(n+1)=(2n!)/(n!*(n+1)!)=2n*(2n-1)*···*(n+2)/n!
func catalanFormula(n int) int {
	cata := 1
	for i := 1; i <= n; i++ {
		cata *= n + i
		cata /= i
	}
	return cata / (n + 1)
}

// bell numbers: number of ways to partition a set.
// let S(n, k) be the number of ways to partition n elements into k sets,
// then S(n, k) = k*S(n-1, k) + S(n-1, k-1), base case: S(n,1)=1, S(n,n)=1
// Bell(n) = Σ(k=1,···,n)S(n,k)
// sequence: 1, 1, 2, 5, 15, 52, 203, ...
// return n-th catalan number (0-indexed)
func bell(n int) int {
	if n == 0 {
		return 1
	}
	s := make([]int, n+1) // s[j] to store S(i, j)
	for i := 1; i <= n; i++ {
		prev := 1
		for j := 1; j <= i; j++ {
			if j == 1 || j == i {
				s[j] = 1
				continue
			}
			prev, s[j] = s[j], prev+j*s[j]
		}
		//fmt.Println(s[1:])
	}
	sum := 0
	for j := 1; j <= n; j++ {
		sum += s[j] // Bell(n) = S(n,1) + S(n,2) + ··· + S(n,n)
	}
	return sum
}
