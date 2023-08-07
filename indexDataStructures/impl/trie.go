package impl

import "github.com/rahularcota/address-book/constants"

type TrieNode struct {
	children map[rune]*TrieNode
	ids      []int
}

type Trie struct {
	root *TrieNode
}

func NewTrie() *Trie {
	return &Trie{
		root: &TrieNode{
			children: make(map[rune]*TrieNode),
			ids:      make([]int, 0),
		},
	}
}

func (t *Trie) Insert(word string, id int) {
	node := t.root
	for _, ch := range word {
		child, ok := node.children[ch]
		if !ok {
			child = &TrieNode{
				children: make(map[rune]*TrieNode),
				ids:      make([]int, 0),
			}
			node.children[ch] = child
		}
		node = child
	}
	node.ids = append(node.ids, id)
}

//func (t *Trie) Search(word string) []int {
//	node := t.root
//	for _, ch := range word {
//		child, ok := node.children[ch]
//		if !ok {
//			return nil
//		}
//		node = child
//	}
//	return node.ids
//}

func (t *Trie) startsWithTrieNode(prefix string) *TrieNode {
	node := t.root
	for _, ch := range prefix {
		child, ok := node.children[ch]
		if !ok {
			return nil
		}
		node = child
	}
	return node
}

func (t *Trie) prefixRec(node *TrieNode, count int) []int {
	ids := node.ids
	for _, value := range node.children {
		ids = append(ids, t.prefixRec(value, count+len(ids))...)
		if len(ids) > constants.MaxResults {
			break
		}
	}
	return ids
}

func (t *Trie) Search(prefix string) []int {
	node := t.startsWithTrieNode(prefix)
	if node == nil {
		return nil
	}
	ids := t.prefixRec(node, 0)
	if len(ids) > constants.MaxResults {
		return ids[:constants.MaxResults]
	}
	return ids
}
