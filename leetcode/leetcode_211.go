package main

type TrieNode struct {
	children [26]*TrieNode
	isEnd    bool
}

type WordDictionary struct {
	root *TrieNode
}

func Constructor() WordDictionary {
	return WordDictionary{root: &TrieNode{}}
}

func (this *WordDictionary) AddWord(word string) {
	node := this.root

	for _, ch := range word {
		if node.children[ch-'a'] == nil {
			node.children[ch-'a'] = &TrieNode{children: [26]*TrieNode{}}
		}

		node = node.children[ch-'a']
	}

	node.isEnd = true
}

func (this *WordDictionary) Search(word string) bool {
	var dfs func(node *TrieNode, idx int) bool

	dfs = func(node *TrieNode, idx int) bool {
		if idx == len(word) {
			return node.isEnd
		}

		ch := word[idx]

		if ch == '.' {
			for _, child := range node.children {
				if child != nil && dfs(child, idx+1) {
					return true
				}
			}

			return false
		}

		next := node.children[ch-'a']

		if next == nil {
			return false
		}

		return dfs(next, idx+1)
	}

	return dfs(this.root, 0)
}
