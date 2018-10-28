package v2

import (
	"fmt"
	"strconv"
	"testing"
)

func TestProofOfWork(t *testing.T) {
	bc := NewBlockchain()

	bc.AddBlock("lgq mine a block")
	bc.AddBlock("lol mine a block")

	for _, block := range bc.blocks {
		block.Print()
		pow := NewProofOfWork(block)
		fmt.Printf("PoW: %s\n", strconv.FormatBool(pow.Validate()))
		fmt.Println()
	}
}
