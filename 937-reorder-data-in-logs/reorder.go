package main

import (
	"fmt"
	"regexp"
	"sort"
	"strings"
)

func Reorder(logs []string) []string {
	var result []string
	var letterLogs []Log
	var digitLogs []Log
	for _, log := range logs {
		l := Parse(log)
		if l.LogType == Digit {
			digitLogs = append(digitLogs, l)
		} else {
			letterLogs = append(letterLogs, l)
		}
	}

	sort.Sort(ByWords(letterLogs))

	for _, l := range letterLogs {
		result = append(result, l.String())
	}

	for _, l := range digitLogs {
		result = append(result, l.String())
	}

	return result
}

type LogType int

const (
	Letter LogType = iota
	Digit
)

type Log struct {
	LogType    LogType
	Identifier string
	Words      string
}

func (l *Log) String() string {
	return fmt.Sprintf("%s %s", l.Identifier, l.Words)
}

func Parse(log string) Log {
	tokens := strings.SplitN(log, " ", 2)
	identifier := tokens[0]
	words := tokens[1]

	re := regexp.MustCompile("\\d")
	isDigit := re.MatchString(words)

	logType := Letter
	if isDigit {
		logType = Digit
	}

	return Log{
		Identifier: identifier,
		LogType:    logType,
		Words:      words,
	}
}

type ByWords []Log

func (b ByWords) Len() int {
	return len(b)
}

func (b ByWords) Less(i, j int) bool {
	l1 := b[i]
	l2 := b[j]

	if l1.Words == l2.Words {
		return l1.Identifier < l2.Identifier
	}

	return l1.Words < l2.Words
}

func (b ByWords) Swap(i, j int) {
	b[i], b[j] = b[j], b[i]
}
