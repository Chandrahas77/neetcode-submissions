func maxProfit(prices []int) int {
	minPrice := prices[0]
	maxProfit := 0
	for i := range prices{
		currPrice := prices[i]
		if currPrice < minPrice{
			minPrice = currPrice
		}
		if currPrice - minPrice > maxProfit{
			maxProfit = currPrice - minPrice
		}
	}
	return maxProfit
}
