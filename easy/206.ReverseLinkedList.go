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

func (listNode *ListNode) AddNodeAfter(newVal int) {
	newNode := &ListNode{Val: newVal}
	listNode.Next = newNode
}

func reverseList(head *ListNode) *ListNode {
	var prev *ListNode

	for head != nil {
		tmp := head.Next
		head.Next = prev
		prev = head
		head = tmp
	}

	return prev
}

func main() {
	head := &ListNode{
		Val:  1,
		Next: nil,
	}
	n := 4
	curr := head
	for i := 2; i <= n; i++ {
		newNode := &ListNode{
			Val:  i,
			Next: nil,
		}
		curr.Next = newNode
		curr = newNode
	}
	reverseList(head)
}
