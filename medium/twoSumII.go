package main

import "fmt"

func twoSum(numbers []int, target int) []int {
	m := make(map[int]int)

	for i, v := range numbers {
		m[v] = i
	}

	for i, v := range numbers {
		if _, ok := m[target-v]; ok {
			return []int{i + 1, m[target-v] + 1}
		}
	}

	return []int{}
}

func main() {
	nums := []int{2, 7, 11, 15}
	target := 9
	fmt.Println(twoSumIInumbers(nums, target))
}
