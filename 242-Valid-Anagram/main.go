package p242

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	m := make(map[string]int)
	for _, v := range s {
		m[string(v)]++
	}

	for _, v := range t {
		m[string(v)]--
		if m[string(v)] < 0 {
			return false
		}
	}

	return true
}
