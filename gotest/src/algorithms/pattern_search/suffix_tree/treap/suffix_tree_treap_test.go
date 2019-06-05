package treap

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

func TestNewSuffixTree(t *testing.T) {
	str := "CCAAACCCGATTA"
	testIterativeDfsTraverse(str, t)
}

// empty string, root still has a child represent $
func TestEmptyString(t *testing.T) {
	tree := NewSuffixTreeTreap("")
	fmt.Println(tree.Root.children.root.value)
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

func TestPBOC(t *testing.T) {
	txt := `中国人民银行（The People's Bank Of China，英文简称PBOC），简称央行，
是中华人民共和国的中央银行，中华人民共和国国务院组成部门。在国务院领导下，制定和执行货币政策，防范和化解金融风险，
维护金融稳定。1948年12月1日，在华北银行、北海银行、西北农民银行的基础上在河北省石家庄市合并组成中国人民银行。
1983年9月，国务院决定中国人民银行专门行使中国国家中央银行职能。1995年3月18日，第八届全国人民代表大会第三次会议通过
了《中华人民共和国中国人民银行法》，至此，中国人民银行作为中央银行以法律形式被确定下来。[1]中国人民银行根据
《中华人民共和国中国人民银行法》的规定，在国务院的领导下依法独立执行货币政策，
履行职责，开展业务，不受地方政府、社会团体和个人的干涉。`
	tree := NewSuffixTreeTreap(txt)
	height := make(map[*treap]string)
	getHeight(tree.Root.children.root, height)
	for _, v := range height {
		fmt.Println(v)
	}
	fmt.Println("root: ", height[tree.Root.children.root])
	queue := list.New()
	queue.PushBack(tree.Root.children.root)
	for queue.Len() > 0 {
		cur := queue.Remove(queue.Front()).(*treap)
		if cur.parent != nil {
			if cur.parent.left == cur && cur.key > cur.parent.key {
				t.Error("bst left violate")
			}
			if cur.parent.right == cur && cur.key < cur.parent.key {
				t.Error("bst right violate")
			}
			if cur.parent.ticket < cur.ticket {
				t.Error("heap violate")
			}
		}
		if cur.left != nil {
			queue.PushBack(cur.left)
		}
		if cur.right != nil {
			queue.PushBack(cur.right)
		}
	}
}

var max = func(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func getHeight(t *treap, mapper map[*treap]string) int {
	if t == nil {
		return 0
	} else {
		left := getHeight(t.left, mapper)
		right := getHeight(t.right, mapper)
		mapper[t] = fmt.Sprintf("%d %d %d", left-right, left, right)
		return 1 + max(left, right)
	}
}

func testIterativeDfsTraverse(text string, t *testing.T) {
	tree := NewSuffixTreeTreap(text)

	fmt.Println("===", text)
	appeared := make([]int, len(tree.Runes))
	str := make([]rune, len(tree.Runes))
	curLen := 0
	stack := list.New()
	it := tree.Root.children.ReverseIter()
	for it.hasNext() {
		stack.PushBack(it.prev().value)
	}
	visited := make(map[*SuffixTreeNode]struct{})
	for stack.Len() > 0 {
		cur := stack.Back().Value.(*SuffixTreeNode)
		if _, ok := visited[cur]; !ok { // not visited, peek and add children
			visited[cur] = struct{}{}
			if cur.children != nil { // non leaf
				it = cur.children.ReverseIter()
				for it.hasNext() {
					stack.PushBack(it.prev().value)
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

func TestLeftRotate(t *testing.T) {
	n := make([]*treap, 8)
	for i := 0; i < 8; i++ {
		n[i] = &treap{key: rune(i)}
	}
	getKey := func(t *treap) rune {
		if t != nil {
			return t.key
		}
		return -1
	}
	tr := &treapRoot{n[7]}
	n[7].right = n[3]
	n[3].parent = n[7]
	n[3].left = n[1]
	n[1].parent = n[3]
	n[3].right = n[5]
	n[5].parent = n[3]
	n[1].left = n[0]
	n[0].parent = n[1]
	n[1].right = n[2]
	n[2].parent = n[1]
	n[5].left = n[4]
	n[4].parent = n[5]
	n[5].right = n[6]
	n[6].parent = n[5]
	for i := 0; i < 8; i++ {
		fmt.Println(n[i].key, getKey(n[i].left), getKey(n[i].right), getKey(n[i].parent))
	}
	fmt.Println()

	tr.rightRotate(n[3])
	for i := 0; i < 8; i++ {
		fmt.Println(n[i].key, getKey(n[i].left), getKey(n[i].right), getKey(n[i].parent))
	}
	fmt.Println()

	tr.leftRotate(n[3])
	for i := 0; i < 8; i++ {
		fmt.Println(n[i].key, getKey(n[i].left), getKey(n[i].right), getKey(n[i].parent))
	}
}
