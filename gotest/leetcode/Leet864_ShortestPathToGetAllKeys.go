package main

import (
    "fmt"
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

// brute force: K! permutation of keys, and we use bfs to find the shortest path between keys.
// time complexity O(m*n*K!)
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
    mask := (1<<uint(K))-1
    allmin := 0x7fffffff
    q := newQueue((m+1)*(n+1))
    findKeys(grid, q, si, sj, 0, 0, mask, &allmin)
    if allmin==0x7fffffff {
        return -1
    }
    return allmin
}

func findKeys(grid []string, q *queue, si, sj, step, curmask, wantmask int, allmin *int) {
    q.reset()
    m, n := len(grid), len(grid[0])
    q.push([3]int{si, sj, step})
    visited := make([]bool, m*n)
    set2d(visited, si, sj, n)
    reachableKeys := make([][4]int, 0)
    for q.length > 0 {
        tuple := q.pop().([3]int)
        i, j, step := tuple[0], tuple[1], tuple[2]
        for _, v := range [4][2]int{{-1,0}, {1,0}, {0,1}, {0,-1}} {
            ni, nj := i+v[0], j+v[1]
            if ni>=0 && ni<m && nj>=0 && nj<n && !get2d(visited, ni, nj, n) {
                set2d(visited, ni, nj, n)
                ch := grid[ni][nj]
                if ch == '#' || (ch >='A' && ch <= 'F' && curmask&(1<<uint(ch-'A'))==0) {
                    // found a wall or a lock but we don't hold a key
                    continue   // set visited so we don't bother to check
                }
                if ch>='a' && ch<='f' && curmask&(1<<uint(ch-'a'))==0 {
                    // find a unvisited key
                    reachableKeys = append(reachableKeys, [4]int{ni, nj, step+1, curmask | (1<<uint(ch-'a'))})
                    if curmask | (1<<uint(ch-'a')) == wantmask {   // find all keys
                        *allmin = min(*allmin, step+1)
                    }
                }
                // otherwise, we can move one step more to current place
                q.push([3]int{ni, nj, step+1})
            }
        }
    }
    for _, v := range reachableKeys {
        findKeys(grid, q, v[0], v[1], v[2], v[3], wantmask, allmin)
    }
}

func min(a, b int) int {
    if a<b {
        return a
    }
    return b
}

func get2d(arr []bool, i, j, col int) bool {
    return arr[i*col+j]
}

func set2d(arr []bool, i, j, col int) {
    arr[i*col+j] = true
}

type queue struct {
    eles []interface{}
    front, back, length int
}

func newQueue(n int) *queue {
    return &queue{
        eles: make([]interface{}, n),
        front: 0,
        back: 0,
        length: 0,
    }
}

func (q *queue) reset() {
    q.front, q.back, q.length = 0, 0, 0
}

func (q *queue) push(x interface{}) {
    q.eles[q.back] = x
    q.length++
    q.back = (q.back+1)%len(q.eles)
}

func (q *queue) pop() interface{} {
    x := q.eles[q.front]
    q.length--
    q.front = (q.front+1)%len(q.eles)
    return x
}

func main() {
    fmt.Println(shortestPathAllKeys([]string{"@.a.#","###.#","b.A.B"}))
    fmt.Println(shortestPathAllKeys([]string{"@..aA","..B#.","....b"}))
}