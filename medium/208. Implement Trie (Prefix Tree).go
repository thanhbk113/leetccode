package main

type Trie struct {
	isEndOfWord bool
	children    [26]*Trie
}

func Constructor() Trie {
	return Trie{}
}

func (this *Trie) Insert(word string) {
	node := this

	for _, char := range word {
		index := char - 'a'
		if node.children[index] == nil {
			node.children[index] = &Trie{}
		}
		node = node.children[index]
	}

	node.isEndOfWord = true
}

func (this *Trie) Search(word string) bool {
	node := this

	for _, char := range word {
		index := char - 'a'
		if node.children[index] == nil {
			return false
		}
		node = node.children[index]
	}

	return node.isEndOfWord
}

func (this *Trie) StartsWith(prefix string) bool {
	node := this

	for _, char := range prefix {
		index := char - 'a'
		if node.children[index] == nil {
			return false
		}
		node = node.children[index]
	}
	return true
}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */
