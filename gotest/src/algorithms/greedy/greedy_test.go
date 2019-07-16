package greedy

import (
	"fmt"
	"testing"
)

func TestFractionOne(t *testing.T) {
	cases := [][]int{{6, 49}, {41, 73}, {6, 48}, {1, 43}, {14, 9}, {4, 14}, {2, 7}}
	for i := range cases {
		a, b := cases[i][0], cases[i][1]
		ans := egyptianFraction(a, b)
		fmt.Printf("%d/%d = %v\n", a, b, ans)
	}
}
