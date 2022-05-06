package main

import "log"

const (
	AlphabetSize = 26
)

type Trie struct {
	root *Node
}

type Node struct {
	value    string
	end      bool
	Children map[string]*Node
}

func (t *Trie) Insert(word string) {
	if t.root == nil {
		t.root = &Node{}
	}
	current := t.root
	for index, l := range word {
		letter := string(l)
		if child, exist := current.Children[letter]; exist {
			current = child
			continue
		}
		childToAdd := &Node{value: letter, Children: map[string]*Node{}}
		if index == len(word)-1 {
			childToAdd.end = true
		}
		current.Children[letter] = childToAdd
		current = childToAdd
	}
}

func (t *Trie) Search(word string) bool {
	if t.root == nil {
		return false
	}
	current := t.root
	for _, l := range word {
		letter := string(l)
		if child, exist := current.Children[letter]; exist {
			current = child
			continue
		} else {
			return false
		}
	}
	if !current.end {
		return false
	}
	return true
}

func initTrie() {
	t := Trie{root: &Node{Children: map[string]*Node{}}}
	t.Insert("Hello")
	log.Print(t.Search("Hello"))
	log.Print(t.Search("Hell"))

}
