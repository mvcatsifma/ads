package p121

import "math"

func maxProfit(prices []int) int {
	var profit int
	lowest := math.MaxInt
	for _, price := range prices {
		if price < lowest {
			lowest = price
		} else {
			profit = max(profit, price-lowest)
		}
	}
	return profit
}
