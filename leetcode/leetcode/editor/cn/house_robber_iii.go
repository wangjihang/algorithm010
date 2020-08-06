//在上次打劫完一条街道之后和一圈房屋后，小偷又发现了一个新的可行窃的地区。这个地区只有一个入口，我们称之为“根”。 除了“根”之外，每栋房子有且只有一个“父“
//房子与之相连。一番侦察之后，聪明的小偷意识到“这个地方的所有房屋的排列类似于一棵二叉树”。 如果两个直接相连的房子在同一天晚上被打劫，房屋将自动报警。
//
// 计算在不触动警报的情况下，小偷一晚能够盗取的最高金额。
//
// 示例 1:
//
// 输入: [3,2,3,null,3,null,1]
//
//     3
//    / \
//   2   3
//    \   \
//     3   1
//
//输出: 7
//解释: 小偷一晚能够盗取的最高金额 = 3 + 3 + 1 = 7.
//
// 示例 2:
//
// 输入: [3,4,5,1,3,null,1]
//
//     3
//    / \
//   4   5
//  / \   \
// 1   3   1
//
//输出: 9
//解释: 小偷一晚能够盗取的最高金额 = 4 + 5 = 9.
//
// Related Topics 树 深度优先搜索
// 👍 508 👎 0

package main

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     l *TreeNode
 *     r *TreeNode
 * }
 */
//type TreeNode struct {
//	Val int
//	l   *TreeNode
//	r   *TreeNode
//}

// dp[node] = max(dp[l]+dp[r], dp[node]+dp[不包含l]+dp[不包含r])
// 升维：使用0,1 分别表示抢或者不抢
// i不抢 dp[0][i] = max(dp[0][l], dp[1][l])+max(dp[0][r], dp[1][r])
// i抢  dp[1][i] = dp[0][l]+dp[0][r]+value
// dp[i] = max(dp[0][i], dp[1][i])
//func rob(root *TreeNode) int {
//	ret := dfs(root)
//	return max(ret[0], ret[1])
//}
//
//func dfs(node *TreeNode) (ret [2]int) {
//	if node == nil {
//		return
//	}
//	l, r := dfs(node.Left), dfs(node.Right)
//	ret[0] = max(l[0], l[1]) + max(r[0], r[1])
//	ret[1] = l[0] + r[0] + node.Val
//	return
//}

//func max(a, b int) int {
//	if a > b {
//		return a
//	}
//	return b
//}

//leetcode submit region end(Prohibit modification and deletion)
