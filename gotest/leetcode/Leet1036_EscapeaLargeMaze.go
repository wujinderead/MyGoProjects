package main

import "fmt"

// https://leetcode.com/problems/escape-a-large-maze/

// In a 1 million by 1 million grid, the coordinates of each grid square are (x,y)
// with 0 <= x, y < 10^6.
// We start at the source square and want to reach the target square. Each move,
// we can walk to a 4-directionally adjacent square in the grid that isn't in the
// given list of blocked squares.
// Return true if and only if it is possible to reach the target square through
// a sequence of moves.
// Example 1:
//   Input: blocked = [[0,1],[1,0]], source = [0,0], target = [0,2]
//   Output: false
//   Explanation:
//     The target square is inaccessible starting from the source square, because we
//     can't walk outside the grid.
// Example 2:
//   Input: blocked = [], source = [0,0], target = [999999,999999]
//   Output: true
//   Explanation:
//     Because there are no blocked cells, it's possible to reach the target square.
// Note:
//   0 <= blocked.length <= 200
//   blocked[i].length == 2
//   0 <= blocked[i][j] < 10^6
//   source.length == target.length == 2
//   0 <= source[i][j], target[i][j] < 10^6
//   source != target

func isEscapePossible(blocked [][]int, source []int, target []int) bool {
	// through the board is very large, the blocks are very sparse.
	// if source can't reach the target, either is source circled by blocks or target.
	// so the area circled won't be larger than 20000
	limit := 25000 // a conservative estimate of largest area
	mask := (1 << 32) - 1
	stack := make([]int, 1, 5000)
	blocks := make(map[int]struct{})
	for i := range blocked {
		blocks[(blocked[i][0]<<32)+blocked[i][1]] = struct{}{}
	}
	// test if source blocked
	stack[0] = (source[0] << 32) + source[1]
	size1 := 1
	visited := make(map[int]struct{})
	visited[stack[0]] = struct{}{}
	for len(stack) > 0 && size1 < limit {
		cur := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		i, j := cur>>32, cur&mask
		for _, v := range [][2]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}} {
			ni, nj := i+v[0], j+v[1]
			if ni >= 0 && nj >= 0 && ni < 1000000 && nj < 1000000 {
				if _, ok := blocks[(ni<<32)+nj]; ok {
					continue // ignore blocked
				}
				if _, ok := visited[(ni<<32)+nj]; !ok {
					if ni == target[0] && nj == target[1] {
						return true
					}
					visited[(ni<<32)+nj] = struct{}{}
					stack = append(stack, (ni<<32)+nj)
					size1++
				}
			}
		}
	}

	// test if target blocked
	stack = stack[:1]
	stack[0] = (target[0] << 32) + target[1]
	size2 := 1
	visited = make(map[int]struct{})
	visited[stack[0]] = struct{}{}
	for len(stack) > 0 && size2 < limit {
		cur := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		i, j := cur>>32, cur&mask
		for _, v := range [][2]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}} {
			ni, nj := i+v[0], j+v[1]
			if ni >= 0 && nj >= 0 && ni < 1000000 && nj < 1000000 {
				if _, ok := blocks[(ni<<32)+nj]; ok {
					continue // ignore blocked
				}
				if _, ok := visited[(ni<<32)+nj]; !ok {
					if ni == source[0] && nj == source[1] {
						return true
					}
					visited[(ni<<32)+nj] = struct{}{}
					stack = append(stack, (ni<<32)+nj)
					size2++
				}
			}
		}
	}
	// if source and target are both non-blocked, they can finally meat
	// if they are blocked in the same area, they have already met during traverse
	return !(size1 < limit) && !(size2 < limit)
}

func main() {
	fmt.Println(isEscapePossible([][]int{{0, 1}, {1, 0}}, []int{0, 0}, []int{0, 2}))
	fmt.Println(isEscapePossible([][]int{}, []int{0, 0}, []int{999999, 999999}))
	fmt.Println(isEscapePossible([][]int{{691938, 300406}, {710196, 624190}, {858790, 609485}, {268029, 225806}, {200010, 188664}, {132599, 612099}, {329444, 633495}, {196657, 757958}, {628509, 883388}}, []int{655988, 180910}, []int{267728, 840949}))
}
