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