package week1

/*
	主体思路，如果存在环，用不同速率指针去遍历，一定会相遇
*/

func hasCycle(head *ListNode) bool {
	slow, fast := head, head

	for slow != nil && fast != nil {
		//注意细节，如果fast.Next为nil，程序会panic，所以需要先判断一下fast.Next是否为nil
		if fast.Next == nil {
			return false
		}
		slow = slow.Next
		fast = fast.Next.Next

		//如果相遇，表示fast多走一圈，那就一定有环
		if slow == fast {
			return true
		}
	}
	//如果fast先走到nil，表示无环
	return false
}
