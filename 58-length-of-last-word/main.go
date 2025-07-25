package p58

import "strings"

func main() {

}

func lengthOfLastWord(s string) int {
	strs := strings.Split(s, " ")
	for i := len(strs) - 1; i >= 0; i-- {
		if strs[i] != "" {
			return len(strs[i])
		}
	}
	return 0
}
