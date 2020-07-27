//给定一个整数 n，求以 1 ... n 为节点组成的二叉搜索树有多少种？
//
// 示例:
//
// 输入: 3
//输出: 5
//解释:
//给定 n = 3, 一共有 5 种不同结构的二叉搜索树:
//
//   1         3     3      2      1
//    \       /     /      / \      \
//     3     2     1      1   3      2
//    /     /       \                 \
//   2     1         2                 3
// Related Topics 树 动态规划
// 👍 711 👎 0

package main

// 二叉搜索树的概念是：左子树都小于根节点，右子树都大于根节点。
// 假设n个节点存在二叉树的个数是G(n)，令f(i)为以i为根的二叉搜索树的个数，则
// G(n) = f(1)+f(2)+f(3)+f(4)+...+f(n)
// 当i为根节点时，其左子树节点个数为i-1，右子树节点个数为n-i，则
// f(i) = G(i-1)*G(n-i)
// 综合可得[卡特兰数]
// G(n) = G(0)*G(n-1)+G(1)*G(n-2)+...+G(n-1)*G(0)
// https://leetcode-cn.com/problems/unique-binary-search-trees/solution/hua-jie-suan-fa-96-bu-tong-de-er-cha-sou-suo-shu-b/
//leetcode submit region begin(Prohibit modification and deletion)
func numTrees(n int) int {
	var res = 1
	for i := 0; i < n; i++ {
		res = res * 2 * (1 + 2*i) / (i + 2)
	}
	return res
}

//leetcode submit region end(Prohibit modification and deletion)
