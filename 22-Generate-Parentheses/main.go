package p22

const (
	ParenthesesOpen int = iota
)

func generateParenthesis(n int) []string {
	var pairs string
	for i := 0; i < n; i++ {
		pairs += "()"
	}

	perms := generatePermutations(pairs)

	var valid []string
	for _, perm := range perms {
		if isValid(perm) {
			valid = append(valid, perm)
		}
	}

	return valid
}

func generatePermutations(chars string) []string {
	if len(chars) <= 1 {
		return []string{chars}
	}

	seen := make(map[string]bool)
	bytes := []byte(chars)
	n := len(bytes)

	var heapPermute func(k int)
	heapPermute = func(k int) {
		if k == 1 {
			perm := string(bytes)
			if !seen[perm] {
				seen[perm] = true
			}
			return
		}

		for i := 0; i < k; i++ {
			heapPermute(k - 1)

			if k%2 == 1 {
				bytes[0], bytes[k-1] = bytes[k-1], bytes[0]
			} else {
				bytes[i], bytes[k-1] = bytes[k-1], bytes[i]
			}
		}
	}

	heapPermute(n)

	// Convert map keys to slice
	result := make([]string, 0, len(seen))
	for perm := range seen {
		result = append(result, perm)
	}

	return result
}

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
