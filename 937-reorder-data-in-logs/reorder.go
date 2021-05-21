package main

import (
	"regexp"
	"sort"
	"strings"
)

var re = regexp.MustCompile("^\\w* \\d")

func reorderLogFiles(logs []string) []string {
	sort.Stable(ByContents(logs))
	return logs
}

type LogType int

const (
	Letter LogType = iota
	Digit
)

type ByContents []string

func (b ByContents) Len() int {
	return len(b)
}

func (b ByContents) Less(i, j int) bool {
	l1 := b[i]
	l2 := b[j]

	logType1 := logType(l1)

	if logType1 == Digit {
		return false
	}

	logType2 := logType(l2)
	if logType1 < logType2 {
		return true
	}

	id1, words1 := tokens(l1)
	id2, words2 := tokens(l2)

	if words1 == words2 {
		return id1 < id2
	}

	return words1 < words2
}

func (b ByContents) Swap(i, j int) {
	b[i], b[j] = b[j], b[i]
}

func tokens(log string) (string, string) {
	tokens := strings.SplitN(log, " ", 2)

	return tokens[0], tokens[1]
}

func logType(log string) LogType {
	isDigit := re.MatchString(log)

	logType := Letter
	if isDigit {
		logType = Digit
	}

	return logType
}
