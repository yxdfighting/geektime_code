package week4

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

type MyNode struct {
	key  int
	node *ListNode
}

/*
    自己实现一个小根堆，每次比较返回最小的元素
    需要实现几个方法
    1、取堆顶
    2、构建二叉堆
    3、移除堆顶


    从0开始
       0
    1     2
  3   4  5   6
*/

type MyListNode []MyNode

func (nl MyListNode) Top() *ListNode {
	if len(nl) == 0 {
		return nil
	}
	return nl[0].node
}

//移除堆顶
func Pop(nl MyListNode) MyListNode {
	//先将最后的节点放到堆顶
	swap(0, len(nl)-1, nl)
	nl = nl[:len(nl)-1]
	//每次与左右子节点比较，跟较小的节点交换
	idx := 0
	for 2*idx+1 <= len(nl)-1 {
		minIdx := getMin(2*idx+1, 2*idx+2, nl)
		// fmt.Printf("idx,minIdx:%v.%v\n",idx,minIdx)
		if nl[minIdx].key < nl[idx].key {
			swap(idx, minIdx, nl)
			idx = minIdx
		} else {
			break
		}
	}
	return nl
}

//build
func build(nl MyListNode, node MyNode) MyListNode {
	//先将node节点放在最后
	nl = append(nl, node)
	idx := len(nl) - 1
	//向上交换
	for (idx-1)/2 >= 0 {
		// fmt.Printf("nllll:%v\n",nl)
		if nl[idx].key < nl[(idx-1)/2].key {
			swap(idx, (idx-1)/2, nl)
			idx = (idx - 1) / 2
		} else {
			break
		}
	}
	return nl
	// fmt.Printf("nllll:%v\n",nl)
}

func swap(i, j int, m MyListNode) {
	m[i], m[j] = m[j], m[i]
}

func getMin(l, r int, m MyListNode) int {
	if r >= len(m) {
		return l
	}
	if m[l].key > m[r].key {
		return r
	} else {
		return l
	}
}

func mergeKLists(lists []*ListNode) *ListNode {
	nl := MyListNode{}
	//构建二叉堆
	for _, list := range lists {
		for list != nil {
			m := MyNode{
				list.Val, list,
			}
			nl = build(nl, m)
			list = list.Next
		}
	}

	// fmt.Printf("nl:%v\n",nl)

	//定义链表的保护节点
	var head *ListNode
	head = nl.Top()
	dummy := ListNode{}
	dummy.Next = head
	//不断取二叉堆的堆顶
	for head != nil {
		nl = Pop(nl)
		// fmt.Printf("nl:%v\n",nl)
		temp := nl.Top()
		head.Next = temp
		head = head.Next
		// fmt.Printf("res:%v\n",temp.Val)
	}
	return dummy.Next
}
