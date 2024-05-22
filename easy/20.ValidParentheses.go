package main

import "fmt"

func isValid(s string) bool {
	stack := []rune{}
	check := map[rune]rune{
		')': '(',
		']': '[',
		'}': '{',
	}

	for _, v := range s {
		if v == '(' || v == '{' || v == '[' {
			stack = append(stack, v)
		} else if v == ')' || v == '}' || v == ']' {
			if len(stack) == 0 || stack[len(stack)-1] != check[v] {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}

	return len(stack) == 0
}

func main() {
	fmt.Println(isValid("{}[]()"))
}
