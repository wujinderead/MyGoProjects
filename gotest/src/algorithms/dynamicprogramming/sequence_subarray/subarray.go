package longest

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

// get the length of the longest common subarray of two strings
func longestCommonSubarray(a, b string) string {
	if len(a) == 0 || len(b) == 0 {
		return ""
	}
	// let lcs(i, j) be the length of longest common subarray of a[0...i] and b[0...j] that ends at i, j.
	// then, if a[i]==b[j], lcs(i, j)=lcs(i-1, j-1)+1; if not lcs(i, j)=0
	// base case: lcs(0, x)=1 if a[0]==b[x], lcs(x, 0)=1 if a[x]=b[0]
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
