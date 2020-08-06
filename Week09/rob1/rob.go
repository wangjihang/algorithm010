package rob1

// 198 打家劫舍
// dp[i] = max(dp[i-1], dp[i-2]+nums[i])
func rob(nums []int) int {
	cur, pre := 0, 0
	for _, v := range nums {
		cur, pre = max(cur, pre+v), cur
	}
	return cur
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
