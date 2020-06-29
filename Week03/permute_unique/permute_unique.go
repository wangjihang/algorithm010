package permute_unique

func permuteUnique(nums []int) [][]int {
	var (
		result    = make([][]int, 0)
		length    = len(nums)
		used      = make([]bool, length)
		backtrack func(path []int, used []bool)
	)
	backtrack = func(path []int, used []bool) {
		if len(path) == length {
			result = append(result, Copy(path, length))
			return
		}
		usedM := make(map[int]bool) // 同层遍历时 判断相同值是否已经做过排列 剪枝
		for i, v := range nums {
			if !used[i] && !usedM[v] {
				used[i] = true
				usedM[v] = true
				path = append(path, v)

				backtrack(path, used)

				used[i] = false
				path = path[:len(path)-1]
			}
		}
	}
	backtrack(nil, used)
	return result
}

func Copy(track []int, length int) []int {
	dup := make([]int, length)
	copy(dup, track)
	return dup
}
