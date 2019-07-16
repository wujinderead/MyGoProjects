package greedy

// make a fraction a/b to the form of 1/x1 + 1/x2 + 1/x3 + ···
// e.g.  2/7 = 1/4 + 1/28
func egyptianFraction(a, b int) []int {
	facs := make([]int, 0)
	if a > b {
		return facs
	}
	for i := 0; i < 10; i++ {
		if b%a == 0 {
			facs = append(facs, b/a)
			return facs
		}
		x := b/a + 1
		facs = append(facs, x)
		a = a*x - b
		b = b * x
	}
	return facs
}
