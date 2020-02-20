package main

import "fmt"

func maxAreaOfIsland(grid [][]int) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}
	m, n := len(grid), len(grid[0])
	visited := make([]bool, m*n)
	max := 0
	stack := make([]int, m*n)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 && !visited[i*n+j] {
				area := getArea(m, n, i, j, grid, visited, stack)
				if area > max {
					max = area
				}
			}
		}
	}
	return max
}

func getArea(m, n, i, j int, grid [][]int, visited []bool, stack []int) int {
	area := 1
	index := 0
	stack[index] = i*n + j
	visited[i*n+j] = true
	for index >= 0 {
		cur := stack[index]
		index--
		curi, curj := cur/n, cur%n
		if curi-1 >= 0 && grid[curi-1][curj] == 1 && !visited[cur-n] {
			visited[cur-n] = true
			index++
			stack[index] = cur - n
			area++
		}
		if curi+1 < m && grid[curi+1][curj] == 1 && !visited[cur+n] {
			visited[cur+n] = true
			index++
			stack[index] = cur + n
			area++
		}
		if curj-1 >= 0 && grid[curi][curj-1] == 1 && !visited[cur-1] {
			visited[cur-1] = true
			index++
			stack[index] = cur - 1
			area++
		}
		if curj+1 < n && grid[curi][curj+1] == 1 && !visited[cur+1] {
			visited[cur+1] = true
			index++
			stack[index] = cur + 1
			area++
		}
	}
	return area
}

func main() {
	matrix := [][]int{
		{1, 0, 1, 0, 0},
		{1, 0, 1, 1, 1},
		{1, 1, 1, 1, 1},
		{1, 0, 0, 1, 0},
	}
	fmt.Println(maxAreaOfIsland(matrix))
	matrix = [][]int{
		{1, 1, 1, 0, 0},
		{0, 0, 1, 0, 1},
		{0, 0, 0, 1, 1},
		{1, 0, 0, 1, 1},
	}
	fmt.Println(maxAreaOfIsland(matrix))
	matrix = [][]int{
		{1, 0, 1, 0, 0, 1, 1, 1, 0},
		{1, 1, 1, 0, 0, 0, 0, 0, 1},
		{0, 0, 1, 1, 0, 0, 0, 1, 1},
		{0, 1, 1, 0, 0, 1, 0, 0, 1},
		{1, 1, 0, 1, 1, 0, 0, 1, 0},
		{0, 1, 1, 1, 1, 1, 1, 0, 1},
		{1, 0, 1, 1, 1, 0, 0, 1, 0},
		{1, 1, 1, 0, 1, 0, 0, 0, 1},
		{0, 1, 1, 1, 1, 0, 0, 1, 0},
		{1, 0, 0, 1, 1, 1, 0, 0, 0},
	}
	fmt.Println(maxAreaOfIsland(matrix))
}
