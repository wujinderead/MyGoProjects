package sort

import "fmt"

type stringSorter struct {
	runes []rune
}

func newStringSorter(str string) *stringSorter {
	return &stringSorter{[]rune(str)}
}

func (runes *stringSorter) Less(i, j int) bool {
	return runes.runes[i] < runes.runes[j]
}

func (runes *stringSorter) Swap(i, j int) {
	runes.runes[i], runes.runes[j] = runes.runes[j], runes.runes[i]
}

func (runes *stringSorter) Len() int {
	return len(runes.runes)
}

func (runes *stringSorter) String() string {
	return string(runes.runes) + " " + fmt.Sprint(runes.runes)
}

func intSlieceEqual(s1, s2 []int) bool {
	if len(s1) != len(s2) {
		return false
	}
	for i:=0; i<len(s1); i++ {
		if s1[i] != s2[i] {
			return false
		}
	}
	return true
}
