package permute

// 解决一个回溯问题，实际上就是一个决策树的遍历过程。你只需要思考 3 个问题：
// 1、路径：也就是已经做出的选择。
// 2、选择列表：也就是你当前可以做的选择。
// 3、结束条件：也就是到达决策树底层，无法再做选择的条件。

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

func Permute(nums []int) [][]int {
	var (
		length    = len(nums)
		used      = make([]bool, length)
		result    = make([][]int, 0)
		backtrack func(path []int, used []bool)
	)
	backtrack = func(path []int, used []bool) {
		if len(path) == length {
			result = append(result, Copy(path, length))
			return
		}
		for i, isused := range used {
			if !isused {
				used[i] = true
				path = append(path, nums[i])

				backtrack(path, used)

				used[i] = false
				path = path[:len(path)-1]
			}
		}
	}
	backtrack(make([]int, 0, length), used)
	return result
}

func Copy(path []int, length int) []int {
	dup := make([]int, length)
	copy(dup, path)
	return dup
}
