package main

func deleteNode(node *ListNode) {
	if node.Next == nil {
		node = nil
	}

	node.Val = node.Next.Val
	node.Next = node.Next.Next
}
