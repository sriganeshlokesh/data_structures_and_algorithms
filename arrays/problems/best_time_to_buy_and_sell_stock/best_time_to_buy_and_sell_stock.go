package best_time_to_buy_and_sell_stock

func MaxProfit(prices []int) int {
	maximumProfit, minPrice := 0, prices[0]

	for i := 1; i < len(prices); i++ {
		if prices[i] < minPrice {
			minPrice = prices[i]
		}

		maximumProfit = max(maximumProfit, prices[i]-minPrice)
	}

	return maximumProfit
}
