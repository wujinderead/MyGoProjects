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
	getHeight(t, tree.Root)
	fmt.Println("root left, right height:", tree.Root.left.height, tree.Root.right.height)
}

// export GOPATH=/Users/xh/golang/gotest:/Users/xh/go
// go test -run ./aaa -test.cpuprofile ./cpu.out -v algorithms/data_structure/bitree/avl_tree
// go tool pprof -pdf avl_tree.test cpu.out
func TestAvlTreeRemoveBatch(t *testing.T) {
	n := (1 << 20) - 1
	rander := rand.New(rand.NewSource(time.Now().UnixNano()))
	ints := rander.Perm(n)
	tree := NewAvlTree()
	for i := range ints {
		tree.Set(ints[i], nil)
	}
	getHeight(t, tree.Root)
	for i := 100; i < 200; i++ {
		tree.Remove(ints[i])
		getHeight(t, tree.Root)
	}
	tree.Remove(tree.Root.key)
	getHeight(t, tree.Root)
}

func getHeight(t *testing.T, node *AvlNode) int {
	if node == nil {
		return 0
	} else {
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
		leftHeight := getHeight(t, node.left)
		rightHeight := getHeight(t, node.right)
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
