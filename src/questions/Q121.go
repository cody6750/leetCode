func maxProfit(prices []int) int {
	var max int
	low := prices[0]
	for i := 1; i < len(prices); i++ {
		if low > prices[i] {
			low = prices[i]
			continue
		}
		difference := prices[i] - low
		if difference > max {
			max = difference
		}
	}
	return max
}