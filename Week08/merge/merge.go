package merge

import "sort"

// 时间复杂度：O(nlogn) 主要是sort开销
// 空间复杂度：O(nlogn) 主要是sort开销
func merge(intervals [][]int) [][]int {
	if len(intervals) <= 0 {
		return nil
	}

	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	index, length := 1, len(intervals)

	for i := 1; i < length; i++ {
		if intervals[i][0] > intervals[index-1][1] {
			intervals[index] = intervals[i]
			index++
		} else if intervals[i][1] > intervals[index-1][1] {
			intervals[index-1][1] = intervals[i][1]
		}
	}
	return intervals[:index]
}
