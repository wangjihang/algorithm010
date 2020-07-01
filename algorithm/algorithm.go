package algorithm

func permuteUnique(nums []int) [][]int {
	var (
		length    = len(nums)
		result    = make([][]int, 0)
		backtrack func(path []int, used []bool)
	)
	backtrack = func(path []int, used []bool) {
		if len(path) == length {
			result = append(result, Copy(path, length))
			return
		}
		usedM := make(map[int]bool) // map[value]bool
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
	backtrack(make([]int, 0, length), make([]bool, length))
	return result
}

func Copy(track []int, length int) []int {
	dup := make([]int, length)
	copy(dup, track)
	return dup
}
