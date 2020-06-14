//给定一个整数数组 nums 和一个目标值 target，请你在该数组中找出和为目标值的那 两个 整数，并返回他们的数组下标。
//
//你可以假设每种输入只会对应一个答案。但是，数组中同一个元素不能使用两遍。
//
//示例:
//
//给定 nums = [2, 7, 11, 15], target = 9
//
//因为 nums[0] + nums[1] = 2 + 7 = 9
//所以返回 [0, 1]

package two_sum

// 一遍hash表解法：

func twoSum(nums []int, target int) []int {
	maps := make(map[int]int)
	for k, v := range nums {
		if index, ok := maps[target-v]; ok {
			return []int{index, k}
		}
		maps[v] = k
	}
	return nil
}

// 暴力求解
// 双层遍历 找到即返回
// 时间复杂度：O(N^2)
// 空间复杂度：O(1)
func twoSum2(nums []int, target int) []int {
	length := len(nums)
	for i := 0; i < length-1; i++ {
		for j := i + 1; j < length; j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return nil
}
