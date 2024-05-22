package main

import (
	"fmt"
)

func searchMatrix(matrix [][]int, target int) bool {
	r, c := len(matrix), len(matrix[0])
	left, right := 0, r*c-1

	for left <= right {
		mid := left + (right-left)/2
		mid_val := matrix[mid/c][mid%c]

		if mid_val == target {
			return true
		} else if mid_val > target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return false
}

func main() {
	matrix := [][]int{{1, 1}}
	target := 2

	a := searchMatrix(matrix, target)
	fmt.Println("🚀 ~ funcmain ~ a : ", a)
}
