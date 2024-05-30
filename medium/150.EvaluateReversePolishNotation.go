package main

import "strconv"

func evalRPN(tokens []string) int {
	stack := make([]int, 0)

	for _, val := range tokens {
		if IsPunct(val) && len(stack) > 1 {
			n := len(stack) - 1
			b := stack[n]
			a := stack[n-1]
			stack = stack[:n-1]
			stack = append(stack, calcNum(a, b, val))
		} else {
			num, _ := strconv.Atoi(val)
			stack = append(stack, num)
		}
	}

	return stack[0]
}

func IsPunct(o string) bool {
	return o == "+" || o == "-" || o == "*" || o == "/"
}

func calcNum(a, b int, operation string) int {
	switch operation {
	case "*":
		return a * b
	case "-":
		return a - b
	case "+":
		return a + b
	case "/":
		return a / b
	}

	return 0
}
