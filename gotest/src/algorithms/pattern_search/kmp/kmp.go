package kmp

// search pattern in txt and return start positions if match
func search(txt, pattern string) []int {
	matched := make([]int, 0)
	if len(txt) < len(pattern) {
		return matched
	}
	lps := getLongestProperPrefixAlsoSuffixBytes(pattern)
	i, j := 0, 0
	lenp := len(pattern)
	for i < len(txt) {
		if txt[i] == pattern[j] {
			i++
			j++
			if j == lenp { // has a match
				matched = append(matched, i-lenp) // append this match with start index in txt
				j = lps[j-1]                      // j reset to lps[j-1], i.e., we can skip lps[j-1] chars in next compare
			}
		} else {
			if j > 0 {
				j = lps[j-1] // reduce j is actually move the pattern forward
			} else { // j==0, move txt forward
				i++
			}
		}
	}
	return matched
}

func getLongestProperPrefixAlsoSuffixBytes(pattern string) []int {
	lps := make([]int, len(pattern))
	lps[0] = 0
	i := 1
	length := 0
	for i < len(pattern) {
		if pattern[i] == pattern[length] {
			length++
			lps[i] = length
			i++
		} else { // use a tricky way
			if length > 0 {
				// set length to lps[length-1], we check if there is symmetric in the prefix
				length = lps[length-1] // no i++ here, continue with new length
			} else {
				lps[i] = 0
				i++
			}
		}
	}
	return lps
}

func searchRunes(txt, pattern []rune) []int {
	matched := make([]int, 0)
	if len(txt) < len(pattern) {
		return matched
	}
	lps := getLongestProperPrefixAlsoSuffix(pattern)
	i, j := 0, 0
	lenp := len(pattern)
	for i < len(txt) {
		if txt[i] == pattern[j] {
			i++
			j++
			if j == lenp { // has a match
				matched = append(matched, i-lenp) // append this match with start index in txt
				j = lps[j-1]                      // j reset to lps[j-1], i.e., we can skip lps[j-1] chars in next compare
			}
		} else {
			if j > 0 {
				j = lps[j-1] // reduce j is actually move the pattern forward
			} else { // j==0, move txt forward
				i++
			}
		}
	}
	return matched
}

func getLongestProperPrefixAlsoSuffix(pattern []rune) []int {
	lps := make([]int, len(pattern))
	lps[0] = 0
	i := 1
	length := 0
	for i < len(pattern) {
		if pattern[i] == pattern[length] {
			length++
			lps[i] = length
			i++
		} else { // use a tricky way
			if length > 0 {
				// set length to lps[length-1], we check if there is symmetric in the prefix
				length = lps[length-1] // no i++ here, continue with new length
			} else {
				lps[i] = 0
				i++
			}
		}
	}
	return lps
}
