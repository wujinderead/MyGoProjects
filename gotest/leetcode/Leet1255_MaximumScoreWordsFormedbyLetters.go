package main

import "fmt"

// https://leetcode.com/problems/maximum-score-words-formed-by-letters

// Given a list of words, list of single letters (might be repeating) and score
// of every character.
// Return the maximum score of any valid set of words formed by using the given 
// letters (words[i] cannot be used two or more times).
// It is not necessary to use all characters in letters and each letter can only
// be used once. Score of letters 'a', 'b', 'c', ... ,'z' is given by
// score[0], score[1], ... , score[25] respectively.
// Example 1:
//   Input: words = ["dog","cat","dad","good"],
//     letters = ["a","a","c","d","d","d","g","o","o"],
//     score = [1,0,9,5,0,0,3,0,0,0,0,0,0,0,2,0,0,0,0,0,0,0,0,0,0,0]
//   Output: 23
//   Explanation:
//     Score  a=1, c=9, d=5, g=3, o=2
//     Given letters, we can form the words "dad" (5+1+5) and "good" (3+2+2+5) with a
//     score of 23. Words "dad" and "dog" only get a score of 21.
// Example 2:
//   Input: words = ["xxxz","ax","bx","cx"],
//     letters = ["z","a","b","c","x","x","x"],
//     score = [4,4,4,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,5,0,10]
//   Output: 27
//   Explanation:
//     Score  a=4, b=4, c=4, x=5, z=10
//     Given letters, we can form the words "ax" (4+5), "bx" (4+5) and "cx" (4+5)
//     with a score of 27. Word "xxxz" only get a score of 25.
// Example 3:
//   Input: words = ["leetcode"],
//     letters = ["l","e","t","c","o","d"],
//     score = [0,0,1,1,1,0,0,0,0,0,0,1,0,0,1,0,0,0,0,1,0,0,0,0,0,0]
//   Output: 0
//   Explanation:
//     Letter "e" can only be used once.
// Constraints:
//   1 <= words.length <= 14
//   1 <= words[i].length <= 15
//   1 <= letters.length <= 100
//   letters[i].length == 1
//   score.length == 26
//   0 <= score[i] <= 10
//   words[i], letters[i] contains only lower case English letters.

// IMPROVEMENT:
// since n is small, we can use bit manipulate to generate all subsets of words,
// and calculate each score.

func maxScoreWords(words []string, letters []byte, score []int) int {
    // let F(i, letters) be the max score for words[i:] and letters, then,
    // F(i, letters) has two candidate, we need the max one:
    //    score(words[i]) + F(i+1, letters-words[i])
    //    F(i+1, letters)
    // so, to compute F(0, x), we need compute two F(1, xx), four F(2, xx),
    // the complexity is O(2^n), since n<=14 in this problem, 2^n is acceptable.
    let := [26]int8{}     // byte is uint8, we need int8
    for _, v := range letters {
    	let[int(v-'a')]++
	}
	return F(words, 0, let, score)
}

func F(words []string, ind int, letters [26]int8, score []int) int {
	if ind == len(words) {
		return 0
	}
	cand := F(words, ind+1, letters, score)   // array as func parameter is copied
	sc := 0
	for _, ch := range words[ind] {
		chind := int(ch-'a')
		sc += score[chind]
		letters[chind]--
		if letters[chind]<0 {
			return cand   // current letters can't form words[ind]
		}
	}
	return max(cand, sc+F(words, ind+1, letters, score))
}

func max(a, b int) int {
	if a>b {
		return a
	}
	return b
}

func main() {
    fmt.Println(maxScoreWords([]string{"dog","cat","dad","good"},
    	[]byte{'a','a','c','d','d','d','g','o','o'},
    	[]int{1,0,9,5,0,0,3,0,0,0,0,0,0,0,2,0,0,0,0,0,0,0,0,0,0,0}))
	fmt.Println(maxScoreWords([]string{"xxxz","ax","bx","cx"},
		[]byte{'z','a','b','c','x','x','x'},
		[]int{4,4,4,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,5,0,10}))
	fmt.Println(maxScoreWords([]string{"leetcode"},
		[]byte{'l','e','t','c','o','d'},
		[]int{0,0,1,1,1,0,0,0,0,0,0,1,0,0,1,0,0,0,0,1,0,0,0,0,0,0}))
}