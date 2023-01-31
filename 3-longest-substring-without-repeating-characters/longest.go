package main

func main() {
	lengthOfLongestSubstring("abcabcbb")
}

func lengthOfLongestSubstring(s string) int {
	m := make(map[string]bool)
	var w []string
	var l int
	for _, v := range s {
		w = append(w, string(v))
		for m[string(v)] {
			first := w[0]
			delete(m, first)
			w = w[1:]
		}
		m[string(v)] = true
		if len(w) > l {
			l = len(w)
		}
	}
	return l
}
