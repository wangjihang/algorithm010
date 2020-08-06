package main

func main() {
	canFinish(2, [][]int{{0, 1}})
}

// dfs
func canFinish(numCourses int, prerequisites [][]int) bool {
	var (
		edges   = make([][]int, numCourses)
		visited = make([]int, numCourses) // 0 未开始, 1 开始搜索, 2 已完成
		dfs     func(node int) bool
	)
	for _, vs := range prerequisites {
		edges[vs[1]] = append(edges[vs[1]], vs[0])
	}
	dfs = func(node int) bool {
		if visited[node] == 2 {
			return true
		}
		visited[node] = 1
		for _, v := range edges[node] {
			if visited[v] == 1 || !dfs(v) {
				return false
			}
		}
		visited[node] = 2
		return true
	}
	for i := 0; i < numCourses; i++ {
		if visited[i] == 0 && !dfs(i) {
			return false
		}
	}
	return true
}
