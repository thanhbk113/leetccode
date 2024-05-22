package main

import (
	"fmt"
)

// func quickSort(arr []int, low, high int) {
// 	if low < high {
// 		pivotIndex := partition(arr, low, high)
// 		quickSort(arr, low, pivotIndex-1)
// 		quickSort(arr, pivotIndex+1, high)
// 	}
// }

// func partition(arr []int, low, high int) int {
// 	pivot := arr[high]
// 	i := low - 1

// 	for j := low; j < high; j++ {
// 		if arr[j] < pivot {
// 			i++
// 			arr[i], arr[j] = arr[j], arr[i]
// 		}
// 	}

// 	arr[i+1], arr[high] = arr[high], arr[i+1]
// 	return i + 1
// }

// func sortColors(nums []int) {
// 	quickSort(nums, 0, len(nums)-1)
// }

func sortColors(nums []int) {
	low, mid, high := 0, 0, len(nums)-1

	for mid <= high {
		switch nums[mid] {
		case 0:
			nums[low], nums[mid] = nums[mid], nums[low]
			low++
			mid++
		case 1:
			mid++
		case 2:
			nums[mid], nums[high] = nums[high], nums[mid]
			high--
		}
	}
}

func main() {
	s := []int{2, 0, 2, 1, 1, 0}

	sortColors(s)
	fmt.Println(s)
}
