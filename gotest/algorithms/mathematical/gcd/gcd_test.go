package gcd

import (
	"testing"
)

var testdata = [][]int{
	{13, 13, 13},
	{37, 600, 1},
	{20, 100, 20},
	{2061517, 624129, 18913},
	{33, 101, 1},
	{1234, 0, 1234},
	{0, 56, 56},
}

func TestGcd(t *testing.T) {
	for i:=0; i<len(testdata); i++ {
		a, b, eg := testdata[i][0], testdata[i][1], testdata[i][2]
		g := gcd(a, b)
		if g != eg {
			t.Errorf("gcd(%d, %d), expected %d, result %d", a, b, eg, g)
		}
		t.Logf("gcd(%d, %d) = %d", a, b, g)
	}
}

func TestExtendedGcd(t *testing.T) {
	for i:=0; i<len(testdata); i++ {
		a, b, eg := testdata[i][0], testdata[i][1], testdata[i][2]
		x, y, g := extendedGcd(a, b)
		if g != eg || a*x+b*y != g {
			t.Errorf("gcd(%d, %d) != %d*%d+%d*%d", a, b, a, x, b, y)
		}
		t.Logf("gcd(%d, %d) = %d = %d*%d+%d*%d", a, b, g, a, x, b, y)
	}
}