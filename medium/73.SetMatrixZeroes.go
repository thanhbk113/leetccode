package main

func setZeroes(matrix [][]int) {
	n := len(matrix)
	m := len(matrix[0])
	checkZeroIndexs := make(map[int][]int)

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if matrix[i][j] == 0 {
				checkZeroIndexs[i] = append(checkZeroIndexs[i], j)
			}
		}
	}

	for key, val := range checkZeroIndexs {
		for _, val := range val {
			replaceMatrix(matrix, key, val)
		}
	}

}

func replaceMatrix(matrix [][]int, i, j int) {
	for k := 0; k < len(matrix); k++ {
		matrix[k][j] = 0
	}
	for l := 0; l < len((matrix)[0]); l++ {
		matrix[i][l] = 0
	}
}

func main() {
	matrix := [][]int{{0, 1, 2, 0}, {3, 4, 5, 2}, {1, 3, 1, 5}}
	setZeroes(matrix)
}
