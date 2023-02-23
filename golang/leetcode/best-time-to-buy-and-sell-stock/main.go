package best_time_to_buy_and_sell_stock

func maxProfit(prices []int) int {
	var (
		profit int
		prev   int
	)
	for i := range prices {
		if i == 0 {
			prev = prices[i]
			continue
		}
		if prices[i]-prev > profit {
			profit = prices[i] - prev
		} else if prices[i] < prev {
			prev = prices[i]
		}
	}
	return profit
}
