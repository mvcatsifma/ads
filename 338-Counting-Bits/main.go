package p358

import (
	"math/bits"
)

func countBits(n int) []int {
	var ans []int
	for i := 0; i <= n; i++ {
		ans = append(ans, bits.OnesCount(uint(i)))
	}
	return ans
}
