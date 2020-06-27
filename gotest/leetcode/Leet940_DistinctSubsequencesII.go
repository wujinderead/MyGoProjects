package main

import "fmt"

// Given a string S, count the number of distinct, non-empty subsequences of S . 
// Since the result may be large, return the answer modulo 10^9 + 7. 
// Example 1: 
//   Input: "abc"
//   Output: 7
//   Explanation: The 7 distinct subsequences are "a", "b", "c", "ab", "ac", "bc", and "abc".
// Example 2: 
//   Input: "aba"
//   Output: 6
//   Explanation: The 6 distinct subsequences are "a", "b", "ab", "ba", "aa" and "aba".
// Example 3: 
//   Input: "aaa"
//   Output: 3
//   Explanation: The 3 distinct subsequences are "a", "aa" and "aaa".
// Note: 
//   S contains only lowercase letters. 
//   1 <= S.length <= 2000 

// let dp[i] be the number of subsequences that end at S[i-1] 
// for example, S="abaacbc"
// dp("")=1
// dp("a")=dp("")=1
// dp("ab")=dp("")+dp("a")                // {"", "a"} plus 'b'
// dp("aba")=dp("a")+dp("ab")             // {"a", "ab"} plus 'a', {""} plus 'a' duplicated
// dp("abaa")=dp("a")+dp("ab")+dp("aba")  // {"a", "ab", "aba"} plus 'a', {""} plus 'a' duplicated
// dp("abaac")                       // {"", "a", "ab", "aba", "abaa"} plus 'c'
// dp("abaacb")                      // {"ab", "aba", "abaa", "abaac"} plus 'b', {"", "a"} plus 'b' duplicated
// dp("abaacbc")                     // {"abaac", "abaacb"} plus 'c'
func distinctSubseqII(S string) int {
	dp := make([]int, len(S)+1)
	dp[0] = 1     // 1 for "" 
	
	// last occurence of letter
	lastoccur := [26]int{}

	// dp 
	for i:=1; i<=len(S); i++ {
		last := lastoccur[int(S[i-1]-'a')]  // last occurence of letter
		// dp[i]=sum(dp[last:i]), we can avoid loop by use prefix array
		// and dp array can be further optimized to O(1) space
		for j:=last; j<i; j++ {
			dp[i] += dp[j]
		}
		lastoccur[int(S[i-1]-'a')] = i      // update last occurence
		dp[i] = dp[i]%int(1e9+7)
	}

	// sum all dp
	sum := 0
	for i:=1; i<=len(S); i++ {
		sum += dp[i]
	}
	return sum%int(1e9+7)
}

func main() {
	fmt.Println(distinctSubseqII("abc"), 7)
	fmt.Println(distinctSubseqII("aba"), 6)
	fmt.Println(distinctSubseqII("aaa"), 3)
	fmt.Println(distinctSubseqII("abaacbc"), 65)
	fmt.Println(distinctSubseqII("ababbacbcbbabghgdbcvdbsdgasgdjhadjhg"+
		"sdgsgdjasgdjhsgdfbbhfsvbdsadvsdv"), 579115464)
}