package directed

// IntQueue represents a FIFO (First In, First Out) queue data structure for integers.
// Provides efficient O(1) enqueue and dequeue operations.
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
// Assumes queue is not empty (caller should check IsEmpty() first).
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
	return q.head == len(q.items)
}
