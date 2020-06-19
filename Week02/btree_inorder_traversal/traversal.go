package btree_inorder_traversal

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 中序遍历
// left-root-right
func inorderTraversal(root *TreeNode) []int {
	return inorderRecursive(root)
}

// recursive

func inorderRecursive(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	result := append(inorderRecursive(root.Left), root.Val)
	result = append(result, inorderRecursive(root.Right)...)
	return result
}

// stack iterative

func inorderIterative(root *TreeNode) []int {
	curr := root
	stack := Stack{}
	result := make([]int, 0)
	for curr != nil || !stack.IsEmpty() {
		for curr != nil {
			stack.Push(curr)
			curr = curr.Left
		}
		curr = stack.Pop()
		result = append(result, curr.Val)
		curr = curr.Right
	}
	return result
}

type Stack []*TreeNode

func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}

func (s *Stack) Push(node *TreeNode) {
	*s = append(*s, node)
}

func (s *Stack) Pop() *TreeNode {
	n := []*TreeNode(*s)[len(*s)-1]
	*s = []*TreeNode(*s)[:len(*s)-1]
	return n
}

// morris
// curr 为 root
// 1. 存在左子节点，并将其作为左子节点的右节点的最右子节点
// 2. 不存在左子节点，就遍历其右节点，如果遇到上面情况，就换一下
// 结局是线性树
// 时间复杂度: O(n)
// 空间复杂度: O(n)
func inorderMorris(root *TreeNode) []int {
	result := make([]int, 0)
	for root != nil {
		if root.Left != nil {
			node := root.Left // 寻找左树最右节点
			for node.Right != nil {
				node = node.Right
			}
			node.Right = root
			root, root.Left = root.Left, nil
		} else { // 说明自己已经是最左节点
			result = append(result, root.Val)
			root = root.Right
		}
	}
	return result
}
