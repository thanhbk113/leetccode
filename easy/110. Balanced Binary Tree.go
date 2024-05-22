package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}

	heightL := getHeight(root.Left)
	heightR := getHeight(root.Right)

	diffHeight := heightL - heightR

	if diffHeight < -1 || diffHeight > 1 {
		return false
	}

	isBalancedLeft := true
	isBalancedRight := true

	isBalancedLeft = isBalanced(root.Left)
	isBalancedRight = isBalanced(root.Right)

	return isBalancedLeft && isBalancedRight
}

func getHeight(node *TreeNode) int {
	if node == nil {
		return 0
	}

	return 1 + max(getHeight(node.Left), getHeight(node.Right))
}
