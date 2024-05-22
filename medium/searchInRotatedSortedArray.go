package main

import (
	"fmt"
)

func binarySearch(arr []int, target int) int {
	left, right := 0, len(arr)-1

	for left <= right {
		mid := (left + right) / 2

		if arr[mid] == target {
			return mid
		} else if arr[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return -1
}

func search(nums []int, target int) int {
	left, right := 0, len(nums)-1

	for left <= right {
		mid := (left + right) / 2

		if nums[mid] == target {
			return mid
		}

		// Nếu nửa đầu mảng là đã sắp xếp và target nằm trong khoảng của nửa đầu.
		if nums[left] <= nums[mid] && nums[left] <= target && target < nums[mid] {
			right = mid - 1
		} else if nums[mid] <= nums[right] && (nums[mid] < target && target <= nums[right]) {
			// Nếu nửa sau mảng là đã sắp xếp và target nằm trong khoảng của nửa sau.
			left = mid + 1
		} else if nums[left] > nums[mid] {
			// Nếu nửa đầu mảng bị xoay, di chuyển về nửa đầu.
			right = mid - 1
		} else if nums[mid] > nums[right] {
			// Nếu nửa sau mảng bị xoay, di chuyển về nửa sau.
			left = mid + 1
		} else {
			// Nếu không có trường hợp nào khớp, không tìm thấy target trong mảng.
			return -1
		}
	}

	return -1
}

func main() {
	arr := []int{4, 5, 6, 7, 0, 1, 2}
	target := 4

	fmt.Println(search(arr, target))
}
