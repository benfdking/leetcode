package _go

func maxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}

	profit := 0
	buyingValue := prices[0]
	for _, p := range prices {
		if p < buyingValue {
			buyingValue = p
		} else {
			newProfit := p - buyingValue
			if newProfit > profit {
				profit = newProfit
			}
		}
	}
	return profit
}

// maxProfitSimple
func maxProfitSimple(prices []int) int {
	max := 0
	for i := 0; i < len(prices); i++ {
		for j := i + 1; j < len(prices); j++ {
			value := prices[j] - prices[i]
			if max < value {
				max = value
			}
		}
	}
	return max
}
