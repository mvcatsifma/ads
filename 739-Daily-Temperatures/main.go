package p739

import "errors"

func dailyTemperatures(temperatures []int) []int {
	return nil
}

// Stack represents a LIFO (Last In, First Out) data structure for integers
type Stack struct {
	items []int
}

// New creates a new stack with pre-allocated capacity
func New(capacity int) *Stack {
	return &Stack{
		items: make([]int, 0, capacity),
	}
}

// Push adds an item to the top of the stack
func (s *Stack) Push(item int) {
	s.items = append(s.items, item)
}

// Pop removes and returns the top item from the stack
func (s *Stack) Pop() (int, error) {
	if s.IsEmpty() {
		return 0, errors.New("stack is empty")
	}

	index := len(s.items) - 1
	item := s.items[index]
	s.items = s.items[:index]

	return item, nil
}

// Peek returns the top item without removing it
func (s *Stack) Peek() (int, error) {
	if s.IsEmpty() {
		return 0, errors.New("stack is empty")
	}

	return s.items[len(s.items)-1], nil
}

// IsEmpty returns true if the stack is empty
func (s *Stack) IsEmpty() bool {
	return len(s.items) == 0
}
