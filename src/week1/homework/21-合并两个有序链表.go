package homework

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

/*
    最开始没有想到递归。。。。

	考虑不增加空间复杂度，就用现在的链表节点修改指向
	1. 主体思路：依次遍历两个链表，比较节点元素大小，较小的留下，min.Next 与 max继续递归
*/

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil && l2 == nil {
		return nil
	}

	if l1 == nil && l2 != nil {
		return l2
	}

	if l2 == nil && l1 != nil {
		return l1
	}

	if l1.Val <= l2.Val {
		l1.Next = mergeTwoLists(l1.Next, l2)
		return l1
	} else {
		l2.Next = mergeTwoLists(l1, l2.Next)
		return l2
	}
}
