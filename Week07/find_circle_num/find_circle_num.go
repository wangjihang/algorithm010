package find_circle_num

// 547
// dfs
// 时间复杂度：O(N^2)
// 空间复杂度：O(N)
//func findCircleNum(M [][]int) int {
//	if len(M) == 0 {
//		return 0
//	}
//
//	var (
//		ret     int
//		length  = len(M)
//		visited = make([]bool, length)
//		dfs     func(begin int) // 染色
//	)
//	dfs = func(i int) {
//		for j := 0; j < length; j++ {
//			if M[i][j] == 1 && !visited[j] {
//				visited[j] = true
//				dfs(j)
//			}
//		}
//	}
//
//	for i := 0; i < length; i++ {
//		if !visited[i] {
//			dfs(i)
//			ret++
//		}
//	}
//	return ret
//}

// 使用并查集
// 时间复杂度：O(N^2)
// 空间复杂度：O(N)
func findCircleNum(M [][]int) int {
	if len(M) == 0 {
		return 0
	}
	length := len(M)
	ufind := NewUnionFind(length)

	for i := 0; i < length; i++ {
		for j := i + 1; j < length; j++ {
			if M[i][j] == 1 {
				ufind.Union(i, j)
			}
		}
	}
	return ufind.count
}

type UnionFind struct {
	count  int
	parent []int
}

func NewUnionFind(n int) *UnionFind {
	parent := make([]int, n)
	for i := 0; i < n; i++ {
		parent[i] = i
	}
	return &UnionFind{
		count:  n,
		parent: parent,
	}
}

func (u *UnionFind) Union(i, j int) {
	pi := u.Find(i)
	pj := u.Find(j)
	if pi != pj {
		u.parent[pi] = pj
		u.count--
	}
}

func (u *UnionFind) Find(i int) int {
	root := i
	for u.parent[root] != root {
		root = u.parent[root]
	}
	for u.parent[i] != i { // 路径压缩
		i, u.parent[i] = u.parent[i], root
	}
	return root
}
