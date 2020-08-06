package main

import (
	"sort"
)

func main() {
}

// m*(n+1)+2k
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
