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

// get the longest palindromic subarray of a string
func longestPalindromicSubarray(str string) string {
	// lps(i, j)=true if str[i...j] is a palindromic string
	// then lps(i, j)=(lps(i+1, j-1) && str[i]==str[j])
	// base case lps(i, i)=true, lps(i, i+1)=(str[i]==str[i+1])
	// time complexity O(n²), space O(n²) can be reduced to O(n) since updated diagonally
	if len(str) < 2 {
		return str
	}
	longest := 1
	maxl := 0
	lps := make([][]bool, len(str))
	// base case
	for i := 0; i < len(str)-1; i++ {
		lps[i] = make([]bool, len(str))
		lps[i][i] = true
		lps[i][i+1] = str[i] == str[i+1]
		if lps[i][i+1] {
			longest = 2
			maxl = i
		}
	}
	lps[len(str)-1] = make([]bool, len(str))
	lps[len(str)-1][len(str)-1] = true
	// dp, update diagonally, start from (i, i+2)
	for diff := 2; diff < len(str); diff++ {
		for i := 0; i+diff < len(str); i++ {
			j := i + diff
			lps[i][j] = lps[i+1][j-1] && str[i] == str[j]
			if lps[i][j] && j-i+1 > longest { // if palindromic, update longest
				longest = j - i + 1
				maxl = i
			}
		}
	}
	return str[maxl : maxl+longest]
}

// get the length of the longest common sub sequence of two strings
func longestCommonSubsequence(a, b string) int {
	if len(a) == 0 || len(b) == 0 {
		return 0
	}
	// let lcs(i, j) be the length of longest common sub sequence of a[0...i] and b[0...j].
	// then, if a[i]==b[j], lcs(i, j)=lcs(i-1, j-1)+1; if not, lcs(i, j)=max( lcs(i-1, j), lcs(i, j-1) )
	// base case: lcs(0, 0)=1 if a[0]==b[0]
	// time complexity O(mn), space O(mn) can be reduced to O(min(m,n)) since updated line by line
	if len(b) > len(a) {
		b, a = a, b // let b be shorter
	}
	lcs := make([]int, len(b))
	// set up lcs(0, j)
	if a[0] == b[0] {
		lcs[0] = 1
	}
	for j := 1; j < len(b); j++ {
		if a[0] == b[j] {
			lcs[j] = 1
		} else {
			lcs[j] = lcs[j-1]
		}
	}
	// dp
	for i := 1; i < len(a); i++ {
		tmp := lcs[0] // tmp to store lcs(i-1, j-1)
		if a[i] == b[0] {
			lcs[0] = 1 // if a[i]==b[0], lcs(i, 0)=1; else lcs(i, 0)=lcs(i-1, 0) which is still lcs[0]
		}
		for j := 1; j < len(b); j++ {
			tmp2 := lcs[j]
			if a[i] == b[j] {
				lcs[j] = tmp + 1 // if a[i]==b[j], lcs(i, j)=lcs(i-1, j-1)+1
			} else {
				lcs[j] = max(lcs[j-1], lcs[j]) // else, lcs(i, j)=max( lcs(i-1, j), lcs(i, j-1) )
			}
			tmp = tmp2
		}
	}
	return lcs[len(b)-1]
}

// find the the longest common subarray of two strings
func longestCommonSubarray(a, b string) string {
	if len(a) == 0 || len(b) == 0 {
		return ""
	}
	// let lcs(i, j) be the length of longest common subarray of a[0...i] and b[0...j] that ends at i, j.
	// then, if a[i]==b[j], lcs(i, j)=lcs(i-1, j-1)+1; if not lcs(i, j)=0
	// base case: lcs(0, x)=1 if a[0]==b[x], lcs(x, 0)=1 if a[x]==b[0]
	// time complexity O(mn), space O(mn) can be reduced to O(min(m,n)) since updated line by line
	if len(b) > len(a) {
		b, a = a, b // let b be shorter
	}
	maxl := 0
	maxj := 0
	lcs := make([]int, len(b))
	// set up lcs(0, j)
	for j := 0; j < len(b); j++ {
		if a[0] == b[j] {
			maxl = 1
			maxj = j
			lcs[j] = 1
		}
	}
	for i := 1; i < len(a); i++ {
		tmp := lcs[0] // tmp to store lcs(i-1, j-1)
		// set up lcs(i, 0)
		if a[i] == b[0] {
			lcs[0] = 1
			if lcs[0] > maxl {
				maxj = 0
				maxl = lcs[0]
			}
		} else {
			lcs[0] = 0
		}
		for j := 1; j < len(b); j++ {
			tmp2 := lcs[j]
			if a[i] == b[j] {
				lcs[j] = tmp + 1
				if lcs[j] > maxl {
					maxj = j
					maxl = lcs[j]
				}
			} else {
				lcs[j] = 0
			}
			tmp = tmp2
		}
	}
	return b[maxj-maxl+1 : maxj+1]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
