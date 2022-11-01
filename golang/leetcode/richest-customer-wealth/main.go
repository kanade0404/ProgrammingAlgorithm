package richest_customer_wealth

func maximumWealth(accounts [][]int) int {
	var result int
	for _, account := range accounts {
		var total int
		for _, a := range account {
			total = total + a
		}
		if total > result {
			result = total
		}
	}
	return result
}
