//给定一个字符串 S，返回 “反转后的” 字符串，其中不是字母的字符都保留在原地，而所有字母的位置发生反转。
//
//
//
//
//
//
// 示例 1：
//
// 输入："ab-cd"
//输出："dc-ba"
//
//
// 示例 2：
//
// 输入："a-bC-dEf-ghIj"
//输出："j-Ih-gfE-dCba"
//
//
// 示例 3：
//
// 输入："Test1ng-Leet=code-Q!"
//输出："Qedo1ct-eeLg=ntse-T!"
//
//
//
//
// 提示：
//
//
// S.length <= 100
// 33 <= S[i].ASCIIcode <= 122
// S 中不包含 \ or "
//
// Related Topics 字符串
// 👍 52 👎 0

package main

// 时间复杂度：O(N)
// 空间复杂度：O(N)
//leetcode submit region begin(Prohibit modification and deletion)
func reverseOnlyLetters(S string) string {
	buf := []byte(S)
	for i, j := 0, len(buf)-1; i < j; {
		for i < j && !isLetter(buf[i]) {
			i++
		}
		for i < j && !isLetter(buf[j]) {
			j--
		}
		buf[i], buf[j] = buf[j], buf[i]
		i++
		j--
	}
	return string(buf)
}

func isLetter(c byte) bool {
	if c >= 'a' && c <= 'z' || c >= 'A' && c <= 'Z' {
		return true
	}
	return false
}

//leetcode submit region end(Prohibit modification and deletion)
