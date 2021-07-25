package week4

//考虑一个问题，如何避免窗口移动时，新元素加进来，旧元素出堆不好处理的问题
/*
   每次只进新的，不除旧的 会有什么问题
   如果算堆顶的时候正好在某个旧元素，则移除堆顶
*/

func maxSlidingWindow(nums []int, k int) []int {
	i := 0
	var res []int
	var mylist Mylist

	for i <= len(nums)-1 {
		//依次将元素加进来，构建一个二叉堆，到k个时，将堆顶入结果
		mylist = build(mylist, nums[i], i)
		// fmt.Printf("i:%v,mylist:%v\n",i,mylist)
		if i >= k-1 {
			top := mylist.Top()
			for top.idx <= i-k {
				mylist = Pop(mylist)
				top = mylist.Top()
			}
			res = append(res, top.val)
		}
		i++
	}
	return res
}

type WindowObj struct {
	val int
	idx int
}

type Mylist []WindowObj

func (list Mylist) Top() WindowObj {
	if len(list) == 0 {
		return WindowObj{}
	}
	return list[0]
}

//移除堆顶
func Pop(list Mylist) Mylist {
	//先将最后的节点放到堆顶
	swap(0, len(list)-1, list)
	list = list[:len(list)-1]
	//每次与左右子节点比较，跟较小的节点交换
	idx := 0
	for 2*idx+1 <= len(list)-1 {
		minIdx := getMax(2*idx+1, 2*idx+2, list)
		// fmt.Printf("idx,minIdx:%v.%v\n",idx,minIdx)
		if list[minIdx].val > list[idx].val {
			swap(idx, minIdx, list)
			idx = minIdx
		} else {
			break
		}
	}
	return list
}

//build
func build(list Mylist, new int, index int) Mylist {
	//先将node节点放在最后
	node := WindowObj{
		new, index,
	}
	list = append(list, node)
	idx := len(list) - 1
	//向上交换
	for (idx-1)/2 >= 0 {
		// fmt.Printf("nllll:%v\n",nl)
		if list[idx].val > list[(idx-1)/2].val {
			swap(idx, (idx-1)/2, list)
			idx = (idx - 1) / 2
		} else {
			break
		}
	}
	return list
	// fmt.Printf("nllll:%v\n",nl)
}

func swap(i, j int, m Mylist) {
	m[i], m[j] = m[j], m[i]
}

func getMax(l, r int, m Mylist) int {
	if r >= len(m) {
		return l
	}
	if m[l].val < m[r].val {
		return r
	} else {
		return l
	}
}
