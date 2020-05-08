package main

import (
	"fmt"
	"container/list"
)

// https://leetcode.com/problems/shortest-path-to-get-all-keys/

// We are given a 2-dimensional grid. "." is an empty cell, "#" is a wall, "@" is
// the starting point, ("a", "b", ...) are keys, and ("A", "B", ...) are locks.
// We start at the starting point, and one move consists of walking one space in
// one of the 4 cardinal directions. We cannot walk outside the grid, or walk into
// a wall. If we walk over a key, we pick it up. We can't walk over a lock unless
// we have the corresponding key.
// For some 1 <= K <= 6, there is exactly one lowercase and one uppercase letter
// of the first K letters of the English alphabet in the grid. This means that there
// is exactly one key for each lock, and one lock for each key; and also that the
// letters used to represent the keys and locks were chosen in the same order as
// the English alphabet.
// Return the lowest number of moves to acquire all keys. If it's impossible, return -1.
// Example 1:
//   Input: ["@.a.#","###.#","b.A.B"]
//      @.a.#
//      ###.#
//      b.A.B
//   Output: 8
// Example 2:
//   Input: ["@..aA","..B#.","....b"]
//      @..aA
//      ..B#.
//      ....b
//   Output: 6
// Note:
//   1 <= grid.length <= 30
//   1 <= grid[0].length <= 30
//   grid[i][j] contains only '.', '#', '@', 'a'-'f' and 'A'-'F'
//   The number of keys is in [1, 6]. Each key has a different letter and opens exactly one lock.

// bfs: use (row, col, current_key_bits) as key for visit, we would have at most row*col*2^K states.
func shortestPathAllKeys(grid []string) int {
	// prep-precessing
	si, sj := 0, 0
	m, n := len(grid), len(grid[0])
	K := 0
	for i:=0; i<m; i++ {
		for j:=0; j<n; j++ {
			if grid[i][j] == '@' {
				si, sj = i, j
			}
			if grid[i][j]>='A' && grid[i][j]<='F' {
				K++
			}
		}
	}
	allkeys := (1<<uint(K))-1
	allmin := 0x7fffffff
	l := list.New()
	visited := make(map[[3]int]struct{})  // (row, col, keys) as visited key
	l.PushBack([4]int{si, sj, 0, 0})   // (row, col, keys, step) pair
	visited[[3]int{si, sj, 0}] = struct{}{}
	for l.Len() > 0 {
		pair := l.Remove(l.Front()).([4]int)
		i, j, keys, step := pair[0], pair[1], pair[2], pair[3]
		for _, v := range [4][2]int{{-1,0}, {1,0}, {0,1}, {0,-1}} {
			ni, nj := i+v[0], j+v[1]
			if _, ok := visited[[3]int{ni, nj, keys}]; ni>=0 && ni<m && nj>=0 && nj<n && !ok {
				visited[[3]int{ni, nj, keys}] = struct{}{}  // set visited so we don't bother to check again
				ch := grid[ni][nj]
				if ch == '#' || (ch >='A' && ch <= 'F' && keys&(1<<uint(ch-'A'))==0) {
					// found a wall or a lock but we don't hold a key
					continue
				}
				if ch>='a' && ch<='f' && keys&(1<<uint(ch-'a'))==0 {
					// find a unvisited key
					newkeys := keys | (1<<uint(ch-'a'))
					if newkeys == allkeys {   // find all keys
						allmin = min(allmin, step+1)   // all keys are find, no need to push to queue
					} else {
						l.PushBack([4]int{ni, nj, newkeys, step + 1})
					}
					continue
				}
				// otherwise, we can move one step more to current place
				l.PushBack([4]int{ni, nj, keys, step+1})
			}
		}
	}
	if allmin==0x7fffffff {
		return -1
	}
	return allmin
}

func min(a, b int) int {
	if a<b {
		return a
	}
	return b
}

func main() {
	fmt.Println(shortestPathAllKeys([]string{"@.a.#","###.#","b.A.B"}))
	fmt.Println(shortestPathAllKeys([]string{"@..aA","..B#.","....b"}))
	fmt.Println(shortestPathAllKeys([]string{".#.b.","A.#aB","#d...","@.cC.","D...#"}))
}