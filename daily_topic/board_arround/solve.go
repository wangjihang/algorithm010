package board_arround

// 130. 被围绕的区域
// dfs 从边缘向其他方向扩散
// 时间复杂度: O(M*N)
func solve(board [][]byte) {
	if len(board) <= 2 { // 不存在需要替换的O
		return
	}

	// 1. 遍历，从边缘找到O,然后dfs所有连接的O为#
	m, n := len(board), len(board[0])
	for i := 0; i < m; i++ {
		dfs(i, 0, board)
		dfs(i, n-1, board)
	}
	for j := 0; j < n; j++ {
		dfs(0, j, board)
		dfs(m-1, j, board)
	}

	// 2. 将剩余的O变为X，将剩余的#变成O
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if board[i][j] == 'O' {
				board[i][j] = 'X'
			} else if board[i][j] == '#' {
				board[i][j] = 'O'
			}
		}
	}
}

func dfs(i, j int, board [][]byte) {
	// terminal
	if i < 0 || j < 0 || i >= len(board) || j >= len(board[0]) || board[i][j] != 'O' {
		return
	}

	// process
	board[i][j] = '#'

	// drill down
	dfs(i-1, j, board) // 上
	dfs(i+1, j, board) // 下
	dfs(i, j-1, board) // 左
	dfs(i, j+1, board) // 右
}

// 并查集
// 思路：通过并查集，将满足条件的均连通起来，最后遍历并从并查集里面查询是否同一个parent，
// 如果是，则连通，否则需要替换为X
//func solve(board [][]byte) {
//	if len(board) <= 2 {
//		return
//	}
//
//	var (
//		m, n         = len(board), len(board[0])
//		dummy, ufind = m * n, NewUnionFind(m*n + 1)
//		position     func(i, j int) int
//	)
//	position = func(i, j int) int {
//		return i*n + j
//	}
//
//	for i := 0; i < m; i++ {
//		for j := 0; j < n; j++ {
//			if board[i][j] == 'O' {
//				if i == 0 || j == 0 || i == m-1 || j == n-1 {
//					ufind.Union(position(i, j), dummy)
//				}
//				if i > 0 && board[i-1][j] == 'O' {
//					ufind.Union(position(i, j), position(i-1, j))
//				}
//				if j > 0 && board[i][j-1] == 'O' {
//					ufind.Union(position(i, j), position(i, j-1))
//				}
//			}
//		}
//	}
//
//	for i := 0; i < m; i++ {
//		for j := 0; j < n; j++ {
//			if board[i][j] == 'O' && !ufind.IsConnected(position(i, j), dummy) {
//				board[i][j] = 'X'
//			}
//		}
//	}
//}
//
//type UnionFind struct {
//	parent []int
//}
//
//func NewUnionFind(n int) *UnionFind {
//	parent := make([]int, n)
//	for i := 0; i < n; i++ {
//		parent[i] = i
//	}
//	return &UnionFind{
//		parent: parent,
//	}
//}
//
//func (u *UnionFind) Union(i, j int) {
//	pi := u.Find(i)
//	pj := u.Find(j)
//	if pi != pj {
//		u.parent[pi] = pj
//	}
//}
//
//func (u *UnionFind) Find(i int) int {
//	root := i
//	for u.parent[root] != root {
//		root = u.parent[root]
//	}
//	for u.parent[i] != root {
//		i, u.parent[i] = u.parent[i], root
//	}
//	return root
//}
//
//func (u *UnionFind) IsConnected(i, j int) bool {
//	return u.Find(i) == u.Find(j)
//}
