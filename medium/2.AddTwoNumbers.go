package main

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	dummyNode := &ListNode{Val: 0}
	carry := 0

	for node := dummyNode; l1 != nil || l2 != nil || carry > 0; node = node.Next {
		if l1 != nil {
			carry += l1.Val
			l1 = l1.Next
		}

		if l2 != nil {
			carry += l2.Val
			l2 = l2.Next
		}
		node.Next = &ListNode(carry % 10)
		carry /= 10
	}

	return dummyNode.Next
}
