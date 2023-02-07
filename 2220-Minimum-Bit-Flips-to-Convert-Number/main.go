package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(minBitFlips(10, 7))
	fmt.Println(minBitFlips(3, 4))
}

func minBitFlips(start int, goal int) int {
	s := strconv.FormatInt(int64(start^goal), 2)

	return strings.Count(s, "1")
}
