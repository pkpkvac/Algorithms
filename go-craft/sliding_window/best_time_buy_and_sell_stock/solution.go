package besttimebuyandsellstock

func MaxProfit(prices []int) int {

	minPrice := prices[0]
	maxProfit := 0
	i := 1
	for i < len(prices) {

		profit := prices[i] - minPrice

		if profit > maxProfit {
			maxProfit = profit
		}

		if prices[i] < minPrice {
			minPrice = prices[i]
		}

		i++
	}

	return maxProfit
}
