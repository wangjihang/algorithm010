//给定一个三角形，找出自顶向下的最小路径和。每一步只能移动到下一行中相邻的结点上。
//
// 相邻的结点 在这里指的是 下标 与 上一层结点下标 相同或者等于 上一层结点下标 + 1 的两个结点。
//
//
//
// 例如，给定三角形：
//
// [
//     [2],
//    [3,4],  [i][j]  [i+1][j]  [i+1][j+1]
//   [6,5,7],
//  [4,1,8,3]
//]
//
//
// 自顶向下的最小路径和为 11（即，2 + 3 + 5 + 1 = 11）。
//
//
//
// 说明：
//
// 如果你可以只使用 O(n) 的额外空间（n 为三角形的总行数）来解决这个问题，那么你的算法会很加分。
// Related Topics 数组 动态规划
// 👍 514 👎 0

package main

import (
	"math"
)

//leetcode submit region begin(Prohibit modification and deletion)
func minimumTotal(triangle [][]int) int {
	var (
		depth = len(triangle)
		res   = math.MaxInt64
		dfs   func(row, col, sum int)
	)
	dfs = func(row, col, sum int) { // [0][0]
		if row >= depth || col >= len(triangle[row]) {
			res = min(res, sum)
			return
		}
		sum += triangle[row][col]
		dfs(row+1, col, sum)
		dfs(row+1, col+1, sum)
	}
	dfs(0, 0, 0)
	return res
}

//func min(a, b int) int {
//	if a < b {
//		return a
//	}
//	return b
//}

//leetcode submit region end(Prohibit modification and deletion)
