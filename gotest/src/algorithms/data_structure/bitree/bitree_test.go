package bitree

import (
	"fmt"
	"testing"
)

func makeTree() *BiTree {
	bt := NewBiTree()
	bt.Set(6, "a")
	bt.Set(3, "b")
	bt.Set(7, "c")
	bt.Set(4, "d")
	bt.Set(9, "e")
	bt.Set(8, "f")
	bt.Set(1, "g")
	bt.Set(5, "h")
	return bt
}

func TestBitree(t *testing.T) {
	bt := NewBiTree()
	bt.Set(6, "a")
	bt.Set(3, "b")
	bt.Set(7, "c")
	bt.Set(4, "d")
	bt.Set(9, "e")
	bt.Set(8, "f")
	bt.Set(1, "g")
	bt.Set(5, "h")
	fmt.Println(bt)

	fmt.Println(bt.Get(1))
	fmt.Println(bt.Get(2))
	fmt.Println(bt.Get(3))
	fmt.Println(bt.Get(4))
	fmt.Println(bt.Get(5))
	fmt.Println(bt.Get(6))

	bt.Set(4, "x")
	bt.Set(6, "y")
	bt.Set(8, "z")
	fmt.Println(bt)

	bt.Traverse(func(node *BiTreeNode) {
		fmt.Println(node.key, node.value)
	})

	bt.Traverse(func(node *BiTreeNode) {
		predecessor := node.predecessor()
		successor := node.successor()
		fmt.Printf("key: %d, value:%s, ", node.key, node.value)
		if predecessor != nil {
			fmt.Printf("pre: %d, ", predecessor.key)
		}
		if successor != nil {
			fmt.Printf("suc: %d, ", successor.key)
		}
		fmt.Println()
	})
}

func TestBiTree_Remove(t *testing.T) {
	var bt *BiTree
	bt = makeTree()
	fmt.Println(bt.Remove(1))
	fmt.Println(bt, "\n")

	bt = makeTree()
	fmt.Println(bt.Remove(2))
	fmt.Println(bt, "\n")

	bt = makeTree()
	fmt.Println(bt.Remove(3))
	fmt.Println(bt, "\n")

	bt = makeTree()
	fmt.Println(bt.Remove(4))
	fmt.Println(bt, "\n")

	bt = makeTree()
	fmt.Println(bt.Remove(5))
	fmt.Println(bt, "\n")

	bt = makeTree()
	fmt.Println(bt.Remove(6))
	fmt.Println(bt, "\n")

	bt = makeTree()
	fmt.Println(bt.Remove(7))
	fmt.Println(bt, "\n")

	bt = makeTree()
	fmt.Println(bt.Remove(8))
	fmt.Println(bt, "\n")

	bt = makeTree()
	fmt.Println(bt.Remove(9))
	fmt.Println(bt, "\n")

	bt = makeTree()
	fmt.Println(bt.Remove(4))
	fmt.Println(bt)
	fmt.Println(bt.Remove(6))
	fmt.Println(bt)
	fmt.Println(bt.Remove(7))
	fmt.Println(bt)
	fmt.Println(bt.Remove(5))
	fmt.Println(bt)
	fmt.Println(bt.Remove(3))
	fmt.Println(bt)
	fmt.Println(bt.Remove(1))
	fmt.Println(bt)
	fmt.Println(bt.Remove(8))
	fmt.Println(bt)
	fmt.Println(bt.Remove(7))
	fmt.Println(bt)
}

func TestBiTree_Rotate(t *testing.T) {
	var bt *BiTree
	bt = makeTree()
	bt.rotateLeft(bt.getNode(3))
	fmt.Println("left rotate on 3")
	fmt.Println(bt, "\n")

	bt = makeTree()
	bt.rotateLeft(bt.getNode(7))
	fmt.Println("left rotate on 7")
	fmt.Println(bt, "\n")

	bt = makeTree()
	bt.rotateLeft(bt.getNode(6))
	fmt.Println("left rotate on 6")
	fmt.Println(bt, "\n")

	bt = makeTree()
	bt.rotateRight(bt.getNode(3))
	fmt.Println("right rotate on 3")
	fmt.Println(bt, "\n")

	bt = makeTree()
	bt.rotateRight(bt.getNode(6))
	fmt.Println("right rotate on 6")
	fmt.Println(bt, "\n")

	bt = makeTree()
	bt.rotateRight(bt.getNode(9))
	fmt.Println("right rotate on 9")
	fmt.Println(bt, "\n")
}

func TestBiTree_getNode(t *testing.T) {
	var bt *BiTree
	bt = makeTree()
	bt.Traverse(func(node *BiTreeNode) {
		fmt.Printf("%p \n", bt.getNode(node.key))
		fmt.Printf("%p, key: %d, value:%s, ", node, node.key, node.value)
		fmt.Println()
	})
}
