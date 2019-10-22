package leetcode

import "fmt"

func uniqueMorseRepresentations(words []string) int {
	var morse = []string{".-", "-...", "-.-.", "-..", ".", "..-.", "--.", "....", "..", ".---", "-.-", ".-..", "--", "-.", "---", ".--.", "--.-", ".-.", "...", "-", "..-", "...-", ".--", "-..-", "-.--", "--.."}
	var dict = make(map[string]int)
	for _, str := range words {
		var cur = ""
		for _, ch := range str {
			cur += morse[ch-'a']
		}
		dict[cur] = 1
	}
	return len(dict)
}

func main() {
	fmt.Println(uniqueMorseRepresentations([]string{"gin", "zen", "gig", "msg"}))
}
