package suffix_tree

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

func NewSuffixTreeUkkonen(text string) *SuffixTree {
	// todo build a suffix tree
	root := new(SuffixTreeNode)
	activeNode := root
	activeEdge := -1
	activeLength := 0
	end := new(int)
	remainingSuffixCount := 0
	for i := 0; i < len(text); i++ {
		cur := text[i]
		remainingSuffixCount++
		for j := 0; j < remainingSuffixCount; j++ {

		}
	}
	return &SuffixTree{Root: root, Text: text}
}
