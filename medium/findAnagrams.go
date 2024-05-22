package main

import (
	"fmt"
)

func findAnagrams(s string, p string) []int {
	res := make([]int, 0)

	if len(s) < len(p) {
		return res
	}

	pCount, sCount := [26]int{}, [26]int{}

	window := len(p)

	for _, c := range p {
		pCount[c-'a']++
	}

	right := 0

	for right < len(s) {
		sCount[s[right]-'a']++
		if right >= window {
			sCount[s[right-window]-'a']--
		}

		if sCount == pCount {
			res = append(res, right-window+1)
		}

		right++
	}

	return res
}

func main() {
	s := "abab"
	p := "ab"

	fmt.Println(findAnagrams(s, p))
}
