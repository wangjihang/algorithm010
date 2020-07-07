package combine

// 回溯：
// result = []
// def backtrack(路径, 选择列表):
// 	if 满足结束条件:
// 		result.add(路径)
// 		return
//
// 	for 选择 in 选择列表:
// 		做选择
// 		backtrack(路径, 选择列表)
// 		撤销选择

func combine(n int, k int) [][]int {
	if n <= 0 || k <= 0 || n < k {
		return nil
	}

	var (
		result    = make([][]int, 0)
		backtrack func(path []int, begin int)
	)
	backtrack = func(path []int, begin int) {
		if len(path) == k {
			result = append(result, Copy(path, k))
			return
		}

		for i := begin; i <= n-(k-len(path))+1; i++ {
			path = append(path, i)

			backtrack(path, i+1)

			path = path[:len(path)-1]
		}
	}
	backtrack(nil, 1)
	return result
}

func Copy(path []int, length int) []int {
	dup := make([]int, length)
	copy(dup, path)
	return dup
}
