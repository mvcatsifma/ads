package main

func canConstruct(ransomNote string, magazine string) bool {
	letters := make(map[rune]int)
	for _, r := range magazine {
		letters[r]++
	}

	for _, r := range ransomNote {
		letters[r]--
		if letters[r] < 0 {
			return false
		}
	}

	return true
}
