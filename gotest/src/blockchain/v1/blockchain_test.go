package v1

import (
	"testing"
)

func TestBlockchain(t *testing.T) {
	bc := NewBlockchain()

	bc.AddBlock("Send 1 BTC to Ivan")
	bc.AddBlock("Send 2 more BTC to Ivan")

	for _, block := range bc.blocks {
		t.Logf("Prev. hash: %x\n", block.PrevBlockHash)
		t.Logf("Data: %s\n", block.Data)
		t.Logf("Hash: %x\n", block.Hash)
		t.Logf("")
	}
}
