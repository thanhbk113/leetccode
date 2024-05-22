package main

import (
	"bytes"
	"strconv"
	"strings"
)

type Codec struct {
}

func Constructor297() Codec {
	return Codec{}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	var res bytes.Buffer

	var dfs func(node *TreeNode)

	dfs = func(node *TreeNode) {
		if node == nil {
			res.WriteString("N,")
		} else {
			res.WriteString(strconv.Itoa(node.Val))
			res.WriteString(",")
			dfs(node.Left)
			dfs(node.Right)
		}
	}
	dfs(root)
	return res.String()
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	var root *TreeNode
	index := 0

	tokens := strings.Split(data, ",")

	var dfs func(node *TreeNode) *TreeNode

	dfs = func(node *TreeNode) *TreeNode {
		if index >= len(tokens) {
			return nil
		}

		token := tokens[index]
		index++
		if token == "N" {
			return nil
		}

		val, _ := strconv.Atoi(token)
		node = &TreeNode{Val: val}
		node.Left = dfs(node.Left)
		node.Right = dfs(node.Right)

		return node
	}

	root = dfs(nil)

	return root
}
