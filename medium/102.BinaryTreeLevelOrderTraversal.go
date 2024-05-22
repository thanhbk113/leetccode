package main

func levelOrder(root *TreeNode) [][]int {
	res := [][]int{}
	if root == nil {
		return res
	}

	queue := []*TreeNode{root}

	for len(queue) > 0 {
		size := len(queue) //size before push left and right child

		curLevel := []int{}
		for i := 0; i < size; i++ {
			curr := queue[0]
			queue = queue[1:]
			curLevel = append(curLevel, curr.Val)

			if curr.Left != nil {
				queue = append(queue, curr.Left)
			}

			if curr.Right != nil {
				queue = append(queue, curr.Right)
			}
		}
		res = append(res, curLevel)
	}

	return res
}
