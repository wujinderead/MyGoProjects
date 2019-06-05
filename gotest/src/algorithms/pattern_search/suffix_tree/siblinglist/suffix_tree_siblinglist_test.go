package siblinglist

import (
	"container/list"
	"fmt"
	"testing"
)

var strs = []string{
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
	"A",
}

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
