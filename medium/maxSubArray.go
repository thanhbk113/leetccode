package main

import (
	"fmt"
	"math"
)

func maxSubArray(nums []int) int {
	maxSubArray := math.MinInt
	curMax := 0

	for i := 0; i < len(nums); i++ {
		curMax += nums[i]
		if maxSubArray < curMax {
			maxSubArray = curMax
		}
		if curMax < 0 {
			curMax = 0
		}
	}

	return maxSubArray
}

func main() {

	fmt.Println(maxSubArray([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}))
}
