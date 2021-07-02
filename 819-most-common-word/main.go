package main

import (
	"strings"
)

const space = " "

func mostCommonWord(p string, b []string) string {
	// lowercase paragraph
	p = strings.ToLower(p)

	// remove symbols
	r := strings.NewReplacer("!", " ", "?", " ", "'", " ", ",", " ", ";", " ", ".", " ")
	p = r.Replace(p)

	// split paragraph on space
	words := strings.Split(p, space)

	// count word frequencies
	freq := make(map[string]int)
WordsLoop:
	for _, word := range words {
		if word == "" {
			continue WordsLoop
		}

		// ignore banned words
		for _, s := range b {
			if word == s {
				continue WordsLoop
			}
		}

		if f, ok := freq[word]; ok { // word already counted
			f += 1
			freq[word] = f
		} else { // word not yet counted
			freq[word] = 1
		}
	}

	// determine word with highest frequency
	var hw string
	var hf int
	for k, v := range freq {
		if v > hf {
			hf = v
			hw = k
		}
	}

	return hw
}
