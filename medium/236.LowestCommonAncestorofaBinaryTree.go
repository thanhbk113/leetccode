package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func lowestCommonAncestorBinaryTree(root, p, q *TreeNode) *TreeNode {
	if root == nil || root == p || root == q {
		return root
	}

	l := dfs(root.Left, p, q)
	r := dfs(root.Right, p, q)

	if r == nil {
		return l
	} else if l == nil {
		return r
	}

	return root
}

func dfs(root, p, q *TreeNode) *TreeNode {
	if root == nil || root == p || root == q {
		return root
	}

	return dfs(root.Left, p, q)
}
