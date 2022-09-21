package trie

type TrieNode struct {
	children  map[rune]*TrieNode
	endOfWord bool
}

func (t *TrieNode) addWord(word string) {
	for _, ch := range word {
		if _, found := t.children[ch]; !found {
			t.children[ch] = &TrieNode{
				children: make(map[rune]*TrieNode),
			}
		}
		t = t.children[ch]
	}
	t.endOfWord = true
}

func (t *TrieNode) search(word string) bool {
	for _, ch := range word {
		if _, found := t.children[ch]; !found {
			return false
		}
		t = t.children[ch]
	}
	return t.endOfWord
}
