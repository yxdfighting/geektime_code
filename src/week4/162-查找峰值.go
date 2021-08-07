package week4

/*
	三分查找极值
    考虑场景，峰值元素是指大于左右相邻值的元素
*/
func findPeakElement(nums []int) int {
	l, r := 0, len(nums)-1

	for l < r {
		lmid := (l + r) / 2
		rmid := lmid + 1

		//确定搜索条件
		if nums[lmid] <= nums[rmid] {
			l = lmid + 1
		} else {
			r = rmid - 1
		}
	}

	return r
}
