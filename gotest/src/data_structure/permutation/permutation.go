package permutation

import (
	"math/rand"
	"time"
)

var rander = rand.New(rand.NewSource(time.Now().Unix()))

func PermN(n int) []int {
	if n<1 {
		return []int{}
	}
	ans := make([]int, n)
	for i:=0; i<n; i++ {
		j := rander.Intn(i+1)
		ans[i] = ans[j]
		ans[j] = i
	}
	return ans
}

func Shuffle(n int, swap func(i, j int)) {
	if n<1 {
		return
	}
	for i:=n-1; i>0; i-- {
		swap(i, rander.Intn(i+1))
	}
}
