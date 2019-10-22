package leetcode

import (
	"fmt"
	"strings"
)

func toGoatLatin(S string) string {
	strs := strings.Split(S, " ")
	ans := make([]string, 0)
	for i := 0; i < len(strs); i++ {
		str := strs[i]
		if str[0] == 'a' || str[0] == 'A' || str[0] == 'e' || str[0] == 'E' ||
			str[0] == 'i' || str[0] == 'I' || str[0] == 'o' || str[0] == 'O' ||
			str[0] == 'u' || str[0] == 'U' {
			str = str + "ma"
		} else {
			str = str[1:] + string(str[0]) + "ma"
		}
		for j := 0; j < i+1; j++ {
			str = str + "a"
		}
		ans = append(ans, str)
	}
	return strings.Join(ans, " ")
}

func main() {
	fmt.Println(toGoatLatin("I speak Goat Latin"))
	fmt.Println(toGoatLatin("The quick brown fox jumped over the lazy dog"))
}
