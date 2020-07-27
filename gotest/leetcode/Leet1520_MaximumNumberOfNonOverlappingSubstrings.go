package main

import (
    "fmt"
)

// https://leetcode.com/problems/maximum-number-of-non-overlapping-substrings/

// Given a string s of lowercase letters, you need to find the maximum number of non-empty substrings of s 
// that meet the following conditions:
//   The substrings do not overlap, that is for any two substrings s[i..j] and s[k..l], either j < k or i > l is true.
//   A substring that contains a certain character c must also contain all occurrences of c.
// Find the maximum number of substrings that meet the above conditions. If there are multiple solutions with 
// the same number of substrings, return the one with minimum total length. It can be shown that there exists 
// a unique solution of minimum total length. Notice that you can return the substrings in any order.
// Example 1:
//   Input: s = "adefaddaccc"
//   Output: ["e","f","ccc"]
//   Explanation: The following are all the possible substrings that meet the conditions:
//       [
//         "adefaddaccc"
//         "adefadda",
//         "ef",
//         "e",
//         "f",
//         "ccc",
//       ]
//     If we choose the first string, we cannot choose anything else and we'd get only 1. If we choose "adefadda", 
//     we are left with "ccc" which is the only one that doesn't overlap, thus obtaining 2 substrings. 
//     Notice also, that it's not optimal to choose "ef" since it can be split into two. Therefore, 
//     the optimal way is to choose ["e","f","ccc"] which gives us 3 substrings. 
//     No other solution of the same number of substrings exist.
// Example 2:
//   Input: s = "abbaccd"
//   Output: ["d","bb","cc"]
//   Explanation: Notice that while the set of substrings ["d","abba","cc"] also has length 3, 
//     it's considered incorrect since it has larger total length.
// Constraints:
//   1 <= s.length <= 10^5
//   s contains only lowercase English letters.

func maxNumOfSubstrings(s string) []string {
	// first and last occurrence of a character
	first, last := [26]int{}, [26]int{}  
    for i := range s {
    	ind := int(s[i]-'a')
    	if first[ind]>0 {
    		last[ind] = i+1
    	} else {
    		first[ind] = i+1
    		last[ind] = i+1
    	}
    }

    // extend intervals
    for i:=1; i<=len(s); i++ {   // for each c in s
    	ind := int(s[i-1]-'a')
    	for j:=0; j<26; j++ {    // for ch j
    		if first[j]>0 && first[j]<i && i<last[j] {   // if [j  c  j], j need merge c 
    			f, l := first[j], last[j]
    			ff, ll := first[ind], last[ind]
    			if f<=ff && ll<=l {
    				continue           //     [ff  ll]
    				                   // [f           l] 
    			}
    			if f<=ff && ff<=l {    // [f     l]   
    				                   //    [ff     ll]
    				last[j] = ll       // change to [f    ll]
    			}
    			if f<=ll && ll<=l {    //       [f      l]
    				                   // [ff       ll]
    				first[j] = ff      // change to [ff    l]
    			}
    			if ff<=f && l<=ll { 
    				first[j] = ff      //      [f  l]
    				last[j] = ll       // [ff          ll] 
    			}
    		}
    	}
    }

    // select shortest intervals
    set := make(map[[2]int]struct{}, 26)
    outer: for i:=0; i<26; i++ {
    	if first[i]==0 {
    		continue
    	}
    	f, l := first[i], last[i]  // interval of current character
    	for k := range set {
    		ff, ll := k[0], k[1]
    		// if this interval shorter than existed, delete existed
    		if ff<f && l<ll {
    			delete(set, k)        //      [f  l]
    			break                 // [ff          ll] 
    		} 
    		// if this interval longer than existed, ignore
    		if f<ff && ll<l {
    			continue outer        //     [ff  ll]
    			                      // [f            l] 
    		}
    	}
    	set[[2]int{f, l}] = struct{}{} 
    }

    // make answer
    ans := make([]string, 0, len(set))
    for k := range set {
    	ans = append(ans, s[k[0]-1: k[1]])
    }
    return ans
}

func main() {
	fmt.Println(maxNumOfSubstrings("adefaddaccc"))	
	fmt.Println(maxNumOfSubstrings("abbaccd"))	
	fmt.Println(maxNumOfSubstrings("acbeffeacb"))
	fmt.Println(maxNumOfSubstrings("acbeefgacb"))
	fmt.Println(maxNumOfSubstrings("abacdefgg"))
	fmt.Println(maxNumOfSubstrings("ababa"))
}