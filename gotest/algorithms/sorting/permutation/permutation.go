package permutation

import (
	"fmt"
	"math/rand"
	"time"
)

var rander = rand.New(rand.NewSource(time.Now().Unix()))

func PermN(n int) []int {
	if n < 1 {
		return []int{}
	}
	ans := make([]int, n)
	for i := 0; i < n; i++ {
		j := rander.Intn(i + 1)
		ans[i] = ans[j]
		ans[j] = i
	}
	return ans
}

func Shuffle(n int, swap func(i, j int)) {
	if n < 1 {
		return
	}
	for i := n - 1; i > 0; i-- {
		swap(i, rander.Intn(i+1))
	}
}

func PermuteNonDuplicated(arr []interface{}) {
	// to detect duplicated
	isDuplicated := func(arr []interface{}, left, right int) bool {
		for i := left; i < right; i++ {
			if arr[i] == arr[right] {
				return true
			}
		}
		return false
	}
	var permute func([]interface{}, int)
	permute = func(arr []interface{}, step int) {
		if step == len(arr)-1 {
			fmt.Println(arr)
			return
		}
		for i := step; i < len(arr); i++ {
			if isDuplicated(arr, step, i) {
				continue
			}
			swap(arr, step, i)
			permute(arr, step+1)
			swap(arr, step, i)
		}
	}
	permute(arr, 0)
}

func swap(arr []interface{}, i, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}

func Combine(arr []interface{}, num int) {
	if num < 0 || num > len(arr) {
		panic("n is out of range")
	}
	var combine func([]interface{}, int, int)
	com := make([]interface{}, num)
	combine = func(arr []interface{}, start, n int) {
		if n == 0 {
			fmt.Println(com)
			return
		}
		for i := start; i <= len(arr)-n; i++ {
			com[num-n] = arr[i]
			combine(arr, i+1, n-1)
		}
	}
	combine(arr, 0, num)
}
