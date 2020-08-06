//你这个学期必须选修 numCourse 门课程，记为 0 到 numCourse-1 。
//
// 在选修某些课程之前需要一些先修课程。 例如，想要学习课程 0 ，你需要先完成课程 1 ，我们用一个匹配来表示他们：[0,1]
//
// 给定课程总量以及它们的先决条件，请你判断是否可能完成所有课程的学习？
//
//
//
// 示例 1:
//
// 输入: 2, [[1,0]]
//输出: true
//解释: 总共有 2 门课程。学习课程 1 之前，你需要完成课程 0。所以这是可能的。
//
// 示例 2:
//
// 输入: 2, [[1,0],[0,1]]
//输出: false
//解释: 总共有 2 门课程。学习课程 1 之前，你需要先完成​课程 0；并且学习课程 0 之前，你还应先完成课程 1。这是不可能的。
//
//
//
// 提示：
//
//
// 输入的先决条件是由 边缘列表 表示的图形，而不是 邻接矩阵 。详情请参见图的表示法。
// 你可以假定输入的先决条件中没有重复的边。
// 1 <= numCourses <= 10^5
//
// Related Topics 深度优先搜索 广度优先搜索 图 拓扑排序
// 👍 480 👎 0

package main

func main() {
	canFinish(2, [][]int{{1, 0}, {0, 1}})
}

//leetcode submit region begin(Prohibit modification and deletion)
// 广度优先 DAG
// 拓扑排序的长度 等于节点长度，说明没有闭环
func canFinish(numCourses int, prerequisites [][]int) bool {
	var (
		edges   = make([][]int, numCourses)
		innodes = make([]int, numCourses) // 入节点的个数
		queue   = make([]int, 0, numCourses)
	)

	// 初始化
	for _, requisites := range prerequisites {
		edges[requisites[1]] = append(edges[requisites[1]], requisites[0])
		innodes[requisites[0]]++
	}

	for i, v := range innodes {
		if v == 0 {
			queue = append(queue, i)
		}
	}
	for ; len(queue) > 0 && numCourses > 0; queue, numCourses = queue[1:], numCourses-1 {
		node := queue[0] // pop
		for _, v := range edges[node] {
			if innodes[v]--; innodes[v] == 0 {
				queue = append(queue, v)
			}
		}
	}
	return len(queue) == 0 && numCourses == 0
}

// 深度优先
// DAG 有向无环图
// 我们将当前搜索的节点 u 标记为「搜索中」，遍历该节点的每一个相邻节点 v：
// 如果 v 为「未搜索」，那么我们开始搜索 v，待搜索完成回溯到 u； 0
// 如果 v 为「搜索中」，那么我们就找到了图中的一个环，因此是不存在拓扑排序的； 1
// 如果 v 为「已完成」，那么说明 v 已经在栈中了，而 u 还不在栈中，因此 u 无论何时入栈都不会影响到 (u, v) 之前的拓扑关系，以及不用进行任何操作。 2
// 当 u 的所有相邻节点都为「已完成」时，我们将 u 放入栈中，并将其标记为「已完成」。
// 时间复杂度：O(M+N)
//func canFinish(numCourses int, prerequisites [][]int) bool {
//	var (
//		edges   = make([][]int, numCourses)
//		visited = make([]int, numCourses) // 0[未搜索]  1[搜索中]  2[已完成]
//		dfs     func(node int) bool
//	)
//	dfs = func(node int) bool {
//		if visited[node] == 2 {
//			return true
//		}
//
//		visited[node] = 1
//		for _, v := range edges[node] {
//			if visited[v] == 1 || !dfs(v) {
//				return false
//			}
//		}
//		visited[node] = 2
//		return true
//	}
//
//	// 初始化
//	for _, requisites := range prerequisites {
//		edges[requisites[1]] = append(edges[requisites[1]], requisites[0])
//	}
//	for i := 0; i < numCourses; i++ {
//		if visited[i] == 0 && !dfs(i) {
//			return false
//		}
//	}
//
//	return true
//}

//leetcode submit region end(Prohibit modification and deletion)
