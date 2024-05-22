package main

import "fmt"

func generateParenthesis(n int) []string {
	ans := make([]string, 0)

	recursionParenthesis(&ans, n, 0, 0, "")

	return ans
}

func recursionParenthesis(ans *[]string, n, open, close int, cur_string string) {
	if len(cur_string) == n*2 {
		*ans = append(*ans, cur_string)
		return
	}

	if open < n {
		recursionParenthesis(ans, n, open+1, close, cur_string+"(")
	}
	if close < open {
		recursionParenthesis(ans, n, open, close+1, cur_string+")")
	}
}
func main() {
	fmt.Println(generateParenthesis(3))
}
