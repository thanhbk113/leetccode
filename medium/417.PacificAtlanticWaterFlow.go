package main

func pacificAtlantic(heights [][]int) [][]int {
	pac, atl := make(map[pair]bool), make(map[pair]bool)

	ROWS := len(heights)
	COLS := len(heights[0])

	for r := 0; r < ROWS; r++ {
		for c := 0; c < COLS; c++ {
			if r == 0 || c == 0 {
				dfs417(r, c, pac, heights[r][c], ROWS, COLS, heights)
			}
			if r == ROWS-1 || c == COLS-1 {
				dfs417(r, c, atl, heights[r][c], ROWS, COLS, heights)
			}
		}
	}

	res := make([][]int, 0)

	for r := 0; r < ROWS; r++ {
		for c := 0; c < COLS; c++ {
			PAIR := pair{r, c}
			_, ok1 := pac[PAIR]
			_, ok2 := atl[PAIR]

			if ok1 && ok2 {
				res = append(res, []int{r, c})
			}
		}
	}

	return res
}

func dfs417(r, c int, visit map[pair]bool, prevHeight, ROWS, COLS int, heights [][]int) {
	if r < 0 || c < 0 || r >= ROWS || c >= COLS {
		return
	}

	if heights[r][c] < prevHeight {
		return
	}

	PAIR := pair{r, c}

	if _, ok := visit[PAIR]; ok {
		return
	}

	visit[PAIR] = true

	dfs417(r-1, c, visit, heights[r][c], ROWS, COLS, heights)
	dfs417(r+1, c, visit, heights[r][c], ROWS, COLS, heights)
	dfs417(r, c-1, visit, heights[r][c], ROWS, COLS, heights)
	dfs417(r, c+1, visit, heights[r][c], ROWS, COLS, heights)
}
