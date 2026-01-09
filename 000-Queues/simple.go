package p000Queues

// SimpleFIFOIntQueue implements a simple First-In-First-Out queue for integers using a slice.
// This is a basic implementation suitable for small queues or educational purposes.
// Note: Dequeue operation is O(n) due to slice shifting, making it inefficient for large queues.
type SimpleFIFOIntQueue struct {
	items []int // Underlying slice storing queue elements (front at index 0, back at end)
}

// NewIntQueue creates and returns a new empty integer queue.
// Initializes with zero capacity that grows dynamically as items are added.
func NewIntQueue() *SimpleFIFOIntQueue {
	return &SimpleFIFOIntQueue{items: make([]int, 0)}
}

// Enqueue adds an item to the back of the queue.
// Time Complexity: O(1) amortized (may trigger slice reallocation).
// Space Complexity: O(1) per operation.
func (q *SimpleFIFOIntQueue) Enqueue(item int) {
	q.items = append(q.items, item)
}

// Dequeue removes and returns the item from the front of the queue.
// Returns the item and true if successful, or 0 and false if queue is empty.
// Time Complexity: O(n) due to slice shifting - all remaining elements must be moved.
// Space Complexity: O(1).
func (q *SimpleFIFOIntQueue) Dequeue() (int, bool) {
	if len(q.items) == 0 {
		return 0, false // Queue is empty
	}
	item := q.items[0]
	q.items = q.items[1:] // Shift all elements left (expensive O(n) operation)
	return item, true
}

// IsEmpty returns true if the queue contains no elements, false otherwise.
// Time Complexity: O(1).
func (q *SimpleFIFOIntQueue) IsEmpty() bool {
	return len(q.items) == 0
}

// Size returns the number of elements currently in the queue.
// Time Complexity: O(1).
func (q *SimpleFIFOIntQueue) Size() int {
	return len(q.items)
}
