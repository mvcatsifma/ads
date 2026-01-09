package p000Queues

import "container/list"

// ListFIFOIntQueue implements a First-In-First-Out queue for integers using Go's container/list.
// This implementation provides O(1) enqueue and dequeue operations, making it efficient for large queues.
// Uses a doubly-linked list internally for optimal performance at both ends.
type ListFIFOIntQueue struct {
	list *list.List // Underlying doubly-linked list (front = head, back = tail)
}

// NewListFIFOIntQueue creates and returns a new empty integer queue backed by container/list.
// The queue is ready to use immediately with no initial capacity limitations.
func NewListFIFOIntQueue() *ListFIFOIntQueue {
	return &ListFIFOIntQueue{list: list.New()}
}

// Enqueue adds an item to the back of the queue.
// Items are added to the tail of the linked list for FIFO ordering.
// Time Complexity: O(1) - constant time insertion at tail.
// Space Complexity: O(1) per operation.
func (q *ListFIFOIntQueue) Enqueue(item int) {
	q.list.PushBack(item)
}

// Dequeue removes and returns the item from the front of the queue.
// Returns the item and true if successful, or 0 and false if queue is empty.
// Items are removed from the head of the linked list for FIFO ordering.
// Time Complexity: O(1) - constant time removal from head.
// Space Complexity: O(1).
func (q *ListFIFOIntQueue) Dequeue() (int, bool) {
	front := q.list.Front()
	if front == nil {
		return 0, false // Queue is empty
	}
	value := front.Value.(int) // Type assertion from interface{} to int
	q.list.Remove(front)       // O(1) removal since we have the element reference
	return value, true
}

// IsEmpty returns true if the queue contains no elements, false otherwise.
// Time Complexity: O(1) - list maintains internal length counter.
func (q *ListFIFOIntQueue) IsEmpty() bool {
	return q.list.Len() == 0
}

// Size returns the number of elements currently in the queue.
// Time Complexity: O(1) - list maintains internal length counter.
func (q *ListFIFOIntQueue) Size() int {
	return q.list.Len()
}
