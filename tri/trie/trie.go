package trie

import (
	"unicode/utf8"
)

type node struct {
	char     rune
	children []*node
}

func find(r rune, children []*node) *node {
	for _, n := range children {
		if n.char == r {
			return n
		}
	}
	return nil
}

func newNode(char rune) node {
	return node{char: char, children: []*node{}}
}

func (n *node) insert(fragment string) {
	firstRune, index := utf8.DecodeRuneInString(fragment)
	rootNode := find(firstRune, n.children)
	if nil == rootNode {
		newNode := newNode(firstRune)
		rootNode = &newNode
		n.children = append(n.children, rootNode)
	}
	remainder := fragment[index:]
	if 0 < len(remainder) {
		rootNode.insert(remainder)
	}
}

func (n *node) find(fragment string) *node {
	firstRune, index := utf8.DecodeRuneInString(fragment)
	rootNode := find(firstRune, n.children)
	if nil == rootNode {
		return nil
	}

	remainder := fragment[index:]
	if 0 == len(remainder) {
		return rootNode
	}
	return rootNode.find(remainder)
}

func (n *node) getAllWords(partialWord string, answer *[]string) {
	partialWord += string(n.char)
	if 0 == len(n.children) {
		*answer = append(*answer, partialWord)
		return
	}
	for _, child := range n.children {
		child.getAllWords(partialWord, answer)
	}
}

type Trie struct {
	rootNode node
}

func NewTrie() Trie {
	return Trie{}
}

func (t *Trie) Insert(word string) {
	t.rootNode.insert(word)
}

func (t *Trie) Search(word string) bool {
	lastNode := t.rootNode.find(word)
	return nil != lastNode
}

func (t *Trie) AutoComplete(word string) []string {
	lastNode := t.rootNode.find(word)
	if nil == lastNode {
		return nil
	}
	result := []string{}
	for _, child := range lastNode.children {
		child.getAllWords(word, &result)
	}
	return result
}
