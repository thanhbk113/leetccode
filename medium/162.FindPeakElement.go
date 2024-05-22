package main

import (
	"fmt"
)

func findPeakElement(nums []int) int {
	left, right := 1, len(nums)-2

	if len(nums) == 1 {
		return 0
	}

	if len(nums) >= 2 && nums[0] > nums[1] {
		return 0
	} else if len(nums) >= 2 && nums[len(nums)-1] > nums[len(nums)-2] {
		return len(nums) - 1
	}

	for left <= right {
		mid := left + (right-left)/2
		if nums[mid-1] < nums[mid] && nums[mid] > nums[mid+1] {
			return mid
		} else if nums[mid-1] < nums[mid] && nums[mid] < nums[mid+1] {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return -1
}

func main() {
	nums := []int{3, 4, 3, 2, 1}

	a := findPeakElement(nums)
	fmt.Println("🚀 ~ funcmain ~ a : ", a)
}
