package main

func maxProfit(prices []int) int {
	var profit int
	var lowest int
	for i, price := range prices {
		if i == 0 {
			lowest = price
			continue
		}
		if price < lowest {
			lowest = price
		}
		if price-lowest > profit {
			profit = price - lowest
		}
	}
	return profit
}
