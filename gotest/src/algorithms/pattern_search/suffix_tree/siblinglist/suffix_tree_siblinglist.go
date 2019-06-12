package siblinglist

import (
	"container/list"
)

// for large character set for example utf8, it has much more characters than ascii.
// so it's not practical to use array 'children [alphabet_size]*SuffixTreeNode'. instead,
// a variety of data structure can be used, which have different time and space costs.
//                        |  Lookup     |  Insert     |  Travel
// -----------------------+-------------+-------------+------------
//   sibling list         |  O(σ)       |  Θ(1)       |  Θ(1)
//   hashmap              |  Θ(1)       |  Θ(1)       |  O(σ)
//   balanced search tree |  O(logσ)    |  O(logσ)    |  O(1)
//   sorted array         |  O(logσ)    |  O(σ)       |  O(1)
//   hashmap+sibling list |  O(1)       |  O(1)       |  O(1)
//
// base on: https://en.wikipedia.org/wiki/Suffix_tree

type SuffixTreeNode struct {
	sibling     *SuffixTreeNode
	children    *SuffixTreeNode
	suffixLink  *SuffixTreeNode
	start       int
	end         *int // use pointer to make all ends expand with O(1) time
	suffixIndex int
}

type SuffixTree struct {
	Root  *SuffixTreeNode
	Text  string
	Runes []rune
}

func NewSuffixTreeSiblingList(text string) *SuffixTree {
	root := new(SuffixTreeNode)
	activeNode := root
	activeEdgeIndex := -1
	activeLength := 0
	end := new(int)
	remainCount := 0
	runes := []rune(text)
	for i := 0; i <= len(runes); i++ { // i, current index
		// get current character
		var curRune rune
		if i == len(runes) {
			curRune = 0
		} else {
			curRune = runes[i]
		}

		*end = i                                  // increment all end
		remainCount++                             // increment remain suffix count and loop
		var preInternalNode *SuffixTreeNode = nil // for suffix link

		for remainCount > 0 {
			// walk down to the active point
			// mistake 4
			for activeLength > 0 {
				// if active length>0, there must be an edge
				curEdge := activeNode.getChildForRune(runes, runes[activeEdgeIndex])
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
				if activeNode.getChildForRune(runes, curRune) == nil { // use curByte for last phase
					// active edge not present, create a new edge
					newNode := new(SuffixTreeNode)
					newNode.start = i
					newNode.end = end
					activeNode.addChild(newNode)
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
				// if active length>0, check whether current runes character present after active point
				curEdge := activeNode.getChildForRune(runes, runes[activeEdgeIndex])
				activePointIndex := curEdge.start + activeLength - 1
				if runes[activePointIndex+1] == curRune {
					// current runes character present after active point,
					// suffix won't be added explicitly in current phase
					// increment active length and exit current phase
					activeLength++
					break
				} else {
					// mistake 2
					// split current edge, the trick here is:
					// use current edge to contain remain characters,
					// create a new node as current node's father
					newNode := new(SuffixTreeNode)
					newNode.start = curEdge.start
					newNode.end = new(int)
					*newNode.end = activePointIndex
					curEdge.start = activePointIndex + 1 // modify start and end
					// newNode become activeNode's child
					activeNode.setChild(curEdge, newNode)
					curEdge.sibling = nil // unlink curEdge's sibling
					// curEdge (the edge to split) become newNode's child
					newNode.addChild(curEdge)

					// create new leaf for current runes character
					newLeaf := new(SuffixTreeNode)
					newLeaf.start = i
					newLeaf.end = end         // leaf end equals to global end
					newNode.addChild(newLeaf) // leaf added to new node

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

	dfsToSetSuffixIndex(root)
	return &SuffixTree{Root: root, Text: text, Runes: runes}
}

func (node *SuffixTreeNode) getChildForRune(runes []rune, r rune) *SuffixTreeNode {
	for n := node.children; n != nil; n = n.sibling {
		if runes[n.start] == r {
			return n
		}
	}
	return nil
}

// change node's child old to new, old must be present
func (node *SuffixTreeNode) setChild(old, new *SuffixTreeNode) {
	if old == node.children {
		new.sibling = old.sibling
		node.children = new
		return
	}
	prev := node.children
	cur := prev.sibling
	for cur != old {
		prev = cur
		cur = cur.sibling
	}
	prev.sibling = new
	new.sibling = old.sibling
}

func (node *SuffixTreeNode) addChild(child *SuffixTreeNode) {
	if node.children == nil {
		node.children = child
	} else {
		child.sibling = node.children // add as first child
		node.children = child
	}
}

func dfsToSetSuffixIndex(root *SuffixTreeNode) {
	curLen := 0
	stack := list.New()
	for n := root.children; n != nil; n = n.sibling { // first add root's children
		stack.PushBack(n)
	}
	visited := make(map[*SuffixTreeNode]struct{})
	for stack.Len() > 0 {
		cur := stack.Back().Value.(*SuffixTreeNode)
		if _, ok := visited[cur]; !ok { // not visited, peek and add children
			visited[cur] = struct{}{}
			curLen += *cur.end - cur.start + 1
			if cur.children != nil { // internal node
				for n := cur.children; n != nil; n = n.sibling {
					stack.PushBack(n)
				}
			} else { // leaf
				cur.suffixIndex = *cur.end - curLen + 1
			}
		} else { // visited, pop
			stack.Remove(stack.Back())
			curLen -= *cur.end - cur.start + 1
		}
	}
}

func (tree *SuffixTree) containSubstring(sub string) bool {
	return true
}

func (tree *SuffixTree) findAllSubstring(sub string) []int {
	subs := make([]int, 0)
	return subs
}

func (tree *SuffixTree) longestRepeatedSubstring() (astart, bstart, lenght int) {
	return 0, 0, 0
}

func (tree *SuffixTree) longestPalindromicSubstring() (start, lenght int) {
	return 0, 0
}

func longestCommonSubstring(stra, strb string) (astart, bstart, length int) {
	return 0, 0, 0
}
