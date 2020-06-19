package n_tree_preorder

type Node struct {
	Val      int
	Children []*Node
}

// 递归

func preorder(root *Node) []int {
	return PreorderIterative(root)
}

func PreorderRecursive(root *Node) []int {
	if root == nil {
		return nil
	}
	result := []int{root.Val}
	for _, node := range root.Children {
		result = append(result, PreorderRecursive(node)...)
	}
	return result
}

// 性能更佳 内存分配次数更少
// 时间复杂度：O(n)
// 空间复杂度：O(n)
func PreorderRecursive2(root *Node) []int {
	var (
		result []int
		f      func(*Node)
	)
	f = func(root *Node) {
		if root == nil {
			return
		}
		result = append(result, root.Val)
		for _, node := range root.Children {
			f(node)
		}
	}
	f(root)
	return result
}

// 时间复杂度：O(n) 所有节点 有且只计算一次
// 空间复杂度：O(n) 使用栈
func PreorderIterative(root *Node) []int {
	if root == nil {
		return nil
	}
	var (
		result []int
		stack  Stack
	)
	stack.Push(root)
	for !stack.IsEmpty() {
		curr := stack.Pop()
		result = append(result, curr.Val)
		for i := len(curr.Children) - 1; i >= 0; i-- { // 逆序入栈 出栈才是顺序
			stack.Push(curr.Children[i])
		}
	}
	return result
}

type Stack []*Node

func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}

func (s *Stack) Push(node *Node) {
	*s = append(*s, node)
}

func (s *Stack) Pop() *Node {
	n := []*Node(*s)[len(*s)-1]
	*s = []*Node(*s)[:len(*s)-1]
	return n
}
