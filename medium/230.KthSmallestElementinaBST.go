package main

func kthSmallest230(root *TreeNode, k int) int {
	leftNode := countNode(root.Left)

	if k == leftNode+1 {
		return root.Val
	} else if k < leftNode+1 {
		return kthSmallest230(root.Left, k)
	} else {
		return kthSmallest230(root.Right, k-leftNode-1)
	}
}

func countNode(node *TreeNode) int {
	if node == nil {
		return 0
	}

	return 1 + countNode(node.Left) + countNode(node.Right)
}
