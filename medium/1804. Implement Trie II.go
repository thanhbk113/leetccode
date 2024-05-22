package main

func NewTrie() *Trie {
	return &Trie{}
}

func (t *Trie) Insert(word string) {
	node := t

	for _, char := range word {
		index := char - 'a'
		if node.children[index] == nil {
			node.children[index] = &Trie{}
		}
		node = node.children[index]
		node.cp++
	}
	node.cw++
}

func (t *Trie) CountWordsEqualTo(word string) int {
	node := t

	for _, char := range word {
		index := char - 'a'
		if node.children[index] == nil {
			return 0
		}
		node = node.children[index]
	}

	return node.cw
}

func (t *Trie) CountWordsStartingWith(prefix string) int {
	node := t

	for _, char := range prefix {
		index := char - 'a'
		if node.children[index] == nil {
			return 0
		}
		node = node.children[index]
	}

	return node.cp
}

func (t *Trie) Erase(word string) {
	node := t

	for _, char := range word {
		index := char - 'a'
		if node.children[index] == nil {
			return
		}
		node = node.children[index]
		node.cp--
	}
	node.cw--
}
