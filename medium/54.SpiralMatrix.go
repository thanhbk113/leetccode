package main

const (
	RIGHT = "RIGHT"
	LEFT  = "LEFT"
	UP    = "UP"
	DOWN  = "DOWN"
)

func spiralOrder(matrix [][]int) []int {
	i, j, k := 0, 0, 0
	m, n := len(matrix), len(matrix[0])
	ans := make([]int, m*n)

	direction := RIGHT

	for k < m*n {
		ans[k] = matrix[i][j]
		switch direction {
		case RIGHT:
			if i+j == n-1 {
				direction = DOWN
				i++
			} else {
				j++
			}
		case DOWN:
			if n-j == m-i {
				direction = LEFT
				j--
			} else {
				i++
			}
		case LEFT:
			if i+j == m-1 {
				direction = UP
				i--
			} else {
				j--
			}
		case UP:
			if i-j == 1 {
				direction = RIGHT
				j++
			} else {
				i--
			}
		}
		k++
	}

	return ans
}

func main() {
	matrix := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	spiralOrder(matrix)
}
