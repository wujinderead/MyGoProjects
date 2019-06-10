package backtracking

import (
	"fmt"
	"testing"
)

//export GOPATH=/home/xzy/go:$(pwd)
//go test -c -o /tmp/aaa algorithms/backtracking
//go test -run /tmp/aaa algorithms/backtracking
//go tool test2json -t /tmp/aaa -test.v -test.run ^TestKnightRun$
func TestKnightRun(t *testing.T) {
	for i := 3; i <= 8; i++ {
		fmt.Println("size", i, ":")
		knightTour(i)
	}
}
