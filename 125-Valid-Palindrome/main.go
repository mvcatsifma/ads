package main

import (
	"bytes"
	"strings"
)

const allowed = "abcdefghijklmnopqrstuvwxyz0123456789"

func isPalindrome(s string) bool {
	var buf bytes.Buffer
	for _, c := range strings.ToLower(s) {
		if strings.Index(allowed, string(c)) != -1 {
			buf.WriteString(string(c))
		}
	}
	s = buf.String()
	l := len(s)
	if l == 0 { // empty string
		return true
	}

	j := l - 1
	for i := 0; i < l-1; i++ {
		if s[i] != s[j] {
			return false
		}
		j--
	}

	return true
}
