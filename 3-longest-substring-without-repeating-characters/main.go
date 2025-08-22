package p3

func lengthOfLongestSubstring(s string) int {
	var answer int
	for i := 0; i < len(s); i++ {
		count := 1
		chars := make(map[uint8]bool)
		chars[s[i]] = true
	SubLoop:
		for j := i + 1; j < len(s); j++ {
			if _, ok := chars[s[j]]; ok {
				break SubLoop
			}
			chars[s[j]] = true
			count++
		}
		if count > answer {
			answer = count
		}
	}

	return answer
}
