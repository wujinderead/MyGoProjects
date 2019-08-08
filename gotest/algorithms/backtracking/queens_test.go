package backtracking

import (
	"fmt"
	"testing"
)

func TestQueens(t *testing.T) {
	for i := 4; i <= 10; i++ {
		fmt.Println("queens:", i)
		queens(i)
		fmt.Println()
	}
}

func TestQueensTrack(t *testing.T) {
	b := newBoard(6)
	b.setColTrack(0)
}

func TestQueensNoStruct(t *testing.T) {
	queensNoStruct(6)
	fmt.Println()
	queens(6)
}
