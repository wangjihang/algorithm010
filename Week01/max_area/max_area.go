package max_area

// 暴力
// 时间复杂度: O(N^2)
// 空间复杂度: O(1)
func maxArea(height []int) int {
	var max int
	length := len(height)
	for i := 0; i < length-1; i++ {
		for j := i + 1; j < length; j++ {
			area := area(j-i, min(height[i], height[j]))
			if area > max {
				max = area
			}
		}
	}
	return max
}

// 双指针法
// 时间复杂度：O(n)
// 空间复杂度：O(1)
func maxArea2(height []int) int {
	i, j, ret := 0, len(height)-1, 0
	for i != j {
		if height[i] < height[j] {
			ret = max(ret, area(j-i, height[i]))
			i++
		} else {
			ret = max(ret, area(j-i, height[j]))
			j--
		}
	}
	return ret
}

func area(weight, height int) int {
	return weight * height
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
