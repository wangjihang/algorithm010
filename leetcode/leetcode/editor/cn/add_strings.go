//给定两个字符串形式的非负整数 num1 和num2 ，计算它们的和。
//
// 注意：
//
//
// num1 和num2 的长度都小于 5100.
// num1 和num2 都只包含数字 0-9.
// num1 和num2 都不包含任何前导零。
// 你不能使用任何內建 BigInteger 库， 也不能直接将输入的字符串转换为整数形式。
//
// Related Topics 字符串
// 👍 215 👎 0

package main

//leetcode submit region begin(Prohibit modification and deletion)
func addStrings(num1 string, num2 string) string {
	nb1, nb2 := []byte(num1), []byte(num2)
	if len(nb1) < len(nb2) {
		nb1, nb2 = nb2, nb1
	}

	sum := byte(0)
	for i, j := len(nb1)-1, len(nb2)-1; i >= 0; i, sum = i-1, sum/10 {
		if j >= 0 {
			sum += nb2[j] - '0'
			j--
		}
		sum += nb1[i] - '0'
		nb1[i] = (sum % 10) + '0'
	}
	if sum != 0 {
		nb1 = append([]byte{'1'}, nb1...)
	}
	return string(nb1)
}

//leetcode submit region end(Prohibit modification and deletion)
