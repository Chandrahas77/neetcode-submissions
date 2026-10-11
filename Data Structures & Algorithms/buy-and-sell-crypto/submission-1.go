func maxProfit(prices []int) int {
	maxProfit := 0
	left := 0
	right := 1
	for right < len(prices){
		if prices[left] < prices[right]{
			profit := prices[right] - prices[left]
			maxProfit = max(maxProfit,profit)
		}else{
			left = right
		}
		right++
	}
	return maxProfit
}
