package boyer_moore

const alphabets = 256

func search1(txt, pattern string) []int {
	return boyerMooreBadCharacter(txt, pattern)
}

func boyerMooreBadCharacter(txt, pattern string) []int {
	matched := make([]int, 0)
	lastOccur := make([]int, alphabets)
	for i := 0; i < alphabets; i++ {
		lastOccur[i] = -1 // initialized as -1; actually no need, but for concise
	}
	for i := 0; i < len(pattern); i++ {
		lastOccur[pattern[i]] = i // get last occur position of a char in pattern
	}
	N, M := len(txt), len(pattern)
	s := 0
	for s <= N-M {
		j := M - 1
		for j >= 0 && txt[s+j] == pattern[j] {
			j--
		}
		if j < 0 { // find a match
			matched = append(matched, s)
			if s+M < N {
				s += M - lastOccur[txt[s+M]]
			} else {
				s += 1
			}
		} else {
			s += max(1, j-lastOccur[txt[s+j]])
		}

	}
	return matched
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
