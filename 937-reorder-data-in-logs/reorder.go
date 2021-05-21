package main

import (
	"fmt"
	"regexp"
	"sort"
	"strings"
)

func Reorder(logs []string) []string {
	var result []string
	var parsedLogs []Log
	for _, log := range logs {
		parsedLogs = append(parsedLogs, Parse(log))
	}

	sort.Stable(ByWords(parsedLogs))

	for _, l := range parsedLogs {
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

	if l1.LogType < l2.LogType {
		return true
	}

	if l1.LogType == Digit {
		return false
	}

	if l1.Words == l2.Words {
		return l1.Identifier < l2.Identifier
	}

	return l1.Words < l2.Words
}

func (b ByWords) Swap(i, j int) {
	b[i], b[j] = b[j], b[i]
}
