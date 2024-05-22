package main

import (
	"fmt"
)

func canConstruct(ransomNote string, magazine string) bool {
	magazineTable := make(map[rune]int)

	for _, c := range magazine {
		magazineTable[c]++
	}

	for _, c := range ransomNote {
		if magazineTable[c] == 0 {
			return false
		}
		magazineTable[c]--
	}

	return true
}

func main() {
	ransomNote := "aac"
	magazine := "baa"

	fmt.Println(canConstruct(ransomNote, magazine))
}
