package p2220

import (
	"strconv"
	"strings"
)

func minBitFlips(start int, goal int) int {
	s := strconv.FormatInt(int64(start^goal), 2)

	return strings.Count(s, "1")
}
