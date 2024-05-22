package main

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxPathSum(root *TreeNode) int {
	if root == nil {
		return 0
	}

	maxSum := math.MinInt

	maxSumPath(root, &maxSum)

	return maxSum
}

func maxSumPath(root *TreeNode, maxSum *int) int {
	if root == nil {
		return 0
	}

	maxLeft := max(0, maxSumPath(root.Left, maxSum))
	maxRight := max(0, maxSumPath(root.Right, maxSum))

	m := max(maxRight, maxLeft)
	*maxSum = max(*maxSum, root.Val+maxLeft, root.Val+maxRight, root.Val+maxLeft+maxRight)
	return root.Val + m
}
