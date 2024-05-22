package main

import "math"

func findPermutationDifference(s string, t string) int {
	if len(s) != len(t) {
		return -1
	}
	posS := make(map[rune]int)
	ans := 0
	for i, c := range s {
		posS[c] = i
	}
	for i, c := range t {
		val, _ := posS[c]
		ans += int(math.Abs(float64(val - i)))
	}
	return ans
}
