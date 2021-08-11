package sortAlgol

func HeapSort(nums []int) {
	//对未排序部分进行堆排序 选出堆顶 放到已排序末尾
	myHeap := &MyHeap{}
	for i := 0; i < len(nums); i++ {
		myHeap.Push(nums[i])
	}
	idx := 0
	for !myHeap.Empty() && idx < len(nums) {
		nums[idx] = myHeap.Top()
		myHeap.Pop()
		idx++
	}
}

//先手动实现一个堆
type MyHeap struct {
	sliceForHeap []int
}

func (h *MyHeap) Top() int {
	sliceForHeap := h.sliceForHeap
	return sliceForHeap[0]
}

func (h *MyHeap) Empty() bool {
	return len(h.sliceForHeap) == 0
}

/*
    0
  1   2
3  4 5  6
*/
func (h *MyHeap) Push(val int) {
	sliceForHeap := h.sliceForHeap
	//push之后每次放到最后
	sliceForHeap = append(sliceForHeap, val)
	//依次向上交换
	nowIdx := len(sliceForHeap) - 1
	for nowIdx > 0 {
		//与父节点比较，交换
		if sliceForHeap[nowIdx] < sliceForHeap[(nowIdx-1)/2] {
			sliceForHeap[nowIdx], sliceForHeap[(nowIdx-1)/2] =
				sliceForHeap[(nowIdx-1)/2], sliceForHeap[nowIdx]
			nowIdx = (nowIdx - 1) / 2
		} else {
			break
		}
	}
	h.sliceForHeap = sliceForHeap
}

func (h *MyHeap) Pop() {
	sliceForHeap := h.sliceForHeap
	//把头去掉
	sliceForHeap[0], sliceForHeap[len(sliceForHeap)-1] =
		sliceForHeap[len(sliceForHeap)-1], sliceForHeap[0]

	sliceForHeap = sliceForHeap[0 : len(sliceForHeap)-1]

	//从上往下交换
	idx := 0
	for 2*idx+2 < len(sliceForHeap) {
		midIdx := findMin(sliceForHeap, 2*idx+1, 2*idx+2)
		if sliceForHeap[idx] > sliceForHeap[midIdx] {
			sliceForHeap[midIdx], sliceForHeap[idx] =
				sliceForHeap[idx], sliceForHeap[midIdx]
			idx = midIdx
		} else {
			break
		}
	}
	h.sliceForHeap = sliceForHeap
}

func findMin(nums []int, aIdx, bIdx int) int {
	if nums[aIdx] > nums[bIdx] {
		return bIdx
	} else {
		return aIdx
	}
}
