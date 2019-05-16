package permutation

import (
	"fmt"
	"testing"
)

func TestPermN(t *testing.T) {
	for i:=0; i<=10; i++ {
		fmt.Println(PermN(i))
	}
}

func TestShuffle(t *testing.T) {
	{
		ints := []int{}
		Shuffle(len(ints), func(i, j int) {
			ints[i], ints[j] = ints[j], ints[i]
		})
		fmt.Println(ints)
	}
	{
		ints := []int{1}
		Shuffle(len(ints), func(i, j int) {
			ints[i], ints[j] = ints[j], ints[i]
		})
		fmt.Println(ints)
	}
	{
		ints := []int{1, 3}
		Shuffle(len(ints), func(i, j int) {
			ints[i], ints[j] = ints[j], ints[i]
		})
		fmt.Println(ints)
	}
	{
		ints := []int{1,2,3}
		Shuffle(len(ints), func(i, j int) {
			ints[i], ints[j] = ints[j], ints[i]
		})
		fmt.Println(ints)
	}
	{
		ints := []int{2,2,2,3,3,4,4,5,5,5}
		Shuffle(len(ints), func(i, j int) {
			ints[i], ints[j] = ints[j], ints[i]
		})
		fmt.Println(ints)
	}
	{
		ints := []rune("文体两开花")
		Shuffle(len(ints), func(i, j int) {
			ints[i], ints[j] = ints[j], ints[i]
		})
		fmt.Println(string(ints))
	}
}

func TestPermute(t *testing.T) {
	PermuteNonDuplicated([]interface{}{1, 2, 3, 4})
	fmt.Println()
	PermuteNonDuplicated([]interface{}{1, 2, 3, 3})
	fmt.Println()
	PermuteNonDuplicated([]interface{}{1, 2, 3})
	fmt.Println()
	PermuteNonDuplicated([]interface{}{1})
	fmt.Println()
	PermuteNonDuplicated([]interface{}{1, 1})
	fmt.Println()
	PermuteNonDuplicated([]interface{}{1, 2})
	fmt.Println()
	PermuteNonDuplicated([]interface{}{})
	fmt.Println()
}

func TestCombine(t *testing.T) {
	Combine([]interface{}{1, 2, 3, 4}, 2)
	fmt.Println()
	Combine([]interface{}{1, 2, 3, 4}, 1)
	fmt.Println()
	Combine([]interface{}{1, 2, 3, 4, 5}, 3)
	fmt.Println()
	Combine([]interface{}{1, 2}, 2)
	fmt.Println()
	Combine([]interface{}{1, 2}, 1)
	fmt.Println()
	Combine([]interface{}{}, 0)
	fmt.Println()
}