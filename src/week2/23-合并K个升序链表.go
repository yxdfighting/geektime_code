package week2

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}

	if len(lists) == 1 {
		return lists[0]
	}

	halfLen := len(lists) / 2
	//前半段合并
	preLists := mergeKLists(lists[:halfLen])
	postLists := mergeKLists(lists[halfLen:])

	return mergeTwoLists(preLists, postLists)
}

//合并两个有序链表
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

	if l1.Val < l2.Val {
		l1.Next = mergeTwoLists(l1.Next, l2)
		return l1
	} else {
		l2.Next = mergeTwoLists(l2.Next, l1)
		return l2
	}
}
