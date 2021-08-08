package week4

/*
   以末尾元素为基准，大于记录为1 小于等于记为0
   11100
   本题可以改为小于等于末尾元素的最小元素
*/

func findMin(nums []int) int {
	l, r := 0, len(nums)-1

	for l < r {
		mid := (l + r) / 2
		//条件
		if nums[mid] <= nums[len(nums)-1] {
			//往左
			r = mid
		} else {
			l = mid + 1
		}
	}
	return nums[r]
}
