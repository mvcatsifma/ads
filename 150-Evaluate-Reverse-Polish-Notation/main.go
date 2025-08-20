package p50

import (
	"strconv"
)

func evalRPN(tokens []string) int {
	stack := NewStack(len(tokens))
	for _, token := range tokens {
		switch token {
		case "+":
			v1 := stack.Pop()
			v2 := stack.Pop()
			stack.Push(v2 + v1)
		case "-":
			v1 := stack.Pop()
			v2 := stack.Pop()
			stack.Push(v2 - v1)
		case "*":
			v1 := stack.Pop()
			v2 := stack.Pop()
			stack.Push(v2 * v1)
		case "/":
			v1 := stack.Pop()
			v2 := stack.Pop()
			stack.Push(v2 / v1)
		default:
			v, _ := strconv.ParseInt(token, 10, 32)
			stack.Push(int(v))
		}
	}
	return stack.Pop()
}

// Stack represents a fixed-size stack for strings
type Stack struct {
	data     []int
	top      int // index of the top element (-1 when empty)
	capacity int
}

// NewStack creates a new stack with the specified capacity
func NewStack(capacity int) *Stack {
	if capacity <= 0 {
		panic("capacity must be positive")
	}
	return &Stack{
		data:     make([]int, capacity),
		top:      -1,
		capacity: capacity,
	}
}

// Push adds an element to the top of the stack
func (s *Stack) Push(item int) {
	s.top++
	s.data[s.top] = item
}

// Pop removes and returns the top element from the stack
func (s *Stack) Pop() int {
	item := s.data[s.top]
	s.data[s.top] = 0 // clear the reference
	s.top--
	return item
}
