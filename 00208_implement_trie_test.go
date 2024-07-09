package leetcode

import (
	"testing"
)

func TestTrieConstructor(t *testing.T) {
	trie := TrieConstructor()
	trie.Insert("apple")
	findApple := trie.Search("apple")
	if !findApple {
		t.Fatal()
	}
	trie.Search("app")
	trie.StartsWith("app")
	trie.Insert("app")
	trie.Search("app")

	newTrie := TrieConstructor()
	newTrie.Insert("hotdog")
	newTrie.StartsWith("dog")
}
