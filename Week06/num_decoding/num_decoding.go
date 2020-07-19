package num_decoding

// if s[i] == '0' && (s[i-1]=='1' || s[i-1] == '2')  dp[i] = dp[i-2] s[i-1~i] 必须要连起来
// dp[i] = dp[i-1] + dp[i-2]  不连起来 + 连起来 的总数
// 时间复杂度：O(N) N为s的长度
// 空间复杂度：O(1)
func numDecodings(s string) int {
	if s[0] == '0' {
		return 0
	}
	pre, cur, length := 1, 1, len(s)
	for i := 1; i < length; i++ {
		tmp := cur
		if s[i] == '0' {
			if s[i-1] == '1' || s[i-1] == '2' {
				cur = pre
				continue
			}
			return 0
		} else if s[i-1] == '1' || (s[i-1] == '2' && s[i] >= '1' && s[i] <= 6) {
			cur += pre
		}
		pre = tmp
	}
	return cur
}
