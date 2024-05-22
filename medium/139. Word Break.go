package main

import "fmt"

type TrieNode struct {
	children [26]*TrieNode
	isEnd    bool
}

func insert139(root *TrieNode, word string) {
	node := root
	for _, ch := range word {
		index := ch - 'a'
		if node.children[index] == nil {
			node.children[index] = &TrieNode{}
		}
		node = node.children[index]
	}
	node.isEnd = true
}

func search139(root *TrieNode, word string) bool {
	node := root
	for _, ch := range word {
		index := ch - 'a'
		if node.children[index] == nil {
			return false
		}
		node = node.children[index]
	}
	return node != nil && node.isEnd
}

func wordBreak(s string, wordDict []string) bool {
	root := &TrieNode{}
	for _, word := range wordDict {
		insert139(root, word)
	}
	dp := make([]bool, len(s)+1)
	dp[0] = true
	for i := 1; i <= len(s); i++ {
		for j := 0; j < i; j++ {
			fmt.Println("🚀 ~ ifdp[j]&&search ~ s[j:i] : ", s[j:i])
			if dp[j] && search139(root, s[j:i]) {
				dp[i] = true
				break
			}
		}
	}
	return dp[len(s)]
}
