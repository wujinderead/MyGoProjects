package rabin_karp

// hash pattern and txt, only compare when hash equals
func search(txt, pattern string) []int {
	matched := make([]int, 0)
	N, M := len(txt), len(pattern)
	if N < M {
		return matched // txt shorter than pattern, no match
	}

	// set number base d=256, then multiply d can computed by left shift 8 bits
	var i, j int
	sf := uint(8)
	h := 1
	q := 211       // a prime number for modular power
	th, ph := 0, 0 // txt hash, pattern hash
	for i := 0; i < M-1; i++ {
		h = (h << sf) % q // h = (h*d) mod q
	} // finally, h = d^(M-1) mod q

	for i = 0; i < M; i++ { // get has of pattern and txt[0...M-1]
		th = ((th << sf) + int(txt[i])) % q
		ph = ((ph << sf) + int(pattern[i])) % q
	}

	for i = 0; i <= N-M; i++ {
		if ph == th { // only compare when hash equals
			for j = 0; j < M; j++ {
				if pattern[j] != txt[i+j] {
					break
				}
			}
			if j == M { // found a match
				matched = append(matched, i)
			}
		}
		// update hash
		if i < N-M {
			th = (((th - int(txt[i])*h) << sf) + int(txt[i+M])) % q
			if th < 0 {
				th = th + q
			}
		}
	}
	return matched
}

func searchRunes(txt, pattern string) []int {
	return searchRunesHelper([]rune(txt), []rune(pattern))
}

func searchRunesHelper(txt, pattern []rune) []int {
	matched := make([]int, 0)
	N, M := len(txt), len(pattern)
	if N < M {
		return matched // txt shorter than pattern, no match
	}

	// set number base d=256, then multiply d can computed by left shift 8 bits
	var i, j int
	sf := uint(8)
	h := 1
	q := 211       // a prime number for modular power
	th, ph := 0, 0 // txt hash, pattern hash
	for i := 0; i < M-1; i++ {
		h = (h << sf) % q // h = (h*d) mod q
	} // finally, h = d^(M-1) mod q

	for i = 0; i < M; i++ { // get has of pattern and txt[0...M-1]
		th = ((th << sf) + int(txt[i])) % q
		ph = ((ph << sf) + int(pattern[i])) % q
	}

	for i = 0; i <= N-M; i++ {
		if ph == th { // only compare when hash equals
			for j = 0; j < M; j++ {
				if pattern[j] != txt[i+j] {
					break
				}
			}
			if j == M { // found a match
				matched = append(matched, i)
			}
		}
		// update hash
		if i < N-M {
			th = (((th - int(txt[i])*h) << sf) + int(txt[i+M])) % q
			if th < 0 {
				th = th + q
			}
		}
	}
	return matched
}
