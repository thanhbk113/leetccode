package main

func rightSideView(root *TreeNode) []int {
	res := []int{}
	if root == nil {
		return res
	}
	// dfs(root, &res)
	bfs(root, &res)

	return res
}

func bfs(root *TreeNode, res *[]int) {
	queue := []*TreeNode{root}

	var curr *TreeNode

	for len(queue) > 0 {
		size := len(queue)
		for i := 0; i < size; i++ {
			curr = queue[0]
			queue = queue[1:]
			if curr.Left != nil {
				queue = append(queue, curr.Left)
			}
			if curr.Right != nil {
				queue = append(queue, curr.Right)
			}
		}
		*res = append(*res, curr.Val)
	}
}
