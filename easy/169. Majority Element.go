package main

func majorityElement(nums []int) int {
	m := make(map[int]int)
	for _, v := range nums {
		if _, ok := m[v]; ok {
			m[v]++
		} else {
			m[v] = 1
		}
	}

	for k, v := range m {
		if v > len(nums)/2 {
			return k
		}
	}

	return 0
}
