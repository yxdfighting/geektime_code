package week1


/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

/*
1-2-3-4

1. 对于链表的处理，首先主体是遍历链表,head != nil,这样可以遍历到链表的每一个元素
2. 如果是反转，那么需要修改n条边 4-3-2-1-nil  1-2-3-4-nil
3. 对于每一个节点，要实现反转，完成一个工作，将自己指向前一个节点，单链表没有这个信息，所以需要存下来
4. 遍历到最后依次循环，head为最后一个节点，继续向后，head=nil pre=head，所以应该返回pre
5. 注意go里面的初始化，var pre *ListNode 此时pre是一个指针类型的零值，表示该指针为指向另外一个地址，初始化为nil
pre := *ListNode{},此时给pre赋值了一个空的结构体，空结构的零值{0,nil}
 */

type ListNode struct {
	Val int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	//定义保护节点执行head
	var pre *ListNode
	for head != nil{
		//先保存一下next和pre，因为后面会覆盖
		nextNode := head.Next
		//将自己指向前一个
		head.Next = pre

		//主体对链表的遍历,head向后，pre同时也向后
		pre = head
		head = nextNode
	}
	return pre
}