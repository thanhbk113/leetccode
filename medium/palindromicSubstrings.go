package main

import "fmt"

func countSubstrings(s string) int {
	begin, end, count := 0, 0, 0

	for end < len(s) {
		if s[end] == s[begin] {
			count++
			for begin < end {
				if s[begin] != s[end] {
					count--
					break
				}
				begin++
			}
		}
		end++
	}

	return count
}

func main() {
	s := "abc"
	fmt.Println(countSubstrings(s))
}
