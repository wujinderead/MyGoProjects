package suffix_tree

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
}

func TestNewSuffixTreeUkkonen(t *testing.T) {
	str := "CCAAACCCGATTA"
	tree := NewSuffixTreeUkkonen(str)
	tree.DfsTraversal(func(node *SuffixTreeNode) {
		if node == tree.Root {
			fmt.Printf("root: %p\n", node)
			return
		}
		end := *node.end
		leaf := false
		if *node.end == len(tree.Text) {
			end--
			leaf = true
		}
		fmt.Printf("p :%p edge: %s", node, tree.Text[node.start:end+1])
		if leaf {
			fmt.Print("$, suffixIndex: ", node.suffixIndex)
		}
		if node.suffixLink != nil {
			fmt.Printf(", suffixLink: %p", node.suffixLink)
		}
		fmt.Println()
	})
	testIterativeDfsTraverse(str, t)
}

// empty string, root still has a child represent $
func TestEmptyString(t *testing.T) {
	tree := NewSuffixTreeUkkonen("")
	fmt.Println(tree.Root.children[0].start)
	fmt.Println(*tree.Root.children[0].end)
	fmt.Println(tree.Root.children[0].children)
}

func TestIterativeDfsTraverse(t *testing.T) {
	for i := range strs {
		testIterativeDfsTraverse(strs[i], t)
	}
	str := "CCAAACCCGATTA"
	for i := 1; i <= len(str); i++ {
		testIterativeDfsTraverse(str[:i], t)
	}
}

// dfs to traverse the suffix tree makes the suffixes sorted naturally
func testIterativeDfsTraverse(text string, t *testing.T) {
	tree := NewSuffixTreeUkkonen(text)

	fmt.Println("===", text)
	appeared := make([]int, len(tree.Text))
	str := make([]byte, len(tree.Text))
	curLen := 0
	stack := list.New()
	for i := alphabet - 1; i >= 0; i-- {
		if tree.Root.children[i] != nil {
			stack.PushBack(tree.Root.children[i])
		}
	}
	visited := make(map[*SuffixTreeNode]struct{})
	for stack.Len() > 0 {
		cur := stack.Back().Value.(*SuffixTreeNode)
		if _, ok := visited[cur]; !ok { // not visited, peek and add children
			visited[cur] = struct{}{}
			if cur.suffixIndex == -1 { // non leaf
				for i := alphabet - 1; i >= 0; i-- {
					if cur.children[i] != nil {
						stack.PushBack(cur.children[i])
					}
				}
				copy(str[curLen:], tree.Text[cur.start:*cur.end+1])
				curLen += *cur.end - cur.start + 1
			} else if cur.suffixIndex != len(tree.Text) { // leaf
				copy(str[curLen:], tree.Text[cur.start:*cur.end])
				curLen += *cur.end - cur.start
				fmt.Println(cur.suffixIndex, string(str[:curLen]))
				if tree.Text[cur.suffixIndex:] != string(str[:curLen]) {
					t.Fatal("suffix index do not equal")
				}
				appeared[cur.suffixIndex] = cur.suffixIndex + 1
			}
		} else { // visited, pop
			stack.Remove(stack.Back())
			if cur.suffixIndex == -1 { // non leaf
				curLen -= *cur.end - cur.start + 1
			} else if cur.suffixIndex != len(tree.Text) { // leaf
				curLen -= *cur.end - cur.start
			}
		}
	}

	for i := 0; i < len(tree.Text); i++ {
		if appeared[i] != i+1 {
			t.Error("suffix index", i, "not appear")
		}
	}
	fmt.Println()
}
