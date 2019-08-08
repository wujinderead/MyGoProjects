package stdlib

import (
	crand "crypto/rand"
	"encoding/hex"
	"fmt"
	"math/big"
	"math/rand"
	"testing"
	"time"
)

func TestRand(t *testing.T) {
	rand.Seed(time.Now().Unix())
	{
		slicer := make([]float64, 0)
		for i := 0; i < 10; i++ {
			// 指数分布
			// exponentially distributed float64 in the range (0, +math.MaxFloat64]
			// with rate parameter (lambda) is 1 and whose mean is 1/lambda (1)
			slicer = append(slicer, rand.ExpFloat64())
		}
		fmt.Println("exp float64: ", slicer)
		fmt.Println()
	}
	{
		slicer := make([]float64, 0)
		for i := 0; i < 10; i++ {
			// 标准正态分布
			// normally distributed float64 in the range [-math.MaxFloat64, +math.MaxFloat64]
			// with standard normal distribution (mean = 0, stddev = 1)
			slicer = append(slicer, rand.NormFloat64())
		}
		fmt.Println("nom float64: ", slicer)
		fmt.Println()
	}
	{
		slicer := make([]float32, 0)
		for i := 0; i < 10; i++ {
			// float32 in [0.0, 1.0]
			slicer = append(slicer, rand.Float32())
		}
		fmt.Println("float32: ", slicer)
		fmt.Println()
	}
	{
		slicer := make([]float64, 0)
		for i := 0; i < 10; i++ {
			// float64 in [0.0, 1.0]
			slicer = append(slicer, rand.Float64())
		}
		fmt.Println("float64: ", slicer)
		fmt.Println()
	}
	{
		slicer := make([]int, 0)
		for i := 0; i < 10; i++ {
			// non-negative int
			slicer = append(slicer, rand.Int())
		}
		fmt.Println("int: ", slicer)
		fmt.Println()
	}
	{
		slicer := make([]int, 0)
		for i := 0; i < 10; i++ {
			// non-negative int in [0,n)
			slicer = append(slicer, rand.Intn(100))
		}
		fmt.Println("int 100: ", slicer)
		fmt.Println()
	}
	{
		slicer := make([]int32, 0)
		for i := 0; i < 10; i++ {
			// non-negative int32
			slicer = append(slicer, rand.Int31())
		}
		fmt.Println("int32: ", slicer)
		fmt.Println()
	}
	{
		slicer := make([]int32, 0)
		for i := 0; i < 10; i++ {
			// non-negative int32 in [0,n)
			slicer = append(slicer, rand.Int31n(100))
		}
		fmt.Println("int32 100: ", slicer)
		fmt.Println()
	}
	{
		slicer := make([]int64, 0)
		for i := 0; i < 10; i++ {
			// non-negative int64
			slicer = append(slicer, rand.Int63())
		}
		fmt.Println("int64: ", slicer)
		fmt.Println()
	}
	{
		slicer := make([]int64, 0)
		for i := 0; i < 10; i++ {
			// non-negative int64 in [0,n)
			slicer = append(slicer, rand.Int63n(100))
		}
		fmt.Println("int64 100: ", slicer)
		fmt.Println()
	}
	{
		slicer := make([]uint32, 0)
		for i := 0; i < 10; i++ {
			// uint32
			slicer = append(slicer, rand.Uint32())
		}
		fmt.Println("uint32: ", slicer)
		fmt.Println()
	}
	{
		slicer := make([]uint64, 0)
		for i := 0; i < 10; i++ {
			// uint64
			slicer = append(slicer, rand.Uint64())
		}
		fmt.Println("uint64: ", slicer)
		fmt.Println()
	}
}

func TestReadPermShuffle(t *testing.T) {
	rand.Seed(time.Now().Unix())
	for i:=0; i<10; i++ {
		fmt.Println(rand.Perm(20))
	}
	byter := make([]byte, 10)
	for i:=0; i<10; i++ {
		n, err := rand.Read(byter)
		if err != nil {
			fmt.Println("read err: ", err.Error())
		}
		fmt.Println(n, hex.EncodeToString(byter))
	}
	words := []rune("中美合拍的西游记即将正式开机")
	for i:=0; i<10; i++ {
		rand.Shuffle(len(words), func(i, j int) {
			words[i], words[j] = words[j], words[i]
		})
		fmt.Println(string(words))
	}
}

func TestPermImplementation(t *testing.T) {
	n := 10
	m := make([]int, n)
	rander := rand.New(rand.NewSource(time.Now().Unix()))
	// for example, if numbers 0 to 7 are already permuted,
	// to insert number 8, random select a position from a[0] to a[8], like a[3],
	// set a[8]=a[3], a[3]=8
	// to insert number 9, random select a position from a[0] to a[9], like a[5],
	// set a[9]=a[5], a[5]=9
	for i := 0; i < n; i++ {
		j := rander.Intn(i + 1)
		// more human logical
		if false {
			m[i] = i
			m[i], m[j] = m[j], m[i]
		}

		// more efficient
		if true {
			m[i] = m[j]
			m[j] = i
		}
		fmt.Println(m)
	}
}

func TestShuffleImplementation(t *testing.T) {
	n := 10
	swap := func(arr []int, i, j int) {
		arr[i], arr[j] = arr[j], arr[i]
	}
	rander := rand.New(rand.NewSource(time.Now().Unix()))
	arr := []int{0,1,2,3,4,5,6,7,8,9}
	// Fisher-Yates shuffle: https://en.wikipedia.org/wiki/Fisher%E2%80%93Yates_shuffle
	// for each position i from n to 0, random select a position j from i to 0, swap a[i] and a[j]
	// for example, for place 9, swap a[9], a[rand.Int(9+1)]
	// for example, for place 8, swap a[8], a[rand.Int(8+1)]
	// ...
	for i:=n-1; i>0; i-- {
		j := rander.Intn(i+1)
		swap(arr, i, j)
		fmt.Println(i, j)
		fmt.Println(arr)
	}
}

func TestCryptoRand(t *testing.T) {
	// cryptography random
	byter := make([]byte, 10)
	n, err := crand.Read(byter)
	if err != nil {
		fmt.Println("read err: ", err.Error())
	}
	fmt.Println(n, hex.EncodeToString(byter))

	// cryptography random reader
	creader := crand.Reader
	n, err = creader.Read(byter)
	if err != nil {
		fmt.Println("read err: ", err.Error())
	}
	fmt.Println(n, hex.EncodeToString(byter))

	// uniform random value in [0, max)
	// read random bytes, construct a big.Int, and check if it is smaller than max.
	// if true, return; if false, continue the loop
	max := new(big.Int).SetInt64(1)
	max.Lsh(max, 100)
	rbi, err := crand.Int(creader, max)
	if err != nil {
		fmt.Println("rand big int err: ", err.Error())
	}
	fmt.Println(rbi.Text(16))

	// read random bytes, construct a big.Int, and check if it is probable prime.
	// if true, return; if false, continue the loop
	p, err := crand.Prime(creader, 100)
	if err != nil {
		fmt.Println("rand prime err: ", err.Error())
	}
	fmt.Println(p.Text(16))
}