package main

import (
	"fmt"
	"math"
)

func minSubArrayLen(target int, nums []int) int {
	var start, end, curSum, minLength, size = 0, 0, 0, math.MaxInt, len(nums)

	for end < size {
		if curSum+nums[end] >= target {
			minLength = min(minLength, end-start+1)
			curSum -= nums[start]
			start++
		} else {
			curSum += nums[end]
			end++
		}
	}

	if minLength == math.MaxInt {
		return 0
	}

	return minLength
}

func main() {
	nums := []int{2, 3, 1, 2, 4, 3}
	target := 7

	fmt.Println(minSubArrayLen(target, nums))
}

//1. start = 0 , end = 0 , length = 0 , sum = 0
//s = 2 < ta(7) == > end + 1 = 1
// start = 0 , end = 1 , length = 0 , sum = 2
//sum = 2 +3 = 5 > ta , end = 2
// start = 0 , end = 2 , length = 0 , sum = 5
//sum 2+3+1 = 6 < 7 , end = 3
// start = 0 , end = 3 , length = 0 , sum = 6
//sum = 6+2 = 8 > ta , sum =6 - 2 , sta = 1 ,
