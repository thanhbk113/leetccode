package main

func orangesRotting(grid [][]int) int {
	rots := make([]pair, 0)
	freshOrange := 0
	minutes := 0
	m := len(grid)
	n := len(grid[0])

	for r := 0; r < m; r++ {
		for c := 0; c < n; c++ {
			if grid[r][c] == 2 {
				rots = append(rots, pair{r, c})
			} else if grid[r][c] == 1 {
				freshOrange++
			}
		}
	}

	directions := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

	for len(rots) > 0 && freshOrange > 0 {
		newRots := make([]pair, 0)
		for _, rot := range rots {
			for _, dir := range directions {
				nr := rot.row + dir[0]
				nc := rot.col + dir[1]

				if nr >= 0 && nr < m && nc >= 0 && nc < n && grid[nr][nc] == 1 {
					grid[nr][nc] = 2
					freshOrange--
					newRots = append(newRots, pair{nr, nc})
				}
			}
		}
		rots = newRots
		minutes++
	}

	if freshOrange > 0 {
		return -1
	}
	return minutes
}
