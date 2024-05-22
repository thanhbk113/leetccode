package main

func middleNode(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	slow := head
	fast := head

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	return slow
}
