package week1

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
/*
	1. 大体分三步： 将链表分组，对每个分组内的链表进行反转，处理相邻分组间的连接
    2. 处理相邻分组间的连接：1-2-3-4 2-1-4-3 找到上一个分组反转后的end 指向下一个分组反转后的head
	3. 注意题目要求，如果最后一个分组链表节点数小于k，则不反转

	整体：遍历分组的链表，对于每一个小组，先反转，然后找到上一个分组的end，指向自己
*/
func reverseKGroup(head *ListNode, k int) *ListNode {
	nodeGroupList := groupListNodes(head,k)

	res := reverseNodeList(nodeGroupList[0],k)

	for i := 1;i < len(nodeGroupList);i++{
		end := getEnd(nodeGroupList[i-1],k)
		if i == len(nodeGroupList) - 1 && getLength(nodeGroupList[i]) < k{
			end.Next = nodeGroupList[i]
		}else{
			end.Next = reverseNodeList(nodeGroupList[i],k)
		}
	}

	return res
}


//返回一个列表
func groupListNodes(head *ListNode,k int) []*ListNode{
	var result []*ListNode
	idx := 0
	//遍历前n个
	for head != nil{
		if idx % k == 0{
			result  = append(result,head)
		}
		idx++
		head = head.Next
	}
	return result
}

//反转链表前k个
func reverseNodeList(head *ListNode,k int) *ListNode{
	var pre *ListNode
	for head != nil && k > 0{
		//先保存一下next和pre，因为后面会覆盖
		nextNode := head.Next
		//将自己指向前一个
		head.Next = pre

		//主体对链表的遍历,head向后，pre同时也向后
		pre = head
		head = nextNode
		k--
	}
	return pre
}

func getEnd(head *ListNode,k int) *ListNode{
	//定义保护节点执行head
	for head.Next != nil && k > 0{
		head = head.Next
		k--
	}
	return head
}

func getLength(head *ListNode) int{
	//定义保护节点执行head
	k := 0
	for head != nil{
		head = head.Next
		k++
	}
	return k
}

