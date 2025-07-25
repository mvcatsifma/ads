package p14

import (
	"fmt"
	"strings"
)

func main() {
	prefix := longestCommonPrefix([]string{"flower", "flow", "flight"})
	fmt.Println(prefix)
}

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	if len(strs) == 1 {
		return strs[0]
	}

	// Get the shortest string from the list
	s1 := strs[0]
	for _, str := range strs {
		if len(str) < len(s1) {
			s1 = str
		}
	}

	l1 := len(s1)
	prefix := ""
PrefixLoop:
	for i := 1; i <= l1; i++ {
		p1 := s1[:i]
		for _, s2 := range strs {
			if !strings.HasPrefix(s2, p1) {
				break PrefixLoop
			}
		}
		prefix = p1
	}

	return prefix
}
