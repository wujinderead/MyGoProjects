package catalan

import (
	"testing"
)

func TestCatalanDp(t *testing.T) {
	catalans := []int{1, 1, 2, 5, 14, 42, 132, 429, 1430, 4862}
	for i:=0; i<len(catalans); i++ {
		c := getCatalanDp(i)
		if c != catalans[i] {
			t.Errorf("error get catalan, exptect %d, got %d", catalans[i], c)
		}
	}
}

func TestCatalanFormula(t *testing.T) {
	catalans := []int{1, 1, 2, 5, 14, 42, 132, 429, 1430, 4862}
	for i:=0; i<len(catalans); i++ {
		c := getCatalanFormula(i)
		if c != catalans[i] {
			t.Errorf("error get catalan, exptect %d, got %d", catalans[i], c)
		}
	}
}

func TestPrintBalancedParenthesis(t *testing.T) {
	printBalancedParenthesis(2)
	printBalancedParenthesis(3)
	printBalancedParenthesis(4)
}