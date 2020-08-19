package stock1

import "math"

//121
// dp[i][0] 表示第i天卖出时最大利润
// dp[i][1] 表示第i天买入时最大利润
func maxProfit(prices []int) int {
	dp_i_0, dp_i_1 := 0, -math.MaxInt64
	for _, v := range prices {
		dp_i_0 = max(dp_i_0, dp_i_1+v)
		dp_i_1 = max(dp_i_1, -v)
	}
	return dp_i_0
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
