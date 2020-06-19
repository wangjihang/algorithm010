package ugly_number

// 动态规划：
// 时间复杂度：O(n)
// 空间复杂度：O(n)
// x[a]*2 > x[i] > x[a-1]*2
// x[b]*3 > x[i] > x[b-1]*3
// x[c]*5 > x[i] > x[c-1]*5
func nthUglyNumber(n int) int {
	a, b, c := 0, 0, 0
	arr := make([]int, n)
	arr[0] = 1
	for i := 1; i < n; i++ {
		na, nb, nc := arr[a]*2, arr[b]*3, arr[c]*5
		arr[i] = min(na, min(nb, nc))
		if arr[i] == na {
			a++
		}
		if arr[i] == nb {
			b++
		}
		if arr[i] == nc {
			c++
		}
	}
	return arr[n-1]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
