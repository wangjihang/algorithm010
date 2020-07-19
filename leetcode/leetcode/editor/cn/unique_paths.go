//一个机器人位于一个 m x n 网格的左上角 （起始点在下图中标记为“Start” ）。
//
// 机器人每次只能向下或者向右移动一步。机器人试图达到网格的右下角（在下图中标记为“Finish”）。
//
// 问总共有多少条不同的路径？
//
//
//
// 例如，上图是一个7 x 3 的网格。有多少可能的路径？
//
//
//
// 示例 1:
//
// 输入: m = 3, n = 2
//输出: 3
//解释:
//从左上角开始，总共有 3 条路径可以到达右下角。
//1. 向右 -> 向右 -> 向下
//2. 向右 -> 向下 -> 向右
//3. 向下 -> 向右 -> 向右
//
//
// 示例 2:
//
// 输入: m = 7, n = 3
//输出: 28
//
//
//
// 提示：
//
//
// 1 <= m, n <= 100
// 题目数据保证答案小于等于 2 * 10 ^ 9
//
// Related Topics 数组 动态规划
// 👍 605 👎 0

package main

// 重复性(分治)
// 定义状态数组
// DP方程
// 当前的所有走法 = 左边所有走法 + 上边面所有走法
// dp[i,j] = dp[i, j-1] + dp[i-1, j]
// O(M*N)
//leetcode submit region begin(Prohibit modification and deletion)
//func uniquePaths(m int, n int) int {
//	dp := make([][]int, m)
//	for i := 0; i < m; i++ {
//		dp[i] = make([]int, n)
//		for j := 0; j < n; j++ {
//			if i > 0 && j > 0 {
//				dp[i][j] = dp[i][j-1] + dp[i-1][j]
//			} else {
//				dp[i][j] = 1
//			}
//		}
//	}
//	return dp[m-1][n-1]
//}

// 减少一维空间
// cur[j] = cur[j] + cur[j-1]  就是一维数据从上往下累加
func uniquePaths(m int, n int) int {
	cur := make([]int, n)
	for i := 0; i < n; i++ {
		cur[i] = 1
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ { // 最左侧列 只能是1 所以从第二列开始
			cur[j] += cur[j-1]
		}
	}
	return cur[n-1]
}

//leetcode submit region end(Prohibit modification and deletion)
