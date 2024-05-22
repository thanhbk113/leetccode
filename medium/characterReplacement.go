package main

import (
	"fmt"
)

func characterReplacement(s string, k int) int {
	freq := make(map[byte]int)
	maxFreq := 0
	start := 0
	maxLength := 0

	for end := 0; end < len(s); end++ {
		freq[s[end]]++
		maxFreq = max(maxFreq, freq[s[end]])

		for end-start+1-maxFreq > k {
			freq[s[start]]--
			start++
		}

		maxLength = max(maxLength, end-start+1)
	}

	return maxLength
}

func main() {
	s := "ACCAACA"
	k := 2

	fmt.Println(characterReplacement(s, k))
}
