package main

import (
	"fmt"
	"strconv"
	"unicode"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	ans := [][]int{}
	helperLevelOrder(&ans, root, 0)

	return ans
}

// pre order : root -> left -> right
func helperLevelOrder(ans *[][]int, root *TreeNode, level int) {
	if root == nil {
		return
	}

	if len(*ans) == level {
		*ans = append(*ans, []int{})
	}

	(*ans)[level] = append((*ans)[level], root.Val)

	helperLevelOrder(ans, root.Left, level+1)
	helperLevelOrder(ans, root.Right, level+1)
}

func calcNum(a, b, operation string) int {
	a1, _ := strconv.Atoi(a)
	b1, _ := strconv.Atoi(b)

	switch operation {
	case "*":
		return a1 * b1
	case "-":
		return a1 - b1
	case "+":
		return a1 + b1
	case "/":
		return a1 / b1
	}

	return 0
}

func main() {
	fmt.Println(calcNum("14", "5", "/"))
	unicode.IsPunct('a')
}
