package main

import (
	"fmt"
)

func countSubstrings(s string) int {

	res, size := 0, len(s)

	for i := 0; i < size; i++ {

		res += countPal(i, i, size, s)

		res += countPal(i, i+1, size, s)
	}
	return res
}

func countPal(l, r, size int, s string) int {
	res := 0
	for l >= 0 && r < size && s[l] == s[r] {
		res++
		l--
		r++
	}

	return res
}

func main() {
	s := "aaa"
	fmt.Println(countSubstrings(s))
}
