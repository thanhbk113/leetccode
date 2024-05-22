package main

import "fmt"

func twoSum(nums []int, target int) []int {
	m := make(map[int]int, len(nums))
	for i, v := range nums {
		if j, ok := m[target-v]; ok {
			return []int{j, i}
		}
		m[v] = i
	}

	return nil
}

func main() {
	fmt.Println(twoSum([]int{3, 2, 4}, 6))
}
