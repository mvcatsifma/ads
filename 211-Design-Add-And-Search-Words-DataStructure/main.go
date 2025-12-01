package p211

// Node represents a single node in the Trie data structure for WordDictionary.
// Each node maps characters (runes) to child nodes, forming a tree of words.
// IsWord marks whether this node represents the end of a valid word.
type Node struct {
	Children map[rune]*Node // Maps characters to child nodes
	Value    string         // Stores the character(s) at this node
	IsWord   bool           // Marks if this node is the end of a valid word
}

// WordDictionary implements a word search structure supporting both exact matches and wildcard searches.
// The dot character '.' acts as a wildcard matching any single character.
// Example: After adding "bad", "dad", "mad", searching ".ad" returns true (matches all three).
type WordDictionary struct {
	Root *Node // Root node of the trie (starting point for all searches)
}

// Constructor creates and returns a new empty WordDictionary with an initialized root node.
func Constructor() WordDictionary {
	node := createNode(0)
	node.Value = ""
	return WordDictionary{
		Root: node,
	}
}

// AddWord adds a word to the dictionary by creating nodes for each character if they don't exist.
// Marks the final node as IsWord=true to indicate a valid word endpoint.
// Time Complexity: O(m) where m is the length of the word.
func (w *WordDictionary) AddWord(word string) {
	cur := w.Root

	// Traverse/create path for each character in word
	for _, c := range word {
		if cur.Children[c] == nil {
			cur.Children[c] = createNode(c)
		}
		cur = cur.Children[c]
	}

	// Mark end of word
	cur.Value = word
	cur.IsWord = true
}

// Search returns true if the word exists in the dictionary, supporting wildcard matching.
// The dot character '.' matches any single character at that position.
// Uses DFS with backtracking to explore all possible paths when encountering wildcards.
// Time Complexity: O(n * 26^m) worst case where n is number of nodes, m is wildcard count.
// Space Complexity: O(m) for recursion stack where m is word length.
func (w *WordDictionary) Search(word string) bool {
	// Recursive DFS helper: j is current position in word, root is current trie node
	var dfs func(j int, root *Node) bool
	dfs = func(j int, root *Node) bool {
		cur := root

		// Traverse through remaining characters in word
		for i := j; i < len(word); i++ {
			c := word[i]

			if c == '.' {
				// Wildcard: try all possible children via backtracking
				for _, child := range cur.Children {
					if dfs(i+1, child) {
						return true
					}
				}
				return false // No child path led to a valid word
			} else {
				// Exact character: must exist in trie
				if cur.Children[rune(c)] == nil {
					return false
				}
				cur = cur.Children[rune(c)]
			}
		}

		// Reached end of word: check if current node marks a valid word
		return cur.IsWord
	}

	return dfs(0, w.Root)
}

// createNode creates and returns a new trie node with initialized children map.
func createNode(c rune) *Node {
	return &Node{
		Children: make(map[rune]*Node),
		Value:    string(c),
	}
}
