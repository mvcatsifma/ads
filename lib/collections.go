package lib

// IntQueue represents a FIFO (First In, First Out) queue data structure for integers.
// Provides efficient O(1) enqueue and dequeue operations.
//
// Implementation uses a slice with a head pointer to avoid expensive slice reallocations
// on every dequeue. The head pointer tracks the index of the first element.
//
// Time Complexity:
//   - Enqueue: O(1) amortized
//   - Dequeue: O(1)
//   - IsEmpty: O(1)
//
// Space Complexity: O(n) where n is the number of elements
type IntQueue struct {
	items []int // Underlying slice storing queue elements
	head  int   // Index of first element (front of queue)
}

// Enqueue adds an item to the back of the queue.
// Time Complexity: O(1) amortized.
func (q *IntQueue) Enqueue(item int) {
	q.items = append(q.items, item)
}

// Dequeue removes and returns the item from the front of the queue.
// Panics if the queue is empty (caller should check IsEmpty() first).
// Time Complexity: O(1).
func (q *IntQueue) Dequeue() int {
	if q.IsEmpty() {
		panic("dequeue from empty queue")
	}
	item := q.items[q.head]
	q.head++
	return item
}

// IsEmpty returns true if the queue contains no elements.
// Time Complexity: O(1).
func (q *IntQueue) IsEmpty() bool {
	return q.head >= len(q.items)
}

// IntStack represents a LIFO (Last In, First Out) stack data structure for integers.
// Provides efficient O(1) push and pop operations.
//
// Time Complexity:
//   - Push: O(1) amortized
//   - Pop: O(1)
//   - Peek: O(1)
//   - IsEmpty: O(1)
//
// Space Complexity: O(n) where n is the number of elements
type IntStack struct {
	items []int // Underlying slice storing stack elements
}

// NewIntStack creates a new stack with optional pre-allocated capacity.
// Pre-allocating capacity can improve performance if the maximum size is known.
func NewIntStack(capacity int) *IntStack {
	return &IntStack{
		items: make([]int, 0, capacity),
	}
}

// Push adds an item to the top of the stack.
// Time Complexity: O(1) amortized.
func (s *IntStack) Push(item int) {
	s.items = append(s.items, item)
}

// Pop removes and returns the top item from the stack.
// Returns an error if the stack is empty.
// Time Complexity: O(1).
func (s *IntStack) Pop() (int, bool) {
	if s.IsEmpty() {
		return 0, false
	}

	index := len(s.items) - 1
	item := s.items[index]
	s.items = s.items[:index]

	return item, true
}

// Peek returns the top item without removing it.
// Returns an error if the stack is empty.
// Time Complexity: O(1).
func (s *IntStack) Peek() (int, bool) {
	if s.IsEmpty() {
		return 0, false
	}

	return s.items[len(s.items)-1], true
}

// IsEmpty returns true if the stack is empty.
// Time Complexity: O(1).
func (s *IntStack) IsEmpty() bool {
	return len(s.items) == 0
}
