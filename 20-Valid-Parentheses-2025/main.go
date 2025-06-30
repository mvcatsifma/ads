package main

const (
	LeftBracket  = '['
	RightBracket = ']'
	LeftParen    = '('
	RightParen   = ')'
	LeftBrace    = '{'
	RightBrace   = '}'
)

func isValid(s string) bool {
	if s == "" {
		return true
	}

	st := &stack{}
	for _, str := range s {
		switch str {
		case LeftBracket:
			fallthrough
		case LeftParen:
			fallthrough
		case LeftBrace:
			st.push(str)
		case RightBracket:
			str2 := st.pop()
			if str2 != LeftBracket {
				return false
			}
		case RightParen:
			str2 := st.pop()
			if str2 != LeftParen {
				return false
			}
		case RightBrace:
			str2 := st.pop()
			if str2 != LeftBrace {
				return false
			}
		}
	}

	return len(st.vals) == 0
}

type stack struct {
	vals []int32
}

func (s *stack) pop() int32 {
	l := len(s.vals)
	if l == 0 {
		return -1
	}
	v := s.vals[l-1]
	s.vals = s.vals[:l-1]
	return v
}

func (s *stack) push(str int32) {
	s.vals = append(s.vals, str)
}
