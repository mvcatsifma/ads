package p208

// Node represents a single node in the Trie data structure.
// Each node maps characters (as rune) to child nodes, forming a tree of prefixes.
// The Value field stores the complete word if this node marks the end of a valid word.
type Node struct {
	Children map[rune]*Node // Maps characters to child nodes
	Value    string         // Stores the complete word if this node is a word endpoint
}

// Trie implements a prefix tree data structure for efficient string storage and retrieval.
// Supports insertion, exact word search, and prefix matching in O(m) time where m is string length.
type Trie struct {
	Root *Node // Root node of the trie (empty node)
}

// Constructor creates and returns a new empty Trie with an initialized root node.
func Constructor() Trie {
	return Trie{
		Root: createNewNode(),
	}
}

// Insert adds a word to the Trie by creating nodes for each character if they don't exist.
// Time Complexity: O(m) where m is the length of the word.
func (t *Trie) Insert(word string) {
	n := t.Root

	// Traverse/create path for each character in word
	for _, c := range word {
		if n.Children[c] == nil {
			n.Children[c] = createNewNode()
		}
		n = n.Children[c]
	}

	// Mark the end of word by storing the complete word
	n.Value = word
}

// Search returns true if the exact word exists in the Trie, false otherwise.
// Traverses the trie following the word's characters and checks if the final node marks a valid word.
// Time Complexity: O(m) where m is the length of the word.
func (t *Trie) Search(word string) bool {
	n := t.Root

	// Traverse trie following word's characters
	for _, c := range word {
		if n.Children[c] == nil {
			return false // Character not found, word doesn't exist
		}
		n = n.Children[c]
	}

	// Check if final node marks a valid word (not just a prefix)
	return n.Value == word
}

// StartsWith returns true if any word in the Trie starts with the given prefix.
// Only checks if the prefix path exists; doesn't require a complete word.
// Time Complexity: O(m) where m is the length of the prefix.
func (t *Trie) StartsWith(prefix string) bool {
	n := t.Root

	// Traverse trie following prefix's characters
	for _, c := range prefix {
		if n.Children[c] == nil {
			return false // Character not found, prefix doesn't exist
		}
		n = n.Children[c]
	}

	return true // Prefix path exists
}

// createNewNode creates and returns a new empty Trie node.
func createNewNode() *Node {
	return &Node{
		Children: make(map[rune]*Node),
		Value:    "",
	}
}
