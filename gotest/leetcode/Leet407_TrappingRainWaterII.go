package main

import (
	"container/heap"
    "fmt"
)

// https://leetcode.com/problems/trapping-rain-water-ii/

// Given an m x n matrix of positive integers representing the height of each unit cell in 
// a 2D elevation map, compute the volume of water it is able to trap after raining.
// Example:
//   Given the following 3x6 height map:
//     [
//       [1,4,3,1,3,2],
//       [3,2,1,3,2,4],
//       [2,3,3,2,3,1]
//     ]
//   Return 4.
// The above image represents the elevation map [[1,4,3,1,3,2],[3,2,1,3,2,4],[2,3,3,2,3,1]] before the rain.
// After the rain, water is trapped between the blocks. The total volume of water trapped is 4.
// Constraints:
//   1 <= m, n <= 110
//   0 <= heightMap[i][j] <= 20000

func trapRainWater(heightMap [][]int) int {
    m, n := len(heightMap), len(heightMap[0])
	if m<=2 || n<=2 {
		return 0
	}
    // water on the border cells must leak; add border to a heap and pop the minimal border.
    // any unvisited cell lower than min-border will trap water. after that, we also add the trapped cell as new border. 
    visited := [110][110]bool{} 
    h := tuples([][3]int{})

    // add border to heap
    for i:=0; i<m; i++ {
    	heap.Push(&h, [3]int{i, 0, heightMap[i][0]})
    	heap.Push(&h, [3]int{i, n-1, heightMap[i][n-1]})
    	visited[i][0], visited[i][n-1] = true, true
    }
    for j:=1; j<n-1; j++ {
    	heap.Push(&h, [3]int{0, j, heightMap[0][j]})
    	heap.Push(&h, [3]int{m-1, j, heightMap[m-1][j]})
    	visited[0][j], visited[m-1][j] = true, true
    }

    trapped := 0
	for h.Len()>0 {
		tuple := heap.Pop(&h).([3]int)        // get the minimal border
		i, j, height := tuple[0], tuple[1], tuple[2]
		for _, v := range [][2]int{{1,0}, {-1,0}, {0,1}, {0,-1}} {
			ni, nj := i+v[0], j+v[1]
			if ni>=0 && ni<m && nj>=0 && nj<n && !visited[ni][nj] {
				visited[ni][nj] = true
				if heightMap[ni][nj]<height {  // lower than minimal border
					trapped += height-heightMap[ni][nj]
					heap.Push(&h, [3]int{ni, nj, height})  // deem trapped water as concrete
				} else {
					heap.Push(&h, [3]int{ni, nj, heightMap[ni][nj]})
				}
			}
		}
	}
    return trapped
}

type tuples [][3]int // (row, col, height) tuple

func (t tuples) Len() int {
	return len(t)
}

func (t tuples) Swap(i, j int) {
	t[i], t[j] = t[j], t[i]
}

func (t tuples) Less(i, j int) bool {
	return t[i][2]<t[j][2]
}

func (t *tuples) Push(x interface{}) {
	*t = append(*t, x.([3]int))
}

func (t *tuples) Pop() interface{} {
	x := (*t)[len(*t)-1]
	*t = (*t)[:len(*t)-1]
	return x
}

func main() {
	fmt.Println(trapRainWater([][]int{
		{1,4,3,1,3,2},
		{3,2,1,3,2,4},
		{2,3,3,2,3,1},
	}), 4)
	fmt.Println(trapRainWater([][]int{
		{4,4,4,4,4,4,4},
		{4,2,2,2,2,2,4},
		{4,2,3,3,3,2,4},
		{4,2,3,1,3,2,4},
		{4,2,3,3,3,2,4},
		{4,2,2,2,2,2,4},
		{4,4,4,4,4,4,4},
	}), 43)
	fmt.Println(trapRainWater([][]int{
		{8,8,8,4,8,8,8,8},
		{8,2,2,2,2,2,2,8},
		{8,2,6,6,6,6,2,8},
		{8,5,6,5,5,6,5,8},
		{8,2,6,6,6,6,2,8},
		{8,2,2,2,2,2,2,8},
		{8,8,8,3,8,8,8,8},
	}), 26)
	fmt.Println(trapRainWater([][]int{
		{4,4,4,3,4,4,4},
		{4,2,2,2,2,2,4},
		{4,2,3,3,3,2,4},
		{4,2,3,1,3,2,4},
		{4,2,3,3,3,2,4},
		{4,2,2,2,2,2,4},
		{4,4,4,4,4,4,4},
	}), 18)
	fmt.Println(trapRainWater([][]int{
		{4,4,4,2,4,4,4},
		{4,2,2,2,2,2,4},
		{4,2,3,3,3,2,4},
		{4,2,3,1,3,2,4},
		{4,2,3,3,3,2,4},
		{4,2,2,2,2,2,4},
		{4,4,4,4,4,4,4},
	}), 2)
}