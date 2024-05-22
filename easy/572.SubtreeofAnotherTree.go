package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isSubtree(root *TreeNode, subRoot *TreeNode) bool {
	if root == nil && subRoot == nil {
		return true
	}

	if root == nil || subRoot == nil {
		return false
	}

	queue := []*TreeNode{root}

	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]

		if isSame(curr, subRoot) {
			return true
		}

		if curr.Left != nil {
			queue = append(queue, curr.Left)
		}
		if curr.Right != nil {
			queue = append(queue, curr.Right)
		}
	}

	return false
}

func isSame(l, r *TreeNode) bool {
	if l == nil && r == nil {
		return true
	}
	if l == nil || r == nil || l.Val != r.Val {
		return false
	}

	return isSame(l.Left, r.Left) && isSame(l.Right, r.Right)
}
