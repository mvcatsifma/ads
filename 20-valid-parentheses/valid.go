package main

import (
	"fmt"
)

const (
	ParenthesesOpen int = iota
	ParenthesesClose
	BraceOpen
	BraceClose
	SquareOpen
	SquareClose
)

func isValid(strs string) bool {
	// verify there is an even number of brackets.
	if len(strs)%2 != 0 {
		return false
	}

	st := &stack{}
	for _, s := range strs {
		switch string(s) {
		case "(":
			st.push(ParenthesesOpen)
		case "{":
			st.push(BraceOpen)
		case "[":
			st.push(SquareOpen)
		case ")":
			b := st.pop()
			if b != ParenthesesOpen {
				return false
			}
		case "}":
			bracket := st.pop()
			if bracket != BraceOpen {
				return false
			}
		case "]":
			bracket := st.pop()
			if bracket != SquareOpen {
				return false
			}
		}
	}

	return st.size() == 0
}

type stack struct {
	values []int
}

func (s *stack) push(v int) {
	s.values = append(s.values, v)
}

func (s *stack) pop() int {
	l := len(s.values)
	if l == 0 {
		return -1
	}
	v := s.values[l-1]
	s.values = s.values[:l-1]
	return v
}

func (s *stack) size() int {
	return len(s.values)
}

func main() {
	s := &stack{}
	s.push(ParenthesesOpen)
	s.push(ParenthesesClose)
	fmt.Println(s.pop())
	fmt.Println(s.pop())
	fmt.Println(s.pop())
}
