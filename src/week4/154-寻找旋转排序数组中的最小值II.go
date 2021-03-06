package week4

/*
   以最末尾的数作为判定基准
	小于末尾元素往左走，大于往右走，
	等于可能往左，也可能往右，分治
*/

func findMin(nums []int) int {
	//因为解一定存在，所以0->n-1
	l, r := 0, len(nums)-1
	for l < r {
		//因为是往左找 所以l+r 而不是l+r+1
		mid := (l + r) / 2
		if nums[mid] == nums[r] {
			//此处考虑为什么要把mid放在前半部分，考虑到最后只有两个元素，
			//此时两个元素都会如果都放在后半部分的话，就会出现死循环
			return min(findMin(nums[l:mid+1]), findMin(nums[mid+1:r+1]))
		}
		if nums[mid] < nums[r] {
			r = mid
		} else {
			l = mid + 1
		}
	}
	return nums[r]
}

func min(x, y int) int {
	if x < y {
		return x
	} else {
		return y
	}
}
