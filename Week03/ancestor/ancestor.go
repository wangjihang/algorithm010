package ancestor

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// (1） 如果当前结点 root 等于 NULL，则直接返回 NULL
//（2） 如果 root 等于 p 或者 q ，那这棵树一定返回 p 或者 q
//（3） 然后递归左右子树，因为是递归，使用函数后可认为左右子树已经算出结果，用 left 和 right 表示
//（4） 此时若left为空，那最终结果只要看 right；若 right 为空，那最终结果只要看 left
//（5） 如果 left 和 right 都非空，因为只给了 p 和 q 两个结点，都非空，说明一边一个，因此 root 是他们的最近公共祖先
//（6） 如果 left 和 right 都为空，则返回空（其实已经包含在前面的情况中了）
// 时间复杂度: O(n)
// 空间复杂度: O(n)
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Val == p.Val || root.Val == q.Val {
		return root
	}

	left := lowestCommonAncestor(root.Left, p, q)
	right := lowestCommonAncestor(root.Right, p, q)

	if left == nil {
		return right
	}
	if right == nil {
		return left
	}
	return root
}
