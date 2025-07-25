package p572

import "strings"

const symbols = "!?',;. "

// TODO: Approach 2: Character Processing in One-Pass
func mostCommonWord(p string, b []string) string {
	freq := make(map[string]int)
	var hf int    // keeps track of the highest frequency so far
	var w string  // keeps track of current word
	var hw string // keeps track of word with highest frequency

LettersLoop:
	for _, l := range p {
		letter := strings.ToLower(string(l))
		if strings.Index(symbols, letter) != -1 {
			if w == "" {
				continue LettersLoop
			}

			// ignore banned words
			for _, s := range b {
				if w == s {
					w = ""
					continue LettersLoop
				}
			}

			// increase the word's frequency count
			freq[w]++
			if freq[w] > hf {
				hf = freq[w]
				hw = w
			}
			w = ""
		} else {
			w += letter
		}
	}

	// last word
	if hw == "" {
		hw = w
	}

	return hw
}
