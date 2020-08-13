package main

import "fmt"

// https://leetcode.com/problems/filling-bookcase-shelves/

// We have a sequence of books: the i-th book has thickness books[i][0] and height books[i][1]. 
// We want to place these books in order onto bookcase shelves that have total width shelf_width. 
// We choose some of the books to place on this shelf (such that the sum of their thickness is <= shelf_width), 
// then build another level of shelf of the bookcase so that the total height of the bookcase has increased 
// by the maximum height of the books we just put down. We repeat this process until there are no more 
// books to place. 
// Note again that at each step of the above process, the order of the books we place is the same order 
// as the given sequence of books. For example, if we have an ordered list of 5 books, we might place 
// the first and second book onto the first shelf, the third book on the second shelf, and the fourth 
// and fifth book on the last shelf. 
// Return the minimum possible height that the total bookshelf can be after placing shelves in this manner. 
// Example 1: 
//        1
//        2233
//        2233
//        2233
//           7
//        4567  
//   Input: books = [[1,1],[2,3],[2,3],[1,1],[1,1],[1,1],[1,2]], shelf_width = 4
//   Output: 6
//   Explanation:
//     The sum of the heights of the 3 shelves are 1 + 3 + 2 = 6.
//     Notice that book number 2 does not have to be on the first shelf.
// Constraints: 
//   1 <= books.length <= 1000 
//   1 <= books[i][0] <= shelf_width <= 1000 
//   1 <= books[i][1] <= 1000 

// time O(n^2), space o(n)
func minHeightShelves(books [][]int, shelf_width int) int {
	// let dp[i] be the minimal height to arrange book[i:], then 
	// dp[i] = min(height[i]+dp[i+1], max(height[i...i+1])+dp[i+2], max(height[i...j]+dp[j+1], ...),
	// under the condition that sum(width[i...j]) < shelf_width
	dp := make([]int, len(books)+1)
	dp[len(books)-1] = books[len(books)-1][1]
	for i:=len(books)-2; i>=0; i-- {
		maxh := books[i][1]
		width := books[i][0]
		dp[i] = maxh + dp[i+1]
		for j:=i+1; j<len(books); j++ {
			if width+books[j][0] > shelf_width {
				break
			}
			width += books[j][0]
			maxh = max(maxh, books[j][1])
			dp[i] = min(dp[i], dp[j+1]+maxh)
		}
	}
	return dp[0]
}

func min(a, b int) int {
	if a<b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a>b {
		return a
	}
	return b
}

func main() {
	fmt.Println(minHeightShelves([][]int{{1,1},{2,3},{2,3},{1,1},{1,1},{1,1},{1,2}}, 4), 6)
	fmt.Println(minHeightShelves([][]int{{1,2}}, 4), 2)
	fmt.Println(minHeightShelves([][]int{{1,2},{2,3}}, 4), 3)
	fmt.Println(minHeightShelves([][]int{{7,3},{8,7},{2,7},{2,5}}, 10), 15)	
}