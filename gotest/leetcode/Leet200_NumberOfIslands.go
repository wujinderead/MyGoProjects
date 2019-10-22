package leetcode

import "fmt"

func numIslands(grid [][]byte) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}
	m, n := len(grid), len(grid[0])
	visited := make([]bool, m*n)
	count := 0
	stack := make([]int, m*n)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == '1' && !visited[i*n+j] {
				visit(m, n, i, j, grid, visited, stack, &count)
			}
		}
	}
	return count
}

func visit(m, n, i, j int, grid [][]byte, visited []bool, stack []int, count *int) {
	*count++
	index := 0
	stack[index] = i*n + j
	visited[i*n+j] = true
	for index >= 0 {
		cur := stack[index]
		index--
		curi, curj := cur/n, cur%n
		if curi-1 >= 0 && grid[curi-1][curj] == '1' && !visited[cur-n] {
			visited[cur-n] = true
			index++
			stack[index] = cur - n
		}
		if curi+1 < m && grid[curi+1][curj] == '1' && !visited[cur+n] {
			visited[cur+n] = true
			index++
			stack[index] = cur + n
		}
		if curj-1 >= 0 && grid[curi][curj-1] == '1' && !visited[cur-1] {
			visited[cur-1] = true
			index++
			stack[index] = cur - 1
		}
		if curj+1 < n && grid[curi][curj+1] == '1' && !visited[cur+1] {
			visited[cur+1] = true
			index++
			stack[index] = cur + 1
		}
	}
}

func main() {
	matrix := [][]byte{
		{'1', '0', '1', '0', '0'},
		{'1', '0', '1', '1', '1'},
		{'1', '1', '1', '1', '1'},
		{'1', '0', '0', '1', '0'},
	}
	fmt.Println(numIslands(matrix))
	matrix = [][]byte{
		{'1', '1', '1', '0', '0'},
		{'0', '0', '1', '0', '1'},
		{'0', '0', '0', '1', '1'},
		{'1', '0', '0', '1', '1'},
	}
	fmt.Println(numIslands(matrix))
	matrix = [][]byte{
		{'1', '0', '1', '0', '0', '1', '1', '1', '0'},
		{'1', '1', '1', '0', '0', '0', '0', '0', '1'},
		{'0', '0', '1', '1', '0', '0', '0', '1', '1'},
		{'0', '1', '1', '0', '0', '1', '0', '0', '1'},
		{'1', '1', '0', '1', '1', '0', '0', '1', '0'},
		{'0', '1', '1', '1', '1', '1', '1', '0', '1'},
		{'1', '0', '1', '1', '1', '0', '0', '1', '0'},
		{'1', '1', '1', '0', '1', '0', '0', '0', '1'},
		{'0', '1', '1', '1', '1', '0', '0', '1', '0'},
		{'1', '0', '0', '1', '1', '1', '0', '0', '0'},
	}
	fmt.Println(numIslands(matrix))
}
