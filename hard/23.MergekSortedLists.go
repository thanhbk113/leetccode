package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeKLists(lists []*ListNode) *ListNode {
	n := len(lists)

	if n == 0 {
		return nil
	}

	return mergeLists(lists, 0, n-1)
}

func mergeLists(lists []*ListNode, start, end int) *ListNode {
	if start == end {
		return lists[start]
	}

	mid := start + (end-start)/2

	left := mergeLists(lists, start, mid)
	right := mergeLists(lists, mid+1, end)

	return merge(left, right)
}

func merge(l1, l2 *ListNode) *ListNode {
	dummy := &ListNode{}
	curr := dummy

	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			curr.Next = l1
			l1 = l1.Next
		} else {
			curr.Next = l2
			l2 = l2.Next
		}
		curr = curr.Next
	}

	if l1 != nil {
		curr.Next = l1
	}
	if l2 != nil {
		curr.Next = l2
	}

	return dummy.Next
}
