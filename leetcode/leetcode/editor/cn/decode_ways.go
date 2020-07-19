//一条包含字母 A-Z 的消息通过以下方式进行了编码：
//
// 'A' -> 1
//'B' -> 2
//...
//'Z' -> 26
//
//
// 给定一个只包含数字的非空字符串，请计算解码方法的总数。
//
// 示例 1:
//
// 输入: "12"
//输出: 2
//解释: 它可以解码为 "AB"（1 2）或者 "L"（12）。
//
//
// 示例 2:
//
// 输入: "226"
//输出: 3
//解释: 它可以解码为 "BZ" (2 26), "VF" (22 6), 或者 "BBF" (2 2 6) 。
//
// Related Topics 字符串 动态规划
// 👍 442 👎 0

package main

// dp[i] = dp[i-1] + dp[i-2] 当前数字的解码方式 = i和i-1分开的解码 + i和i-1相连的解码
//leetcode submit region begin(Prohibit modification and deletion)
func numDecodings(s string) int {
	if s[0] == '0' {
		return 0
	}
	pre, cur, length := 1, 1, len(s)
	for i := 1; i < length; i++ {
		tmp := cur // 替换pre
		if s[i] == '0' {
			if s[i-1] == '1' || s[i-1] == '2' {
				cur = pre
			} else {
				return 0
			}
		} else if s[i-1] == '1' || s[i-1] == '2' && s[i] >= '1' && s[i] <= '6' {
			cur += pre // dp[i] = dp[i-1] + dp[i-2]
		}
		pre = tmp
	}
	return cur
}

//leetcode submit region end(Prohibit modification and deletion)
