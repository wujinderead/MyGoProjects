package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"

	"gotest/ostest/assembly"
)

/*
go build -o /tmp/getg /home/xzy/golang/gotest/ostest/main/testAsm.go
/tmp/getg 2>/tmp/log1
# $11 g.m.g0, $12 g.m, $13 g.m.p; g0 and m are always one-one
cat /tmp/log1 | grep g0p | wc -l
cat /tmp/log1 | grep g0p | awk '{print $11}' | uniq | wc -l
cat /tmp/log1 | grep g0p | awk '{print $12}' | uniq | wc -l
cat /tmp/log1 | grep g0p | awk '{print $13}' | uniq | wc -l
cat /tmp/log1 | grep g0p | awk '{print $11 $12}' | uniq | wc -l
*/
func main() {
	runtime.GOMAXPROCS(5)
	fmt.Println("gomaxprocs:", runtime.GOMAXPROCS(0))
	time.Sleep(time.Millisecond)
	n := 1000
	var wg sync.WaitGroup
	wg.Add(n)
	for i := 0; i < n; i++ {
		if i%3 == 2 {
			go func(wg *sync.WaitGroup) {
				runtime.Gosched()
				assembly.Display()
				wg.Done()
			}(&wg)
		}
		if i%3 != 2 {
			go func(wg *sync.WaitGroup) {
				assembly.Display()
				time.Sleep(30 * time.Millisecond)
				wg.Done()
			}(&wg)
		}
	}
	wg.Wait()
}
