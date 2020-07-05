package max_profit

// 一次遍历：
func maxProfit(prices []int) int {
	var (
		max    int
		length = len(prices)
	)
	for i := 1; i < length-1; i++ {
		if value := prices[i] - prices[i-1]; value > 0 {
			max += value
		}
	}
	return max
}

// 波峰、谷底
// 时间复杂度：O(n)
// 时间复杂度：O(1)
func maxProfit2(prices []int) int {
	var (
		length            = len(prices)
		max, valley, peek int
	)

	for i := 0; i < length-1; {
		for i < length-1 && prices[i] >= prices[i+1] { // 寻找谷底
			i++
		}
		valley = prices[i]
		for i < length-1 && prices[i] <= prices[i+1] { // 寻找波峰
			i++
		}
		peek = prices[i]
		max += peek - valley
	}
	return max
}

// 一次遍历法
