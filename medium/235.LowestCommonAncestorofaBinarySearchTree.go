package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return root
	}
	if q.Val <= root.Val && p.Val >= root.Val || q.Val >= root.Val && p.Val <= root.Val {
		return root
	} else if q.Val < root.Val && p.Val < root.Val {
		return lowestCommonAncestor(root.Left, p, q)
	} else if q.Val > root.Val && p.Val > root.Val {
		return lowestCommonAncestor(root.Right, p, q)
	}
	return root
}
