package main

import (
	"fmt"
	"strconv"
)

func countFreqChar(s string) string {
	count := make([]int, 26)

	for _, val := range s {
		count[val-'a']++
	}

	key := ""

	for _, c := range count {
		key += strconv.Itoa(c) + "/"
	}
	return key
}

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	return countFreqChar(s) == countFreqChar(t)
}

func main() {
	s := "aacc"
	t := "ccac"

	fmt.Println(isAnagram(s, t))
}
