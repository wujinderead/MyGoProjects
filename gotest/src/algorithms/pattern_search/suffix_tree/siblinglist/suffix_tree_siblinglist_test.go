package siblinglist

import (
	"algorithms/pattern_search/kmp"
	"container/list"
	"fmt"
	"strings"
	"testing"
)

var (
	strs = []string{
		"banana",
		"GEEKSFORGEEKS",
		"AAAAAAAAAA",
		"ABCDEFG",
		"ABABABA",
		"abcabxabcd",
		"abc",
		"xabxac",
		"xabxa",
		"THIS IS A TEST TEXT",
		"AABAACAADAABAAABAA",
		"A"}
	txt = `中国人民银行（The People's Bank Of China，英文简称PBOC），简称央行，
是中华人民共和国的中央银行，中华人民共和国国务院组成部门。在国务院领导下，制定和执行货币政策，防范和化解金融风险，
维护金融稳定。1948年12月1日，在华北银行、北海银行、西北农民银行的基础上在河北省石家庄市合并组成中国人民银行。
1983年9月，国务院决定中国人民银行专门行使中国国家中央银行职能。1995年3月18日，第八届全国人民代表大会第三次会议通过
了《中华人民共和国中国人民银行法》，至此，中国人民银行作为中央银行以法律形式被确定下来。[1]中国人民银行根据
《中华人民共和国中国人民银行法》的规定，在国务院的领导下依法独立执行货币政策，
履行职责，开展业务，不受地方政府、社会团体和个人的干涉。`
)

func TestNewSuffixTree(t *testing.T) {
	str := "CCAAACCCGATTA"
	testIterativeDfsTraverse(str, t)
}

// empty string, root still has a child represent $
func TestEmptyString(t *testing.T) {
	tree := NewSuffixTreeSiblingList("")
	fmt.Println(tree.Root.children.start)
	fmt.Println(*tree.Root.children.end)
	fmt.Println(tree.Root.children.children)
	fmt.Println(tree.Root.children.sibling)
}

func TestIterativeDfsTraverse(t *testing.T) {
	for i := range strs {
		testIterativeDfsTraverse(strs[i], t)
	}
	str := []rune("我我的的的我我我天的哪哪的")
	for i := 1; i <= len(str); i++ {
		testIterativeDfsTraverse(string(str[:i]), t)
	}
	testIterativeDfsTraverse("CCAAACCCGATTA", t)
}

func testIterativeDfsTraverse(text string, t *testing.T) {
	tree := NewSuffixTreeSiblingList(text)

	fmt.Println("===", text)
	appeared := make([]int, len(tree.Runes))
	str := make([]rune, len(tree.Runes))
	curLen := 0
	stack := list.New()
	for n := tree.Root.children; n != nil; n = n.sibling { // first add root's children
		stack.PushBack(n)
	}
	visited := make(map[*SuffixTreeNode]struct{})
	for stack.Len() > 0 {
		cur := stack.Back().Value.(*SuffixTreeNode)
		if _, ok := visited[cur]; !ok { // not visited, peek and add children
			visited[cur] = struct{}{}
			if cur.children != nil { // non leaf
				for n := cur.children; n != nil; n = n.sibling { // first add root's children
					stack.PushBack(n)
				}
				copy(str[curLen:], tree.Runes[cur.start:*cur.end+1])
				curLen += *cur.end - cur.start + 1
			} else if cur.suffixIndex != len(tree.Runes) { // leaf
				copy(str[curLen:], tree.Runes[cur.start:*cur.end])
				curLen += *cur.end - cur.start
				fmt.Println(cur.suffixIndex, string(str[:curLen]))
				if string(tree.Runes[cur.suffixIndex:]) != string(str[:curLen]) {
					t.Fatal("suffix index do not equal")
				}
				appeared[cur.suffixIndex] = cur.suffixIndex + 1
			}
		} else { // visited, pop
			stack.Remove(stack.Back())
			if cur.children != nil { // non leaf
				curLen -= *cur.end - cur.start + 1
			} else if cur.suffixIndex != len(tree.Runes) { // leaf
				curLen -= *cur.end - cur.start
			}
		}
	}

	for i := 0; i < len(tree.Runes); i++ {
		if appeared[i] != i+1 {
			t.Error("suffix index", i, "not appear")
		}
	}
	fmt.Println()
}

func TestContainSub(t *testing.T) {
	str := "THIS IS A TEST TEXT"
	tree := NewSuffixTreeSiblingList(str)
	subs := []string{"TEST", "A", " ", "IS A", " IS A ", "TEST1", "THIS IS GOOD", "TES", "TESA", "ISB"}
	for i := range subs {
		sub := subs[i]
		my := tree.containSubstring(sub)
		rt := strings.Contains(str, sub)
		if my != rt {
			t.Errorf("contain '%s'? expect: %v, got %v", sub, rt, my)
		}
	}
}

func TestAllSub(t *testing.T) {
	strs := [][]string{{"GEEKSFORGEEKS", "GEEKS", "GEEK1", "FOR"},
		{"AABAACAADAABAAABAA", "AABA", "AA", "AAE", "ABAA"},
		{"AAAAAAAAA", "AAAA", "AA", "A", "AB"},
		{txt, "中国人民银行"}}
	for i := range strs {
		str := strs[i][0]
		tree := NewSuffixTreeSiblingList(str)
		for j := 1; j < len(strs[i]); j++ {
			sub := strs[i][j]
			occurs := tree.findAllSubstring(sub)
			ref := kmp.Search(str, sub)
			if !equalInts(ref, occurs) {
				t.Errorf("sub error for '%s' '%s', expect: %v, got: %v",
					str, sub, ref, occurs)
			}
			fmt.Println(str)
			for _, v := range occurs {
				fmt.Println(v, ":", str[v:v+len(sub)])
			}
			fmt.Println()
		}
	}
}

func TestLongestRepeated(t *testing.T) {
	strs := [][]string{
		{"GEEKSFORGEEKS", "GEEKS"},
		{"AAAAAAAAAA", "AAAAAAAAA"},
		{"ABCDEFG", nil},
		{"ABABABA", "ABABA"},
		{"你好你好你好你", "你好你好你"},
		{"ATCGATCGA", "ATCGA"},
		{"banana", "ana"},
		{"abcpqrabpqpq", "ab"},
		{"pqrpqpqabab", "ab"}}
	for i := range strs {
		str, exp := strs[i][0], strs[i][1]
		tree := NewSuffixTreeSiblingList(str)
		as, bs, length := tree.longestRepeatedSubstring()
		if str[as:as+length] != exp || str[bs:bs+length] != exp {
			t.Errorf("error for '%s', expect '%s', got '%s'",
				str, exp, str[as:as+length])
		}
		fmt.Println(as, bs, length)
		fmt.Println(str, str[as:as+length], str[bs:bs+length])
		fmt.Println()
	}
}

func TestLongestCommon(t *testing.T) {
	strs := [][]string{{"xabxac", "abcabxabcd", "abxa"},
		{"xabxaabxa", "babxba", "abx"},
		{"GeeksforGeeks", "GeeksQuiz", "Geeks"},
		{"OldSite:GeeksforGeeks.org", "NewSite:GeeksQuiz.com", "Site:Geeks"},
		{"abcde", "fgefg", "e"},
		{"pqrst", "uvwxyz", nil}}
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
		{"", ""}}
	for i := range strs {
		str, exp := strs[i][0], strs[i][1]
		tree := NewSuffixTreeSiblingList(str)
		start, length := tree.longestPalindromicSubstring()
		if str[start:start+length] != exp {
			t.Errorf("error for '%s', expect '%s', got '%s'",
				str, exp, str[start:start+length])
		}
		fmt.Println(start, length)
		fmt.Println(str, str[start:start+length])
		fmt.Println()
	}
}

func equalInts(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
