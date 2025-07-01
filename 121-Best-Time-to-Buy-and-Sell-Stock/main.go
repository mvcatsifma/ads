package main

func maxProfit(prices []int) int {
	var profit int
	for i := 0; i <= len(prices)-1; i++ {
		for j := i + 1; j <= len(prices)-1; j++ {
			if prices[j] > prices[i] {
				profit = max(profit, prices[j]-prices[i])
			}
		}
	}
	return profit
}
