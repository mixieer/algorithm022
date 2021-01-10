package Week_04


func maxProfit(prices []int) int {
	total := 0

	for i, v := range prices {
		if i+2 > len(prices) {
			break
		}
		if v <= prices[i+1] {
			total = total + prices[i+1] - prices[i]
		}
	}
	return total
}
