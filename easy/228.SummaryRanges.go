package main

import (
	"fmt"
	"strconv"
)

func summaryRanges(nums []int) []string {
	ans := []string{}

	if len(nums) == 0 {
		return ans
	}

	curr := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] > nums[i-1]+1 {
			ans = appendRange(ans, curr, nums[i-1])
			curr = nums[i]
		}
	}

	ans = appendRange(ans, curr, nums[len(nums)-1])
	return ans
}

func appendRange(ans []string, start, end int) []string {
	if start == end {
		ans = append(ans, strconv.Itoa(start))
	} else {
		ans = append(ans, strconv.Itoa(start)+"->"+strconv.Itoa(end))
	}
	return ans
}

func main() {
	nums := []int{0, 2, 3, 4, 5, 6, 8}
	a := summaryRanges(nums)
	fmt.Println("🚀 ~ funcmain ~ a : ", a)
}
