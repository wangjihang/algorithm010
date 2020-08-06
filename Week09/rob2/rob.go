package rob2

// dp
// nums[1:] 第一个不偷
// nums[:len-1] 最后一个不偷
// 退化为rob1的解法
// 时间复杂度: O(N)
func rob(nums []int) int {
	return max(myRob(nums[1:]), myRob(nums[:len(nums)-1]))
}

func myRob(nums []int) int {
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
