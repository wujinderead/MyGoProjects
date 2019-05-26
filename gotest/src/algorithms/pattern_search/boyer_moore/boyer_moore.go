package boyer_moore

const alphabets = 256

func search(txt, pattern string) []int {
	matched := make([]int, 0)
	lastOccur := make([]int, alphabets)
	for i := 0; i < len(pattern); i++ {
		lastOccur[pattern[i]] = i + 1 // get last occur position of a char in pattern
	}
	N, M := len(txt), len(pattern)
	i, j := M-1, M-1
	for i < N {
		if txt[i] == pattern[j] {
			if j == 0 { // find a match
				matched = append(matched, i)
				if i+M < N && lastOccur[txt[i+M]] != 0 {
					i += M + (M - lastOccur[txt[i+M]])
				} else {
					i += M
				}
				j = M - 1
			} else {
				i--
				j--
			}
		} else {
			occur := lastOccur[txt[i]] - 1
			if occur == -1 {
				i += M
				j = M - 1
			} else {
				i += j - occur
				j = M - 1
			}
		}
	}
	return matched
}
