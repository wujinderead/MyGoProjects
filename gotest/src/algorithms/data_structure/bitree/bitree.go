package bitree

import "container/list"

type BiTreeNode struct {
	left, right *BiTreeNode
	key          int
}

type BiTree struct {
	Root *BiTreeNode
}

func NewBiTree() *BiTree {
	return &BiTree{}
}

func (t *BiTree) TraverseBFS(todo func (tnode *BiTreeNode)) {
	if t.Root == nil {
		return
	}
	queue := list.New()
	queue.PushBack(t.Root)
	for queue.Len()>0 {
		cur := queue.Remove(queue.Front()).(*BiTreeNode)
		todo(cur)
		if cur.left != nil {
			queue.PushBack(cur.left)
		}
		if cur.right != nil {
			queue.PushBack(cur.right)
		}
	}

}

func (t *BiTree) TraversePreOrder(todo func (tnode *BiTreeNode)) {
	preOrder(t.Root, todo)
}

func (t *BiTree) TraverseInOrder(todo func (tnode *BiTreeNode)) {
	inOrder(t.Root, todo)
}

func (t *BiTree) TraversePostOrder(todo func (tnode *BiTreeNode)) {
	postOrder(t.Root, todo)
}

func preOrder(tnode * BiTreeNode, todo func (tnode *BiTreeNode)) {
	if tnode != nil {
		todo(tnode)
		preOrder(tnode.left, todo)
		preOrder(tnode.right, todo)
	}
}

func inOrder(tnode * BiTreeNode, todo func (tnode *BiTreeNode)) {
	if tnode != nil {
		inOrder(tnode.left, todo)
		todo(tnode)
		inOrder(tnode.right, todo)
	}
}

func postOrder(tnode * BiTreeNode, todo func (tnode *BiTreeNode)) {
	if tnode != nil {
		postOrder(tnode.left, todo)
		postOrder(tnode.right, todo)
		todo(tnode)
	}
}

func (t *BiTree) TraversePreOrderIterative(todo func (tnode *BiTreeNode)) {
	// use stack to store parent nodes
	cur := t.Root
	if cur == nil {
		return
	}
	stack := list.New()
	stack.PushBack(cur)
	for stack.Len()>0 {
		cur = stack.Remove(stack.Back()).(*BiTreeNode)
		todo(cur)
		if cur.right != nil {
			stack.PushBack(cur.right)   // push right first, so we can pop left first
		}
		if cur.left != nil {
			stack.PushBack(cur.left)
		}
	}
}

func (t *BiTree) TraverseInOrderIterative(todo func (tnode *BiTreeNode)) {
	cur := t.Root
	stack := list.New()
	for stack.Len()>0 || cur != nil {
		if cur != nil {
			stack.PushBack(cur)
			cur = cur.left
		} else {
			cur = stack.Remove(stack.Back()).(*BiTreeNode)
			todo(cur)
			cur = cur.right
		}
	}
}

func (t *BiTree) TraversePostOrderIterative(todo func (tnode *BiTreeNode)) {
	cur := t.Root
	stack := list.New()
	var lastVisited *BiTreeNode = nil
	for stack.Len()>0 || cur != nil {
		if cur != nil {
			stack.PushBack(cur)
			cur = cur.left
		} else {
			peek := stack.Back().Value.(*BiTreeNode)
			if peek.right != nil && lastVisited != peek.right {
				cur = peek.right
			} else {        // peek.right=nil, we can certainly visit peek
				todo(peek)  // peek.right=lastVisit, we have visit peek.right, we can visit peek
				lastVisited = stack.Remove(stack.Back()).(*BiTreeNode)
			}
		}
	}
}