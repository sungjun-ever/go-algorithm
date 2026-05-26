package main

// 배열을 활용한 방식 - 소문자 알파벳으로만 이루어져 있는 경우
type TrieNode struct {
	children [26]*TrieNode
	isEnd    bool
}

type Trie struct {
	root *TrieNode
}

func Constructor() Trie {
	return Trie{root: &TrieNode{}}
}

func (this *Trie) Insert(word string) {
	node := this.root
	for _, w := range word {
		if node.children[w-'a'] == nil {
			node.children[w-'a'] = &TrieNode{
				children: [26]*TrieNode{},
			}
		}

		node = node.children[w-'a']
	}

	node.isEnd = true
}

func (this *Trie) Search(word string) bool {
	node := this.root

	for _, w := range word {
		if node.children[w-'a'] == nil {
			return false
		}

		node = node.children[w-'a']
	}

	return node.isEnd
}

func (this *Trie) StartsWith(prefix string) bool {
	node := this.root

	for _, w := range prefix {
		if node.children[w-'a'] == nil {
			return false
		}

		node = node.children[w-'a']
	}

	return true
}

// map을 활용한 방식
// type TrieNode struct {
//     children map[rune]*TrieNode
//     isEnd bool
// }

// type Trie struct {
//     root *TrieNode
// }

// func Constructor() Trie {
//     return Trie{
//         root: &TrieNode{
//             children: make(map[rune]*TrieNode),
//         },
//     }
// }

// func (this *Trie) Insert(word string)  {
//     node := this.root
//     for _, w := range word {
//         if node.children[w] == nil {
//             node.children[w] = &TrieNode{
//                 children: make(map[rune]*TrieNode),
//             }
//         }

//         node = node.children[w]
//     }

//     node.isEnd = true
// }

// func (this *Trie) Search(word string) bool {
//     node := this.root
//     for _, w := range word {
//         if node.children[w] == nil {
//             return false
//         }

//         node = node.children[w]
//     }

//     return node.isEnd
// }

// func (this *Trie) StartsWith(prefix string) bool {
//     node := this.root

//     for _, w := range prefix {
//         if node.children[w] == nil {
//             return false
//         }

//         node = node.children[w]
//     }

//     return true
// }
