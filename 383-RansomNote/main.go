package main

import (
	"strings"
)

func canConstruct(ransomNote string, magazine string) bool {
	if ransomNote == magazine {
		return true
	}
	for _, v := range ransomNote {
		i := strings.Index(magazine, string(v))
		if i == -1 {
			return false
		}
		magazine = magazine[:i] + magazine[i+1:]
	}
	return true
}
