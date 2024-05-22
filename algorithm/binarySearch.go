package main

import "fmt"

func binarySearch(arr []int, target int) int {
	low, high := 0, len(arr)-1

	for low <= high {
		mid := (low + high) / 2
		if arr[mid] == target {
			return mid
		} else if arr[mid] < target {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	return -1
}

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8}
	target := 5
	fmt.Println("Array:", arr)
	fmt.Println("Target:", target)
	index := binarySearch(arr, target)
	if index != -1 {
		fmt.Println("Target found at index:", index)
	} else {
		fmt.Println("Target not found in the array.")
	}
}
