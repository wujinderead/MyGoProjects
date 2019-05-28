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
}

func NewSuffixTree(text string) *SuffixTree {
	// todo build a suffix tree
	return nil
}
