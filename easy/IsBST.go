package main

import (
	"fmt"
	"math"
)

type Node struct {
	val   int
	left  *Node
	right *Node
}

type Tree struct {
	root *Node
}

func (t *Tree) Insert(data int) {
	if t.root == nil {
		t.root = &Node{val: data}
	} else {
		t.root.Insert(data)
	}
}

func (n *Node) Insert(data int) {
	if data <= n.val {
		if n.left == nil {
			n.left = &Node{val: data}
		} else {
			n.left.Insert(data)
		}
	} else {
		if n.right == nil {
			n.right = &Node{val: data}
		} else {
			n.right.Insert(data)
		}
	}
}

func MaxVal(node *Node) int {
	if node == nil {
		return math.MinInt
	}
	value := node.val
	leftMax := MaxVal(node.left)
	rightMax := MaxVal(node.right)

	return max(value, max(leftMax, rightMax))
}

func MinVal(node *Node) int {
	if node == nil {
		return math.MinInt
	}
	value := node.val
	leftMax := MaxVal(node.left)
	rightMax := MaxVal(node.right)

	return min(value, min(leftMax, rightMax))
}

func IsBST(node *Node) bool {
	if node == nil {
		return true
	}

	if node.left != nil && MaxVal(node.left) >= node.val {
		return false
	}

	if node.right != nil && MinVal(node.right) <= node.val {
		return false
	}

	if !IsBST(node.left) || !IsBST(node.right) {
		return false
	}

	return true
}
func main() {
	var t Tree

	t.Insert(5)
	t.Insert(2)
	t.Insert(3)
	t.Insert(1)

	if IsBST(t.root) {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
