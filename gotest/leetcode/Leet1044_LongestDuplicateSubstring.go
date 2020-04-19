package main

import (
	"fmt"
	"math/big"
)

// https://leetcode.com/problems/longest-duplicate-substring/

// Given a string S, consider all duplicated substrings: (contiguous) substrings
// of S that occur 2 or more times. (The occurrences may overlap.)
// Return any duplicated substring that has the longest possible length. (If S does
// not have a duplicated substring, the answer is "".)
// Example 1:
//   Input: "banana"
//   Output: "ana"
// Example 2:
//   Input: "abcd"
//   Output: ""
// Note:
//   2 <= S.length <= 10^5
//   S consists of lowercase English letters.

func longestDupSubstring(S string) string {
	// use binary search to find the max length. use rabin-karp to match string
	l, h := 1, len(S)-1
	ans := ""
	p := 1000000007
	pinv := getModInverse(26, p)
outer:
	for l <= h {
		m := (l + h) / 2
		fmt.Println(l, h, m)
		mapp := make(map[int][]int)
		hash := 0
		pow := pinv

		// get initial hash
		for i := 0; i < m; i++ {
			hash = (hash*26 + int(S[i]-'a')) % p
			pow = (pow * 26) % p
		}
		mapp[hash] = []int{0} // save hash and sub array start
		// iterate following sub strings
		for i := 1; i+m <= len(S); i++ {
			hash = (hash - int(S[i-1]-'a')*pow) % p
			if hash < 0 {
				hash += p
			}
			hash = (hash*26 + int(S[i+m-1]-'a')) % p
			if subs, ok := mapp[hash]; ok { // got a match of hash
				for _, s := range subs {
					if S[s:s+m] == S[i:i+m] { // find a match, we set l=m+1 to search
						ans = S[i : i+m] // save candidate result
						l = m + 1
						continue outer
					}
				}
				mapp[hash] = append(subs, i)
			} else {
				mapp[hash] = []int{i}
			}
		}
		// no match, set h=m-1 to search
		h = m - 1
	}
	return ans
}

func getModInverse(a, p int) int {
	pp := new(big.Int).SetInt64(int64(p))
	aa := new(big.Int).SetInt64(int64(a))
	aa.ModInverse(aa, pp)
	return int(aa.Int64())
}

func main() {
	fmt.Println(longestDupSubstring("qwertyuhgqwerty"))
	fmt.Println(longestDupSubstring("banana"))
	fmt.Println(longestDupSubstring("abcd"))
	fmt.Println(longestDupSubstring("aaaaaaaa"))
	a := longestDupSubstring("okmzpmxzwjbfssktjtebhhxfphcxefhonkncnrumgduoaeltjvwqwydpdsrbxsgmc" +
		"dxrthilniqxkqzuuqzqhlccmqcmccfqddncchadnthtxjruvwsmazlzhijygmtabbzelslebyrfpyyvcwnaiqkkzlyillxmk" +
		"fggyfwgzhhvyzfvnltjfxskdarvugagmnrzomkhldgqtqnghsddgrjmuhpgkfcjkkkaywkzsikptkrvbnvuyamegwempuwfp" +
		"aypmuhhpuqrufsgpiojhblbihbrpwxdxzolgqmzoyeblpvvrnbnsdnonhpmbrqissifpdavvscezqzclvukfgmrmbmmwvzfp" +
		"xcgecyxneipexrzqgfwzdqeeqrugeiupukpveufmnceetilfsqjprcygitjefwgcvqlsxrasvxkifeasofcdvhvrpmxvjevu" +
		"pqtgqfgkqjmhtkyfsjkrdczmnettzdxcqexenpxbsharuapjmdvmfygeytyqfcqigrovhzbxqxidjzxfbrlpjxibtbndgubw" +
		"gihdzwoywqxegvxvdgaoarlauurxpwmxqjkidwmfuuhcqtljsvruinflvkyiiuwiiveplnxlviszwkjrvyxijqrulchzkerb" +
		"dyrdhecyhscuojbecgokythwwdulgnfwvdptzdvgamoublzxdxsogqpunbtoixfnkgbdrgknvcydmphuaxqpsofmylyijpzh" +
		"bqsxryqusjnqfikvoikwthrmdwrwqzrdmlugfglmlngjhpspvnfddqsvrajvielokmzpmxzwjbfssktjtebhhxfphcxefhon" +
		"kncnrumgduoaeltjvwqwydpdsrbxsgmcdxrthilniqxkqzuuqzqhlccmqcmccfqddncchadnthtxjruvwsmazlzhijygmtab" +
		"bzelslebyrfpyyvcwnaiqkkzlyillxmkfggyfwgzhhvyzfvnltjfxskdarvugagmnrzomkhldgqtqnghsddgrjmuhpgkfcjk" +
		"kkaywkzsikptkrvbnvuyamegwempuwfpaypmuhhpuqrufsgpiojhblbihbrpwxdxzolgqmzoyeblpvvrnbnsdnonhpmbrqis" +
		"sifpdavvscezqzclvukfgmrmbmmwvzfpxcgecyxneipexrzqgfwzdqeeqrugeiupukpveufmnceetilfsqjprcygitjefwgc" +
		"vqlsxrasvxkifeasofcdvhvrpmxvjevupqtgqfgkqjmhtkyfsjkrdczmnettzdxcqexenpxbsharuapjmdvmfygeytyqfcqi" +
		"grovhzbxqxidjzxfbrlpjxibtbndgubwgihdzwoywqxegvxvdgaoarlauurxpwmxqjkidwmfuuhcqtljsvruinflvkyiiuwi" +
		"iveplnxlviszwkjrvyxijqrulchzkerbdyrdhecyhscuojbecgokythwwdulgnfwvdptzdvgamoublzxdxsogqpunbtoixfn" +
		"kgbdrgknvcydmphuaxqpsofmylyijpzhbqsxryqusjnqfikvoikwthrmdwrwqzrdmlugfglmlngjhpspvnfddqsvrajviel")
	exp := "okmzpmxzwjbfssktjtebhhxfphcxefhonkncnrumgduoaeltjvwqwydpdsrbxsgmcdxrthilniqxkq" +
		"zuuqzqhlccmqcmccfqddncchadnthtxjruvwsmazlzhijygmtabbzelslebyrfpyyvcwnaiqkkzlyillx" +
		"mkfggyfwgzhhvyzfvnltjfxskdarvugagmnrzomkhldgqtqnghsddgrjmuhpgkfcjkkkaywkzsikptkrv" +
		"bnvuyamegwempuwfpaypmuhhpuqrufsgpiojhblbihbrpwxdxzolgqmzoyeblpvvrnbnsdnonhpmbrqis" +
		"sifpdavvscezqzclvukfgmrmbmmwvzfpxcgecyxneipexrzqgfwzdqeeqrugeiupukpveufmnceetilfs" +
		"qjprcygitjefwgcvqlsxrasvxkifeasofcdvhvrpmxvjevupqtgqfgkqjmhtkyfsjkrdczmnettzdxcqe" +
		"xenpxbsharuapjmdvmfygeytyqfcqigrovhzbxqxidjzxfbrlpjxibtbndgubwgihdzwoywqxegvxvdga" +
		"oarlauurxpwmxqjkidwmfuuhcqtljsvruinflvkyiiuwiiveplnxlviszwkjrvyxijqrulchzkerbdyrd" +
		"hecyhscuojbecgokythwwdulgnfwvdptzdvgamoublzxdxsogqpunbtoixfnkgbdrgknvcydmphuaxqps" +
		"ofmylyijpzhbqsxryqusjnqfikvoikwthrmdwrwqzrdmlugfglmlngjhpspvnfddqsvrajviel"
	fmt.Println(a == exp, len(exp))
}
