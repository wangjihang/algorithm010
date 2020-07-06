package find_content

import "sort"

func findContentChildren(g []int, s []int) int {
	var (
		i, j    int
		sLength = len(s)
		gLength = len(g)
	)

	sort.Ints(g)
	sort.Ints(s)

	for i < sLength && j < gLength {
		if s[i] >= g[j] { // 满足
			j++
		}
		i++
	}
	return j
}
