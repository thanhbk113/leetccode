package main

import "slices"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}

	mid := slices.Index(inorder, preorder[0])

	root := &TreeNode{
		Val:   preorder[0],
		Left:  buildTree(preorder[1:mid+1], inorder[:mid]),
		Right: buildTree(preorder[mid+1:], inorder[mid+1:]),
	}

	return root
}
