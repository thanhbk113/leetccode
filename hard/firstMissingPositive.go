package main

import (
	"fmt"
)

func firstMissingPositive(nums []int) int {
	n := len(nums)

	if n == 0 {
		return 0
	}
	m := make(map[int]int)

	for i := 0; i < n; i++ {
		m[nums[i]] = i
	}

	for i := 1; i <= n; i++ {
		if _, ok := m[i]; !ok {
			return i
		}
	}

	return n + 1
}

func main() {
	nums := []int{2}
	fmt.Println(firstMissingPositive(nums))

}
