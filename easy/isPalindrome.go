package main

import (
	"fmt"
	"strings"
	"unicode"
)

func removeWhitespaceAndPunctuation(s string) string {
	var builder strings.Builder

	for _, char := range s {
		if !unicode.IsSpace(char) && !unicode.IsPunct(char) && !unicode.IsSymbol(char) {
			builder.WriteRune(char)
		}
	}

	return strings.ToLower(builder.String())
}

func isPalindrome(s string) bool {
	s = removeWhitespaceAndPunctuation(s)

	l, r := 0, len(s)-1

	for l <= r {
		if s[l] != s[r] {
			return false
		}
		l++
		r--
	}

	return true
}

func main() {
	s := "`l;`` 1o1 ??;l`"

	fmt.Println(isPalindrome(s))
}
