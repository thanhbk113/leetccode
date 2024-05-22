package main

import "fmt"

func lengthOfLongestSubstring(s string) int {
	length, start, end, max := 0, 0, 0, 0

	m := make(map[byte]int, len(s))
	for i := 0; i < len(s); i++ {

		if _, ok := m[s[i]]; ok {
			if m[s[i]] >= start {
				start = m[s[i]] + 1
			}
		}
		m[s[i]] = i
		end = i
		length = end - start + 1

		if length > max {
			max = length
		}
	}

	return max
}

func main() {
	s := " "
	fmt.Println(lengthOfLongestSubstring(s))
}
