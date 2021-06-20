package week1

/*
	用两个栈，origin存原始数据，min每次对要入栈的元素与栈顶比较，每次放一个较小值在栈顶
*/

type MinStack struct {
	min    []int
	origin []int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	min := make([]int, 0)
	origin := make([]int, 0)

	minStack := &MinStack{
		min, origin,
	}

	return *minStack
}

func (this *MinStack) Push(val int) {
	this.origin = append(this.origin, val)
	if len(this.min) == 0 {
		this.min = append(this.min, val)
		return
	}
	if val < this.min[len(this.min)-1] {
		this.min = append(this.min, val)
	} else {
		this.min = append(this.min, this.min[len(this.min)-1])
	}
}

func (this *MinStack) Pop() {
	this.origin = this.origin[0 : len(this.origin)-1]
	this.min = this.min[0 : len(this.min)-1]
}

func (this *MinStack) Top() int {
	return this.origin[len(this.origin)-1]
}

func (this *MinStack) GetMin() int {
	return this.min[len(this.min)-1]
}
