package main

import (
	"fmt"
)

func longestPalindrome(s string) string {
	n := len(s)

	startStr, endStr := 0, 0

	for i := 0; i < n; i++ {
		findPali(s, i, i, &startStr, &endStr)
		findPali(s, i, i+1, &startStr, &endStr)
	}

	return s[startStr : endStr+1]
}

func findPali(s string, l, r int, startStr, endStr *int) {
	for l >= 0 && r < len(s) && s[l] == s[r] {
		if *endStr-*startStr < r-l {
			*startStr, *endStr = l, r
		}
		l--
		r++
	}
}

func main() {
	s := "babad"

	fmt.Println(longestPalindrome(s))

}
