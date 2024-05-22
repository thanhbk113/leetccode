package main

import (
	"math"
	"strconv"
)

func evalRPN(tokens []string) int {
	stack := []int{}

	for _, token := range tokens {
		if len(stack) >= 2 && isOperatorDigit(token) {
			oneVal := stack[len(stack)-1]
			twoVal := stack[len(stack)-2]
			stack = stack[:len(stack)-2]
			ansVal := caculatorDigit(twoVal, oneVal, token)
			stack = append(stack, ansVal)
		} else {
			tokenVal, _ := strconv.Atoi(token)
			stack = append(stack, tokenVal)
		}
	}
	return stack[len(stack)-1]
}

func isOperatorDigit(digit string) bool {
	return digit == "+" || digit == "-" || digit == "*" || digit == "/"
}

func caculatorDigit(val1, val2 int, digit string) int {
	var ans int
	if digit == "+" {
		ans = val1 + val2
	}
	if digit == "-" {
		ans = val1 - val2
	}
	if digit == "*" {
		ans = val1 * val2
	}
	if digit == "/" {
		ans = val1 / val2
	}

	return int(math.Round(float64(ans)))
}
