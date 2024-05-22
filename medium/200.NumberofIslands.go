package main

import "container/list"

func bfs200(row, col int, vis *map[[2]int]bool, grid [][]byte) {
	q := list.New()
	q.PushBack([]int{row, col})
	(*vis)[[2]int{row, col}] = true

	for q.Len() > 0 {
		pop := q.Remove(q.Front()).([]int)
		row := pop[0]
		col := pop[1]

		directions := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}} // Các hướng đi lên, xuống, trái, phải 👈👉👆👇

		for _, dir := range directions {
			nrow := row + dir[0]
			ncol := col + dir[1]

			if nrow >= 0 && nrow < len(grid) && ncol >= 0 && ncol < len(grid[0]) && !(*vis)[[2]int{nrow, ncol}] && grid[nrow][ncol] == '1' {
				(*vis)[[2]int{nrow, ncol}] = true
				q.PushBack([]int{nrow, ncol})
			}
		}
	}
}

func dfs200(grid [][]byte, r, c int) {
	if r < 0 || r >= len(grid) || c < 0 || c >= len(grid[r]) || grid[r][c] != '1' {
		return
	}

	grid[r][c] = '2'
	dfs200(grid, r+1, c)
	dfs200(grid, r-1, c)
	dfs200(grid, r, c+1)
	dfs200(grid, r, c-1)
}

func numIslands(grid [][]byte) int {
	n := len(grid)
	m := len(grid[0])
	vis := make(map[[2]int]bool)
	cnt := 0

	for row := 0; row < n; row++ {
		for col := 0; col < m; col++ {
			if !vis[[2]int{row, col}] && grid[row][col] == '1' {
				cnt++
				bfs200(row, col, &vis, grid)
			}
		}
	}

	return cnt
}
