package suffix_tree

import (
	"container/list"
	"fmt"
)

const alphabet = 256

type SuffixTreeNode struct {
	children    [alphabet]*SuffixTreeNode
	suffixLink  *SuffixTreeNode
	start       int
	end         *int // use pointer to make all ends expand with O(1) time
	suffixIndex int
}

type SuffixTree struct {
	Root *SuffixTreeNode
	Text string
}

func NewSuffixTreeNode() *SuffixTreeNode {
	node := new(SuffixTreeNode)
	node.suffixIndex = -1
	node.end = new(int)
	node.start = -1
	*node.end = -1
	return node
}

// mistake I have made:
// 1. when process suffix link after AddLeaf, need to set preInternalNode to nil.
//    e.g, preInternalNode=A, activeNode=B, add a leaf in B, then suffixLink, A->B.
//    next turn, it may be preInternalNode=A, activeNode=root,
//    if not set preInternalNode to nil, suffix link will be A->root, which is wrong.
// 2. the node to split can be a leaf or internal node, not just leaf.
//    so the trick is to create a new node as old node's father and change links.
// 3. the last suffix to add must be a leaf of root. the previous suffixes are added
//    not only by split, but also by AddLeaf. after add leaf in middle, also need to
//    find next active point by suffix link
// 4. a loop to the active point is needed
// 5. if activeLength == 0 && activeNode.children[curByte]!=nil, it's the end of
//    current phase, so need to link suffix link for previous internal node.
func NewSuffixTreeUkkonen(text string) *SuffixTree {
	root := NewSuffixTreeNode()
	activeNode := root
	activeEdgeIndex := -1
	activeLength := 0
	end := new(int)
	remainCount := 0
	for i := 0; i <= len(text); i++ { // i, current index
		// get current character
		var curByte byte
		if i == len(text) {
			curByte = 0
		} else {
			curByte = text[i]
		}

		*end = i                                  // increment all end
		remainCount++                             // increment remain suffix count and loop
		var preInternalNode *SuffixTreeNode = nil // for suffix link

		for remainCount > 0 {
			// walk down to the active point
			// mistake 4
			for activeLength > 0 {
				// if active length>0, there must be an edge
				curEdge := activeNode.children[text[activeEdgeIndex]]
				curEdgeLen := *curEdge.end - curEdge.start + 1
				// if active len == edge len (can only be equal)
				// walk down to reset active node
				if activeLength >= curEdgeLen {
					activeLength -= curEdgeLen
					activeEdgeIndex += curEdgeLen
					activeNode = curEdge
				} else {
					break
				}
			}

			// active length=0, check if need to add a new leaf
			// active length>0, check if need to split an edge and make an new internal node
			if activeLength == 0 {
				// if active length==0, set active edge to current character
				activeEdgeIndex = i

				// check if active edge going out of active node
				if activeNode.children[curByte] == nil { // use curByte for last phase
					// active edge not present, create a new edge
					newNode := NewSuffixTreeNode()
					newNode.start = i
					newNode.end = end
					activeNode.children[curByte] = newNode
					// create new node (add new suffix), need to decrease remainCount
					remainCount--

					// if previous internal node not null, link suffix link
					if preInternalNode != nil {
						preInternalNode.suffixLink = activeNode
						// set preInternalNode to null to prevent unexpected link
						// mistake 1
						preInternalNode = nil
					}
				} else {
					// active edge present, suffix won't be added explicitly in current phase
					// increment active length and exit current phase
					activeLength = 1
					// mistake 5
					if preInternalNode != nil {
						preInternalNode.suffixLink = activeNode
					}
					break
				}
			} else {
				// if active length>0, check whether current text character present after active point
				curEdge := activeNode.children[text[activeEdgeIndex]]
				activePointIndex := curEdge.start + activeLength - 1
				if text[activePointIndex+1] == curByte {
					// current text character present after active point,
					// suffix won't be added explicitly in current phase
					// increment active length and exit current phase
					activeLength++
					break
				} else {
					// mistake 2
					// split current edge, the trick here is:
					// use current edge to contain remain characters,
					// create a new node as current node's father
					newNode := NewSuffixTreeNode()
					newNode.start = curEdge.start
					*newNode.end = activePointIndex
					curEdge.start = activePointIndex + 1 // modify start and end
					// newNode become activeNode's child
					activeNode.children[text[activeEdgeIndex]] = newNode
					// curEdge (the edge to split) become newNode's child
					newNode.children[text[activePointIndex+1]] = curEdge

					// create new leaf for current text character
					newLeaf := NewSuffixTreeNode()
					newLeaf.start = i
					newLeaf.end = end                   // leaf end equals to global end
					newNode.children[curByte] = newLeaf // leaf added to new node

					// if previous internal node not null, link suffix link
					if preInternalNode != nil {
						preInternalNode.suffixLink = newNode
					}
					preInternalNode = newNode

					// add new suffix, need to decrease remainCount
					remainCount--
				}
			}

			// find next active point
			if activeNode == root && activeLength > 0 {
				// if active node is root, next active node is still root
				// just change active edge and decrement active length
				// if activeLength=0, it must be the last added suffix,
				// i.e., remainCount=0, thus no need to find next
				activeEdgeIndex = i - remainCount + 1
				activeLength--
			} else if activeNode != root {
				// mistake 3
				// if active node is internal node, use suffix link
				// no need to decrement active length when use suffix link
				activeNode = activeNode.suffixLink
			}
		}
	}

	dfsToSetSuffixLength(root, root, 0)
	return &SuffixTree{Root: root, Text: text}
}

func dfsToSetSuffixLength(root, node *SuffixTreeNode, len int) {
	var curLen int
	if node == root {
		curLen = 0
	} else {
		curLen = *node.end - node.start + 1
	}
	isLeaf := true
	for i := 0; i < alphabet; i++ {
		if node.children[i] != nil {
			isLeaf = false
			dfsToSetSuffixLength(root, node.children[i], len+curLen)
		}
	}
	if isLeaf {
		node.suffixIndex = *node.end - (len + curLen) + 1
	}
}

func (tree *SuffixTree) DfsTraversal(f func(*SuffixTreeNode)) {
	tree.dfsTraversal(tree.Root, f)
}

func (tree *SuffixTree) dfsTraversal(node *SuffixTreeNode, f func(*SuffixTreeNode)) {
	f(node)
	for i := 0; i < alphabet; i++ {
		if node.children[i] != nil {
			tree.dfsTraversal(node.children[i], f)
		}
	}
}

func (tree *SuffixTree) BfsTraversal(f func(*SuffixTreeNode)) {
	stack := list.New()
	stack.PushBack(tree.Root)
	for stack.Len() > 0 {
		node := stack.Remove(stack.Front()).(*SuffixTreeNode)
		f(node)
		for i := 0; i < alphabet; i++ {
			if node.children[i] != nil {
				stack.PushBack(node.children[i])
			}
		}
	}
}

func (node *SuffixTreeNode) String() string {
	return fmt.Sprintf("(%d,%d)", node.start, *node.end)
}
