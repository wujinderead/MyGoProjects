package bitree

import (
	"fmt"
	"testing"
)

func makeTree1() *BiTree {
	n1 := &BiTreeNode{nil, nil ,1}
	n3 := &BiTreeNode{nil, nil ,3}
	n2 := &BiTreeNode{n1, n3 ,2}
	n5 := &BiTreeNode{nil, nil ,5}
	n7 := &BiTreeNode{nil, nil ,7}
	n6 := &BiTreeNode{n5, n7 ,6}
	n4 := &BiTreeNode{n2, n6 ,4}
	return &BiTree{n4}
}

func TestBiTreeTraverse(t *testing.T) {
	testTraverse(makeTree1())
}

func testTraverse(tree *BiTree) {
	fmt.Println("bfs:")
	tree.TraverseBFS(func(tnode *BiTreeNode) {
		fmt.Print(tnode.key, ", ")
	})
	fmt.Println()

	fmt.Println("preOrder:")
	tree.TraversePreOrder(func(tnode *BiTreeNode) {
		fmt.Print(tnode.key, ", ")
	})
	fmt.Println()

	fmt.Println("preOrder iterative:")
	tree.TraversePreOrderIterative(func(tnode *BiTreeNode) {
		fmt.Print(tnode.key, ", ")
	})
	fmt.Println()

	fmt.Println("inOrder:")
	tree.TraverseInOrder(func(tnode *BiTreeNode) {
		fmt.Print(tnode.key, ", ")
	})
	fmt.Println()

	fmt.Println("inOrder iterative:")
	tree.TraverseInOrderIterative(func(tnode *BiTreeNode) {
		fmt.Print(tnode.key, ", ")
	})
	fmt.Println()

	fmt.Println("postOrder:")
	tree.TraversePostOrder(func(tnode *BiTreeNode) {
		fmt.Print(tnode.key, ", ")
	})
	fmt.Println()

	fmt.Println("postOrder iterative:")
	tree.TraversePostOrderIterative(func(tnode *BiTreeNode) {
		fmt.Print(tnode.key, ", ")
	})
	fmt.Println()
}