package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func invertTree(root *TreeNode) *TreeNode {
	temp := root
	invertTreeHelpper(root)

	return temp
}

func invertTreeHelpper(root *TreeNode) {
	if root == nil {
		return
	}
	swapTreeNode(root)
	invertTreeHelpper(root.Left)
	invertTreeHelpper(root.Right)
}

func swapTreeNode(node *TreeNode) {
	temp := node.Left
	node.Left = node.Right
	node.Right = temp
}
