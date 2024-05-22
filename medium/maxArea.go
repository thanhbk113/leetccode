package main

import "fmt"

func minHigh(a, b int) int {
	if a < b {
		return a
	}

	return b
}

func maxArea(height []int) int {

	left := 0
	right := len(height) - 1
	maxArea := 0

	for left < right {
		h := minHigh(height[left], height[right])

		w := right - left

		area := h * w

		if maxArea < area {
			maxArea = area
			left++
		} else {
			right--
		}
	}

	return maxArea
}

func main() {
	height := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	fmt.Println(maxArea(height))
}
