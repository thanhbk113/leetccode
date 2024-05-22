package main

import "fmt"

func findKthLargest(nums []int, k int) int {
	minVal, maxVal := nums[0], nums[0]
	for _, num := range nums {
		minVal = min(minVal, num)
		maxVal = max(maxVal, num)
	}

	frequency := make([]int, maxVal-minVal+1)
	for _, num := range nums {
		frequency[num-minVal]++
	}

	count, index := 0, len(frequency)-1
	for count < k {
		for frequency[index] == 0 {
			index--
		}
		frequency[index]--
		count++
	}
	return index + minVal
}

func main() {
	nums := []int{3, 2, 1, 5, 6, 4}
	k := 2
	a := findKthLargest(nums, k)
	fmt.Println("🚀 ~ funcmain ~ a : ", a)
}
