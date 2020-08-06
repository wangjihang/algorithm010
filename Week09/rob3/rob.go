package rob3

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 升维
// 0表示不抢，1表示抢
// i不抢 dp[0][i] = max(dp[0][left], dp[1][left]) + max(dp[0][right], dp[1][right])
// i抢  dp[1][i] = dp[0][left] + dp[0][right] + value
// dp[i] = max(dp[0][i], dp[1][i])
// 后续遍历 时间复杂度：O(N)
func rob(root *TreeNode) int {
	ret := dfs(root)
	return max(ret[0], ret[1])
}

func dfs(node *TreeNode) (ret [2]int) {
	if node == nil {
		return
	}

	l, r := dfs(node.Left), dfs(node.Right)
	ret[0] = max(l[0], l[1]) + max(r[0], r[1])
	ret[1] = l[0] + r[0] + node.Val
	return
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
