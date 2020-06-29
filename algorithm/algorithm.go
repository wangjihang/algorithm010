package algorithm

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(preorder []int, inorder []int) *TreeNode {
	var (
		p, i   int
		length = len(preorder)
		build  func(stop int) *TreeNode
	)
	build = func(stop int) *TreeNode {
		if p >= length {
			return nil
		}
		if inorder[i] != stop {
			root := &TreeNode{preorder[p], nil, nil}
			p++
			root.Left = build(root.Val) // 结束说明碰到stop了
			i++
			root.Right = build(stop)
			return root
		}
		return nil
	}
	return build(-1)
}
