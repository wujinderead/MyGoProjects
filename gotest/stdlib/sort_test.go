package stdlib

import (
	"fmt"
	"math"
	"math/rand"
	"sort"
	"strings"
	"testing"
)

func TestSort(t *testing.T) {
	// sort provides functions to sort []int, []string, and []float64
	{
		s := []float64{math.Inf(1), math.NaN(), 4.7, math.Inf(-1), -2.3, 0.0, 3.5}
		fmt.Println(s)
		fmt.Println(sort.Float64sAreSorted(s))
		sort.Float64s(s)
		fmt.Println(s)
		fmt.Println(sort.Float64sAreSorted(s))
		fmt.Println()
	}
	// use Float64Slice to sort []float64
	{
		s := []float64{math.Inf(1), math.NaN(), 4.7, math.Inf(-1), -2.3, 0.0, 3.5}
		ss := sort.Float64Slice(s)
		fmt.Println(ss)
		fmt.Println(sort.IsSorted(ss))
		sort.Sort(ss)
		fmt.Println(ss)
		fmt.Println(s)
		fmt.Println(sort.IsSorted(ss))
		fmt.Println()
	}
	{
		s := rand.Perm(20)
		fmt.Println(s)
		fmt.Println(sort.IntsAreSorted(s))
		sort.Ints(s)
		fmt.Println(s)
		fmt.Println(sort.IntsAreSorted(s))
		fmt.Println()
	}
	// use IntSlice to sort []int
	{
		s := rand.Perm(20)
		ss := sort.IntSlice(s)
		fmt.Println(ss)
		fmt.Println(sort.IsSorted(ss))
		sort.Sort(ss)
		fmt.Println(s)
		fmt.Println(sort.IsSorted(ss))
		fmt.Println()
	}
	{
		s := strings.Split("The Boeing 787 Dreamliner is an American long-haul, mid-size wide-body, " +
			"twin-engine jet airliner made by Boeing Commercial Airplanes. " +
			"Its variants seat 242 to 335 passengers in typical three-class seating configurations. ", " ")
		fmt.Println(s)
		fmt.Println(sort.StringsAreSorted(s))
		sort.Strings(s)
		fmt.Println(s)
		fmt.Println(sort.StringsAreSorted(s))
		fmt.Println()
	}
	{
		s := strings.Split("The Boeing 787 Dreamliner is an American long-haul, mid-size wide-body, " +
			"twin-engine jet airliner made by Boeing Commercial Airplanes. " +
			"Its variants seat 242 to 335 passengers in typical three-class seating configurations. ", " ")
		ss := sort.StringSlice(s)
		fmt.Println(ss)
		fmt.Println(sort.IsSorted(ss))
		sort.Sort(ss)
		fmt.Println(s)
		fmt.Println(sort.IsSorted(ss))
		fmt.Println()
	}
}

type int8Slice []int8
func (s int8Slice) Less(i, j int) bool {
	return s[i]<s[j]
}
func (s int8Slice) Len() int {
	return len(s)
}
func (s int8Slice) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func TestSortInterface(t *testing.T) {
	{
		s := int8Slice([]int8{12, 4, 2, 13, 10, 0, 19, 11, 7, 5, 15, 18, 9, 14, 6, 8, 1, 16, 17, 3})
		fmt.Println("s: ", s)
		fmt.Println("is sorted: ", sort.IsSorted(s))
		sort.Sort(s)
		fmt.Println("s: ", s)
		fmt.Println("is sorted: ", sort.IsSorted(s))
		fmt.Println()
	}
	{
		s := int8Slice([]int8{12, 4, 2, 13, 10, 0, 19, 11, 7, 5, 15, 18, 9, 14, 6, 8, 1, 16, 17, 3})
		fmt.Println("s: ", s)
		fmt.Println("is sorted: ", sort.IsSorted(s))
		// reverse sort
		sort.Sort(sort.Reverse(s))
		fmt.Println("reverse sort: ", s)
	}
}

func TestSortSlice(t *testing.T) {
	chars := []rune{'中', '国', '行', '政', '地', '图'}
	sort.Slice(chars, func(i, j int) bool {
		return chars[i] < chars[j]
	})
	fmt.Println(chars, string(chars))
}
