package util

import (
	"fmt"
	"math/big"
	"math/rand"
	"sync"
	"testing"
	"time"
)

func Test_NextBytes(t *testing.T) {
	m := make([]byte, 35)
	NextBytes(m)
	NextBytes(m)
}

func Test_NextBytesMulti(t *testing.T) {
	var wg sync.WaitGroup
	for i := 0; i < 3; i++ {
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
	for i := 0; i < 9; i++ {
		b := rand.Intn(253) + 2
		fmt.Printf("#172.%d.%d.1\n", a, b)
		fmt.Printf("#172.%d.%d.0\n", a, b)
		fmt.Printf("#172.%d.%d.0/24\n", a, b)
		a += 2
	}
}

func TestCswap(t *testing.T) {
	cswap := func(b byte, x0, x1 *big.Int) {
		m := ^big.Word(b) + 1 // if b=0, ^b=11111111, ^b+1=0; if b=1, ^b=11111110, ^b+1=11111111
		v := new(big.Int).Xor(x1, x0)
		bits := v.Bits() // changing Bits() will modify the big.Int directly
		for i := range bits {
			bits[i] = bits[i] & m
		}
		x0.Xor(x0, v)
		x1.Xor(x1, v)
	}
	src := rand.New(rand.NewSource(time.Now().UnixNano()))
	a := new(big.Int).SetInt64(1)
	a.Lsh(a, 300)
	b := new(big.Int).SetInt64(1)
	b.Lsh(b, 100)

	x0, x1 := new(big.Int).Rand(src, a), new(big.Int).Rand(src, b)
	fmt.Println(x0, x1)
	cswap(0, x0, x1)
	fmt.Println(x0, x1)
	fmt.Println()

	x0, x1 = new(big.Int).Rand(src, a), new(big.Int).Rand(src, b)
	fmt.Println(x0, x1)
	cswap(1, x0, x1)
	fmt.Println(x0, x1)
	fmt.Println()
}
