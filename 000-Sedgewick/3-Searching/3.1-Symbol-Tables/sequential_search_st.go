package symboltables

// SequentialSearchST implements a symbol table using an unordered linked list.
// Uses sequential search - scans through list linearly to find keys.
//
// Performance characteristics:
// - Search: O(N) - must scan entire list in worst case
// - Insert: O(N) - must search for key first
// - Delete: O(N) - must search to find key
// - Space: O(N) - one node per key-value pair
//
// Best for: Small tables (< 10 items) or when simplicity is priority
type SequentialSearchST struct {
	first *node // head of linked list
	n     int   // number of key-value pairs
}

// node represents a key-value pair in the linked list
type node struct {
	key   string
	value int
	next  *node
}

// NewSequentialSearchST creates a new empty symbol table
func NewSequentialSearchST() *SequentialSearchST {
	return &SequentialSearchST{}
}

// put inserts a key-value pair into the table
// If key already exists, updates its value
// If value is nil/zero, deletes the key
func (st *SequentialSearchST) put(key string, value int) {
	if value == 0 {
		st.delete(key)
		return
	}
	for x := st.first; x != nil; x = x.next {
		if key == x.key {
			x.value = value
			return
		}
	}
	st.first = &node{
		key:   key,
		value: value,
		next:  st.first,
	}
	st.n++
}

// get returns the value associated with key
// Returns (value, true) if key exists, (0, false) if not found
func (st *SequentialSearchST) get(key string) (int, bool) {
	for x := st.first; x != nil; x = x.next {
		if key == x.key {
			return x.value, true
		}
	}
	return 0, false
}

// delete removes key from the table
func (st *SequentialSearchST) delete(key string) {
	var prev *node
	for x := st.first; x != nil; x = x.next {
		if x.key == key {
			if prev == nil {
				st.first = x.next
			} else {
				prev.next = x.next
			}
			st.n--
			return
		}
		prev = x
	}
}

// contains returns true if key is in the table
func (st *SequentialSearchST) contains(key string) bool {
	_, found := st.get(key)
	return found
}

// isEmpty returns true if table is empty
func (st *SequentialSearchST) isEmpty() bool {
	return st.n == 0
}

// size returns the number of key-value pairs
func (st *SequentialSearchST) size() int {
	return st.n
}

// keys returns all keys in the table (unordered)
func (st *SequentialSearchST) keys() []string {
	var retVal []string
	for x := st.first; x != nil; x = x.next {
		retVal = append(retVal, x.key)
	}
	return retVal
}
