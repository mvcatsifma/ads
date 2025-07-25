package p49

import (
	"sort"
	"strings"
)

func groupAnagrams(strs []string) [][]string {
	m := make(map[string][]string)

	for _, str := range strs {
		strSorted := sortString(str)
		if _, ok := m[strSorted]; ok {
			m[strSorted] = append(m[strSorted], str)
		} else {
			m[strSorted] = []string{str}
		}
	}

	var result [][]string
	for _, v := range m {
		result = append(result, v)
	}

	return result
}

func sortString(s string) string {
	runes := make([]rune, 0, len(s))
	for _, r := range s {
		runes = append(runes, r)
	}

	sort.Slice(runes, func(i, j int) bool {
		return runes[i] < runes[j]
	})

	var builder strings.Builder
	builder.Grow(len(runes))
	for _, r := range runes {
		builder.WriteRune(r)
	}

	return builder.String()
}
