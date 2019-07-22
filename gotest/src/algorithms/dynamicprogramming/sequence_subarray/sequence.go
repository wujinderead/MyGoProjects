package longest

// find the length of the longest increasing sub sequence of an integer array
func longestIncreasingSubsequence(seq []int) int {
	// for seq[i], check if seq[i] (0≤k<i) > seq[k],
	// if seq[i] > seq[k], seq[i] = max(seq[i], seq[k]+1)
	// time complexity O(n²), space O(n)
	if len(seq) < 0 {
		return 0
	}
	maxlis := 1
	lis := make([]int, len(seq))
	for i := 0; i < len(seq); i++ {
		lis[i] = 1
		for j := 0; j < i; j++ {
			if seq[j] < seq[i] && lis[i] < lis[j]+1 {
				lis[i] = lis[j] + 1
			}
			if maxlis < lis[i] {
				maxlis = lis[i]
			}
		}
	}
	return maxlis
}

// find the longest palindromic sub sequence of a string
func longestPalindromicSubsequence(str string) int {
	// let lps(i, j) denoting the longest palindromic sub sequence in str[i...j]
	// then, if str[i]==str[j], lps(i, j) = lps(i+1, j-1)+2;
	// if not, lps(i, j) = max( lps(i+1, j), lps(i, j-1) )
	// base case is lps(i, i) = 1; lps(i, i+1) = 2 if str[i]==str[j]
	// time complexity O(n²), space O(n²) can be reduced to O(n) since updated diagonally
	if len(str) < 2 {
		return len(str)
	}
	lps := make([][]int, len(str))
	for i := range lps {
		lps[i] = make([]int, len(str))
	}
	// set lps[i][i] and lps[i][i+1]
	for i := 0; i < len(str)-1; i++ {
		lps[i][i] = 1
		lps[i][i+1] = 1
		if str[i] == str[i+1] {
			lps[i][i+1] = 2
		}
	}
	lps[len(str)-1][len(str)-1] = 1
	// dp
	for diff := 2; diff < len(str); diff++ { // update diagonally
		for i := 0; i+diff < len(str); i++ {
			j := i + diff
			if str[i] == str[j] {
				lps[i][j] = lps[i+1][j-1] + 2
			} else {
				lps[i][j] = max(lps[i+1][j], lps[i][j-1])
			}
		}
	}
	return lps[0][len(str)-1]
}

func longestPalindromicSubsequenceSpaceOn(str string) int {
	if len(str) < 2 {
		return len(str)
	}
	lps := make([][]int, 3) // update a row need last 2 rows, so we need 3 rows
	lps[0] = make([]int, len(str))
	lps[1] = make([]int, len(str))
	lps[2] = make([]int, len(str))
	// set lps[0] and lps[1]
	for i := 0; i < len(str)-1; i++ {
		lps[0][i] = 1
		lps[1][i] = 1
		if str[i] == str[i+1] {
			lps[1][i] = 2
		}
	}
	lps[0][len(str)-1] = 1
	// fmt.Println(lps[0])
	// fmt.Println(lps[1][:len(str)-1])

	for i := 2; i < len(str); i++ {
		for j := 0; j < len(str)-i; j++ {
			// lps[i][j] represent str[j...j+i]
			// if str[j]==str[j+i], lps[i][j] = lps[i-2][j+1]+2
			// else lps[i][j] = max(lps[i-1][j], lps[i-1][j+1])
			// the i index need to mod 3
			if str[j] == str[j+i] {
				lps[i%3][j] = lps[(i-2)%3][j+1] + 2
			} else {
				lps[i%3][j] = max(lps[(i-1)%3][j], lps[(i-1)%3][j+1])
			}
		}
		// fmt.Println(lps[i%3][:len(str)-i])
	}
	return lps[(len(str)-1)%3][0]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
