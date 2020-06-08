package main

import (
    "fmt"
)

// https://leetcode.com/problems/minimum-distance-to-type-a-word-using-two-fingers/

// You have a keyboard layout as shown above in the XY plane, where each English uppercase letter 
// is located at some coordinate, for example, the letter A is located at coordinate (0,0), the 
// letter B is located at coordinate (0,1), the letter P is located at coordinate (2,3) and the 
// letter Z is located at coordinate (4,1). 
//       ABCDEF
//       GHIJKL
//       MNOPQR
//       STUVWX
//       YZ
// Given the string word, return the minimum total distance to type such string using only two fingers. 
// The distance between coordinates (x1,y1) and (x2,y2) is |x1 - x2| + |y1 - y2|. 
// Note that the initial positions of your two fingers are considered free so don't count towards your 
// total distance, also your two fingers do not have to start at the first letter or the first two letters. 
// Example 1: 
//   Input: word = "CAKE"
//   Output: 3
//   Explanation: 
//     Using two fingers, one optimal way to type "CAKE" is: 
//     Finger 1 on letter 'C' -> cost = 0 
//     Finger 1 on letter 'A' -> cost = Distance from letter 'C' to letter 'A' = 2 
//     Finger 2 on letter 'K' -> cost = 0 
//     Finger 2 on letter 'E' -> cost = Distance from letter 'K' to letter 'E' = 1 
//     Total distance = 3
// Example 2: 
//   Input: word = "HAPPY"
//   Output: 6
//   Explanation: 
//     Using two fingers, one optimal way to type "HAPPY" is:
//     Finger 1 on letter 'H' -> cost = 0
//     Finger 1 on letter 'A' -> cost = Distance from letter 'H' to letter 'A' = 2
//     Finger 2 on letter 'P' -> cost = 0
//     Finger 2 on letter 'P' -> cost = Distance from letter 'P' to letter 'P' = 0
//     Finger 1 on letter 'Y' -> cost = Distance from letter 'A' to letter 'Y' = 4
//     Total distance = 6
// Example 3: 
//   Input: word = "NEW"
//   Output: 3
// Example 4: 
//   Input: word = "YEAR"
//   Output: 7
// Constraints: 
//   2 <= word.length <= 300 
//   Each word[i] is an English uppercase letter. 

// time O(26*26*n), space O(26*26)
func minimumDistance(word string) int {
    // let F(α,β,w[k:]) be the cost to type w[k:] with finger1 at letter α and finger2 at β.
    // then F(α,β,w[k:]) = min( dist(α->w[k])+F(w[k],β,w[k+1:]) , dist(β->w[k])+F(α,w[k],w[k+1:]) )
    old, new := [26][26]int{}, [26][26]int{}
    for k:=len(word)-1; k>=0; k-- {
    	w := int(word[k]-'A')
    	for i:=0; i<26; i++ {
    		for j:=0; j<26; j++ {
    			// i still, j->w; j still, i->w; the coordinates of letter i on keyboard is (i/6, i%6)
    			new[i][j] = min(old[i][w]+abs(j/6-w/6)+abs(j%6-w%6), old[w][j]+abs(i/6-w/6)+abs(i%6-w%6))
    		}
    	}
    	old, new = new, old
    }
    
    // find the minimal for F(word[0],x,word)
    allmin := 0x7fffffff
    for i:=0; i<26; i++ {
    	allmin = min(allmin, old[int(word[0]-'A')][i])  // we always use finger1 to type word[0]
    }
    return allmin
}

func min(a, b int) int {
	if a<b {
		return a
	}
	return b
}

func abs(a int) int {
	if a<0 {
		return -a
	}
	return a
}

func main() {
	fmt.Println(minimumDistance("CAKE"), 3)
	fmt.Println(minimumDistance("HAPPY"), 6)
	fmt.Println(minimumDistance("NEW"), 3)
	fmt.Println(minimumDistance("YEAR"), 7)
}