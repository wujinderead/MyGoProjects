package main

import (
	"container/list"
    "fmt"
)

// https://leetcode.com/problems/sliding-puzzle/

// On a 2x3 board, there are 5 tiles represented by the integers 1 through 5, and an empty square represented by 0.
// A move consists of choosing 0 and a 4-directionally adjacent number and swapping it. The state of the board is 
// solved if and only if the board is [[1,2,3],[4,5,0]]. Given a puzzle board, return the least number of moves required 
// so that the state of the board is solved. If it is impossible for the state of the board to be solved, return -1.
// Example 1:
//   Input: board = [[1,2,3],[4,0,5]]
//   Output: 1
//   Explanation: Swap the 0 and the 5 in one move.
// Example 2:
//   Input: board = [[1,2,3],[5,4,0]]
//   Output: -1
//   Explanation: No number of moves will make the board solved.
// Example 3:
//   Input: board = [[4,1,2],[5,0,3]]
//   Output: 5
//   Explanation: 5 is the smallest number of moves that solves the board.
//     An example path:
//     After move 0: [[4,1,2],[5,0,3]]
//     After move 1: [[4,1,2],[0,5,3]]
//     After move 2: [[0,1,2],[4,5,3]]
//     After move 3: [[1,0,2],[4,5,3]]
//     After move 4: [[1,2,0],[4,5,3]]
//     After move 5: [[1,2,3],[4,5,0]]
// Example 4:
//   Input: board = [[3,2,4],[1,5,0]]
//   Output: 14
// Note:
//   board will be a 2 x 3 array as described above.
//   board[i][j] will be a permutation of [0, 1, 2, 3, 4, 5].

func slidingPuzzle(board [][]int) int {
	// initial state
	var i0, j0 byte
    key := [8]byte{}
    for i := range board {
    	for j := range board[i] {
    		key[i*3+j] = byte(board[i][j])
    		if board[i][j]==0 {
    			i0, j0 = byte(i), byte(j)
    		}
    	}
    }
    key[6], key[7] = i0, j0
    if key==[8]byte{1,2,3,4,5,0,1,2} {   // already solved
    	return 0
    }

    visited := make(map[[8]byte]struct{})   // at most 720 states
    visited[key] = struct{}{}
    queue := list.New()
    queue.PushBack(key)
    step := 0
    
    // bfs to find the minimal step
	for queue.Len()>0 {
		L := queue.Len()
		for k:=0; k<L; k++ {
			key = queue.Remove(queue.Front()).([8]byte)
			i, j := int(key[6]), int(key[7])
			for _, v := range [][2]int{{0,1},{0,-1},{1,0},{-1,0}} {
				ni, nj := i+v[0], j+v[1]
				if ni>=0 && ni<2 && nj>=0 && nj<3 {
					key[ni*3+nj], key[i*3+j] = key[i*3+j], key[ni*3+nj]
					key[6], key[7] = byte(ni), byte(nj)
					if _, ok := visited[key]; !ok {
						//fmt.Println(key, step+1)
						if key==[8]byte{1,2,3,4,5,0,1,2} {
							return step+1
						}
						visited[key] = struct{}{}
						queue.PushBack(key)
					}
					key[ni*3+nj], key[i*3+j] = key[i*3+j], key[ni*3+nj]
					key[6], key[7] = byte(i), byte(j)
				}
			}
		} 
		step++
	} 
    return -1
}

func main() {
	fmt.Println(slidingPuzzle([][]int{{1,2,3}, {4,0,5}}), 1)
	fmt.Println(slidingPuzzle([][]int{{1,2,3}, {5,4,0}}), -1)
	fmt.Println(slidingPuzzle([][]int{{4,1,2}, {5,0,3}}), 5)
	fmt.Println(slidingPuzzle([][]int{{3,2,4}, {1,5,0}}), 14)
}