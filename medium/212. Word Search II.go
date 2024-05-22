package main

type Trie212 struct {
	children [26]*Trie212
	isEnd    bool
}

func (this *Trie212) insert(word string) {
	node := this
	for _, char := range word {
		index := char - 'a'
		if node.children[index] == nil {
			node.children[index] = &Trie212{}
		}
		node = node.children[index]
	}
	node.isEnd = true
}

func (this *Trie212) search(word string) bool {
	node := this
	for _, char := range word {
		index := char - 'a'
		if node.children[index] == nil {
			return false
		}
		node = node.children[index]
	}

	return node.isEnd
}

func (this *Trie212) searchPrefixes(word string) bool {
	node := this
	for _, char := range word {
		index := char - 'a'
		if node.children[index] == nil {
			return false
		}
		node = node.children[index]
	}

	return true
}

func findWords(board [][]byte, words []string) []string {
	root := &Trie212{}
	ans := []string{}
	vis := make(map[string]bool)
	m := len(board)
	n := len(board[0])
	for _, word := range words {
		root.insert(word)
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			helper(root, board, i, j, string(board[i][j]), &ans, vis)
		}
	}

	return ans
}

func helper(node *Trie212, board [][]byte, i, j int, word string, ans *[]string, vis map[string]bool) {
	if !node.searchPrefixes(word) {
		return
	}

	if node.search(word) && !vis[word] {
		vis[word] = true
		*ans = append(*ans, word)
	}

	directions := [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}

	temp := board[i][j]
	board[i][j] = '#'
	for _, dir := range directions {
		new_i := i + dir[0]
		new_j := j + dir[1]
		if new_i >= 0 && new_i < len(board) && new_j >= 0 && new_j < len(board[0]) && board[new_i][new_j] != '#' {
			helper(node, board, new_i, new_j, word+string(board[new_i][new_j]), ans, vis)
		}
	}
	board[i][j] = temp
}
