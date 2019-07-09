package numbers

import (
	"testing"
)

func TestFibonacci(t *testing.T) {
	fibs := []int{0, 1, 1, 2, 3, 5, 8, 13, 21, 34, 55, 89, 144}
	for i := range fibs {
		ans := fibonacci(i)
		if ans != fibs[i] {
			t.Errorf("index %d, expect %d, got %d", i, fibs[i], ans)
		}
	}
}

func TestCatalan(t *testing.T) {
	catalans := []int{1, 1, 2, 5, 14, 42, 132, 429, 1430, 4862}
	for i := range catalans {
		ans := catalanDp(i)
		if ans != catalans[i] {
			t.Errorf("index %d, expect %d, got %d", i, catalans[i], ans)
		}
	}
	for i := range catalans {
		ans := catalanFormula(i)
		if ans != catalans[i] {
			t.Errorf("index %d, expect %d, got %d", i, catalans[i], ans)
		}
	}
}

func TestBell(t *testing.T) {
	bells := []int{1, 1, 2, 5, 15, 52, 203}
	for i := range bells {
		ans := bell(i)
		if ans != bells[i] {
			t.Errorf("index %d, expect %d, got %d", i, bells[i], ans)
		}
	}
}
