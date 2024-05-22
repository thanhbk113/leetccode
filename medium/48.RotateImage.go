package main

// func rotate(matrix [][]int) {
// 	n := len(matrix)
// 	ans := make([][]int, n)
// 	for i := range ans {
// 		ans[i] = make([]int, n)
// 	}

// 	for i := 0; i < n; i++ {
// 		for j := 0; j < n; j++ {
// 			ans[j][n-i-1] = matrix[i][j]
// 		}
// 	}

//		copy(matrix, ans)
//	}
//
// TC:(O N^2)
// SC:(O N^2)
// ---------------------------------------------------------------------------
func rotate(matrix [][]int) {
	n := len(matrix)

	// Transpose the matrix
	for i := 0; i < n; i++ {
		for j := i; j < n; j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}

	// Reverse each row
	for i := 0; i < n; i++ {
		for j := 0; j < n/2; j++ {
			matrix[i][j], matrix[i][n-j-1] = matrix[i][n-j-1], matrix[i][j]
		}
	}
}

// 	copy(matrix, ans)
// }
//TC:(O N^2)
//SC:(O 1)

func main() {
	matrix := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	rotate(matrix)
}
