package main

import "fmt"

func search(nums []int, target int) int {
	low, high := 0, len(nums)-1

	for low <= high {
		mid := (low + high) / 2
		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	return -1
}

func main() {

	nums := []int{-1, 0, 3, 5, 9, 12}
	target := 9

	a := search(nums, target)
	fmt.Println("🚀 ~ funcmain ~ a : ", a)
}
