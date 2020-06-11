package main

import (
	"container/list"
    "fmt"
)

// https://leetcode.com/problems/minimum-moves-to-reach-target-with-rotations/

// In an n*n grid, there is a snake that spans 2 cells and starts moving from thetop left corner at (0, 0) and (0, 1). 
// The grid has empty cells represented by zeros and blocked cells represented by ones. The snake wants to reach the 
// lower right corner at (n-1, n-2) and (n-1, n-1). In one move the snake can: 
//   Move one cell to the right if there are no blocked cells there. This move keeps the horizontal/vertical 
//     position of the snake as it is. 
//   Move down one cell if there are no blocked cells there. This move keeps the horizontal/vertical position 
//     of the snake as it is. 
//   Rotate clockwise if it's in a horizontal position and the two cells under it are both empty. In that case 
//     the snake moves from (r, c) and (r, c+1) to (r, c) and (r+1, c). 
//        X X   ->   X 0
//        0 0        X 0
//   Rotate counterclockwise if it's in a vertical position and the two cells to its right are both empty. In that 
//     case the snake moves from (r, c) and (r+1, c) to (r, c) and (r, c+1). 
//        X 0   ->   X X
//        X 0        0 0
// Return the minimum number of moves to reach the target. 
// If there is no way to reach the target, return -1. 
// Example 1: 
//   Input: grid = [[0,0,0,0,0,1],
//               [1,1,0,0,1,0],
//               [0,0,0,0,1,1],
//               [0,0,1,0,1,0],
//               [0,1,1,0,0,0],
//               [0,1,1,0,0,0]]
//   Output: 11
//   Explanation:
//     One possible solution is [right, right, rotate clockwise, right, down, down, down, down, 
//       rotate counterclockwise, right, down].
// Example 2: 
//   Input: grid = [[0,0,1,1,1,1],
//               [0,0,0,0,1,1],
//               [1,1,0,0,0,1],
//               [1,1,1,0,0,1],
//               [1,1,1,0,0,1],
//               [1,1,1,0,0,0]]
//   Output: 9
// Constraints: 
//   2 <= n <= 100 
//   0 <= grid[i][j] <= 1 
//   It is guaranteed that the snake starts at empty cells. 

func minimumMoves(grid [][]int) int {
	n := len(grid)
    queue := list.New()
    queue.PushBack([3]int{0,0,0})  // left-up coordinates and direction
    visited := make(map[[3]int]struct{})
    visited[[3]int{0,0,0}] = struct{}{}
    step := 0
    for queue.Len()>0 {
    	L := queue.Len()
    	for k:=0; k<L; k++ {   // for current layer
    		tmp := queue.Remove(queue.Front()).([3]int)
    		i, j, d := tmp[0], tmp[1], tmp[2]
    		if d==0 {  // horizontal
    			if _, ok := visited[[3]int{i, j+1, 0}]; j+2<n && grid[i][j+2]==0 && !ok {   // right
    				visited[[3]int{i, j+1, 0}] = struct{}{}
    				queue.PushBack([3]int{i, j+1, 0})
    				if i==n-1 && j+1==n-2 {
    					return step+1
    				}
    			}
    			if i+1<n && grid[i+1][j]==0 && grid[i+1][j+1]==0 {    // clock
    				if _, ok := visited[[3]int{i, j, 1}]; !ok {
    					visited[[3]int{i, j, 1}] = struct{}{}
    					queue.PushBack([3]int{i, j, 1})
    				}
    				if _, ok := visited[[3]int{i+1, j, 0}]; !ok {     // down
    					visited[[3]int{i+1, j, 0}] = struct{}{}
    					queue.PushBack([3]int{i+1, j, 0})
    					if i+1==n-1 && j==n-2 {
    						return step+1
    					}
    				}
    			}
    		}
    		if d==1 {  // vertical
    			if _, ok := visited[[3]int{i+1, j, 1}]; i+2<n && grid[i+2][j]==0 && !ok {   // down
    				visited[[3]int{i+1, j, 1}] = struct{}{}
    				queue.PushBack([3]int{i+1, j, 1})
    			}
    			if j+1<n && grid[i][j+1]==0 && grid[i+1][j+1]==0 {      // right
    				if _, ok := visited[[3]int{i, j+1, 1}]; !ok {
    					visited[[3]int{i, j+1, 1}] = struct{}{}
    					queue.PushBack([3]int{i, j+1, 1})
    				}
    				if _, ok := visited[[3]int{i, j, 0}]; !ok {     // anti-clock
    					visited[[3]int{i, j, 0}] = struct{}{}
    					queue.PushBack([3]int{i, j, 0})
    				}
    			}
    		}
    	}
    	step++
    }
    return -1
}

func main() {
	fmt.Println(minimumMoves([][]int{
		{0,0,0,0,0,1},
		{1,1,0,0,1,0},
		{0,0,0,0,1,1},
		{0,0,1,0,1,0},
		{0,1,1,0,0,0},
		{0,1,1,0,0,0},
	}), 11)
	fmt.Println(minimumMoves([][]int{
		{0,0,1,1,1,1},
		{0,0,0,0,1,1},
		{1,1,0,0,0,1},
		{1,1,1,0,0,1},
		{1,1,1,0,0,1},
		{1,1,1,0,0,0},
	}), 9)
	fmt.Println(minimumMoves([][]int{
		{0,0},
		{0,0},
	}), 1)
	fmt.Println(minimumMoves([][]int{
		{0,0,0},
		{0,0,1},
		{0,0,0},
	}), 3)
	fmt.Println(minimumMoves([][]int{
		{0,0,0},
		{0,0,1},
		{0,0,0},
	}), 3)
	fmt.Println(minimumMoves([][]int{
		{0,0,0,0,0,0,0,0,0,1},
		{0,1,0,0,0,0,0,1,0,1},
		{1,0,0,1,0,0,1,0,1,0},
		{0,0,0,1,0,1,0,1,0,0},
		{0,0,0,0,1,0,0,0,0,1},
		{0,0,1,0,0,0,0,0,0,0},
		{1,0,0,1,0,0,0,0,0,0},
		{0,0,0,0,0,0,0,0,0,0},
		{0,0,0,0,0,0,0,0,0,0},
		{1,1,0,0,0,0,0,0,0,0},
	}), -1)
}