package main

import (
	"strings"
)

func checkIfPangram(sentence string) bool {
	if len(sentence) < 26 {
		return false
	}
	var alphabet = "abcdefghijklmnopqrstuvwxyz"
	for i := 0; i < len(sentence); i++ {
		alphabet = strings.Replace(alphabet, string(sentence[i]), "", 1)
		if len(alphabet) == 0 {
			return true
		}
	}

	return false
}
