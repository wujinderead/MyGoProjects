package pattern_search

import (
	"fmt"
	"testing"
)

func TestLps(t *testing.T) {
	for _, str := range []string{"cabbaabb", "forgeeksskeegfor", "abcde",
		"abcdae", "abacd", "xababayz", "xabaabayz"} {
		maxind, maxstr := lpsQuadratic(str)
		fmt.Println(str, maxstr, str[maxind:maxind+len(maxstr)])
	}
}
