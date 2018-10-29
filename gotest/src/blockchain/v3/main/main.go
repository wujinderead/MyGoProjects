package main

import "blockchain/v3"

func main() {
	bc := v3.NewBlockchain()
	defer bc.Close()

	cli := v3.NewCli(bc)
	cli.Run()
}
