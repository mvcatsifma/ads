package p20

const (
	ParenthesesOpen int = iota
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
		case ")":
			b := st.pop()
			if b != ParenthesesOpen {
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
