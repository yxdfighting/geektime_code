package week4

/*
   本质上是两个问题，
   左边界 --- 大于等于target的最小元素
   右边界 --- 小于等于target的最大元素
*/

func searchRange(nums []int, target int) []int {
	l, r := 0, len(nums)
	var res []int

	for l < r {
		mid := (l + r) / 2
		if nums[mid] >= target {
			r = mid
		} else {
			l = mid + 1
		}
	}

	res = append(res, r)

	l, r = -1, len(nums)-1

	for l < r {
		mid := (l + r + 1) / 2
		if nums[mid] <= target {
			l = mid
		} else {
			r = mid - 1
		}
	}
	res = append(res, r)

	if res[0] > res[1] {
		return []int{-1, -1}
	}
	return res
}
