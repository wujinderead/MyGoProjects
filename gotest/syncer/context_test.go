package syncer

import (
	"context"
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestContextTimeout(t *testing.T) {
	num := 10
	timeout := 5
	var wg sync.WaitGroup
	wg.Add(num)
	// when timeout, all <-ctx.Done() can receive from channel
	ctx, _ := context.WithTimeout(context.Background(), time.Duration(timeout*int(time.Second)))
	for i:=1; i<=10; i++ {
		go func(ctx context.Context, i int) {
			select {
			case <-time.After(time.Duration(i * int(time.Second))):
				fmt.Println(i, "done  :", time.Now())
			case <-ctx.Done():
				fmt.Println(i, "cancel:", time.Now())
			}
			wg.Done()
		}(ctx, i)
	}
	wg.Wait()
	fmt.Println("exit", time.Now())
}

func TestContextCancel(t *testing.T) {
	num := 10
	timeout := 5
	var wg sync.WaitGroup
	wg.Add(num)
	// when timeout, all <-ctx.Done() can receive from channel
	ctx, cancel := context.WithCancel(context.Background())
	for i:=1; i<=10; i++ {
		go func(ctx context.Context, i int) {
			select {
			case <-time.After(time.Duration(i * int(time.Second))):
				fmt.Println(i, "done  :", time.Now())
			case <-ctx.Done():
				fmt.Println(i, "cancel:", time.Now())
			}
			wg.Done()
		}(ctx, i)
	}
	<- time.After(time.Duration(timeout*int(time.Second)))
	cancel()  // use cancel() func to trigger all <-ctx.Down() can receive
	wg.Wait()
	fmt.Println("exit", time.Now())
}
