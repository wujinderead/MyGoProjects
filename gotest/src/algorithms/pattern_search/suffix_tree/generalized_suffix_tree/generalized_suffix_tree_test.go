package generalized_suffix_tree

import (
	"fmt"
	"testing"
)

func TestNewGeneralizedSuffixTree(t *testing.T) {
	testIterativeDfsTraverse([]string{"xabxac", "abcabxabcd", "eabxabcxdab"}, t)
	testIterativeDfsTraverse([]string{"xabxaabxa", "babxba", "abx"}, t)
	testIterativeDfsTraverse([]string{"轻轻的我走了", "正如我轻轻的来", "我轻轻的招手", "悄悄的我走了", "正如我悄悄的来"}, t)
	testIterativeDfsTraverse([]string{"GeeksforGeeks", "GeeksQuiz"}, t)
	testIterativeDfsTraverse([]string{"OldSite:GeeksforGeeks.org", "NewSite:GeeksQuiz.com"}, t)
	testIterativeDfsTraverse([]string{"abcde", "fgefg"}, t)
	testIterativeDfsTraverse([]string{"pqrst", "uvwxyz"}, t)
	testIterativeDfsTraverse([]string{"cabbaabb", "bbaabbac"}, t)
	testIterativeDfsTraverse([]string{"forgeeksskeegfor", "rofgeeksskeegrof"}, t)
	testIterativeDfsTraverse([]string{"xababayz", "zyababax"}, t)
	testIterativeDfsTraverse([]string{"xabax", "xabax"}, t)
}

func testIterativeDfsTraverse(texts []string, t *testing.T) {
	tree := NewGeneralizedSuffixTree(texts)

	fmt.Println("===", texts)
	str := make([]rune, getMaxlen(texts)) // to store chars
	appeared := makeAppeared(tree)
	tree.stack.reinit()
	curLen := 0
	cur := tree.Root.children // root do not represent start or end, so start with first child
	for cur != nil {
		if cur.children != nil { // non leaf
			copy(str[curLen:], tree.Runes[cur.textindex][cur.start:*cur.end+1])
			curLen += *cur.end - cur.start + 1
			tree.stack.push(cur)
			cur = cur.children
		} else { // leaf
			copy(str[curLen:], tree.Runes[cur.textindex][cur.start:*cur.end])
			curstr := string(str[:curLen+*cur.end-cur.start])
			fmt.Println(cur.suffixIndex, curstr)
			for ti, si := range cur.suffixIndex {
				if si >= 0 && si < len(tree.Runes[ti])-1 {
					txtstr := string(tree.Runes[ti][si : len(tree.Runes[ti])-1])
					if curstr != txtstr {
						t.Error("suffix index do not equal")
					}
					appeared[ti][si] = si + 1
				}

			}
			// find a non-nil sib
			for tree.stack.len() > 0 && cur.sibling == nil {
				cur = tree.stack.pop()
				curLen -= *cur.end - cur.start + 1
			}
			cur = cur.sibling
		}
	}
	// check appeared
	for i := range appeared {
		for j := range appeared[i] {
			if appeared[i][j] != j+1 {
				t.Error("not appear", i, j)
			}
		}
	}
	fmt.Println()
}

func getMaxlen(texts []string) int {
	max := 0
	for i := range texts {
		if len(texts[i]) > max {
			max = len(texts[i])
		}
	}
	return max
}

func makeAppeared(tree *SuffixTree) [][]int {
	appeared := make([][]int, len(tree.Runes))
	for i := range tree.Runes {
		appeared[i] = make([]int, len(tree.Runes[i])-1)
	}
	return appeared
}

func TestLongestCommon(t *testing.T) {
	strs := [][]string{{"xabxac", "abcabxabcd", "abxa"},
		{"xabxaabxa", "babxba", "abx"},
		{"GeeksforGeeks", "GeeksQuiz", "Geeks"},
		{"OldSite:GeeksforGeeks.org", "NewSite:GeeksQuiz.com", "Site:Geeks"},
		{"abcde", "fgefg", "e"},
		{"pqrst", "uvwxyz", ""}}
	for i := range strs {
		a, b, exp := strs[i][0], strs[i][1], strs[i][2]
		as, bs, length := longestCommonSubstring(a, b)
		if a[as:as+length] != exp || b[bs:bs+length] != exp {
			t.Errorf("error for '%s' and '%s', expect '%s', got '%s'",
				a, b, exp, a[as:as+length])
		}
		fmt.Println(as, bs, length)
		fmt.Println(a, b, exp)
		fmt.Println()
	}
}

func TestLongestPalindromic(t *testing.T) {
	strs := [][]string{{"cabbaabb", "bbaabb"},
		{"forgeeksskeegfor", "geeksskeeg"},
		{"abcde", "a"},
		{"abcdae", "a"},
		{"abacd", "aba"},
		{"abcdc", "cdc"},
		{"abacdfgdcaba", "aba"},
		{"xyabacdfgdcaba", "aba"},
		{"xababayz", "ababa"},
		{"xabax", "xabax"},
	}
	for i := range strs {
		str, exp := strs[i][0], strs[i][1]
		start, length := longestPalindromicSubstring(str)
		if str[start:start+length] != exp {
			t.Errorf("error for '%s', expect '%s', got '%s'",
				str, exp, str[start:start+length])
		}
		fmt.Println(start, length)
		fmt.Println(str, str[start:start+length])
		fmt.Println()
	}
}
