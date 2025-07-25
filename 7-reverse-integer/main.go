package p7

import (
	"fmt"
	"log"
	"math"
	"strconv"
	"strings"
)

// TODO: use approach from solution ("Pop and Push Digits & Check before Overflow")
func reverse(x int) int {
	if x == 0 {
		return 0
	}
	var xStr = fmt.Sprintf("%.f", math.Abs(float64(x)))
	var xArr []string
	for i := len(xStr) - 1; i >= 0; i-- {
		xArr = append(xArr, string(xStr[i]))
	}
	result := strings.Join(xArr, "")
	out, err := strconv.ParseInt(result, 10, 64)
	if err != nil {
		log.Fatal(err)
	}
	if out > math.MaxInt32 || out < math.MinInt32 {
		out = 0
	}
	if x < 0 {
		out = out * -1
	}
	return int(out)
}
