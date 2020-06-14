//合并两个有序链表
//将两个升序链表合并为一个新的 升序 链表并返回。新链表是通过拼接给定的两个链表的所有节点组成的。
//
//示例：
//
//输入：1->2->4, 1->3->4
//输出：1->1->2->3->4->4

package merge_two_lists

type ListNode struct {
	Val  int
	Next *ListNode
}

// 使用哨兵节点降低复杂度
// 时间复杂度: O(m+n)
// 空间复杂度: O(1)
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	var (
		prehead = &ListNode{}
		pre     = prehead
	)

	for l1 != nil && l2 != nil {
		if l1.Val <= l2.Val {
			pre.Next = l1
			l1 = l1.Next
		} else {
			pre.Next = l2
			l2 = l2.Next
		}
		pre = pre.Next
	}

	if l1 != nil {
		pre.Next = l1
	}

	if l2 != nil {
		pre.Next = l2
	}

	return prehead.Next
}

// 递归解法：
// 时间复杂度：O(m+n) 每个节点有且仅有进行一次计算
// 空间复杂度: O(m+n) 栈空间
func mergeTwoLists2(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}

	if l1.Val < l2.Val {
		l1.Next = mergeTwoLists2(l1.Next, l2)
		return l1
	} else {
		l2.Next = mergeTwoLists2(l2.Next, l1)
		return l2
	}
}
