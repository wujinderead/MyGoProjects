package sort

import "fmt"

type stringQuickSorter struct {
	runes []rune
}

func newStringQuickSorter(str string) *stringQuickSorter {
	return &stringQuickSorter{[]rune(str)}
}

func (runes *stringQuickSorter) Less(i, j int) bool {
	return runes.runes[i] < runes.runes[j]
}

func (runes *stringQuickSorter) Swap(i, j int) {
	runes.runes[i], runes.runes[j] = runes.runes[j], runes.runes[i]
}

func (runes *stringQuickSorter) Len() int {
	return len(runes.runes)
}

func (runes *stringQuickSorter) String() string {
	return string(runes.runes) + " " + fmt.Sprint(runes.runes)
}