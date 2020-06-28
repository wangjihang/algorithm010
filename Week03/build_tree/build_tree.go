package build_tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// preorder [root][left][right]
// inorder [left][root][right]
// 时间O(n)
// 空间O(h) 树的深度
func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	root := &TreeNode{preorder[0], nil, nil}
	i := 0
	for key, value := range inorder {
		if preorder[0] == value {
			i = key
			break
		}
	}
	root.Left = buildTree(preorder[1:len(inorder[:i])+1], inorder[:i])
	root.Right = buildTree(preorder[len(inorder[:i])+1:], inorder[i+1:])
	return root
}
