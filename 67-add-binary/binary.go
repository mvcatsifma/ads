package main

import (
	"fmt"
	"math"
	"strings"
)

func main() {
	fmt.Println(addBinary("1", "11"))
}

func addBinary(a string, b string) string {
	maxLen := int(math.Max(float64(len(a)), float64(len(b))))
	a = fmt.Sprintf("%0*s", maxLen, a)
	b = fmt.Sprintf("%0*s", maxLen, b)

	var result []string
	carry := "0"
	for i := maxLen - 1; i >= 0; i-- {
		ac := a[i]
		bc := b[i]
		r := string(ac) + string(bc)
		switch r {
		case "00":
			result = append(result, carry)
			carry = "0"
			break
		case "01":
			fallthrough
		case "10":
			result = append(result, carry)
			carry = "0"
			break
		case "11":
			result = append(result, carry)
			carry = "1"
			break
		}
	}

	result = append(result, carry)

	return strings.Join(result, "")
}
