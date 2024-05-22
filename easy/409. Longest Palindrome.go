package main

import "fmt"

func longestPalindrome(s string) int {
	if s == "" {
		return 0
	}
	check := make(map[rune]int)

	for _, c := range s {
		check[c]++
	}

	ans := 0

	for key, val := range check {
		fmt.Println(string(key), val)
		ans += val / 2 * 2
		if ans%2 == 0 && val%2 == 1 {
			ans += 1
		}
	}

	return ans
}
