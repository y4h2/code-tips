package main

type Trie struct {
	r        rune
	children map[rune]*Trie
}

func NewTrie(r rune) *Trie {
	return &Trie{
		r:        r,
		children: map[rune]*Trie{},
	}
}
