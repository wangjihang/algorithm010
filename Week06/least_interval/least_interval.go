package least_interval

import "sort"

// https://leetcode-cn.com/problems/task-scheduler/solution/tong-zi-by-popopop/ 这个题解不错
// (N-1) * (n+1) + k
// 时间复杂度：O(nlogn)
//leetcode submit region begin(Prohibit modification and deletion)
func leastInterval(tasks []byte, n int) int {
	arr := make([]int, 26)
	for _, v := range tasks {
		arr[v-'A']++
	}
	sort.Ints(arr)

	count := 1
	for i := 24; i >= 0; i-- {
		if arr[i] < arr[25] {
			break
		}
		count++
	}

	return max((arr[25]-1)*(n+1)+count, len(tasks))
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
