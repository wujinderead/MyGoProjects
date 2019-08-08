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

func search2(txt, pattern string) []int {
	return boyerMooreGoodSuffix(txt, pattern)
}

func boyerMooreGoodSuffix(txt, pattern string) []int {
	matched := make([]int, 0)
	m := len(pattern)
	n := len(txt)

	shift := make([]int, len(pattern)+1, len(pattern)+1)
	bpos := make([]int, len(pattern)+1, len(pattern)+1)

	//do preprocessing
	preprocessStrongSuffix(shift, bpos, pattern)
	preprocessCase2(shift, bpos, pattern)

	s := 0
	for s <= n-m {

		j := m - 1

		// keep reducing index j of pattern while characters of
		// pattern and text are matching at this shift s
		for j >= 0 && pattern[j] == txt[s+j] {
			j--
		}

		// if the pattern is present at the current shift, then index j
		// will become -1 after the above loop
		if j < 0 {
			matched = append(matched, s)
			s += shift[0]
		} else {
			// pat[i] != pat[s+j] so shift the pattern shift[j+1] times
			s += shift[j+1]
		}
	}
	return matched
}

func preprocessStrongSuffix(shift, bpos []int, pat string) {
	// m is the length of pattern\
	m := len(pat)
	i, j := m, m+1
	bpos[i] = j

	for i > 0 {
		// if character at position i-1 is not equivalent to
		// character at j-1, then continue searching to right
		// of the pattern for border
		for j <= m && pat[i-1] != pat[j-1] {
			// the character preceding the occurrence of t in
			// pattern P is different than the mismatching character in P,
			// we stop skipping the occurrences and shift the pattern
			// from i to j
			if shift[j] == 0 {
				shift[j] = j - i
			}

			//Update the position of next border
			j = bpos[j]
		}
		// p[i-1] matched with p[j-1], border is found.
		// store the  beginning position of border
		i--
		j--
		bpos[i] = j
	}
}

//Preprocessing for case 2
func preprocessCase2(shift, bpos []int, pat string) {
	m := len(pat)
	j := bpos[0]
	for i := 0; i <= m; i++ {
		// set the border position of the first character of the pattern
		// to all indices in array shift having shift[i] = 0
		if shift[i] == 0 {
			shift[i] = j
		}

		// suffix becomes shorter than bpos[0], use the position of
		// next widest border as value of j
		if i == j {
			j = bpos[j]
		}
	}
}
