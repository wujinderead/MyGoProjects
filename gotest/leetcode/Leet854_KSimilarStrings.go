package main

import (
	"container/list"
	"fmt"
)

// https://leetcode.com/problems/k-similar-strings/

// Strings A and B are K-similar (for some non-negative integer K) if we can swap the 
// positions of two letters in A exactly K times so that the resulting string equals B. 
// Given two anagrams A and B, return the smallest K for which A and B are K-similar. 
// Example 1: 
//   Input: A = "ab", B = "ba"
//   Output: 1
// Example 2: 
//   Input: A = "abc", B = "bca"
//   Output: 2
// Example 3: 
//   Input: A = "abac", B = "baca"
//   Output: 2
// Example 4: 
//   Input: A = "aabc", B = "abca"
//   Output: 2 
// Note: 
//   1 <= A.length == B.length <= 20 
//   A and B contain only lowercase letters from the set {'a', 'b', 'c', 'd', 'e', 'f'} 

// for example, the target is "xAxxxxxxx", we currently have "xBxxAxxxA", 
// the problem is which A we should choose to swap with B.
// we try both, use bfs to find the minimal step. 
func kSimilarity(A string, B string) int {
	if A==B {
		return 0
	}
    visited := make(map[string]struct{})
    queue := list.New()
    queue.PushBack(A)
    visited[A] = struct{}{}

    step := 0
    for queue.Len()> 0 {
    	curlen := queue.Len()
    	for k:=0; k<curlen; k++ {
    		str := []byte(queue.Remove(queue.Front()).(string))
    		for i:=0; i<len(A); i++ {
				if str[i]!=B[i] {   // find only the first wrong place   
					for j:=i+1; j<len(A); j++ {  
						if str[j]==B[i] {   // find all correct char to swap
							str[i], str[j] = str[j], str[i]
							strr := string(str)
							str[i], str[j] = str[j], str[i]
							if strr==B {
								return step+1
							}
							if _, ok := visited[strr]; !ok {
								visited[strr] = struct{}{}
								queue.PushBack(strr)
							}
						}
					}
					break
				}
			}
    	}
    	step++
	}
    return 0
}

func main() {
	fmt.Println(kSimilarity("ab", "ba"), 1)
	fmt.Println(kSimilarity("abc", "bca"), 2)
	fmt.Println(kSimilarity("abac", "baca"), 2)
	fmt.Println(kSimilarity("aabc", "abca"), 2)
	fmt.Println(kSimilarity("abccaacceecdeea", "bcaacceeccdeaae"), 9)
	fmt.Println(kSimilarity("aabbccddee", "cdacbeebad"), 6)
	fmt.Println(kSimilarity("aabcde", "cdaeba"), 4)
	fmt.Println(kSimilarity("bcde", "cebd"), 3)
}