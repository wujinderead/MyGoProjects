package main

import (
	"util"
	"time"
	"golang.org/x/crypto/ripemd160"
	"fmt"
)

func main() {
	go util.Listen()
	time.Sleep(time.Second)
	util.Send()
	fmt.Println(ripemd160.Size)
}