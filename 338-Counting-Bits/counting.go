package main

import (
	"strconv"
	"strings"
)

func countBits(n int) []int {
	var ans []int
	for i := 0; i <= n; i++ {
		fi := strconv.FormatInt(int64(i), 2)
		c := strings.Count(fi, "1")
		ans = append(ans, c)
	}
	return ans
}
