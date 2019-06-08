package avl_tree

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestAvlTreeSet(t *testing.T) {
	rander := rand.New(rand.NewSource(time.Now().UnixNano()))
	tree := NewAvlTree()
	ints := rander.Perm(10)
	for i := range ints {
		tree.Set(ints[i], nil)
		fmt.Println("set", ints[i], ":")
		tree.Print()
		fmt.Println()
	}
}

func TestAvlTreeRemove(t *testing.T) {
	rander := rand.New(rand.NewSource(time.Now().UnixNano()))
	tree := NewAvlTree()
	ints := rander.Perm(15)
	fmt.Println(ints)
	for i := range ints {
		tree.Set(ints[i], nil)
	}
	count := 0
	getHeight(t, tree.Root, &count)
	tree.Print()
	fmt.Println()
	for i := range ints {
		fmt.Println("remove", ints[i], ":")
		tree.Remove(ints[i])
		tree.Print()
		fmt.Println()
	}
}

func TestAvlTreeHeight(t *testing.T) {
	n := (1 << 20) - 1
	rander := rand.New(rand.NewSource(time.Now().UnixNano()))
	ints := rander.Perm(n)
	tree := NewAvlTree()
	for i := range ints {
		tree.Set(ints[i], nil)
	}
	if n < 100 {
		fmt.Println(tree)
	}
	count := 0
	getHeight(t, tree.Root, &count)
	fmt.Println("root left, right height:", tree.Root.left.height, tree.Root.right.height)
	fmt.Println("node count:", count)
}

// export GOPATH=/Users/xh/golang/gotest:/Users/xh/go
// go test -run ./aaa -test.cpuprofile ./cpu.out -v algorithms/data_structure/bitree/avl_tree
// go tool pprof avl_tree.test cpu.out
func TestAvlTreeRemoveBatch(t *testing.T) {
	n := (1 << 20) - 1
	rander := rand.New(rand.NewSource(time.Now().UnixNano()))
	ints := rander.Perm(n)
	tree := NewAvlTree()
	count := 0
	for i := range ints {
		tree.Set(ints[i], nil)
	}
	getHeight(t, tree.Root, &count)
	fmt.Println(count)
	count = 0
	for i := 100; i < 200; i++ {
		tree.Remove(ints[i])
		getHeight(t, tree.Root, &count)
		fmt.Println(count)
		count = 0
	}
	count = 0
	tree.Remove(tree.Root.key)
	getHeight(t, tree.Root, &count)
	fmt.Println(count)
}

func getHeight(t *testing.T, node *AvlNode, count *int) int {
	if node == nil {
		return 0
	} else {
		*count++
		if node.left != nil {
			if node.left.key >= node.key {
				t.Error("left violation")
			}
		}
		if node.right != nil {
			if node.right.key <= node.key {
				t.Error("right violation")
			}
		}
		leftHeight := getHeight(t, node.left, count)
		rightHeight := getHeight(t, node.right, count)
		if leftHeight-rightHeight > 1 || leftHeight-rightHeight < (-1) {
			t.Error("height diff err")
		}
		height := 1 + max(leftHeight, rightHeight)
		if height != node.height {
			t.Error("height != node.height")
		}
		return height
	}
}
