package main

type pair struct {
	row, col int
}

func updateMatrix(mat [][]int) [][]int {
	queue := make([]pair, 0)

	m := len(mat)
	n := len(mat[0])
	for r := 0; r < m; r++ {
		for c := 0; c < n; c++ {
			if mat[r][c] == 0 {
				queue = append(queue, pair{r, c})
			} else {
				mat[r][c] = -1
			}
		}
	}

	var directions = [][]int{{-1, 0}, {0, 1}, {0, -1}, {1, 0}}

	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]

		for _, dir := range directions {
			nr := curr.row + dir[0]
			nc := curr.col + dir[1]

			if nr >= 0 && nr < m && nc >= 0 && nc < n && mat[nr][nc] == -1 {
				mat[nr][nc] = mat[curr.row][curr.col] + 1
				queue = append(queue, pair{nr, nc})
			}
		}
	}

	return mat
}
