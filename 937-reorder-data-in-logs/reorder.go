package main

import (
	"fmt"
	"regexp"
	"strings"
)

func Reorder(logs []string) []string {
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

	fmt.Printf("letter logs: %+v\n", letterLogs)
	fmt.Printf("digit logs: %+v\n", digitLogs)

	return []string{}
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
