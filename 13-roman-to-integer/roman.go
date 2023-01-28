package main

import (
	"strings"
)

type numeral struct {
	letter string
	value  int
}

var numerals []numeral

func init() {
	numerals = append(numerals, numeral{"MMM", 3000})
	numerals = append(numerals, numeral{"MM", 2000})
	numerals = append(numerals, numeral{"CM", 900})
	numerals = append(numerals, numeral{"M", 1000})
	numerals = append(numerals, numeral{"DCCC", 800})
	numerals = append(numerals, numeral{"DCC", 700})
	numerals = append(numerals, numeral{"DC", 600})
	numerals = append(numerals, numeral{"CD", 400})
	numerals = append(numerals, numeral{"D", 500})
	numerals = append(numerals, numeral{"CCC", 300})
	numerals = append(numerals, numeral{"CC", 200})
	numerals = append(numerals, numeral{"XC", 90})
	numerals = append(numerals, numeral{"C", 100})
	numerals = append(numerals, numeral{"LXXX", 80})
	numerals = append(numerals, numeral{"LXX", 70})
	numerals = append(numerals, numeral{"LX", 60})
	numerals = append(numerals, numeral{"XL", 40})
	numerals = append(numerals, numeral{"L", 50})
	numerals = append(numerals, numeral{"XXX", 30})
	numerals = append(numerals, numeral{"XX", 20})
	numerals = append(numerals, numeral{"IX", 9})
	numerals = append(numerals, numeral{"X", 10})
	numerals = append(numerals, numeral{"VIII", 8})
	numerals = append(numerals, numeral{"VII", 7})
	numerals = append(numerals, numeral{"VI", 6})
	numerals = append(numerals, numeral{"IV", 4})
	numerals = append(numerals, numeral{"V", 5})
	numerals = append(numerals, numeral{"III", 3})
	numerals = append(numerals, numeral{"II", 2})
	numerals = append(numerals, numeral{"I", 1})
}

func romanToInt(s string) int {
	result := 0
	for _, v := range numerals {
		pos := strings.Index(s, v.letter)
		if pos != -1 {
			result += v.value
			s = string(s[:pos] + s[pos+len(v.letter):])
		}
	}
	return result
}
