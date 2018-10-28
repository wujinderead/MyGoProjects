package util

import (
	"testing"
	"fmt"
	"sync"
	"math/rand"
	"time"
)

func Test_NextBytes(t *testing.T) {
	m := make([]byte, 35)
	NextBytes(m)
	NextBytes(m)
}

func Test_NextBytesMulti(t *testing.T) {
	var wg sync.WaitGroup
	for i:=0; i<3; i++ {
		j := i
		wg.Add(1)
		go func() {
			m := make([]byte, 9)
			NextBytes(m)
			fmt.Printf("%x=%d\n\n", m, j)
			NextBytes(m)
			fmt.Printf("%x=%d\n\n", m, j)
			wg.Done()
		}()
	}
	wg.Wait()
}

func TestBytes(t *testing.T) {
	a := make([]int, 5)
	fmt.Println(len(a), a)
	a[0] = 5
	fmt.Println(len(a), a)
	a = append(a, 1, 2)
	fmt.Println(len(a), a)
	b := make([]int, 3, 6)
	fmt.Println(len(b))
	b[1] = 5
	b = append(b, 1, 2)
	fmt.Println(len(b), b)
}

func TestRand(t *testing.T) {
	rand.Seed(time.Now().Unix())
	a := 17
	for i:=0;i<9;i++ {
		b := rand.Intn(253)+2
		fmt.Printf("#172.%d.%d.1\n", a, b)
		fmt.Printf("#172.%d.%d.0\n", a, b)
		fmt.Printf("#172.%d.%d.0/24\n", a, b)
		a += 2
	}
}