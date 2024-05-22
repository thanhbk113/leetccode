package main

import (
	"fmt"
)

var mapping = map[string]string{
	"2": "abc",
	"3": "def",
	"4": "ghi",
	"5": "jkl",
	"6": "mno",
	"7": "pqrs",
	"8": "tuv",
	"9": "wxyz",
}

func letterCombinations(digits string) []string {
	if digits == "" {
		return []string{}
	}

	ans := []string{}

	backtrack("", digits, &ans)

	return ans
}

func backtrack(combination string, nextDigits string, ans *[]string) {
	if len(nextDigits) == 0 {
		*ans = append(*ans, combination)
	} else {
		digit := string(nextDigits[0])
		letters := mapping[digit]

		for i := 0; i < len(letters); i++ {
			letters := string(letters[i])
			backtrack(combination+letters, nextDigits[1:], ans)
		}
	}
}

func main() {
	digits := "29"
	fmt.Println(letterCombinations(digits))
}
