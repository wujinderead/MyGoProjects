package red_balck_tree

import (
	"fmt"
	"math"
	"math/rand"
	"testing"
	"time"
)

func TestRbTree(t *testing.T) {
	bt := NewRedBlackTree()
	bt.Set(6, "a")
	fmt.Println("after insert 6:\n", bt)
	bt.Set(3, "b")
	fmt.Println("after insert 3:\n", bt)
	bt.Set(7, "c")
	fmt.Println("after insert 7:\n", bt)

	bt = NewRedBlackTree()
	bt.Set(3, "a")
	fmt.Println("after insert 3:\n", bt)
	bt.Set(6, "b")
	fmt.Println("after insert 6:\n", bt)
	bt.Set(7, "c")
	fmt.Println("after insert 7:\n", bt)

	bt = NewRedBlackTree()
	bt.Set(7, "a")
	fmt.Println("after insert 7:\n", bt)
	bt.Set(6, "b")
	fmt.Println("after insert 6:\n", bt)
	bt.Set(3, "c")
	fmt.Println("after insert 3:\n", bt)
}

func TestRbTreeSet(t *testing.T) {
	test := [][]int{{5, 8, 7, 2, 10, 6, 1, 4, 0, 9, 11, 3, 12},
		{8, 7, 0, 2, 1, 3, 12, 4, 9, 11, 5, 6, 10},
		{7, 20, 0, 17, 16, 4, 1, 2, 13, 11, 15, 3, 6, 5, 12, 8, 14, 19, 9, 10, 18,
			28, 24, 22, 21, 26, 27, 23, 25},
		{16, 20, 13, 5, 10, 2, 18, 8, 6, 11, 14, 4, 9, 15, 7, 3, 19, 17, 0, 12, 1}}
	for i, arr := range test {
		rbt := NewRedBlackTree()
		fmt.Printf("\ntest case %d:\n", i)
		for _, v := range arr {
			rbt.Set(v, "")
			fmt.Printf("after insert %d:\n", v)
			rbt.Print()
		}
	}
}

func TestRbTreeRemove(t *testing.T) {
	test := []int{7, 20, 0, 17, 16, 4, 1, 2, 13, 11, 15, 3, 6, 5, 12, 8, 14, 19, 9, 10, 18,
		28, 24, 22, 21}
	toDelete := []int{17, 1, 18, 15, 6, 7, 10, 3, 14, 2, 8, 0, 19, 11, 16}
	rbt := NewRedBlackTree()
	for _, v := range test {
		rbt.Set(v, "")
	}
	fmt.Println("original tree: ")
	rbt.Print()
	for _, v := range toDelete {
		fmt.Printf("\ndelete %d:\n", v)
		rbt.Remove(v)
	}
}

func TestRedBlackTreePrint(t *testing.T) {
	test1 := [][]int{{1}, {1, 2}, {1, 2, 3}, {1, 2, 3, 4, 5},
		{8, 7, 0, 2, 1, 3, 12, 4, 9, 11, 5, 6, 10},
		{16, 20, 13, 5, 10, 2, 18, 8, 6, 11, 14, 4, 9, 15, 7, 3, 19, 17, 0, 12, 1},
		{7, 20, 0, 17, 16, 4, 1, 2, 13, 11, 15, 3, 6, 5, 12, 8, 14, 19, 9, 10, 18,
			28, 24, 22, 21, 26, 27, 23, 25}}
	for _, arr := range test1 {
		rbt := NewRedBlackTree()
		for _, v := range arr {
			rbt.Set(v, "")
		}
		fmt.Println("original tree: \n", rbt)
		rbt.Print()
	}
}

func TestRedBlackTreeIterator(t *testing.T) {
	test := [][]int{{5, 8, 7, 2, 10, 6, 1, 4, 0, 9, 11, 3, 12},
		{8, 7, 0, 2, 1, 3, 12, 4, 9, 11, 5, 6, 10},
		{7, 20, 0, 17, 16, 4, 1, 2, 13, 11, 15, 3, 6, 5, 12, 8, 14, 19, 9, 10, 18,
			28, 24, 22, 21, 26, 27, 23, 25},
		{16, 20, 13, 5, 10, 2, 18, 8, 6, 11, 14, 4, 9, 15, 7, 3, 19, 17, 0, 12, 1}}
	for i, arr := range test {
		rbt := NewRedBlackTree()
		fmt.Printf("\ntest case %d:\n", i)
		for _, v := range arr {
			rbt.Set(v, "")
		}
		rbt.Print()
		it := rbt.Iter()
		for it.HasNext() {
			fmt.Print(it.Next().key, " ")
		}
		fmt.Println()
		it = rbt.ReverseIter()
		for it.HasPrev() {
			fmt.Print(it.Prev().key, " ")
		}
		fmt.Println()
	}
}

func TestRedBlackTreeHeight(t *testing.T) {
	n := (1 << 20) - 1
	rander := rand.New(rand.NewSource(time.Now().UnixNano()))
	ints := rander.Perm(n)
	tree := NewRedBlackTree()
	for i := range ints {
		tree.Set(ints[i], nil)
	}
	left := make(map[*RedBlackNode]int)
	right := make(map[*RedBlackNode]int)
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

var max = func(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func getHeight(t *RedBlackNode, left, right map[*RedBlackNode]int) int {
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
