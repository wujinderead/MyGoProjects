package gcd

func gcd(a, b int) int {
	if a==0 {
		return b
	}
	return gcd(b%a, a)
}

// not compute only gcd(a,b), but also x, y that make ax+by = gcd(a,b)
func extendedGcd(a, b int) (x, y, g int) {
	if a==0 {
		return 0, 1, b
	}
	x1, y1, g := extendedGcd(b%a, a);
	x = y1 - b/a*x1
	y = x1
	return
}