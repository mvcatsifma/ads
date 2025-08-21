package p739

import (
	"errors"
)

//
// Given a slice of daily temperatures, this function returns a slice where each
// element represents the number of days you have to wait after that day to get
// a warmer temperature. If there is no future day with a warmer temperature,
// the corresponding element is 0.
//
// The function uses a monotonic decreasing stack to efficiently track indices
// of temperatures that haven't found their next warmer day yet.
//
// Time Complexity: O(n) - each element is pushed and popped at most once
// Space Complexity: O(n) - for the stack and result slice
//
// Parameters:
//   - temperatures: slice of integers representing daily temperatures
//
// Returns:
//   - slice of integers where answer[i] is the number of days to wait after day i
//     for a warmer temperature, or 0 if no such day exists
//
// Example:
//   Input:  [73, 74, 75, 71, 69, 72, 76, 73]
//   Output: [1,  1,  4,  2,  1,  1,  0,  0]
//
//   Day 0 (73°F): Next warmer day is day 1 (74°F) → wait 1 day
//   Day 1 (74°F): Next warmer day is day 2 (75°F) → wait 1 day
//   Day 2 (75°F): Next warmer day is day 6 (76°F) → wait 4 days
//   Day 6 (76°F): No warmer day exists → 0
func dailyTemperatures(temperatures []int) []int {
	answer := make([]int, len(temperatures))
	stack := New(len(temperatures))

	for i := range temperatures {
		// Only pop indices where temperature < current temperature
		for !stack.IsEmpty() {
			j, _ := stack.Peek()
			if temperatures[j] < temperatures[i] {
				idx, _ := stack.Pop()
				answer[idx] = i - idx
			} else {
				break // Stop when stack top >= current temp
			}
		}
		stack.Push(i)
	}

	return answer
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
