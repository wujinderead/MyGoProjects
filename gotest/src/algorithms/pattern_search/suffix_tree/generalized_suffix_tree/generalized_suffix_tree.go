package generalized_suffix_tree

type SuffixTreeNode struct {
	sibling     *SuffixTreeNode
	children    *SuffixTreeNode
	suffixLink  *SuffixTreeNode
	textindex   int // this node represent texts[textindex][start:end+1]
	start       int
	end         *int // use pointer to make all ends expand with O(1) time
	suffixIndex []int
}

type SuffixTree struct {
	Root    *SuffixTreeNode
	Texts   []string
	Runes   [][]rune
	indices [][]int
	stack   *stack
}

func NewGeneralizedSuffixTree(texts []string) *SuffixTree {
	tree := new(SuffixTree)
	tree.Root = new(SuffixTreeNode)
	tree.Texts = texts
	tree.Runes = make([][]rune, len(texts))
	tree.indices = make([][]int, len(texts))
	tree.stack = newStack()
	for i := range texts {
		addStrToTree(tree, texts, i)
	}
	return tree
}

func addStrToTree(tree *SuffixTree, texts []string, index int) {
	root := tree.Root
	activeNode := root
	activeEdgeIndex := -1
	activeLength := 0
	end := new(int)
	remainCount := 0
	runes, indices := toRunes(texts[index])
	tree.Runes[index] = runes
	tree.indices[index] = indices
	num := len(texts)

	for i := 0; i < len(runes); i++ { // i, current rune index
		// get current character
		curRune := runes[i]

		*end = i                                  // increment all end
		remainCount++                             // increment remain suffix count and loop
		var preInternalNode *SuffixTreeNode = nil // for suffix link

		for remainCount > 0 {
			// walk down to the active point
			// mistake 4
			for activeLength > 0 {
				// if active length>0, there must be an edge
				curEdge := activeNode.getChildForRune(tree.Runes, runes[activeEdgeIndex])
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
				edge := activeNode.getChildForRune(tree.Runes, curRune)
				if edge == nil { // use curByte for last phase
					// active edge not present, create a new edge
					newNode := NewNode(num)
					newNode.start = i
					newNode.end = end
					newNode.textindex = index
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
					// terminal edge already present
					if curRune == 0 {
						edge.suffixIndex[index] = -2 // mark as -2 to indicate a new terminal
						remainCount--
						continue
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
				}
			} else {
				// if active length>0, check whether current runes character present after active point
				curEdge := activeNode.getChildForRune(tree.Runes, runes[activeEdgeIndex])
				activePointIndex := curEdge.start + activeLength - 1
				if tree.Runes[curEdge.textindex][activePointIndex+1] == curRune {
					// for single string, this situation (last rune is already present) won't happen.
					// for multiple strings, this is possible. when it occurs, we need to mark the
					// suffix index for current string as -2, to indicate that it is the terminal
					// of current string. afterwards, decrease remain count, because we actually have
					// added a suffix. then continue to add other remaining suffixes.
					if curRune == 0 {
						curEdge.suffixIndex[index] = -2 // mark as -2 to indicate a new terminal
						remainCount--
					} else {
						// current runes character present after active point,
						// suffix won't be added explicitly in current phase
						// increment active length and exit current phase
						activeLength++
						break
					}
				} else {
					// mistake 2
					// split current edge, the trick here is:
					// use current edge to contain remain characters,
					// create a new node as current node's father
					newNode := NewNode(num)
					newNode.start = curEdge.start
					newNode.end = new(int)
					newNode.textindex = curEdge.textindex
					*newNode.end = activePointIndex
					curEdge.start = activePointIndex + 1 // modify start and end
					// newNode become activeNode's child
					activeNode.setChild(curEdge, newNode)
					curEdge.sibling = nil // unlink curEdge's sibling
					// curEdge (the edge to split) become newNode's child
					newNode.addChild(curEdge)

					// create new leaf for current runes character
					newLeaf := NewNode(num)
					newLeaf.start = i
					newLeaf.end = end // leaf end equals to global end
					newLeaf.textindex = index
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
	dfsToSetSuffixIndex(tree, index, end)
}

func NewNode(n int) *SuffixTreeNode {
	node := new(SuffixTreeNode)
	node.suffixIndex = make([]int, n)
	for i := 0; i < n; i++ {
		node.suffixIndex[i] = -1 // -1 means not a leaf
	}
	return node
}

func (node *SuffixTreeNode) getChildForRune(runes [][]rune, r rune) *SuffixTreeNode {
	for n := node.children; n != nil; n = n.sibling {
		ti := n.textindex
		if runes[ti][n.start] == r {
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

// use curend to specify whether a leaf is added for current str
func dfsToSetSuffixIndex(tree *SuffixTree, index int, end *int) {
	curLen := 0
	root := tree.Root
	tree.stack.reinit()
	cur := root.children // root do not represent start or end, so start with first child
	for cur != nil {
		if cur.children != nil {
			curLen += *cur.end - cur.start + 1
			tree.stack.push(cur)
			cur = cur.children
		} else {
			// check if this suffix is added for current string
			curLen += *cur.end - cur.start
			if cur.end == end || cur.suffixIndex[index] == -2 {
				cur.suffixIndex[index] = len(tree.Runes[index]) - 1 - curLen
			}
			curLen -= *cur.end - cur.start
			for tree.stack.len() > 0 && cur.sibling == nil {
				cur = tree.stack.pop()
				curLen -= *cur.end - cur.start + 1
			}
			cur = cur.sibling
		}
	}
}

func toRunes(str string) (runes []rune, indices []int) {
	n := 0
	for range str {
		n++
	}
	runes = make([]rune, n+1)
	indices = make([]int, n+1)
	n = 0
	for i, v := range str {
		runes[n] = v
		indices[n] = i
		n++
	}
	runes[n] = 0 // set last rune to 0
	indices[n] = len(str)
	return
}

func longestPalindromicSubstring(str string) (start, length int) {
	return 0, 0
}

func longestCommonSubstring(stra, strb string) (astart, bstart, length int) {
	return 0, 0, 0
}

type stack struct {
	top   int
	slice []*SuffixTreeNode
}

func newStack() *stack {
	return &stack{-1, make([]*SuffixTreeNode, 64)}
}

func (s *stack) push(node *SuffixTreeNode) {
	s.top++
	s.slice[s.top] = node
}

func (s *stack) pop() *SuffixTreeNode {
	s.top--
	return s.slice[s.top+1]
}

func (s *stack) peek() *SuffixTreeNode {
	return s.slice[s.top]
}

func (s *stack) len() int {
	return s.top + 1
}

func (s *stack) reinit() {
	s.top = -1
}
