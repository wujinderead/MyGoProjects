package treap

import (
	"container/list"
	"fmt"
	"math"
	"testing"
)

// export GOROOT=/usr/local/go
// export GOPATH=/home/xzy/golang/gotest:/home/xzy/go
// go test -c -o /tmp/treaper algorithms/data_structure/bitree/treap
// go test -run /tmp/treaper -test.cpuprofile ./cpu.out -v algorithms/data_structure/bitree/treap
// go tool pprof -pdf treap.test cpu.out
func TestTreapHeight(t *testing.T) {
	n := (1 << 20) - 1
	ints := rander.Perm(n)
	tree := NewTreap()
	for i := range ints {
		tree.Set(ints[i], nil)
	}
	validateTreap(t, tree)
	left := make(map[*TreapNode]int)
	right := make(map[*TreapNode]int)
	getHeight(tree.Root, left, right)
	fmt.Println("root left, right:", left[tree.Root], right[tree.Root])
	fmt.Println("left right map len:", len(left), len(right))
	maxdiff := float64(0)
	curdiff := float64(0)
	alldiff := float64(0)
	nonleaf := 0
	for k := range left {
		curdiff = math.Abs(float64(left[k] - right[k]))
		maxdiff = math.Max(maxdiff, curdiff)
		if left[k] == 0 && right[k] == 0 {
			continue // skip leaf
		}
		nonleaf++
		alldiff += curdiff
	}
	fmt.Println("max diff:", maxdiff)
	fmt.Println("nonleaf:", nonleaf, ", avg diff:", alldiff/float64(nonleaf))
}

func TestTreapRemove(t *testing.T) {
	n := (1 << 16) - 1
	ints := rander.Perm(n)
	tree := NewTreap()
	for i := range ints {
		tree.Set(ints[i], nil)
	}
	count := validateTreap(t, tree)
	if count != n {
		t.FailNow()
	}
	for i := 100; i < 200; i++ {
		tree.Remove(ints[i])
		if validateTreap(t, tree) != count-1 {
			t.Error("error remove")
		}
		count--
	}
}

func validateTreap(t *testing.T, tree *Treap) int {
	count := 0
	queue := list.New()
	queue.PushBack(tree.Root)
	for queue.Len() > 0 {
		cur := queue.Remove(queue.Front()).(*TreapNode)
		count++
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
	return count
}

var max = func(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func getHeight(t *TreapNode, left, right map[*TreapNode]int) int {
	if t == nil {
		return 0
	} else {
		leftHeight := getHeight(t.left, left, right)
		rightHeight := getHeight(t.right, left, right)
		left[t] = leftHeight
		right[t] = rightHeight
		return 1 + max(leftHeight, rightHeight)
	}
}
