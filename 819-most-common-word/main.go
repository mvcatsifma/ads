package main

import "strings"

const space = " "

// Approach 1: String Processing in Pipeline
//func mostCommonWord(p string, b []string) string {
//	// lowercase paragraph
//	p = strings.ToLower(p)
//
//	// remove symbols
//	r := strings.NewReplacer("!", " ", "?", " ", "'", " ", ",", " ", ";", " ", ".", " ")
//	p = r.Replace(p)
//
//	// split paragraph on space
//	words := strings.Split(p, space)
//
//	// count word frequencies
//	freq := make(map[string]int)
//WordsLoop:
//	for _, word := range words {
//		if word == "" {
//			continue WordsLoop
//		}
//
//		// ignore banned words
//		for _, s := range b {
//			if word == s {
//				continue WordsLoop
//			}
//		}
//
//		if f, ok := freq[word]; ok { // word already counted
//			f += 1
//			freq[word] = f
//		} else { // word not yet counted
//			freq[word] = 1
//		}
//	}
//
//	// determine word with highest frequency
//	var hw string
//	var hf int
//	for k, v := range freq {
//		if v > hf {
//			hf = v
//			hw = k
//		}
//	}
//
//	return hw
//}

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
