/*
* leetcode[https://leetcode-cn.com/problems/remove-duplicates-from-sorted-array/]
* 删除排序数组中的重复项
 */
package remove_duplicate

// 最优解：

// 双指针法[快慢指针]：
// 时间复杂度 O(n)
// 空间复杂度 O(1)
// 压测下来 这个更好一点

func RemoveDuplicates(nums []int) int {
	var (
		i      = 0
		length = len(nums)
	)
	for j := i + 1; j < length; j++ {
		if nums[i] != nums[j] {
			if j-i > 1 {
				nums[i+1] = nums[j]
			}
			i++
		}
	}
	return i + 1
}

// 不同时，就会进行交换 为了测试性能
func RemoveDuplicates2(nums []int) int {
	var (
		i      = 0
		length = len(nums)
	)
	for j := i + 1; j < length; j++ {
		if nums[i] != nums[j] {
			nums[i+1] = nums[j]
			i++
		}
	}
	return i + 1
}
