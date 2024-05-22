package main

import (
	"fmt"
	"sort"
)

func containsDuplicate(nums []int) bool {

	sort.Ints(nums)

	for i := 0; i < len(nums); i++ {
		if nums[i] == nums[i+1] {
			return true
		}
	}

	return false
}

func main() {
	fmt.Println(containsDuplicate([]int{1, 2, 3, 4}))
}
