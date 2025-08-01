package _go

type Trie struct {
	trie   map[byte]*Trie
	isWord bool
}

func TrieConstructor() Trie {
	return Trie{nil, false}
}

func (trie *Trie) Insert(word string) {
	if trie.trie == nil {
		trie.trie = make(map[byte]*Trie)
	}
	firstCharacter := word[0]
	if insideTrie, ok := trie.trie[firstCharacter]; ok {
		if len(word) == 1 {
			insideTrie.isWord = true
		} else {
			insideTrie.Insert(word[1:])
		}
	} else {
		newTrie := TrieConstructor()
		if len(word) == 1 {
			newTrie.isWord = true
			trie.trie[firstCharacter] = &newTrie
		} else {
			newTrie.Insert(word[1:])
			trie.trie[firstCharacter] = &newTrie
		}

	}
}

func (trie *Trie) Search(word string) bool {
	if len(word) == 0 {
		return false
	}
	if len(word) == 1 {
		if trie.trie == nil {
			return false
		}
		if trie.trie[word[0]] == nil {
			return false
		}
		return trie.trie[word[0]].isWord
	}
	firstCharacter := word[0]
	if t, ok := trie.trie[firstCharacter]; ok {
		return t.Search(word[1:])
	} else {
		return false
	}
}

func (trie *Trie) StartsWith(prefix string) bool {
	if len(prefix) == 0 {
		return true
	}
	firstCharacter := prefix[0]
	if _, ok := trie.trie[firstCharacter]; ok {
		return trie.trie[firstCharacter].StartsWith(prefix[1:])
	} else {
		return false
	}
}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */
