package backtracking

import (
	"fmt"
	"testing"
)

//export GOPATH=/home/xzy/go:$(pwd)
//go test -c -o /tmp/aaa algorithms/backtracking
//go test -run /tmp/aaa algorithms/backtracking
//go tool test2json -t /tmp/aaa -test.v -test.run ^TestKnightRun$

// the order of delta would greatly affect the result
// use the following delta sequence for size 8, it take hours to get the result
// var deltas = [][]int{{2, 1}, {2, -1}, {-2, 1}, {-2, -1}, {1, 2}, {1, -2}, {-1, 2}, {-1, -2}}
func TestKnightTour(t *testing.T) {
	for i := 3; i <= 8; i++ {
		fmt.Println("size", i, ":")
		knightTour(i)
	}
}

func TestKnightTourWarnsdorff(t *testing.T) {
	for i := 3; i <= 20; i++ {
		fmt.Println("size", i, ":")
		knightTourWarnsdorff(i)
	}
}
