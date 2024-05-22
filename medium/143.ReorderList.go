package main

func reorderList(head *ListNode) {
	if head == nil && head.Next == nil {
		return
	}
	slow := head
	fast := head

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	secondHalf := reverseList(slow.Next)
	slow.Next = nil

	p1 := head
	p2 := secondHalf

	for p2 != nil {
		next1 := p1.Next
		next2 := p2.Next
		p1.Next = p2
		p2.Next = next1
		p1 = next1
		p2 = next2
	}
}

func reverseList(head *ListNode) *ListNode {
	var prev *ListNode
	curr := head

	for curr != nil {
		next := curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}

	return prev
}
